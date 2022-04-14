import json
from drf_yasg2 import openapi
from drf_yasg2.utils import swagger_auto_schema
from drf_yasg2.openapi import Parameter, IN_QUERY
# pylint: disable=no-name-in-module
from assets.serializers import ZoneInfoSerializers, DeviceTypeSerializers,\
    DeviceStatusSerializers, IdleAssetsSerializers
from utils.util import swagger_auto_schema2


def zone_info_get():
    response = openapi.Response("response description", ZoneInfoSerializers)
    return swagger_auto_schema(manual_parameters=[
        Parameter('id',
                  IN_QUERY,
                  type=openapi.Schema(type=openapi.TYPE_INTEGER))
    ],
                               responses={200: response})


def zone_info_put():
    return swagger_auto_schema2(
        req={
            "params": {
                "id": "int",
                "name": "string",
                "description": "string"
            },
            "required": ["id", "name", "description"],
            type: "body"
        },
        res={200: json.dumps({
            "code": 200,
            "msg": "更新区信息成功.",
            "data": {}
        })},
        info="更新区域")


def zone_info_post():
    return swagger_auto_schema2(
        req={
            "params": {
                "name": "string",
                "description": "string"
            },
            "required": ["name", "description"],
            type: "body"
        },
        res={200: json.dumps({
            "code": 200,
            "msg": "区域创建成功.",
            "data": {}
        })},
        info="创建区域")


def zone_info_delete():
    return swagger_auto_schema2(
        req={
            "params": {
                "id": "int"
            },
            "required": ["id"],
            type: "body"
        },
        res={200: json.dumps({
            "code": 200,
            "msg": "区域删除成功.",
            "data": {}
        })},
        info="删除区域")


def device_type_get():
    response = openapi.Response("response description", DeviceTypeSerializers)
    return swagger_auto_schema(manual_parameters=[
        Parameter('id',
                  IN_QUERY,
                  type=openapi.Schema(type=openapi.TYPE_INTEGER))
    ],
                               responses={200: response})


def device_type_put():
    return swagger_auto_schema2(
        req={
            "params": {
                "id": "int",
                "name": "string"
            },
            "required": ["id","name"],
            type: "body"
        },
        res={200: json.dumps({
            "code": 200,
            "msg": "更新设备类型成功[资产标签].",
            "data": {}
        })},
        info="更新设备类型")


def device_type_post():
    return swagger_auto_schema2(
        req={
            "params": {
                "name": "string"
            },
            "required": ["name"],
            type: "body"
        },
        res={200: json.dumps({
            "code": 200,
            "msg": "创建设备类型成功.",
            "data": {}
        })},
        info="创建设备类型")


def device_type_delete():
    return swagger_auto_schema2(
        req={
            "params": {
                "id": "int"
            },
            "required": ["id"],
            type: "body"
        },
        res={200: json.dumps({
            "code": 200,
            "msg": "删除设备类型成功.",
            "data": {}
        })},
        info="删除设备类型")


def device_status_get():
    response = openapi.Response("response description",
                                DeviceStatusSerializers)
    return swagger_auto_schema(manual_parameters=[
        Parameter('id',
                  IN_QUERY,
                  type=openapi.Schema(type=openapi.TYPE_INTEGER))
    ],
                               responses={200: response})


def device_status_put():
    return swagger_auto_schema2(
        req={
            "params": {
                "id": "int",
                "name": "string"
            },
            "required": ["id", "name"],
            type: "body"
        },
        res={200: json.dumps({
            "code": 200,
            "msg": "设备状态更新成功.",
            "data": {}
        })},
        info="更新设备状态")


def device_status_post():
    return swagger_auto_schema2(
        req={
            "params": {
                "name": "string"
            },
            "required": ["name"],
            type: "body"
        },
        res={200: json.dumps({
            "code": 200,
            "msg": "新增设备状态成功.",
            "data": {}
        })},
        info="创建设备状态")


def device_status_delete():
    return swagger_auto_schema2(
        req={
            "params": {
                "id": "int"
            },
            "required": ["id"],
            type: "body"
        },
        res={200: json.dumps({
            "code": 200,
            "msg": "设备状态删除成功.",
            "data": {}
        })},
        info="删除设备状态")


def idle_assets_get():
    response = openapi.Response("response description", IdleAssetsSerializers)
    return swagger_auto_schema(manual_parameters=[
        Parameter('id',
                  IN_QUERY,
                  type=openapi.Schema(type=openapi.TYPE_INTEGER))
    ],
                               responses={200: response})


def idle_assets_put():
    return swagger_auto_schema2(
        req={
            "params": {
                "id": "int",
                "idc": "int",
                "name": "string",
                "count": "int",
                "description": "string",
            },
            "required": ["id", "idc", "name", "count", "description"],
            type: "body"
        },
        res={200: json.dumps({
            "code": 200,
            "msg": "更新闲置资产信息成功.",
            "data": {}
        })},
        info="更新闲置资产")


def idle_assets_post():
    return swagger_auto_schema2(
        req={
            "params": {
                "idc": "int",
                "name": "string",
                "count": "int",
                "description": "string",
            },
            "required": ["idc", "name", "count"],
            type: "body"
        },
        res={200: json.dumps({
            "code": 200,
            "msg": "创建闲置资产信息成功.",
            "data": {}
        })},
        info="创建闲置资产")


def idle_assets_delete():
    return swagger_auto_schema2(
        req={
            "params": {
                "id": "int"
            },
            "required": ["id"],
            type: "body"
        },
        res={200: json.dumps({
            "code": 200,
            "msg": "删除闲置资产信息成功.",
            "data": {}
        })},
        info="删除闲置资产")
