# Generated by Django 3.2.7 on 2022-01-20 03:12

from django.db import migrations, models
import django.db.models.deletion


class Migration(migrations.Migration):

    dependencies = [
        ('approval', '0004_fviewers'),
    ]

    operations = [
        migrations.CreateModel(
            name='FAForm',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.CharField(default=None, help_text='控件名称', max_length=150)),
                ('type', models.IntegerField(choices=[('input', '单行文本'), ('textarea', '多行文本'), ('number', '数字'), ('date', '日期')], default='input', help_text='控件类型')),
                ('f_approval', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to='approval.fapproval')),
            ],
        ),
        migrations.CreateModel(
            name='FANode',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.CharField(default=None, help_text='节点名称', max_length=150)),
                ('need_approver', models.BooleanField(default=False, help_text='是否发起人自选节点')),
                ('node_id', models.CharField(default=None, help_text='节点 ID', max_length=150)),
                ('custom_node_id', models.CharField(default=None, help_text='节点自定义 ID', max_length=150)),
                ('node_type', models.CharField(choices=[('AND', '会签'), ('OR', '或签'), ('CC_NODE', '抄送节点')], default='AND', help_text='审批方式', max_length=100)),
                ('f_approval', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to='approval.fapproval')),
            ],
        ),
        migrations.CreateModel(
            name='FAViewers',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('type', models.CharField(choices=[('TENANT', '租户内可见'), ('DEPARTMENT', '指定部门'), ('USER', '指定用户'), ('NONE', '任何人都不可见')], default='AND', help_text='可见人类型', max_length=100)),
                ('open_id', models.CharField(default=None, help_text='open id', max_length=150)),
                ('user_id', models.CharField(default=None, help_text='可见人用户 ID', max_length=150)),
                ('f_approval', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to='approval.fapproval')),
            ],
        ),
        migrations.CreateModel(
            name='FIComment',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('user_id', models.CharField(default=None, help_text='发表评论用户', max_length=150)),
                ('comment', models.CharField(default=None, help_text='评论内容', max_length=250)),
                ('create_time', models.CharField(default=None, help_text='评论时间', max_length=150)),
            ],
        ),
        migrations.CreateModel(
            name='FIForm',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.CharField(default=None, help_text='控件名称', max_length=250)),
                ('en_name', models.CharField(default=None, help_text='控件英文名', max_length=250)),
                ('type', models.IntegerField(choices=[('input', '单行文本'), ('textarea', '多行文本'), ('number', '数字'), ('date', '日期')], default='input', help_text='控件类型')),
            ],
        ),
        migrations.CreateModel(
            name='FITask',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('user_id', models.CharField(default=None, help_text='审批人', max_length=150)),
                ('status', models.IntegerField(choices=[('PENDING', '审批中'), ('APPROVED', '同意'), ('REJECTED', '拒绝'), ('TRANSFERRED', '已转交'), ('DONE', '完成')], default=0, help_text='任务状态')),
                ('type', models.IntegerField(choices=[('AND', '会签'), ('OR', '或签'), ('AUTO_PASS', '自动通过'), ('AUTO_REJECT', '自动拒绝'), ('SEQUENTIAL', '按顺序')], default='input', help_text='审批方式')),
                ('start_time', models.IntegerField(default=0, help_text='task 开始时间')),
                ('end_time', models.IntegerField(default=0, help_text='task 完成时间')),
            ],
        ),
        migrations.DeleteModel(
            name='FComment',
        ),
        migrations.DeleteModel(
            name='FNode',
        ),
        migrations.DeleteModel(
            name='FViewers',
        ),
        migrations.RemoveField(
            model_name='finstancevalue',
            name='form',
        ),
        migrations.RemoveField(
            model_name='finstancevalue',
            name='task_id',
        ),
        migrations.AddField(
            model_name='fitask',
            name='f_instance',
            field=models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to='approval.finstancevalue'),
        ),
        migrations.AddField(
            model_name='fitask',
            name='f_node',
            field=models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to='approval.fanode'),
        ),
        migrations.AddField(
            model_name='fiform',
            name='f_instance',
            field=models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to='approval.finstancevalue'),
        ),
        migrations.AddField(
            model_name='ficomment',
            name='f_instance',
            field=models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to='approval.finstancevalue'),
        ),
    ]