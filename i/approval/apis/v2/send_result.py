from django.http import JsonResponse
from utils import baseview
from utils.feishu_approval import ApprovalT
from approval.models import User, FITask, FInstanceValue, FApproval, FIForm, DeployHistory
from approval.serializers import FIFormSerializer
from utils.feishu_rebot import FeishuRobot
from utils.hello.run import Run
from django.conf import settings
import os
import time

os.environ.setdefault('DJANGO_SETTINGS_MODULE', 'djentry.settings')


def formToObjEn(data):
    result = []
    keys = []
    for item in data:
        if item["service"] not in keys:
            keys.append(item["service"])
    for service in keys:
        obj = {}
        for item in data:
            if item["service"] == service:
                obj[item["en_name"]] = item["value"]
        obj["SERNum"] = service
        result.append(obj)
    return result


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
                if email in result["data"]["email_users"]:
                    open_id = result["data"]["email_users"][email][0]["open_id"]
            if open_id == "":
                msg = {"code": 10002, "data": {}, "msg": "未从飞书上查询到用户信息"}
            else:
                Run(open_id, info)
        return JsonResponse(msg)


class SendResultToGroup(baseview.AnyLogin):
    """群机器人发送消息"""
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
    """生产发布结果通知群"""
    def post(self, request, args=None):
        username = request.data["username"]
        job_name = request.data["job_name"]
        ok = request.data["result"]
        env = request.data["env"]
        project = request.data["project"]
        service = request.data["service"]
        service_type = request.data["service_type"]
        number = request.data["number"]
        BrName = request.data["BrName"]
        fa = FApproval.objects.filter(job_name=job_name).first()
        fi = FInstanceValue.objects.filter(job_number=number,
                                           f_approval_id=fa.id).first()
        end_time = time.strftime("%Y-%m-%d %H:%M:%S", time.localtime())
        jenkins_url = "%s/job/%s/%s/console" % (settings.JENKINS_HOST,
                                                job_name, number)
        now = time.strftime("%Y-%m-%d %H:%M:%S", time.localtime())
        app = ApprovalT()
        app.get_tenant_access_token()
        if ok == "SUCCESS":
            DeployHistory.objects.create(env=env,
                                         project=project,
                                         service=service,
                                         service_type=service_type,
                                         BrName=BrName,
                                         deploy_time=now,
                                         result="SUCCESS: 发布成功",
                                         del_tag=0)
            fiforms = FIForm.objects.filter(f_instance_id=fi.id, del_tag=0)
            data = FIFormSerializer(fiforms, many=True).data
            result = formToObjEn(data)
            serviceTotal = len(result)
            for form in result:
                if form["SERVICE"] == service:
                    if serviceTotal > form["SERNum"]:
                        task = FITask.objects.filter(f_instance_id=fi.id,
                                                     name="经办人").first()
                        keys = list(filter(None, task.comment.split(",")))
                        keys.append(service)
                        strService = ","
                        for k in keys:
                            strService += k + ","
                        FITask.objects.filter(
                            f_instance_id=fi.id,
                            name="经办人").update(comment=strService)
                    else:
                        FITask.objects.filter(f_instance_id=fi.id,
                                              name="经办人").update(
                                                  comment="自动审批通过",
                                                  status="DONE",
                                                  end_time=end_time)
                        FInstanceValue.objects.filter(id=fi.id).update(status="APPROVED",
                                                           descriptions="发布成功")
                        app.send_card_data(settings.DEPLOY_CHAT_ID, username,
                                           env, project, service, jenkins_url,
                                           ok)
                    break
        else:
            DeployHistory.objects.create(env=env,
                                         project=project,
                                         service=service,
                                         service_type=service_type,
                                         BrName=BrName,
                                         deploy_time=now,
                                         result="FAILURE: 发布失败",
                                         del_tag=0)
            FITask.objects.filter(f_instance_id=fi.id,
                                  name="经办人").update(comment="jenkins构建失败",
                                                     end_time=end_time,
                                                     status="REJECTED")
            FInstanceValue.objects.filter(id=fi.id).update(status="REJECTED",
                                                           descriptions="发布失败")
            app.send_card_data(settings.DEPLOY_CHAT_ID, username, env, project,
                               service, jenkins_url, ok)
        msg = {"code": 200, "data": {}, "msg": "成功"}
        return JsonResponse(msg)


class SendResultToSelf(baseview.AnyLogin):
    def post(self, request, args=None):
        username = request.data["username"]
        job_name = request.data["job_name"]
        ok = request.data["result"]
        env = request.data["env"]
        project = request.data["project"]
        service = request.data["service"]
        service_type = request.data["service_type"]
        number = request.data["number"]
        BrName = request.data["BrName"]
        fa = FApproval.objects.filter(job_name=job_name).first()
        fi = FInstanceValue.objects.filter(job_number=number,
                                           f_approval_id=fa.id).first()
        end_time = time.strftime("%Y-%m-%d %H:%M:%S", time.localtime())
        jenkins_url = "%s/job/%s/%s/console" % (settings.JENKINS_HOST,
                                                job_name, number)
        now = time.strftime("%Y-%m-%d %H:%M:%S", time.localtime())
        app = ApprovalT()
        app.get_tenant_access_token()
        admin = User.objects.filter(username="admin").first()
        result = app.get_open_id(admin.email)
        open_id = ""
        if "email_users" in result["data"]:
            if admin.email in result["data"]["email_users"]:
                open_id = result["data"]["email_users"][admin.email][0]["open_id"]
        if open_id == "":
            msg = {"code": 10002, "data": {}, "msg": "未从飞书上查询到用户信息"}
        if ok == "SUCCESS":
            DeployHistory.objects.create(env=env,
                                         project=project,
                                         service=service,
                                         service_type=service_type,
                                         BrName=BrName,
                                         deploy_time=now,
                                         result="SUCCESS: 发布成功",
                                         del_tag=0)
            fiforms = FIForm.objects.filter(f_instance_id=fi.id, del_tag=0)
            data = FIFormSerializer(fiforms, many=True).data
            result = formToObjEn(data)
            serviceTotal = len(result)
            if fi.descriptions == "手动触发构建一次":
                app.send_card_data_test(open_id, username, env,
                                                project, service, jenkins_url,
                                                ok)
            else:
                for form in result:
                    if form["SERVICE"] == service:
                        if serviceTotal > form["SERNum"]:
                            task = FITask.objects.filter(f_instance_id=fi.id,
                                                        name="经办人").first()
                            keys = list(filter(None, task.comment.split(",")))
                            keys.append(service)
                            strService = ","
                            for k in keys:
                                strService += k + ","
                            FITask.objects.filter(
                                f_instance_id=fi.id,
                                name="经办人").update(comment=strService)
                        else:
                            FITask.objects.filter(f_instance_id=fi.id,
                                                name="经办人").update(
                                                    comment="自动审批通过",
                                                    status="DONE",
                                                    end_time=end_time)
                            app.send_card_data_test(open_id, username, env,
                                                    project, service, jenkins_url,
                                                    ok)
                        break
        else:
            DeployHistory.objects.create(env=env,
                                         project=project,
                                         service=service,
                                         service_type=service_type,
                                         BrName=BrName,
                                         deploy_time=now,
                                         result="FAILURE: 发布失败",
                                         del_tag=0)
            FITask.objects.filter(f_instance_id=fi.id,
                                  name="经办人").update(comment="jenkins构建失败",
                                                     end_time=end_time,
                                                     status="REJECTED")
            FInstanceValue.objects.filter(id=fi.id).update(status="REJECTED",
                                                           descriptions="发布失败")
            app.send_card_data_test(open_id, username, env, project, service,
                                    jenkins_url, ok)
        msg = {"code": 200, "data": {}, "msg": "成功"}
        return JsonResponse(msg)


class SendResultToSelf2(baseview.AnyLogin):
    def post(self, request, args=None):
        username = request.data["username"] # 申请人
        operator = request.data["operator"] # 流程操作者
        address = request.data["address"]
        approval_name = request.data["approval_name"]
        app = ApprovalT()
        app.get_tenant_access_token()
        user = User.objects.filter(username=operator).first()
        if user:
            email = user.email
            result = app.get_open_id(email)
            open_id = ""
            if "email_users" in result["data"]:
                if email in result["data"]["email_users"]:
                    open_id = result["data"]["email_users"][email][0]["open_id"]
            if open_id == "":
                msg = {"code": 10002, "data": {}, "msg": "未从飞书上查询到用户信息"}
            else:
                result = app.send_card_data_self("green", "有审批需要你操作", open_id,
                                                 username, approval_name,
                                                 address)
                msg = {"code": 200, "data": {}, "msg": "成功"}
        else:
            msg = {"code": 10002, "data": {}, "msg": "未从飞书上查询到用户信息"}
        return JsonResponse(msg)