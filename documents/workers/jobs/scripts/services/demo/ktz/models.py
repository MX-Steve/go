from django.db import models
from django.contrib.auth.models import AbstractUser


class Accounts(AbstractUser):
    """
       用户表，鉴权使用
    """
    tel = models.CharField('Telephone', max_length=50, null=True, blank=True, default='')
    remark = models.CharField('介绍', max_length=255, null=True, blank=True, default='')
    del_flag = models.IntegerField(blank=True, null=True)

    class Meta:
        db_table = 'auth_accounts'
        verbose_name = '用户列表'
        verbose_name_plural = verbose_name

    @property
    def fullname(self):
        last_name = self.last_name.strip()
        first_name = self.first_name.strip()
        if not last_name.encode().isalnum() and not first_name.startswith(last_name):
            name = '{}{}'.format(last_name, first_name)
        else:
            name = first_name
        if name:
            return name
        else:
            return self.username

    def __str__(self):
        return self.fullname

