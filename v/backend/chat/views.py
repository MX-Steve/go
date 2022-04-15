from ipaddress import ip_address
from django.http import JsonResponse
from django.db.models import Q
from chat.models import Host
import json


def putHostView(request):
    data = json.loads(request.body)
    if "username" not in data or "password" not in data or "port" not in data or "ip_address" not in data:
        return JsonResponse({
            "code": 10000,
            "data": {},
            "msg": "ip_address,username,port,password required"
        })
    hObj = Host.objects.filter(ip_address=data["ip_address"])
    if hObj.count() == 0:
        Host.objects.create(ip_address=data["ip_address"],
                            username=data["username"],
                            pkey=data["password"],
                            port=data["port"])
    h = Host.objects.filter(ip_address=data["ip_address"]).first()
    msg = {"code": 200, "data": {"id": h.id}, "msg": "success"}
    return JsonResponse(msg)


def delHostView(request):
    data = json.loads(request.body)
    if "id" not in data:
        return JsonResponse({"code": 10000, "data": {}, "msg": "id required"})
    host = Host.objects.get(id=data['id'])
    if host:
        host.delete()
        msg = {"code": 200, "data": {}, "msg": "删除成功"}
    else:
        msg = {"code": 10000, "data": {}, "msg": "不存在，无法删除"}
    return JsonResponse(msg)


def getHostView(request):
    id = request.GET.get("id", "")
    ip_address = request.GET.get("ip_address", "")
    username = request.GET.get("username", "")
    port = request.GET.get("port", "")
    password = request.GET.get("password", "")
    q = Q()
    hs_list = []
    if id:
        hs = Host.objects.filter(id=id).first()
        hs_list.append({
            "id": hs.id,
            "ip_address": hs.ip_address,
            "username": hs.username,
            "port": hs.port,
            "password": hs.pkey
        })
    else:
        if ip_address:
            q.children.append(('ip_address', ip_address))
        if username:
            q.children.append(('username', username))
        if port:
            q.children.append(('port', port))
        if password:
            q.children.append(('pkey', password))
        hs = Host.objects.filter(q)
        for h in hs:
            hs_list.append({
                "id": h.id,
                "ip_address": h.ip_address,
                "username": h.username,
                "port": h.port,
                "password": h.pkey
            })
    return JsonResponse({"code": 200, "data": hs_list, "msg": "success"})