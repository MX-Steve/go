from django.http import JsonResponse
from utils import baseview
from utils.feishu_approval import ApprovalT
from approval.models import User, FInstanceValue, FApproval, DeployHistory
from utils.feishu_rebot import FeishuRobot
from utils.hello.run import Run
from django.conf import settings
import os
import time

os.environ.setdefault('DJANGO_SETTINGS_MODULE', 'djentry.settings')


class SendDeployResultView(baseview.AnyLogin):
    def post(self, request, args=None):
        """给发布结果发送通知"""
        username = request.data["username"]
        info = request.data["result"]
        msg = {"code": 200, "data": {}, "msg": "成功"}
        user = User.objects.filter(username=username).first()
        if not user:
            msg = {"code": 10002, "data": {}, "msg": "未查询到用户信息"}
        else:
            email = user.email
            app = ApprovalT()
            app.get_tenant_access_token()
            result = app.get_open_id(email)
            open_id = ""
            if "email_users" in result["data"]:
                open_id = result["data"]["email_users"][email][0]["open_id"]
            if open_id == "":
                msg = {"code": 10002, "data": {}, "msg": "未从飞书上查询到用户信息"}
            else:
                Run(open_id, info)
        return JsonResponse(msg)


class SendResultToGroup(baseview.AnyLogin):
    def post(self, request, args=None):
        msg = {"code": 200, "data": {}, "msg": "成功"}
        info = request.data["result"]
        info += """ .<at user_id="all">所有人</at>"""
        user = User.objects.filter(username="admin").first()
        robot_url = user.robot_url
        robot_secret = user.robot_secret
        robot = FeishuRobot(robot_url, robot_secret)
        robot.sign_generate()
        robot.send(info)
        return JsonResponse(msg)


class SendResultToGroup2(baseview.AnyLogin):
    def post(self, request, args=None):
        username = request.data["username"]
        job_name = request.data["job_name"]
        ok = request.data["result"]
        env = request.data["env"]
        project = request.data["project"]
        service = request.data["service"]
        number = request.data["number"]
        BrName = request.data["BrName"]
        jenkins_url = "%s/job/%s/%s/console" % (settings.JENKINS_HOST,
                                                job_name, number)
        fa = FApproval.objects.filter(job_name=job_name).first()
        now = time.strftime("%Y-%m-%d %H:%M:%S", time.localtime())
        if ok == "SUCCESS":
            FInstanceValue.objects.filter(job_number=number,
                                          f_approval_id=fa.id).update(
                                              descriptions="发布成功",
                                              status="DONE")
            DeployHistory.objects.create(env=env,
                                             project=project,
                                             service=service,
                                             BrName=BrName,
                                             deploy_time=now,
                                             result="SUCCESS: 发布成功",
                                             del_tag=0)
        else:
            FInstanceValue.objects.filter(job_number=number,
                                          f_approval_id=fa.id).update(
                                              descriptions="发布失败",
                                              status="REJECTED")
            DeployHistory.objects.create(env=env,
                                             project=project,
                                             service=service,
                                             BrName=BrName,
                                             deploy_time=now,
                                             result="FAILURE: 发布失败",
                                             del_tag=0)
        app = ApprovalT()
        app.get_tenant_access_token()
        app.send_card_data(settings.DEPLOY_CHAT_ID, username, env, project,
                           service, jenkins_url, ok)
        msg = {"code": 200, "data": {}, "msg": "成功"}
        return JsonResponse(msg)