from django.http import JsonResponse
from django.db.models import Q
from utils import baseview
from approval.models import FAForm, FApproval, FIComment, FIForm, FITask, FIViewers, FInstanceValue, FANode
from approval.serializers import FICommentSerializer, FIFormSerializer, FITaskSerializer, FInstanceValueSerializers


class FInstanceView(baseview.AnyLogin):
    def post(self, request, args=None):
        """工单流程发布状态更新"""
        data = request.data
        msg = {"code": 200, "data": {}, "msg": "状态更新成功"}
        if data["type"] == "status":
            instance_code = data["instance_code"]
            status = data["status"]
            descriptions = data["comment"]
            fi = FInstanceValue.objects.filter(instance_code=instance_code, del_tag=0)
            if not fi.exists():
                msg = {"code": 10001, "data": {}, "msg": "所要更新的工单不存在"}
            else:
                FInstanceValue.objects.filter(
                    instance_code=instance_code).update(
                        status=status, descriptions=descriptions)
        else:
            msg = {"code": 10003, "data": {}, "msg": "所传 type 类型值不存在"}
        return JsonResponse(msg)

    def get(self, request, args=None):
        """工单流程发布状态查询|详细信息查询"""
        type = request.GET.get('type', '')
        if type == "status":
            msg = {"code": 200, "data": {}, "msg": "查询状态成功"}
            instance_code = request.GET.get("instance_code", "")
            q = Q()
            q.children.append(("del_tag", 0))
            if instance_code:
                q.children.append(("instance_code", instance_code))
            fis = FInstanceValue.objects.filter(q)
            msg["data"] = FInstanceValueSerializers(fis, many=True).data
        elif type == "details":
            msg = {"code": 200, "data": {}, "msg": "查询详情成功"}
            instance_code = request.GET.get("instance_code", "")
            if not instance_code:
                msg = {"code": 10001, "data": {}, "msg": "instance_code 参数值不存在"}
            else:
                fi = FInstanceValue.objects.filter(instance_code=instance_code,del_tag=0)
                msg["data"] = FInstanceValueSerializers(fi, many=True).data[0]
                f_instance_id = msg["data"]["id"]
                form = FIForm.objects.filter(del_tag=0, f_instance_id=f_instance_id)
                msg["data"]["form"] = FIFormSerializer(form, many=True).data
                tasks = FITask.objects.filter(del_tag=0, f_instance_id=f_instance_id)
                msg["data"]["task_list"] = FITaskSerializer(tasks, many=True).data
                comments = FIComment.objects.filter(del_tag=0, f_instance_id=f_instance_id)
                msg["data"]["comment_list"] = FICommentSerializer(comments, many=True).data
        else:
            msg = {"code": 10001, "data": {}, "msg": "type 参数值不存在"}
        return JsonResponse(msg)


class FInstanceListView(baseview.AnyLogin):
    def get(self, request, args=None):
        """查询指定审批流下的工单实例列表"""
        approval_code = request.GET.get('approval_code', '')
        status = request.GET.get('status', '')
        if approval_code == "":
            msg = {"code": 10003, "data": {}, "msg": "必须传递参数 approval_code"}
        else:
            approval = FApproval.objects.filter(approval_code=approval_code,
                                                del_tag=0).first()
            instance_codes = []
            if approval:
                q = Q()
                q.children.append(("del_tag", 0))
                q.children.append(("f_approval_id", approval.id))
                if status:
                    status_choices = FInstanceValue.status_choices
                    en_status = ""
                    for k, v in status_choices:
                        if v == status:
                            en_status = k
                            break
                    q.children.append(("status", en_status))
                finstances = FInstanceValue.objects.filter(q)
                data = FInstanceValueSerializers(finstances, many=True).data
                for ins in data:
                    instance_codes.append(ins["instance_code"])
            msg = {"code": 200, "data": instance_codes, "msg": "获取列表成功"}
        return JsonResponse(msg)

class ChoicesView(baseview.AnyLogin):
    def get(self, request, args=None):
        """models choices value"""
        table_name = request.GET.get("table_name", "")
        column = request.GET.get("column", "")
        if table_name == "" or column == "":
            msg = {"code": 10002, "data": {}, "msg": "必须传递表名和choice 列"}
        else:
            choices = []
            if table_name == "FANode":
                for k, v in FANode.node_type_choices:
                    choices.append(v)
            elif table_name == "FAForm":
                for k, v in FAForm.type_choices:
                    choices.append(v)
            elif table_name == "FInstanceValue":
                for k, v in FInstanceValue.status_choices:
                    choices.append(v)
            elif table_name == "FIViewers":
                for k, v in FIViewers.type_choices:
                    choices.append(v)
            elif table_name == "FIForm":
                for k, v in FIForm.type_choices:
                    choices.append(v)
            elif table_name == "FITask":
                if column == "status_choices":
                    for k, v in FITask.status_choices:
                        choices.append(v)
                else:
                    for k, v in FITask.type_choices:
                        choices.append(v)
            msg = {"code": 200, "data": choices, "msg": "获取选项成功"}
        return JsonResponse(msg)
