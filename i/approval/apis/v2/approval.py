import uuid
from django.http import JsonResponse
from django.db.models import Q
from utils import baseview
from approval.models import FApproval, FANode, FAForm, SelectOptions, FInstanceValue
from approval.serializers import FApprovalSerializer, FAFormSerializer, FANodeSerializer
from approval.apis.audit import PutAudit
from utils.auth import auth


class FApprovalView(baseview.BaseView):
    """工单审批流操作"""
    @auth("approvals.manage.add")
    def post(self, request, args=None):
        """新增审批流程"""
        approval_name = request.data.get('approval_name', '')
        job_name = request.data.get('job_name', '')
        descriptions = request.data.get('descriptions', '')
        approval_code = str(uuid.uuid4())
        open_id = "open_" + approval_code
        fa1 = FApproval.objects.filter(approval_name=approval_name).first()
        fa2 = FApproval.objects.filter(job_name=job_name).first()
        if fa1:
            msg = {"code": 100001, "data": {}, "msg": "审批流名称已经存在"}
        elif fa2:
            msg = {"code": 100001, "data": {}, "msg": "jenkins job 名称已经绑定"}
        else:
            FApproval.objects.create(approval_name=approval_name,
                                     approval_code=approval_code,
                                     open_id=open_id,
                                     job_name=job_name,
                                     subscribe=0,
                                     descriptions=descriptions,
                                     old_version=0,
                                     del_tag=0)
            msg = {"code": 200, "data": {}, "msg": "审批流新增成功"}
        PutAudit(request, msg)
        return JsonResponse(msg)

    @auth("approvals.manage.view")
    def get(self, request, args=None):
        """获取指定审批流程的详细信息"""
        print(request.user)
        id = request.GET.get('id', '')
        approval_code = request.GET.get('approval_code', '')
        approval_name = request.GET.get('approval_name', '')
        job_name = request.GET.get('job_name', '')
        subscribe = request.GET.get('subscribe', '')
        old_version = request.GET.get('old_version', '')
        q = Q()
        q.children.append(("del_tag", 0))
        if id:
            q.children.append(("id", id))
        if approval_code:
            q.children.append(("approval_code", approval_code))
        if approval_name:
            q.children.append(("approval_name", approval_name))
        if job_name:
            q.children.append(("job_name", job_name))
        if subscribe:
            q.children.append(("subscribe", subscribe))
        if old_version:
            q.children.append(("old_version", int(old_version)))
        fas = FApproval.objects.filter(q)
        data = FApprovalSerializer(fas, many=True).data
        msg = {"code": 200, "data": data, "msg": "获取审批流信息成功"}
        return JsonResponse(msg)

    @auth("approvals.manage.del")
    def delete(self, request, args=None):
        """删除指定审批流程"""
        id = request.data.get('id', '')
        fa = FApproval.objects.filter(id=id, del_tag=0).first()
        if fa.subscribe == 1:
            msg = {"code": 100001, "data": {}, "msg": "先取消订阅，才能删除"}
        else:
            FApproval.objects.filter(id=id).update(del_tag=1)
            msg = {"code": 200, "data": {}, "msg": "审批流已删除"}
        PutAudit(request, msg)
        return JsonResponse(msg)

    @auth("approvals.manage.edit")
    def put(self, request, args=None):
        """更新审批流程"""
        data = request.data
        id = request.data.get('id', '')
        FApproval.objects.filter(id=id).update(**data)
        msg = {"code": 200, "data": {}, "msg": "更新成功"}
        PutAudit(request, msg)
        return JsonResponse(msg)


class FAFormView(baseview.BaseView):
    """工单表单操作"""
    @auth("approvals.manage.view")
    def get(self, request, args=None):
        """获取表单信息"""
        approval_id = request.GET.get('approval_id', '')
        fas = FAForm.objects.filter(f_approval_id=approval_id, del_tag=0)
        data = FAFormSerializer(fas, many=True).data
        i = 0
        for item in data:
            if item["type"] == "单选项":
                options = []
                o = SelectOptions.objects.filter(form_id=item["id"])
                for item in o:
                    options.append(item.value)
                data[i]["options"] = options
            i += 1
        msg = {"code": 200, "data": data, "msg": "获取表单信息成功"}
        return JsonResponse(msg)

    @auth("approvals.manage.edit")
    def post(self, request, args=None):
        """更新表单信息"""
        approval_id = request.data.get('approval_id', '')
        form = request.data.get('form', '')
        type_choices = FAForm.type_choices
        fas = FAForm.objects.filter(f_approval_id=approval_id)
        if fas.exists():
            for fa in fas:
                SelectOptions.objects.filter(form_id=fa.id).update(del_tag=1)
        FAForm.objects.filter(f_approval_id=approval_id).update(del_tag=1)
        for item in form:
            for k, v in type_choices:
                if v == item["type"]:
                    f = FAForm.objects.create(
                        name=item["name"],
                        en_name=item["en_name"],
                        type=k,
                        serial_number=item["serial_number"],
                        f_approval_id=approval_id,
                        del_tag=0)
                    if v == "单选项":
                        for option in item["options"]:
                            SelectOptions.objects.create(value=option,
                                                         form_id=f.id,
                                                         del_tag=0)
        msg = {"code": 200, "data": {}, "msg": "表单更新成功"}
        PutAudit(request, msg)
        return JsonResponse(msg)


class FANodeView(baseview.BaseView):
    """工单流程节点操作"""
    @auth("approvals.manage.view")
    def get(self, request, args=None):
        """获取流程节点信息"""
        approval_id = request.GET.get('approval_id', '')
        nodes = FANode.objects.filter(f_approval_id=approval_id, del_tag=0)
        data = FANodeSerializer(nodes, many=True).data
        msg = {"code": 200, "data": data, "msg": "获取流程节点信息成功"}
        return JsonResponse(msg)

    @auth("approvals.manage.edit")
    def post(self, request, args=None):
        """更新流程节点信息"""
        approval_id = request.data.get('approval_id', '')
        nodes = request.data.get('nodes', '')
        FANode.objects.filter(f_approval_id=approval_id).update(del_tag=1)
        node_type_choices = FANode.node_type_choices
        approval_type_choices = FANode.approval_type_choices
        for item in nodes:
            for k, v in approval_type_choices:
                if v == item["approval_type"]:
                    approval_type = k
                    break
            for k2, v2 in node_type_choices:
                if v2 == item["node_type"]:
                    node_type = k2
                    break
            if "user_id" not in item:
                item["user_id"] = ""
            FANode.objects.create(name=item["name"],
                                  need_approver=item["need_approver"],
                                  node_type=node_type,
                                  approval_type=approval_type,
                                  user_id=item["user_id"],
                                  serial_number=item["serial_number"],
                                  f_approval_id=approval_id,
                                  del_tag=0)
        msg = {"code": 200, "data": {}, "msg": "更新流程节点信息成功"}
        PutAudit(request, msg)
        return JsonResponse(msg)


class SubscribeApprovalView(baseview.BaseView):
    """subscribe approval view"""
    @auth("approvals.manage.edit")
    def post(self, request, args=None):
        approval_id = request.data.get('approval_id', '')
        type = request.data.get('type', '')
        if type == 'subscribe':
            fa = FApproval.objects.filter(id=approval_id, del_tag=0).first()
            if fa:
                if fa.subscribe == 0:
                    FApproval.objects.filter(id=approval_id).update(
                        subscribe=1)
                    msg = {"code": 200, "data": {}, "msg": "激活成功"}
                else:
                    msg = {"code": 100001, "data": {}, "msg": "已经激活，无需重复操作"}
            else:
                msg = {"code": 100001, "data": {}, "msg": "审批流不存在，无法激活"}
        else:
            fa = FApproval.objects.filter(id=approval_id, del_tag=0).first()
            if fa:
                if fa.subscribe == 1:
                    FApproval.objects.filter(id=approval_id).update(
                        subscribe=0)
                    FInstanceValue.objects.filter(f_approval_id=approval_id).update(del_tag=1)
                    msg = {"code": 200, "data": {}, "msg": "禁用成功"}
                else:
                    msg = {"code": 100001, "data": {}, "msg": "已经禁用，无需重复操作"}
            else:
                msg = {"code": 100001, "data": {}, "msg": "审批流不存在，无法禁用"}
        PutAudit(request, msg)
        return JsonResponse(msg)


class ChangeVersionView(baseview.BaseView):
    """修改版本"""
    def post(self, request, args=None):
        id = request.data.get('id', '')
        direct = request.data.get('direct', '')
        fa = FApproval.objects.filter(id=id).first()
        if direct == "up":
            if fa.subscribe == 1:
                msg = {"code": 10001, "data": {}, "msg": "需要先取消订阅"}
            else:
                FApproval.objects.filter(id=id).update(old_version=0)
                msg = {"code": 200, "data": {}, "msg": "升级版本成功"}
        elif direct == "down":
            if fa.subscribe == 1:
                msg = {"code": 10001, "data": {}, "msg": "需要先取消订阅"}
            else:
                FApproval.objects.filter(id=id).update(old_version=1)
                msg = {"code": 200, "data": {}, "msg": "降级版本成功"}
        PutAudit(request, msg)
        return JsonResponse(msg)
