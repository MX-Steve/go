from rest_framework import serializers
from .models import *


class AccountsSerializer(serializers.ModelSerializer):
    class Meta:
        model = Accounts
        fields = "__all__"
