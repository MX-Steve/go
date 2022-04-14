# Generated by Django 3.2.7 on 2022-01-06 09:31

from django.db import migrations, models


class Migration(migrations.Migration):

    initial = True

    dependencies = [
    ]

    operations = [
        migrations.CreateModel(
            name='AuditHistory',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('operater', models.CharField(blank=True, default='', max_length=24, null=True, verbose_name='操作者')),
                ('operate_time', models.DateTimeField(blank=True, null=True, verbose_name='操作时间')),
                ('req_url', models.CharField(blank=True, default='', max_length=255, null=True, verbose_name='请求url')),
                ('req_method', models.CharField(default='', max_length=25, null=True, verbose_name='请求方法')),
                ('req_data', models.TextField(default='', null=True, verbose_name='请求数据')),
                ('res_data', models.TextField(default='', null=True, verbose_name='相应数据')),
                ('del_tag', models.IntegerField(blank=True, default=0, null=True, verbose_name='删除标识')),
            ],
        ),
    ]
