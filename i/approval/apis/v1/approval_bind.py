from django.http import JsonResponse
from django.db.models import Q
from utils import baseview
from utils.feishu_approval import ApprovalT
from approval.models import FApproval, FInstanceValue, User, FANode
from approval.serializers import FApprovalSerializer


def load_history_instance(approval_code):
    fa = FApproval.objects.filter(approval_code=approval_code,
                                  del_tag=0).first()
    if fa:
        app = ApprovalT()
        app.get_tenant_access_token()
        res = app.get_instance_list(approval_code)
        instance_code_list = res["data"]["instance_code_list"]
        for instance_code in instance_code_list:
            FInstanceValue.objects.update_or_create(
                instance_code=instance_code,
                defaults={
                    "status": "DONE",
                    "user_id": "",
                    "descriptions": "历史数据",
                    "job_number": 0,
                    "f_approval_id": fa.id,
                    "del_tag": 0
                })


def putNodes(f_approval_id, node_list):
    FANode.objects.filter(f_approval_id=f_approval_id).update(del_tag=1)
    user = User.objects.filter(username="admin").first()
    user_id = user.id
    nodes = [{
        "name": "发起",
        "need_approver": False,
        "need_cc_user": False,
        "node_id": "b078ffd28db767c502ac367053f6e0ac",
        "node_type": "AND"
    }]
    for node in node_list:
        if node["name"] not in ["发起", "结束"]:
            nodes.append(node)
    nodes.append({
        "name": "结束",
        "need_approver": False,
        "need_cc_user": False,
        "node_id": "b1a326c06d88bf042f73d70f50197905",
        "node_type": "AND"
    })
    i = 1
    for item in nodes:
        if item["name"] in ["发起", "结束"]:
            approval_type = "AUTO_PASS"
        else:
            approval_type = "MANUAL"
        FANode.objects.create(name=item["name"],
                              need_approver=item["need_approver"],
                              node_type=item["node_type"],
                              user_id=user_id,
                              approval_type=approval_type,
                              serial_number=i,
                              f_approval_id=f_approval_id,
                              del_tag=0)
        i += 1


def unbind_approval(approval_code):
    fa = FApproval.objects.filter(approval_code=approval_code,
                                  del_tag=0).first()
    if not fa:
        msg = {"code": 10002, "data": {}, "msg": "工单审批流不存在，无法取消订阅"}
    else:
        fa = FApproval.objects.filter(approval_code=approval_code,
                                      subscribe=0).first()
        if fa:
            msg = {"code": 10002, "data": {}, "msg": "已经取消订阅，无需重复操作"}
        else:
            app = ApprovalT()
            app.get_tenant_access_token()
            result = app.unsubscribe_events(approval_code)
            if result["code"] != 0:
                msg = {"code": 10001, "data": {}, "msg": "取消订阅失败，请检查原因"}
            else:
                FApproval.objects.filter(approval_code=approval_code).update(
                    subscribe=0)
                msg = {"code": 200, "data": {}, "msg": "取消订阅成功"}
    return msg


def bind_approval(approval_code, approval_name, job_name):
    open_id = "open_" + approval_code
    fa_obj = FApproval.objects.filter(approval_code=approval_code,
                                      subscribe=1).first()
    if fa_obj:
        msg = {
            "code": 10003,
            "data": {},
            "msg": "%s 已经订阅了，无需再次订阅" % (approval_name)
        }
    else:
        app = ApprovalT()
        app.get_tenant_access_token()
        result = app.subscribe_events(approval_code)
        if result["code"] != 0:
            msg = {"code": 10001, "data": {}, "msg": "订阅失败，请检查原因"}
        else:
            FApproval.objects.update_or_create(approval_code=approval_code,
                                               defaults={
                                                   "approval_name":
                                                   approval_name,
                                                   "open_id": open_id,
                                                   "job_name": job_name,
                                                   "subscribe": 1,
                                                   "old_version": 1,
                                                   "descriptions": "初次绑定"
                                               })
            load_history_instance(approval_code)
            result2 = app.approval_details(approval_code)
            if result2["code"] == 0:
                fa = FApproval.objects.filter(approval_code=approval_code).first()
                node_list = result2["data"]["node_list"]
                putNodes(fa.id, node_list)
            msg = {"code": 200, "data": {}, "msg": "订阅成功，开始监听"}
    return msg


class ApprovalBindView(baseview.AnyLogin):
    def post(self, request, args=None):
        """工单与jenkins job绑定接口并从飞书订阅或取消订阅"""
        msg = {"code": 200, "data": {}, "msg": "更改成功"}
        approval_code = request.data["approval_code"]
        approval_name = request.data["approval_name"]
        job_name = request.data["job_name"]
        # 1: 订阅; 0: 取消订阅
        subscribe = int(request.data["subscribe"])
        if subscribe == 0:
            msg = unbind_approval(approval_code)
        else:
            msg = bind_approval(approval_code, approval_name, job_name)
        return JsonResponse(msg)

    def delete(self, request, args=None):
        """删除现有工单审批流"""
        msg = {"code": 200, "data": {}, "msg": "删除成功"}
        approval_code = request.data["approval_code"]
        fa = FApproval.objects.filter(approval_code=approval_code).first()
        if not fa:
            msg = {"code": 10003, "data": {}, "msg": "审批流不存在，无法删除"}
        else:
            fa = FApproval.objects.filter(approval_code=approval_code,
                                          subscribe=1).first()
            if fa:
                msg = {"code": 10003, "data": {}, "msg": "请先取消订阅，否则无法删除"}
            else:
                FApproval.objects.filter(approval_code=approval_code).update(
                    del_tag=1)
        return JsonResponse(msg)

    def get(self, request, args=None):
        """查询已有工单审批流"""
        approval_code = request.GET.get('approval_code', '')
        q = Q()
        q.children.append(("del_tag", 0))
        if approval_code:
            q.children.append(("approval_code", approval_code))
        fas = FApproval.objects.filter(q)
        data = FApprovalSerializer(fas, many=True).data
        msg = {"code": 200, data: data, "msg": "获取工单审批流成功"}
        return JsonResponse(msg)
