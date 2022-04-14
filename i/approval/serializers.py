from rest_framework import serializers
from .models import User, FApproval, FInstanceValue, FAForm, FANode, FIViewers, FIForm, FITask, FIComment, DeployHistory


class FInstanceValueSerializers(serializers.ModelSerializer):
    """FInstanceValue serializers"""
    class Meta:
        model = FInstanceValue
        fields = "__all__"

    def to_representation(self, value):
        data = super().to_representation(value)
        data["status"] = {k: v
                          for k, v in FInstanceValue.status_choices
                          }[data["status"]]
        approval = FApproval.objects.filter(id=data["f_approval_id"]).first()
        data["approval_name"] =approval.approval_name
        data["old_version"] = approval.old_version
        if data["user_id"]:
            user = User.objects.filter(id=data["user_id"]).first()
            if user:
                un = user.last_name.strip() + user.first_name.strip()
                if un:
                    data["username"] = un
                else:
                    data["username"] = user.username
            else:
                data["username"] = "未知"
        else:
            data["username"] = "未知"
        data["job_name"] = approval.job_name
        return data


class FApprovalSerializer(serializers.ModelSerializer):
    """FApproval serializer"""
    class Meta:
        model = FApproval
        fields = "__all__"


class FAFormSerializer(serializers.ModelSerializer):
    """FAForm serializer"""
    class Meta:
        model = FAForm
        fields = "__all__"

    def to_representation(self, value):
        data = super().to_representation(value)
        data["type"] = {k: v for k, v in FAForm.type_choices}[data["type"]]
        return data


class FANodeSerializer(serializers.ModelSerializer):
    """FANode serializer"""
    class Meta:
        model = FANode
        fields = "__all__"

    def to_representation(self, value):
        data = super().to_representation(value)
        data["node_type"] = {k: v
                             for k, v in FANode.node_type_choices
                             }[data["node_type"]]
        data["approval_type"] = {k: v
                             for k, v in FANode.approval_type_choices
                             }[data["approval_type"]]
        if data["user_id"]:
            user = User.objects.filter(id=data["user_id"]).first()
            if user:
                un = user.last_name.strip() + user.first_name.strip()
                if un:
                    data["username"] = un
                else:
                    data["username"] = user.username
            else:
                data["username"] = "未知"
        else:
            if data["name"] in ["开始", "结束"]:
                data["username"] = "admin"
            else:
                data["username"] = "未知"
        return data


class FIViewersSerializer(serializers.ModelSerializer):
    """FIViewers serializer"""
    class Meta:
        model = FIViewers
        fields = "__all__"

    def to_representation(self, value):
        data = super().to_representation(value)
        data["type"] = {k: v for k, v in FIViewers.type_choices}[data["type"]]
        return data

class FIFormSerializer(serializers.ModelSerializer):
    """FIForm serializer"""
    class Meta:
        model = FIForm
        fields = "__all__"

    def to_representation(self, value):
        data = super().to_representation(value)
        data["type"] = {k: v for k, v in FIForm.type_choices}[data["type"]]
        return data

class FITaskSerializer(serializers.ModelSerializer):
    """FITask serializer"""
    class Meta:
        model = FITask
        fields = "__all__"

    def to_representation(self, value):
        data = super().to_representation(value)
        data["type"] = {k: v for k, v in FITask.type_choices}[data["type"]]
        data["status"] = {k: v for k, v in FITask.status_choices}[data["status"]]
        if data["user_id"] == "":
            data["username"] = ""
        else:
            user = User.objects.filter(id=data["user_id"]).first()
            if user is None:
                data["username"] = "未知"
            else:
                un = user.last_name.strip() + user.first_name.strip()
                if un:
                    data["username"] = un
                else:
                    data["username"] = user.username
        return data

class FICommentSerializer(serializers.ModelSerializer):
    """FIComment serializer"""
    class Meta:
        model = FIComment
        fields = "__all__"

class DeployHistorySerializer(serializers.ModelSerializer):
    """DeployHistory serializer"""
    class Meta:
        model = DeployHistory
        fields = "__all__"
