from django.utils.deprecation import MiddlewareMixin
from django.shortcuts import HttpResponse
from rest_framework.response import Response
import sys
import os
import time

def testAuth(request):
    print("in testAuth")
    print(request.path)
    return HttpResponse(status=403)

class MD1(MiddlewareMixin):
    def process_request(self, request):
        try:
            path = request.path
            if  "/ktz/test/" == path and "GET" == request.method:
                testAuth(request)
                # return Response({'code': 200, 'res': '成功!', 'data': []})
            # print("md1  process_request 方法。%s"%(time.time()), id(request),path) #在视图之前执行
        except:
            err = '%s [%s] happend on %s line at %s' % (
                sys.exc_info()[0], sys.exc_info()[1], sys.exc_info()[2].tb_lineno, os.path.basename(__file__))
            print(err)
            return HttpResponse(status=500)
    def process_response(self,request,response):
        print("Md2返回")
        return response