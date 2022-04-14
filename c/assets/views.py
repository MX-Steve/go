import json
import os
import django
import requests
from django.http import JsonResponse
from .connection import WebSocket
from utils import ssh, baseview
from utils.auth import auth
from .models import Machine
from chat.models import PeriodicTask
from .tasks import add2
from audit.apis.audit import PutAudit
from assets import tasks
from django.conf import settings

BaseURL = 'http://127.0.0.1:' + str(settings.CMDB_BACK_PORT)
ITSMURL = settings.ITSMURL

os.environ["DJANGO_ALLOW_ASYNC_UNSAFE"] = "true"
django.setup()


async def websocket_view(socket: WebSocket, id):
    machine = Machine.objects.filter(id=id).first()
    host = ssh.SSH2(hostname=machine.ip_address,
                    port=machine.port,
                    username=machine.username,
                    password=machine.password)
    await socket.accept()
    cmd = ""
    history = []
    now_cd = ""
    last_cd = ""
    up_num = 0
    while True:
        message = await socket.receive_text()
        if "\r" not in message:
            # \177: 删除键;
            # \x1b[D: 左键; \x1b[C: 右键; \x1b[A: 上键; \x1b[B: 下键
            # \003: ctrl + c
            if message == '\177':  # 删除键
                cmd = cmd[:len(cmd) - 1]
            elif message == '\003':  # ctrl + c
                cmd = ""
            elif message == '\x1b[A':  # 上
                print("按了上键")
                up_num += 1
                if up_num == 10:
                    up_num = 0
                    cmd = ""
                else:
                    cmd = history[len(history) - up_num]
                await socket.send_text(cmd)
            elif message == '\x1b[B':  # 下
                print("按了下键")
                up_num -= 1
                if up_num > 0:
                    cmd = history[len(history) - up_num]
                    await socket.send_text(cmd)
            elif message == '\x1b[D':  # 左
                print("按了左键")
            elif message == '\x1b[C':  # 右
                print("按了右键")
            else:
                up_num = 0
                cmd += message
        else:
            try:
                cmd = cmd.strip()
                print("cmd.len: ", len(cmd))
                if cmd != "":
                    history.append(cmd)
                print("history: ", history)
                print("now_cd: ", now_cd)
                if len(history) > 10:
                    history = history[1:]
                if "cd " in cmd:
                    last_cd = now_cd
                    now_cd = cmd
                if now_cd != "":
                    cmd = now_cd + ";" + cmd
                out, err = host.do(cmd)
                if err:
                    if "cd: " in err:
                        now_cd = last_cd
                    err = "\x1B[1;3;31m " + err + " \x1B[0m"
                    await socket.send_text(err + "\r$ ")
                else:
                    await socket.send_text(out + "\r$ ")
            except:
                err = "\x1B[1;3;31m 不能够连接当前机器，请确认信息是否正确 \x1B[0m"
                await socket.send_text(err + " !!! ")
                await socket.close()
            cmd = ""


def test_celery(request):
    add2.delay(3, 9)
    # result = add2.apply_async(args=[3, 8])
    return JsonResponse({
        "code": 200,
        "data": {
            # "task_id": result.task_id
        },
        "msg": "测试 celery 功能"
    })


class go_ssh(baseview.BaseView):
    @auth("projects.ssh.view")
    def get(self, request, args=None):
        id = request.GET.get('id', '')
        print(id)
        return JsonResponse({
            "code": 200,
            "data": {
                "id": id,
            },
            "msg": "进入ssh终端"
        })

class run_one(baseview.BaseView):
    def post(self, request, args=None):
        data = request.data
        name = data.get('name','')
        if name:
            job = PeriodicTask.objects.filter(name=data["name"],del_tag=0).first()
            func_name = job.task.split('.')[-1]
            if func_name == "machine_get":
                url = BaseURL + "/assets/v1/inner/machine"
                r = requests.post(url)
                return JsonResponse(r.json())
            elif func_name == "oss_get":
                url = BaseURL + "/assets/v1/inner/oss"
                r = requests.post(url)
                return JsonResponse(r.json())
            elif func_name == "disk_get":
                url = BaseURL + "/assets/v1/inner/disk"
                r = requests.post(url)
                return JsonResponse(r.json())
            elif func_name == "domain_get":
                url = BaseURL + "/assets/v1/inner/domain"
                r = requests.post(url)
                return JsonResponse(r.json())
            elif func_name == "rds_get":
                url = BaseURL + "/assets/v1/inner/rds"
                r = requests.post(url)
                return JsonResponse(r.json())
            elif func_name == "vpc_get":
                url = BaseURL + "/assets/v1/inner/vpc"
                r = requests.post(url)
                return JsonResponse(r.json())
            elif func_name == "cdn_get":
                url = BaseURL + "/assets/v1/inner/cdn"
                r = requests.post(url)
                return JsonResponse(r.json())
            elif func_name == "slb_get":
                url = BaseURL + "/assets/v1/inner/slb"
                r = requests.post(url)
                return JsonResponse(r.json())
            elif func_name == "approval_get":
                url = ITSMURL
                r = requests.post(url)
                return JsonResponse(r.json)
            elif func_name == "users_update":
                url = BaseURL + "/users/v1/user-update"
                r = requests.post(url)
                return JsonResponse(r.json)
            elif func_name == "dnsrecord_get":
                url = BaseURL + "/assets/v1/inner/dnsrecord"
                r = requests.post(url)
                return JsonResponse(r.json)
            else:
                res = {"code": 10001, "data": {}, "msg": "job不存在"}
                return JsonResponse(res)