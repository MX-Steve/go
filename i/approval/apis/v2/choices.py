from django.http import JsonResponse
from utils import baseview
from approval.models import FAForm, FIForm, FITask, FIViewers, FInstanceValue, FANode, BusinessServices, DnsRecords


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
                if column == "approval_type_choices":
                    for k, v in FANode.approval_type_choices:
                        choices.append(v)
                else:
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
            elif table_name == "BusinessServices":
                for k, v in BusinessServices.type_choices:
                    choices.append(v)
            elif table_name == "DnsRecords":
                if column == "line_choices":
                    for k, v in DnsRecords.line_choices:
                        choices.append(v)
                elif column == "proxy_choices":
                    for k, v in DnsRecords.proxy_choices:
                        choices.append(v)
            msg = {"code": 200, "data": choices, "msg": "获取选项成功"}
        return JsonResponse(msg)
