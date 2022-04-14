import uuid
import json
from django.contrib.auth.models import AbstractUser
from django.db import models
from utils.mixins import ModelMixin


class User(AbstractUser):
    """The basic user model"""
    id = models.UUIDField(primary_key=True,
                          default=uuid.uuid4,
                          editable=False,
                          help_text="用户唯一标识ID")
    roles = models.CharField(default='[]', max_length=50, help_text="用户角色id列表")
    robot_url = models.CharField(max_length=450,
                                 null=True,
                                 blank=True,
                                 verbose_name="机器人地址")
    robot_secret = models.CharField(max_length=200,
                                    null=True,
                                    blank=True,
                                    verbose_name="机器人密钥")
    user_id = models.CharField(max_length=150,
                               null=True,
                               blank=True,
                               verbose_name="飞书user_id")

    class Meta:
        db_table = "users_user"
        managed = False


class AuditHistory(models.Model):
    """audit history"""
    operater = models.CharField(max_length=24,
                                null=True,
                                blank=True,
                                default="",
                                verbose_name="操作者")
    operate_time = models.DateTimeField(null=True,
                                        blank=True,
                                        verbose_name="操作时间")
    req_url = models.CharField(max_length=255,
                               null=True,
                               blank=True,
                               default="",
                               verbose_name="请求url")
    req_method = models.CharField(max_length=25,
                                  null=True,
                                  default="",
                                  verbose_name='请求方法')
    req_data = models.TextField(null=True, default="", verbose_name="请求数据")
    res_data = models.TextField(null=True, default="", verbose_name="相应数据")
    del_tag = models.IntegerField(null=True,
                                  blank=True,
                                  default=0,
                                  verbose_name="删除标识")

    class Meta:
        db_table = "audit_audithistory"
        managed = False


class Role(models.Model, ModelMixin):
    "role"
    name = models.CharField(max_length=50)
    desc = models.CharField(max_length=255, null=True)
    page_perms = models.TextField(null=True)
    created_at = models.CharField(max_length=20, default="")
    created_by = models.ForeignKey(User,
                                   on_delete=models.PROTECT,
                                   related_name='+')
    del_tag = models.SmallIntegerField(default=0, help_text="删除标识")

    def to_dict(self, *args, **kwargs):
        tmp = super().to_dict(*args, **kwargs)
        tmp['page_perms'] = json.loads(
            self.page_perms) if self.page_perms else {}
        return tmp
