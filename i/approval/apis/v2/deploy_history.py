from django.http import JsonResponse
from django.db.models import Q
from utils import baseview
from approval.models import DeployHistory
from approval.serializers import DeployHistorySerializer
from utils.feishu_approval import ApprovalT
from utils.call_jenkins import JenkinsApi
import time
from django.conf import settings
import os

os.environ.setdefault('DJANGO_SETTINGS_MODULE', 'djentry.settings')


class DeployHistoryView(baseview.AnyLogin):
    """deploy history view"""
    def get(self, request, args=None):
        start = request.GET.get('start', '')
        end = request.GET.get('end', '')
        env = request.GET.get('env', '')
        project = request.GET.get('project', '')
        service = request.GET.get('service', '')
        service_type = request.GET.get('service_type', '')
        page_no = int(request.GET.get('page_no', 1))
        page_size = int(request.GET.get('page_size', 10))
        result = request.GET.get('result', '')
        q = Q()
        q.children.append(("del_tag", 0))
        q.children.append(("deploy_time__range", [start, end]))
        if env:
            q.children.append(("env", env))
        if project:
            q.children.append(("project", project))
        if service:
            q.children.append(("service__contains", service))
        if result:
            q.children.append(("result__contains", result))
        if service_type:
            q.children.append(("service_type", service_type))
        history = DeployHistory.objects.filter(q).order_by("-deploy_time")
        serializer = DeployHistorySerializer(history, many=True)
        page_data = serializer.data[(page_no - 1) * page_size:page_no *
                                    page_size]
        total = len(serializer.data)
        msg = {
            "code": 200,
            "data": {
                "histories": page_data,
                "total": total
            },
            "msg": "获取发布历史成功"
        }
        return JsonResponse(msg)

    def post(self, request, args=None):
        env = request.data.get('env', '')
        project = request.data.get('project', '')
        service = request.data.get('service', '')
        BrName = request.data.get('BrName', '')
        result = request.data.get('result', '')
        deploy_time = time.strftime("%Y-%m-%d %H:%M:%S", time.localtime())
        DeployHistory.objects.create(project=project,
                                     service=service,
                                     env=env,
                                     deploy_time=deploy_time,
                                     BrName=BrName,
                                     result=result)
        msg = {"code": 200, "data": {}, "msg": "新增发布成功"}
        return JsonResponse(msg)

    def delete(self, request, args=None):
        ids = request.data.get('ids', '')
        for id in ids:
            DeployHistory.objects.filter(id=id).update(del_tag=1)
        msg = {"code": 200, "data":ids, "msg": "删除历史记录成功"}
        return JsonResponse(msg)
