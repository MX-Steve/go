# encoding:utf-8
from ktz.utils import baseview
from django.shortcuts import HttpResponse
from rest_framework.response import Response
import logging
import time 


# 初始化logger
logger = logging.getLogger("ktzIn.app")


# 登录接口
class TestView(baseview.AnyLogin):
    def get(self, request, args=None):
        print(request.data)
        print(request.path)
        print("view函数...")
        # return Response({'code': 200, 'res': '成功!', 'data': []})
        return HttpResponse("ok,%s"%(time.time()))