from django.http import JsonResponse
from utils import baseview
from approval.models import FApproval, User, FInstanceValue, FITask, FANode, FIForm, DeployHistory, BusinessEnvironment
from approval.serializers import FIFormSerializer
from approval.apis.v2.instance import send_user
from utils.call_jenkins import JenkinsApi
from approval.apis.audit import PutAudit
import time
import os

os.environ.setdefault('DJANGO_SETTINGS_MODULE', 'djentry.settings')

KVS = {
    "项目名称": "PROJECT",
    "仓库名称": "GIT_NAME",
    "分支名称": "BrName",
    "影响评估": "LEVEL",
    "ONES ID": "ONES_ID",
    "服务类型": "SERVICE_TYPE",
    "服务名称": "SERVICE",
    "期望发布时间": "DATE"
}
NEEDS = ["服务类型", "分支名称", "服务名称", "项目名称"]
NEEDS2 = ["SERVICE_TYPE", "BrName", "SERVICE", "PROJECT"]
SKIPS = ["ecoplants-mobile", "ecoplants-web"]


def formToObjEn(data):
    result = []
    keys = []
    for item in data:
        if item["service"] not in keys:
            keys.append(item["service"])
    for service in keys:
        obj = {}
        for item in data:
            if item["service"] == service:
                obj[item["en_name"]] = item["value"]
        obj["SERNum"] = service
        result.append(obj)
    return result


def jenkins_deploy(job_name, instance_id, task_id):
    fiforms = FIForm.objects.filter(f_instance_id=instance_id, del_tag=0)
    data = FIFormSerializer(fiforms, many=True).data
    result = formToObjEn(data)
    serviceTotal = len(result)
    i = 0
    for form in result:
        task = FITask.objects.filter(id=task_id).first()
        now = time.strftime("%Y-%m-%d %H:%M:%S", time.localtime())
        if (now > form["DATE"]) and (','+form["SERVICE"]+',' not in task.comment):
            print("工单编号 %d 开始发布第 %d 个服务，一共有 %d 个服务" %
                  (int(instance_id), int(form["SERNum"]), int(serviceTotal)))
            if form["SERVICE_TYPE"] not in SKIPS:
                obj = {}
                for key in NEEDS2:
                    obj[key] = form[key]
                fi = FInstanceValue.objects.filter(id=instance_id).first()
                user = User.objects.filter(id=fi.user_id).first()
                if user:
                    un = user.last_name.strip() + user.first_name.strip()
                    if un:
                        obj["USERNAME"] = un
                    else:
                        obj["USERNAME"] = user.username
                else:
                    obj["USERNAME"] = "未知<无权获取>"
                if (fi.job_number == 0) and (fi.status != "DEPLOY"):
                    jenkinsApi = JenkinsApi()
                    num = jenkinsApi.get_next_buildnum(job_name)
                    FInstanceValue.objects.filter(id=instance_id).update(
                        status="DEPLOY", job_number=num)
                    jenkinsApi.execute_job(job_name, obj)
                else:
                    if i > 0:
                        lastForm = result[i - 1]
                        lastService = lastForm["SERVICE"]
                        if lastService in task.comment:
                            jenkinsApi = JenkinsApi()
                            num = jenkinsApi.get_next_buildnum(job_name)
                            FInstanceValue.objects.filter(
                                id=instance_id).update(job_number=num)
                            jenkinsApi.execute_job(job_name, obj)
                        else:
                            print("上一个服务还未发布结束, 请稍等")
                break
            else:
                print("%s 无需jenkins构建" % (form["SERVICE"]))
                FITask.objects.filter(id=task_id).update(
                    comment="当前项目不走 jenkins 发布,自动审批通过",
                    end_time=now,
                    status="DONE")
                env = BusinessEnvironment.objects.filter(
                    job_name=job_name).first()
                DeployHistory.objects.create(env=env.name,
                                             project=form["PROJECT"],
                                             service=form["SERVICE"],
                                             BrName=form["BrName"],
                                             deploy_time=now,
                                             result="UNKNOWN: 当前项目不走jenkins",
                                             del_tag=0)
        i += 1


def ends_step(task, fi_id):
    start = time.strftime("%Y-%m-%d %H:%M:%S", time.localtime())
    FITask.objects.filter(id=task.id).update(start_time=start,
                                             status="PENDING")
    time.sleep(1)
    end = time.strftime("%Y-%m-%d %H:%M:%S", time.localtime())
    FITask.objects.filter(id=task.id).update(end_time=end,
                                             status="DONE",
                                             comment="工单结束")
    FInstanceValue.objects.filter(id=fi_id).update(status="DONE",
                                                   descriptions="工单结束")


def between_steps(task, fa, fi):
    user = User.objects.filter(id=task.user_id).first()
    node = FANode.objects.filter(id=task.f_node_id).first()
    if node.approval_type == "AUTO_PASS":
        start = time.strftime("%Y-%m-%d %H:%M:%S", time.localtime())
        FITask.objects.filter(id=task.id).update(start_time=start,
                                                 status="PENDING")
        if task.name == "经办人" and task.end_time == "":
            if task.comment == "":
                FITask.objects.filter(id=task.id).update(comment="到点可发布")
                print("审批已经通过，到点可发布，job名称: %s" % (fa.job_name))
            elif task.comment != "自动审批通过":
                jenkins_deploy(fa.job_name, fi.id, task.id)
        else:
            time.sleep(1)
            end = time.strftime("%Y-%m-%d %H:%M:%S", time.localtime())
            FITask.objects.filter(id=task.id).update(end_time=end,
                                                     status="DONE",
                                                     comment="自动审批通过")
    else:
        user2 = User.objects.filter(id=fi.user_id).first()
        if task.status == "BEFORE":
            un = user2.last_name.strip() + user2.first_name.strip()
            if un:
                uname = un
            else:
                uname = user2.username
            send_user(user.email, uname, fa.approval_name, fi.id, "green",
                      "有审批需要你操作")
            now = time.strftime("%Y-%m-%d %H:%M:%S", time.localtime())
            FITask.objects.filter(id=task.id).update(start_time=now,
                                                     status="PENDING")


class ProcessView(baseview.AnyLogin):
    """process view"""
    def post(self, request, args=None):
        fas = FApproval.objects.filter(subscribe=1, old_version=0, del_tag=0)
        for fa in fas:
            fis = FInstanceValue.objects.filter(
                f_approval_id=fa.id,
                status__in=["PENDING", "DEPLOY", "APPROVED"])
            for fi in fis:
                tasks = FITask.objects.filter(
                    f_instance_id=fi.id).order_by("serial_number")
                for task in tasks:
                    if task.status == "PENDING" or task.start_time == "" and task.end_time == "":
                        if task.name == "结束":
                            ends_step(task, fi.id)
                        else:
                            between_steps(task, fa, fi)
                        break
        msg = {"code": 200, "data": {}, "msg": "轮询实例任务成功"}
        return JsonResponse(msg)


class InfosView(baseview.BaseView):
    """infos view"""
    def get(self, request, args=None):
        username = request.user.username
        user = User.objects.filter(username=username).first()
        fas = FApproval.objects.filter(subscribe=1, del_tag=0, old_version=0)
        data = []
        for fa in fas:
            approval_name = fa.approval_name
            fa_id = fa.id
            fis = FInstanceValue.objects.filter(
                f_approval_id=fa_id,
                status__in=["PENDING", "APPROVED"],
                del_tag=0)
            instance_ids = []
            for fi in fis:
                tasks = FITask.objects.filter(f_instance_id=fi.id)
                for task in tasks:
                    if task.status == "PENDING":
                        user1 = User.objects.filter(id=task.user_id).first()
                        if user.id == user1.id:
                            instance_ids.append(fi.id)
                            break
            data.append({
                "approval_name": approval_name,
                "instance_ids": instance_ids
            })
        msg = {"code": 200, "data": data, "msg": "获取审批id成功"}
        return JsonResponse(msg)


class DeployDirectlyView(baseview.BaseView):
    def post(self, request, args=None):
        """直接触发jenkins构建"""
        msg = {"code": 200, "data": {}, "msg": "触发jenkins成功，请注意机器人消息"}
        instance_id = request.data.get('instance_id', '')
        service = int(request.data.get('服务编号', ''))
        fi = FInstanceValue.objects.filter(id=instance_id).first()
        fa = FApproval.objects.filter(id=fi.f_approval_id).first()
        form = FIForm.objects.filter(f_instance_id=instance_id,
                                     service=service,
                                     del_tag=0)
        form = FIFormSerializer(form, many=True).data
        result = formToObjEn(form)[0]
        if result["SERVICE_TYPE"] not in SKIPS:
            obj = {}
            for key in NEEDS2:
                obj[key] = result[key]
            user = User.objects.filter(id=fi.user_id).first()
            if user:
                un = user.last_name.strip() + user.first_name.strip()
                if un:
                    obj["USERNAME"] = un
                else:
                    obj["USERNAME"] = user.username
            else:
                obj["USERNAME"] = "未知<无权获取>"
            jenkinsApi = JenkinsApi()
            num = jenkinsApi.get_next_buildnum(fa.job_name)
            FInstanceValue.objects.filter(id=instance_id).update(
                        descriptions="手动触发构建一次", job_number=num)
            jenkinsApi.execute_job(fa.job_name, obj)
        else:
            msg = {"code": 10001, "data": {}, "msg": "当前项目不通过jenkins构建"}
        PutAudit(request, msg)
        return JsonResponse(msg)
