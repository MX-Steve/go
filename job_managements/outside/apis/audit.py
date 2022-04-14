from django.http import JsonResponse
from django.db.models import Count, Q
from datetime import timedelta, datetime
from utils import baseview
from outside.models import *
from utils.util import now
import json


def PutAudit(request, res):
    req_method = request.method
    if req_method == "GET":
        return
    operater = request.user.username
    req_data = json.dumps(request.data)
    req_url = request.path
    res_data = json.dumps(res)
    AuditHistory.objects.create(operater=operater,
                                operate_time=now(),
                                req_url=req_url,
                                req_method=req_method,
                                req_data=req_data,
                                res_data=res_data)