# Generated by Django 3.2.7 on 2022-01-27 01:39

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('approval', '0010_auto_20220120_1716'),
    ]

    operations = [
        migrations.AddField(
            model_name='faform',
            name='serial_number',
            field=models.IntegerField(default=0, help_text='表单元素序列号'),
        ),
        migrations.AddField(
            model_name='fanode',
            name='serial_number',
            field=models.IntegerField(default=0, help_text='审批节点序列号'),
        ),
    ]
