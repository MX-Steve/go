# -*- coding:utf-8 -*-
from rest_framework.permissions import (
    IsAdminUser,
    IsAuthenticated,
    BasePermission
)
from rest_framework.views import APIView
from rest_framework import serializers
from ktz.models import Accounts


class UserInfoSerializers(serializers.ModelSerializer):
    """"这是用户信息页面的序列化器"""

    # user_id=serializers.CharField()
    # token=serializers.CharField()
    class Meta:
        model = Accounts
        fields = ('username', 'user_id', 'email')


class BaseView(APIView):
    permission_classes = (IsAuthenticated,)
    serializer_class = UserInfoSerializers

    def get(self, request, args):
        pass

    def post(self, request, args):
        pass

    def put(self, request, args):
        pass

    def delete(self, request, args):
        pass


class SuperUserpermissions(APIView):
    permission_classes = (IsAdminUser,)

    def get(self, request, args):
        pass

    def post(self, request, args):
        pass

    def put(self, request, args):
        pass

    def delete(self, request, args):
        pass


class AnyLogin(APIView):
    permission_classes = ()
    authentication_classes = ()

    def get(self, request, args):
        pass

    def post(self, request, args):
        pass

    def put(self, request, args):
        pass

    def delete(self, request, args):
        pass
