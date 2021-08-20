# encoding:utf-8
from demo.utils import baseview
from rest_framework.response import Response    
import logging


# 初始化logger
logger = logging.getLogger("DemoIn.app")


# 登录接口
class TestView(baseview.AnyLogin):
    def post(self, request, args=None):
        print(request.data)
        return Response({'code': 200, 'res': '成功!', 'data': []})