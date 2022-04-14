from django.http import JsonResponse, HttpResponse
from utils import baseview
from approval.apis.audit import PutAudit
import os


class UploadFilesView(baseview.BaseView):
    """upload files view"""
    def post(self, request, args=None):
        file = request.FILES.get("file")
        data = request.data
        file_name = data["file_name"]
        file_path = "/data/files/" + file_name
        if os.path.exists(file_path):
            msg = {"code": 200, "data": {}, "msg": "文件已经存在"}
        else:
            f = open(file_path, "wb")
            for chunk in file.chunks():
                f.write(chunk)
            f.close()
        msg = {"code": 200, "data": {}, "msg": "文件上传成功"}
        PutAudit(request, msg)
        return JsonResponse(msg)

class DownloadFilesView(baseview.BaseView):
    """download files view"""
    def post(self, request, args=None):
        data = request.data
        file_name = data["file_name"]
        file_path = "/data/files/" + file_name
        file = open(file_path, "rb")
        response = HttpResponse(file)
        response["Content-Type"] = "application/octet-stream"
        response["Access-Control-Expose-Headers"] = "Content-Disposition"
        response["Content-Disposition"] = "attachment;filename={}".format(file_name)
        response["code"] = 200
        msg = {"code": 200, "data": {}, "msg": "文件下载成功"}
        PutAudit(request, msg)
        return response
