import os
import json
import time
from django.http import JsonResponse
from utils import baseview
from utils.encr import AESCipher
from utils.feishu_approval import ApprovalT
from utils.call_jenkins import JenkinsApi
from utils.feishu_rebot import FeishuRobot
from approval.models import FApproval, FInstanceValue
from users.models import User
from approval.serializers import FInstanceValueSerializers
from utils.hello.run import Run


class FInstanceView(baseview.AnyLogin):
    def post(self, request, args=None):
        """工单流程发布状态更新"""
        data = request.data
        msg = {"code": 200, "data": {}, "msg": "状态更新成功"}
        if data["type"] == "status":
            instance_code = data["instance_code"]
            status = data["status"]
            FInstanceValue.objects.filter(instance_code=instance_code).update(
                status=status, descriptions=data["comment"])
        return JsonResponse(msg)

    def get(self, request, args=None):
        """工单流程发布状态查询"""
        msg = {"code": 200, "data": {}, "msg": "查询状态成功"}
        instance_code = request.GET.get("instance_code", "")
        if instance_code:
            fi_obj = FInstanceValue.objects.filter(instance_code=instance_code)
            if not fi_obj.exists():
                msg = {"code": 10001, "data": {}, "msg": "所查询的工单流程不存在"}
            else:
                fi = FInstanceValue.objects.filter(
                    instance_code=instance_code).first()
                data = FInstanceValueSerializers(fi, many=True)
                msg["data"] = data
        else:
            msg = {"code": 10005, "data": {}, "msg": "必须传递 instance_code 参数"}
        return JsonResponse(msg)


class FApprovalView(baseview.AnyLogin):
    def post(self, request, args=None):
        """审批流程绑定 jenkins job, 并订阅 """
        msg = {"code": 200, "data": {}, "msg": "更改成功"}
        approval_code = request.data["approval_code"]
        approval_name = request.data["approval_name"]
        job_name = request.data["job_name"]
        subscribe = int(request.data["subscribe"])
        fa_obj = FApproval.objects.filter(approval_code=approval_code)
        if fa_obj.exists():
            msg = {
                "code": 10001,
                "data": {},
                "msg": "%s 已经订阅了" % (approval_code)
            }
        else:
            app = ApprovalT()
            app.get_tenant_access_token()
            result = app.subscribe_events(approval_code)
            if result["code"] != 0:
                msg = {"code": 10001, "data": {}, "msg": "订阅失败，请检查原因"}
            else:
                fa = FApproval.objects.filter(approval_code=approval_code)
                if not fa.exists():
                    FApproval.objects.create(approval_code=approval_code,
                                             approval_name=approval_name,
                                             open_id="open_" +
                                             str(time.time()).split(".")[0],
                                             job_name=job_name,
                                             subscribe=subscribe)
                else:
                    msg = {
                        'code': 10003,
                        'data': {},
                        "msg": "%s 已经订阅，无需再次订阅" % (approval_code)
                    }
        return JsonResponse(msg)


KVS = {
    "项目名称": "PROJECT",
    "仓库名称": "GIT_NAME",
    "分支名称": "BrName",
    "响应评估": "LEVEL",
    "ONES ID": "ONES_ID",
    "服务类型": "SERVICE_TYPE",
    "服务名称": "SERVICE"
}


class ListenApprovalEventView(baseview.AnyLogin):
    def post(self, request, args=None):
        """事件审批流监听"""
        app = ApprovalT()
        app.get_tenant_access_token()
        msg = {"code": 200, "data": [], "msg": "成功"}
        encrypt = request.data["encrypt"]
        cipher = AESCipher("approvalApply")
        event_data = json.loads(cipher.decrypt_string(encrypt))
        # 连接心跳，绑定飞书 APP 使用
        if "challenge" in event_data:
            return JsonResponse({"challenge": event_data["challenge"]})
        event_obj = event_data["event"]
        approval_code = event_obj["approval_code"]
        instance_code = event_obj["instance_code"]
        approval = FApproval.objects.filter(approval_code=approval_code)
        res = app.get_instance(instance_code)
        data = res['data']
        approval_name = data["approval_name"]
        open_id = data["open_id"]
        if not approval.exists():
            # 审批流未入库
            FApproval.objects.filter(approval_code=approval_code).update(
                approval_name=approval_name,
                open_id=open_id,
                subscribe=0,
                job_name="init...",
                del_tag=0)
            user = User.objects.filter(username="admin").first()
            url = user.robot_url
            secret = user.robot_secret
            msg = """instance_code: %s 还未绑定 job .<at user_id="all">所有人</at>""" % (
                instance_code)
            robot = FeishuRobot(url, secret)
            robot.sign_generate()
            robot.send(msg)
            approval_instance = FApproval.objects.filter(
                approval_code=approval_code).first()
            FInstanceValue.objects.create(instance_code=instance_code,
                                          form=data["form"],
                                          user_id="",
                                          task_id="",
                                          descriptions="初始化...",
                                          f_approval_id=approval_instance.id)
        else:
            approval_instance = FApproval.objects.filter(
                approval_code=approval_code).first()
            instance = FInstanceValue.objects.filter(
                instance_code=instance_code)
            if not instance.exists():
                print(instance_code, " 审批还未入库")
                FInstanceValue.objects.create(
                    instance_code=instance_code,
                    form=data["form"],
                    user_id="",
                    task_id="",
                    descriptions="初始化...",
                    status="PENDING",
                    f_approval_id=approval_instance.id)
            else:
                print(instance_code, " 审批已经入库")
                fi = FInstanceValue.objects.filter(
                    instance_code=instance_code,
                    status__in=["PENDING", "APPROVED"]).first()
                if not fi:
                    return JsonResponse(msg)
                print("接收到审批实例一下参数: ")
                d = json.loads(fi.form)
                d_obj = {}
                for item in d:
                    name = item["name"]
                    if name in ["服务类型", "分支名称", "服务名称", "项目名称"]:
                        print("name: ", item["name"])
                        print("value: ", item["value"])
                        d_obj[KVS[item["name"]]] = item["value"]
                if fi.status == "PENDING":
                    task_list = data["task_list"]
                    for task in task_list:
                        if task["node_name"] == "经办人" and \
                            task["status"] == "DONE":
                            job_name = approval_instance.job_name
                            print("%s 审批已经通过，准备发布了，job名称: %s" %
                                  (instance_code, job_name))
                            jenkinsApi = JenkinsApi()
                            num = jenkinsApi.get_next_buildnum(job_name)
                            FInstanceValue.objects.filter(
                                instance_code=instance_code).update(
                                    status="APPROVED",
                                    user_id=data["user_id"],
                                    task_id=task["id"],
                                    job_number=num)
                            jenkinsApi.execute_job(job_name, d_obj)
                            break
                        if task["status"] == "REJECTED":
                            FInstanceValue.objects.filter(
                                instance_code=instance_code).update(
                                    status="REJECTED",
                                    descriptions="被拒绝发布",
                                    user_id=data["user_id"],
                                    task_id=task["id"],
                                )
                            break
        return JsonResponse(msg)

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