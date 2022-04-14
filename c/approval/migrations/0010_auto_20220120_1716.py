# Generated by Django 3.2.7 on 2022-01-20 09:16

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('approval', '0009_auto_20220120_1445'),
    ]

    operations = [
        migrations.RemoveField(
            model_name='faform',
            name='f_approval',
        ),
        migrations.RemoveField(
            model_name='fanode',
            name='f_approval',
        ),
        migrations.RemoveField(
            model_name='ficomment',
            name='f_instance',
        ),
        migrations.RemoveField(
            model_name='fiform',
            name='f_instance',
        ),
        migrations.RemoveField(
            model_name='finstancevalue',
            name='f_approval',
        ),
        migrations.RemoveField(
            model_name='fitask',
            name='f_instance',
        ),
        migrations.RemoveField(
            model_name='fitask',
            name='f_node',
        ),
        migrations.RemoveField(
            model_name='fiviewers',
            name='f_instance',
        ),
        migrations.AddField(
            model_name='faform',
            name='del_tag',
            field=models.SmallIntegerField(default=0, help_text='删除标识'),
        ),
        migrations.AddField(
            model_name='faform',
            name='f_approval_id',
            field=models.IntegerField(default=0, help_text='FApproval ID'),
        ),
        migrations.AddField(
            model_name='fanode',
            name='del_tag',
            field=models.SmallIntegerField(default=0, help_text='删除标识'),
        ),
        migrations.AddField(
            model_name='fanode',
            name='f_approval_id',
            field=models.IntegerField(default=0, help_text='FApproval ID'),
        ),
        migrations.AddField(
            model_name='fapproval',
            name='del_tag',
            field=models.SmallIntegerField(default=0, help_text='删除标识'),
        ),
        migrations.AddField(
            model_name='ficomment',
            name='del_tag',
            field=models.SmallIntegerField(default=0, help_text='删除标识'),
        ),
        migrations.AddField(
            model_name='ficomment',
            name='f_instance_id',
            field=models.IntegerField(default=0, help_text='FInstanceValue ID'),
        ),
        migrations.AddField(
            model_name='fiform',
            name='del_tag',
            field=models.SmallIntegerField(default=0, help_text='删除标识'),
        ),
        migrations.AddField(
            model_name='fiform',
            name='f_instance_id',
            field=models.IntegerField(default=0, help_text='FInstanceValue ID'),
        ),
        migrations.AddField(
            model_name='finstancevalue',
            name='del_tag',
            field=models.SmallIntegerField(default=0, help_text='删除标识'),
        ),
        migrations.AddField(
            model_name='finstancevalue',
            name='f_approval_id',
            field=models.IntegerField(default=0, help_text='FApproval ID'),
        ),
        migrations.AddField(
            model_name='fitask',
            name='del_tag',
            field=models.SmallIntegerField(default=0, help_text='删除标识'),
        ),
        migrations.AddField(
            model_name='fitask',
            name='f_instance_id',
            field=models.IntegerField(default=0, help_text='FInstanceValue ID'),
        ),
        migrations.AddField(
            model_name='fitask',
            name='f_node_id',
            field=models.IntegerField(default=0, help_text='FANode ID'),
        ),
        migrations.AddField(
            model_name='fiviewers',
            name='del_tag',
            field=models.SmallIntegerField(default=0, help_text='删除标识'),
        ),
        migrations.AddField(
            model_name='fiviewers',
            name='f_instance_id',
            field=models.IntegerField(default=0, help_text='FInstanceValue ID'),
        ),
    ]
