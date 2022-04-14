# pylint: disable=wildcard-import,unused-wildcard-import
import json
from drf_yasg2 import openapi
from drf_yasg2.utils import swagger_auto_schema
from drf_yasg2.openapi import Parameter, IN_QUERY
# pylint: disable=no-name-in-module
from manages.serializers import *
from utils.util import swagger_auto_schema2


def role_get():
    response = openapi.Response("response description", RoleSerializer)
    return swagger_auto_schema(manual_parameters=[
        Parameter('id',
                  IN_QUERY,
                  type=openapi.Schema(type=openapi.TYPE_INTEGER))
    ],
                               responses={200: response})


def role_put():
    return swagger_auto_schema2(
        req={
            "params": {
                "id": "int",
                "name": "string",
                "desc": "string",
            },
            "required": ["id", "name", "desc"],
            type: "body"
        },
        res={200: json.dumps({
            "code": 200,
            "msg": "更新角色成功.",
            "data": {}
        })},
        info="更新角色")


def role_post():
    return swagger_auto_schema2(
        req={
            "params": {
                "name": "string",
                "desc": "string",
            },
            "required": ["name", "desc"],
            type: "body"
        },
        res={200: json.dumps({
            "code": 200,
            "msg": "创建角色成功.",
            "data": {}
        })},
        info="创建角色")


def role_delete():
    return swagger_auto_schema2(
        req={
            "params": {
                "id": "int",
            },
            "required": ["id"],
            type: "body"
        },
        res={200: json.dumps({
            "code": 200,
            "msg": "删除角色成功.",
            "data": {}
        })},
        info="删除角色")
