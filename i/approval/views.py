import json
import uuid
from django.http import JsonResponse
from google.protobuf.json_format import MessageToJson
from backend.stubs import GetStub
from backend.proto.djentry_pb2 import UserInfoRequest
from django.conf import settings
from utils.kafkas import client
import func_timeout


def grpc_test(req):
    stub = GetStub()
    req = UserInfoRequest(user_id="abc-23a")
    print(req)
    r = stub.GetUserInfo(req)
    r = json.loads(MessageToJson(r))
    return JsonResponse({"code": 200, "data": {"user": r}, "msg": "grpc test"})


def kafka_get(req):
    topic = req.GET.get('topic', '')
    key = req.GET.get('key', '')
    offset = req.GET.get('offset', '')
    if topic == "" or key == "" or offset == "":
        return JsonResponse({
            "code": 100001,
            "data": {},
            "msg": "必须提供topic，key，offset值"
        })
    try:
        result, last_offset = client.get_data(settings.KAFKA_ADDR, topic,
                                              str(uuid.uuid4()), key,
                                              int(offset))
    except func_timeout.exceptions.FunctionTimedOut:
        return JsonResponse({"code": 10001, "data": [], "msg": "没有kafka对应数据"})
    return JsonResponse({
        "code": 200,
        "data": {
            "result": result,
            "last_offset": last_offset
        },
        "msg": "获取kafka数据成功"
    })
