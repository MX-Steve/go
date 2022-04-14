from django.http import JsonResponse
from utils import baseview
from approval.models import FITask
from approval.apis.audit import PutAudit

class FITaskUserView(baseview.BaseView):
    """FITask view"""
    def post(self, request, args=None):
        msg = {"code": 200, "data": {}, "msg": "设置审批人成功"}
        id = int(request.data.get('id', ''))
        user_id = request.data.get('user_id', '')
        fi = FITask.objects.filter(id=id).first()
        if fi:
            FITask.objects.filter(id=id).update(user_id=user_id)
        else:
            msg = {"code": 10002, "data": {}, "msg": "审批任务流不存在"}
        PutAudit(request, msg)
        return JsonResponse(msg)