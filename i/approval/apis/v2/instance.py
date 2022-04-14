import uuid
from django.http import JsonResponse
from django.db.models import Q
from utils import baseview
from approval.models import FAForm, FANode, User, FIForm, FInstanceValue, FITask, FApproval
from approval.serializers import FIFormSerializer, FInstanceValueSerializers, FITaskSerializer
from approval.apis.audit import PutAudit
from utils.feishu_approval import ApprovalT
from django.conf import settings
from utils.auth import auth
import time

APPROVAL_INSTANCES_DETAIL = settings.APPROVAL_INSTANCES_DETAIL


def PutTask(approval_id, instance_id):
    FITask.objects.filter(f_instance_id=instance_id).update(del_tag=1)
    nodes = FANode.objects.filter(f_approval_id=approval_id, del_tag=0)
    i = 1
    for node in nodes:
        name = node.name
        node_type = node.node_type
        f_node_id = node.id
        user_id = node.user_id
        if name == "开始":
            start_time = time.strftime("%Y-%m-%d %H:%M:%S", time.localtime())
            time.sleep(1)
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
                              serial_number=i,
                              need_approver=node.need_approver,
                              del_tag=0)
        i += 100


class FInstanceValueView(baseview.BaseView):
    """FInstanceValue view"""
    @auth("approvals.start.view")
    def get(self, request, args=None):
        """获取实例信息"""
        pageNo = int(request.GET.get('page_no', 1))
        pageSize = int(request.GET.get('page_size', 10))
        instance_code = request.GET.get('instance_code', '')
        status = request.GET.get('status', '')
        user_id = request.GET.get('user_id', '')
        f_approval_id = request.GET.get('f_approval_id', '')
        id = request.GET.get('id', '')
        q = Q()
        q.children.append(('del_tag', 0))
        if id:
            q.children.append(("id", id))
        if instance_code:
            q.children.append(("instance_code", instance_code))
        if status:
            en_status = {v: k
                         for k, v in FInstanceValue.status_choices}[status]
            q.children.append(("status", en_status))
        if user_id:
            q.children.append(("user_id", user_id))
        if f_approval_id:
            q.children.append(("f_approval_id", f_approval_id))
        fis = FInstanceValue.objects.filter(q).exclude(status="BEFORE").order_by("-id")
        serializer = FInstanceValueSerializers(fis, many=True)
        total = len(serializer.data)
        instances = serializer.data[(pageNo - 1) * pageSize:pageNo * pageSize]
        msg = {
            "code": 200,
            "data": {
                "instances": instances,
                "total": total
            },
            "msg": "获取实例信息成功"
        }
        return JsonResponse(msg)

    @auth("approvals.start.add")
    def post(self, request, args=None):
        """新增实例信息"""
        instance_code = str(uuid.uuid4())
        user = User.objects.filter(username=request.user.username).first()
        user_id = user.id
        descriptions = request.data.get('descriptions', '')
        approval_id = request.data.get('approval_id', '')
        if not approval_id:
            msg = {"code": 10002, "data": {}, "msg": "必须传递参数 approval_id"}
        else:
            FInstanceValue.objects.create(instance_code=instance_code,
                                          status="BEFORE",
                                          user_id=user_id,
                                          descriptions=descriptions,
                                          job_number=0,
                                          f_approval_id=approval_id,
                                          del_tag=0)
            fi = FInstanceValue.objects.filter(instance_code=instance_code)
            instance = FInstanceValue.objects.filter(
                instance_code=instance_code).first()
            PutTask(approval_id, instance.id)
            data = FInstanceValueSerializers(fi, many=True).data
            msg = {"code": 200, "data": data, "msg": "实例新增成功"}
        PutAudit(request, msg)
        return JsonResponse(msg)

    @auth("approvals.start.edit")
    def put(self, request, args=None):
        """更新实例信息"""
        data = request.data
        id = request.data.get('id', '')
        FInstanceValue.objects.filter(id=id).update(**data)
        msg = {"code": 200, "data": {}, "msg": "实例更新成功"}
        PutAudit(request, msg)
        return JsonResponse(msg)

    @auth("approvals.start.del")
    def delete(self, request, args=None):
        """删除实例信息"""
        id = request.data.get('id', '')
        FIForm.objects.filter(f_instance_id=id).update(del_tag=1)
        FInstanceValue.objects.filter(id=id).update(del_tag=1)
        msg = {"code": 200, "data": {}, "msg": "实例删除成功"}
        PutAudit(request, msg)
        return JsonResponse(msg)


def formToObjZh(data):
    result = []
    keys = []
    for item in data:
        if item["service"] not in keys:
            keys.append(item["service"])
    for service in keys:
        obj = {}
        obj["服务编号"] = service
        for item in data:
            if item["service"] == service:
                obj[item["name"]] = item["value"]
        result.append(obj)
    return result


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


class FIFormView(baseview.BaseView):
    """FIForm view"""
    @auth("approvals.start.edit")
    def post(self, request, args=None):
        """更新form 表单"""
        approval_id = request.data.get('approval_id', '')
        instance_id = request.data.get('instance_id', '')
        service = int(request.data.get('service', 1))
        form = request.data.get('form', '')
        if not approval_id:
            msg = {"code": 10002, "data": {}, "msg": "必须传递参数 approval_id"}
        else:
            for key in form:
                fanode = FAForm.objects.filter(del_tag=0,
                                               f_approval_id=approval_id,
                                               en_name=key).first()
                FIForm.objects.create(name=fanode.name,
                                      en_name=fanode.en_name,
                                      type=fanode.type,
                                      f_instance_id=instance_id,
                                      del_tag=0,
                                      service=service,
                                      value=form[key])
            fis = FIForm.objects.filter(f_instance_id=instance_id)
            data = FIFormSerializer(fis, many=True).data
            result = formToObjZh(data)
            msg = {"code": 200, "data": result, "msg": "实例新增成功"}
        PutAudit(request, msg)
        return JsonResponse(msg)

    @auth("approvals.fi-list.view")
    def get(self, request, args=None):
        """获取form信息"""
        instance_id = request.GET.get('instance_id', '')
        q = Q()
        q.children.append(("del_tag", 0))
        if instance_id:
            q.children.append(("f_instance_id", instance_id))
        fis = FIForm.objects.filter(q)
        data = FIFormSerializer(fis, many=True).data
        result = formToObjZh(data)
        msg = {"code": 200, "data": result, "msg": "获取form信息成功"}
        return JsonResponse(msg)

    @auth("approvals.fi-list.del|approvals.start.del")
    def delete(self, request, args=None):
        instance_id = request.data.get('instance_id', '')
        service = int(request.data.get('service', 1))
        FIForm.objects.filter(f_instance_id=instance_id, del_tag=0, service=service).update(del_tag=1)
        msg = {"code": 200, "data": {}, "msg": "移除服务成功"}
        PutAudit(request, msg)
        return JsonResponse(msg)
    


def send_user(email, username, approval_name, instance_id, color, title):
    """审批提醒使用"""
    app = ApprovalT()
    app.get_tenant_access_token()
    result = app.get_open_id(email)
    open_id = ""
    if "email_users" in result["data"]:
        if email in result["data"]["email_users"]:
            open_id = result["data"]["email_users"][email][0]["open_id"]
    if open_id != "":
        result = app.send_card_data_self(
            color, title, open_id, username, approval_name,
            APPROVAL_INSTANCES_DETAIL + str(instance_id))
        if result["code"] == 0:
            print("已经发送飞书通知")


class FITaskView(baseview.BaseView):
    """FITask view"""
    def get(self, request, args=None):
        instance_id = request.GET.get('instance_id', '')
        q = Q()
        q.children.append(("del_tag", 0))
        if instance_id:
            q.children.append(("f_instance_id", instance_id))
        tasks = FITask.objects.filter(q).order_by("serial_number")
        data = FITaskSerializer(tasks, many=True).data
        msg = {"code": 200, "data": data, "msg": "获取事件流成功"}
        return JsonResponse(msg)

    def post(self, request, args=None):
        """task 更新"""
        id = int(request.data.get('id', ''))
        descriptions = request.data.get('descriptions', '')
        status = request.data.get('status', '')
        now = time.strftime("%Y-%m-%d %H:%M:%S", time.localtime())
        task = FITask.objects.filter(id=id).first()
        msg = {"code": 200, "data": {}, "msg": "task 更新成功"}
        if status == "TRANSFERRED":
            user_id = request.data.get('user_id', '')
            FITask.objects.filter(id=id).update(status=status,
                                                comment=descriptions,
                                                end_time=now)
            FITask.objects.create(name=task.name + "[转]",
                                  user_id=user_id,
                                  status="BEFORE",
                                  type=task.type,
                                  start_time="",
                                  end_time="",
                                  f_node_id=task.f_node_id,
                                  f_instance_id=task.f_instance_id,
                                  comment="",
                                  serial_number=task.serial_number + 1,
                                  need_approver=task.need_approver,
                                  del_tag=0)
            msg = {"code": 200, "data": {}, "msg": "审批任务已经转办"}
        elif status == "REFUTED":
            jenkinsTask = FITask.objects.filter(
                name="经办人", f_instance_id=task.f_instance_id,
                del_tag=0).first()
            if task.serial_number >= jenkinsTask.serial_number:
                msg = {"code": 10001, "data": {}, "msg": "服务已经发布，不能驳回"}
            else:
                FITask.objects.filter(id=id).update(status=status,
                                                    comment=descriptions,
                                                    end_time=now)
                fi = FInstanceValue.objects.filter(
                    id=task.f_instance_id).first()
                nodeId = task.f_node_id
                needNodes = FANode.objects.filter(
                    del_tag=0, f_approval_id=fi.f_approval_id, id__lte=nodeId)
                j = 1
                for node in needNodes:
                    if node.user_id == "" and node.approval_type == "MANUAL":
                        t = FITask.objects.filter(f_instance_id=task.f_instance_id,f_node_id=node.id,del_tag=0).first()
                        user_id = t.user_id
                    else: 
                        user_id = node.user_id
                    FITask.objects.create(name=node.name,
                                          user_id=user_id,
                                          status="BEFORE",
                                          type=node.node_type,
                                          start_time="",
                                          end_time="",
                                          f_node_id=node.id,
                                          f_instance_id=task.f_instance_id,
                                          comment="",
                                          serial_number=task.serial_number +
                                          10 * j,
                                          need_approver=task.need_approver,
                                          del_tag=0)
                    j += 1
                FInstanceValue.objects.filter(id=fi.id).update(
                    status=status, descriptions=descriptions)
                fa = FApproval.objects.filter(id=fi.f_approval_id).first()
                user = User.objects.filter(id=fi.user_id).first()
                send_user(user.email, user.username, fa.approval_name, fi.id,
                          "red", "你的审批被驳回")
        elif status == "REJECTED":
            FITask.objects.filter(id=id).update(status=status,
                                                comment=descriptions,
                                                end_time=now)
            FInstanceValue.objects.filter(id=task.f_instance_id).update(
                status=status, descriptions=descriptions)
            fi = FInstanceValue.objects.filter(id=task.f_instance_id).first()
            fa = FApproval.objects.filter(id=fi.f_approval_id).first()
            user = User.objects.filter(id=fi.user_id).first()
            send_user(user.email, user.username, fa.approval_name, fi.id,
                      "red", "你的审批被拒绝")
            msg = {"code": 200, "data": {}, "msg": "审批任务因某种原因已经终止"}
        else:
            FITask.objects.filter(id=id).update(status=status,
                                                comment=descriptions,
                                                end_time=now)
        PutAudit(request, msg)
        return JsonResponse(msg)
