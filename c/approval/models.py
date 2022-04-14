from django.db import models


class FApproval(models.Model):
    """fei shu approval list"""
    approval_name = models.CharField(max_length=250,
                                     unique=True,
                                     help_text="审批名称",
                                     default=None)
    approval_code = models.CharField(max_length=250,
                                     unique=True,
                                     help_text="审批编码",
                                     default=None)
    open_id = models.CharField(max_length=250,
                               unique=True,
                               help_text="open_id",
                               default=None)
    job_name = models.CharField(max_length=250,
                                unique=True,
                                help_text="jenkins job name",
                                default=None)
    subscribe = models.SmallIntegerField(default=0, help_text="是否订阅")
    del_tag = models.SmallIntegerField(default=0, help_text="删除标识")


class FANode(models.Model):
    """fei shu approval node list"""
    name = models.CharField(max_length=150, help_text="节点名称", default=None)
    need_approver = models.BooleanField(default=False, help_text="是否发起人自选节点")
    node_type_choices = (('AND', '会签'), ('OR', '或签'), ('CC_NODE', '抄送节点'))
    node_type = models.CharField(choices=node_type_choices,
                                 help_text="审批方式",
                                 max_length=100,
                                 default='AND')
    serial_number = models.IntegerField(default=0, help_text="审批节点序列号")
    f_approval_id = models.IntegerField(help_text="FApproval ID", default=0)
    del_tag = models.SmallIntegerField(default=0, help_text="删除标识")


class FAForm(models.Model):
    """fei shu approval form"""
    type_choices = (('input', '单行文本'), ('textarea', '多行文本'), ('number', '数字'),
                    ('date', '日期'))
    name = models.CharField(max_length=150, help_text="控件名称", default=None)
    type = models.IntegerField(choices=type_choices,
                               help_text="控件类型",
                               default='input')
    serial_number = models.IntegerField(default=0, help_text="表单元素序列号")
    f_approval_id = models.IntegerField(help_text="FApproval ID", default=0)
    del_tag = models.SmallIntegerField(default=0, help_text="删除标识")

class FInstanceValue(models.Model):
    """fei shu instance value"""
    instance_code = models.CharField(max_length=250,
                                     unique=True,
                                     help_text="实例编码",
                                     default=None)
    status_choices = (('PENDING', '审批中'), ('APPROVED', '通过'),
                      ('REJECTED', '拒绝'), ('CANCELED', '撤回'), ('DELETED',
                                                               '删除'))
    status = models.CharField(choices=status_choices,
                                 help_text="审批实例状态",
                                 max_length=100,
                                 default='PENDING')
    user_id = models.CharField(max_length=150, help_text="用户ID", default=None)
    descriptions = models.CharField(max_length=250,
                                    help_text="描述信息",
                                    default=None)
    form = models.TextField(default="", help_text="临时使用1")
    task_id = models.CharField(max_length=150, default="", help_text="临时使用2")
    job_number = models.IntegerField(help_text="job 编号", default=0)
    f_approval_id = models.IntegerField(help_text="FApproval ID", default=0)
    del_tag = models.SmallIntegerField(default=0, help_text="删除标识")


class FIViewers(models.Model):
    """fei shu approval viewers"""
    type_choices = (('TENANT', '租户内可见'), ('DEPARTMENT', '指定部门'),
                    ('USER', '指定用户'), ('NONE', '任何人都不可见'))
    type = models.CharField(choices=type_choices,
                            help_text="可见人类型",
                            max_length=100,
                            default='AND')
    user_id = models.CharField(max_length=150,
                               help_text="可见人用户 ID",
                               default=None)
    f_instance_id = models.IntegerField(help_text="FInstanceValue ID", default=0)
    del_tag = models.SmallIntegerField(default=0, help_text="删除标识")


class FIForm(models.Model):
    """fei shu instance form"""
    type_choices = (('input', '单行文本'), ('textarea', '多行文本'), ('number', '数字'),
                    ('date', '日期'))
    name = models.CharField(max_length=250, help_text="控件名称", default=None)
    en_name = models.CharField(max_length=250, help_text="控件英文名", default=None)
    type = models.IntegerField(choices=type_choices,
                               help_text="控件类型",
                               default='input')
    f_instance_id = models.IntegerField(help_text="FInstanceValue ID", default=0)
    del_tag = models.SmallIntegerField(default=0, help_text="删除标识")


class FITask(models.Model):
    """fei shu instance task list"""
    user_id = models.CharField(max_length=150, help_text="审批人", default=None)
    status_choices = (('PENDING', '审批中'), ('APPROVED', '同意'),
                      ('REJECTED', '拒绝'), ('TRANSFERRED', '已转交'), ('DONE',
                                                                   '完成'))
    status = models.IntegerField(choices=status_choices,
                                 help_text="任务状态",
                                 default=0)
    type_choices = (('AND', '会签'), ('OR', '或签'), ('AUTO_PASS', '自动通过'),
                    ('AUTO_REJECT', '自动拒绝'), ('SEQUENTIAL', '按顺序'))
    type = models.IntegerField(choices=type_choices,
                               help_text="审批方式",
                               default='input')
    comment = models.CharField(max_length=250, help_text="评论内容", default=None)
    start_time = models.IntegerField(default=0, help_text="task 开始时间")
    end_time = models.IntegerField(default=0, help_text="task 完成时间")
    f_node_id =models.IntegerField(help_text="FANode ID", default=0)
    f_instance_id = models.IntegerField(help_text="FInstanceValue ID", default=0)
    del_tag = models.SmallIntegerField(default=0, help_text="删除标识")


class FIComment(models.Model):
    """fei shu instance comment list"""
    user_id = models.CharField(max_length=150,
                               help_text="发表评论用户",
                               default=None)
    comment = models.CharField(max_length=250, help_text="评论内容", default=None)
    create_time = models.CharField(max_length=150,
                                   help_text="评论时间",
                                   default=None)
    f_instance_id = models.IntegerField(help_text="FInstanceValue ID", default=0)
    del_tag = models.SmallIntegerField(default=0, help_text="删除标识")
