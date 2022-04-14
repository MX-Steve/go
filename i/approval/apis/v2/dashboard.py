from django.http import JsonResponse
from utils import baseview
from approval.models import FApproval,FITask, FInstanceValue
import datetime

COLORS = ["#c4cddc","#2caf68", "#097bed", "#ffbb00", "#ff6c37", "#fc5531", "#0894fe", "#ffcb0a"]

class DashboardView(baseview.BaseView):
    """获取审批首页数据"""
    def get(self, request, args=None):
        approvals = FApproval.objects.filter(del_tag=0, subscribe=1)
        data = []
        j = 0
        for approval in approvals:
            a = {}
            for i in [7,6,5,4,3,2,1]:
                start = (datetime.datetime.now() - datetime.timedelta(days = i)).strftime("%Y-%m-%d")
                end = (datetime.datetime.now() - datetime.timedelta(days = i-1)).strftime("%Y-%m-%d")
                fis = FInstanceValue.objects.filter(del_tag=0, f_approval_id=approval.id)
                total = 0
                for fi in fis:
                    num = FITask.objects.filter(del_tag=0, f_instance_id=fi.id,name__in=["开始", "发起"], start_time__gte=start, start_time__lt=end).count()
                    total += num
                a[start] = total
            data.append({
                "approval_name": approval.approval_name,
                "color": COLORS[j],
                "data": a
            })
            j += 1
        msg = {"code": 200, "data": data, "msg": "获取数据成功"}
        return JsonResponse(msg)