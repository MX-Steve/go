from rest_framework import serializers
from .models import FInstanceValue


class FInstanceValueSerializers(serializers.ModelSerializer):
    """FInstanceValue serializers"""
    class Meta:
        model = FInstanceValue
        fields = "__all__"
