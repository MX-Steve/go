import json
from django.http import JsonResponse
from utils import baseview
from utils.encr import AESCipher
from utils.feishu_approval import ApprovalT
from utils.call_jenkins import JenkinsApi
from approval.models import FApproval, FIForm, FInstanceValue, FANode, FITask, FAForm, DeployHistory
from approval.models import User
import time

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
SKIPS = ["ecoplants-mobile", "ecoplants-web"]


def input_fiform(form_data, f_instance_id):
    for item in form_data:
        name = item["name"]
        if name in KVS:
            en_name = KVS[name]
            type = item["type"]
            value = item["value"]
            FIForm.objects.create(f_instance_id=f_instance_id,
                                  name=name,
                                  en_name=en_name,
                                  type=type,
                                  value=value,
                                  service=1,
                                  del_tag=0)


def input_faform(form_data, f_approval_id):
    i = 1
    for item in form_data:
        name = item["name"]
        if name in KVS:
            en_name = KVS[name]
            type = item["type"]
            fa = FAForm.objects.filter(name=name,
                                       f_approval_id=f_approval_id).first()
            if not fa:
                FAForm.objects.create(name=name,
                                      en_name=en_name,
                                      type=type,
                                      serial_number=i,
                                      f_approval_id=f_approval_id,
                                      del_tag=0)
        i += 1


def PutTask(approval_id, instance_id):
    FITask.objects.filter(f_instance_id=instance_id).update(del_tag=1)
    nodes = FANode.objects.filter(f_approval_id=approval_id, del_tag=0)
    for node in nodes:
        name = node.name
        node_type = node.node_type
        f_node_id = node.id
        user_id = node.user_id
        if name == "发起":
            start_time = time.strftime("%Y-%m-%d %H:%M:%S", time.localtime())
            time.sleep(0.2)
            end_time = time.strftime("%Y-%m-%d %H:%M:%S", time.localtime())
            comment = "工单创建"
            status = "DONE"
        else:
            start_time = ""
            end_time = ""
            comment = ""
            status = "BEFORE"
        FITask.objects.create(user_id=user_id,
                              status=status,
                              name=name,
                              type=node_type,
                              f_node_id=f_node_id,
                              f_instance_id=instance_id,
                              comment=comment,
                              start_time=start_time,
                              end_time=end_time,
                              del_tag=0)


class ListenApprovalEventView(baseview.AnyLogin):
    def post(self, request, args=None):
        """事件审批流监听"""
        app = ApprovalT()
        app.get_tenant_access_token()
        msg = {"code": 200, "data": [], "msg": "成功"}
        encrypt = request.data["encrypt"]
        cipher = AESCipher("approvalApply")
        event_data = json.loads(cipher.decrypt_string(encrypt))
        # 连接心跳，绑定飞书 APP 使用
        if "challenge" in event_data:
            return JsonResponse({"challenge": event_data["challenge"]})
        event_obj = event_data["event"]
        approval_code = event_obj["approval_code"]
        instance_code = event_obj["instance_code"]
        approval = FApproval.objects.filter(approval_code=approval_code,
                                            del_tag=0,
                                            old_version=1,
                                            subscribe=1).first()
        if approval:
            res = app.get_instance(instance_code)
            data = res['data']
            instance = FInstanceValue.objects.filter(
                instance_code=instance_code, del_tag=0).first()
            if not instance:
                print(instance_code, " 审批还未入库")
                user_id = data["user_id"]
                user = User.objects.filter(user_id=user_id).first()
                if user:
                    user_id = str(user.id).replace("-", "", 4)
                form_data = json.loads(data["form"])
                fi, ok = FInstanceValue.objects.update_or_create(
                    instance_code=instance_code,
                    defaults={
                        "user_id": user_id,
                        "descriptions": "初始化...",
                        "status": "PENDING",
                        "f_approval_id": approval.id,
                        "del_tag": 0
                    })
                input_faform(form_data, approval.id)
                input_fiform(form_data, fi.id)
                PutTask(approval.id, fi.id)
            else:
                fi = FInstanceValue.objects.filter(instance_code=instance_code,
                                                   status="PENDING",
                                                   del_tag=0).first()
                if not fi:
                    return JsonResponse(msg)
                print(instance_code, " 审批已经入库")
                print("接收到审批实例以下参数: ")
                forms = FIForm.objects.filter(f_instance_id=fi.id)
                obj = {}
                for item in forms:
                    name = item.name
                    if name in NEEDS:
                        obj[item.en_name] = item.value
                print(obj)

                task_list = data["task_list"]
                for task in task_list:
                    if task["status"] == "APPROVED":
                        task_status = "DONE"
                    else:
                        task_status = task["status"]
                    FITask.objects.filter(
                        name=task["node_name"],
                        f_instance_id=fi.id).update(status=task_status)
                    if task["node_name"] == "经办人" and \
                        (task["status"] == "DONE" or task["status"] == "APPROVED"):
                        if obj["SERVICE_TYPE"] in SKIPS:
                            print("当前项目不走jenkins 发布")
                            FInstanceValue.objects.filter(
                                instance_code=instance_code).update(
                                    status="DONE",
                                    descriptions="当前项目不走 jenkins 发布")
                            now = time.strftime("%Y-%m-%d %H:%M:%S",
                                                time.localtime())
                            DeployHistory.objects.create(
                                env=approval.approval_name,
                                project=obj["PROJECT"],
                                service=obj["SERVICE"],
                                BrName=obj["BrName"],
                                deploy_time=now,
                                result="UNKNOWN: 当前项目不走jenkins")
                        else:
                            job_name = approval.job_name
                            print("%s 审批已经通过，准备发布了，job名称: %s" %
                                  (instance_code, job_name))
                            jenkinsApi = JenkinsApi()
                            num = jenkinsApi.get_next_buildnum(job_name)
                            user_id = fi.user_id
                            user = User.objects.filter(id=user_id).first()
                            if user:
                                obj["USERNAME"] = user.username
                            else:
                                obj["USERNAME"] = "未知<无权获取>"
                            FInstanceValue.objects.filter(
                                instance_code=instance_code).update(
                                    status="APPROVED", job_number=num)
                            jenkinsApi.execute_job(job_name, obj)
                            now = time.strftime("%Y-%m-%d %H:%M:%S",
                                                time.localtime())
                            break
                    if task["status"] == "REJECTED":
                        FInstanceValue.objects.filter(
                            instance_code=instance_code).update(
                                status="REJECTED", descriptions="被拒绝发布")
                        break
        return JsonResponse(msg)
