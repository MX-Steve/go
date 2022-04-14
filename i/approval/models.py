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
        managed = False
        db_table = 'users_user'


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

    class Meta:
        managed = False
        db_table = 'manages_role'


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
        managed = False
        db_table = 'audit_audithistory'


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
    descriptions = models.CharField(max_length=250,
                                    help_text="描述信息",
                                    default=None)
    old_version = models.SmallIntegerField(default=1, help_text="v1 版本")
    del_tag = models.SmallIntegerField(default=0, help_text="删除标识")


class FANode(models.Model):
    """fei shu approval node list"""
    name = models.CharField(max_length=150, help_text="节点名称", default=None)
    need_approver = models.BooleanField(default=False, help_text="是否发起人自选节点")
    node_type_choices = (('AND', '会签'), ('OR', '或签'))
    node_type = models.CharField(choices=node_type_choices,
                                 help_text="审批方式",
                                 max_length=100,
                                 default='AND')
    user_id = models.CharField(max_length=150, help_text="用户ID", default=None)
    approval_type_choices = (('AUTO_PASS', '自动通过'), ('AUTO_REFUSED', '自动拒绝'),
                             ('MANUAL', '人工操作'))
    approval_type = models.CharField(max_length=150,
                                     choices=approval_type_choices,
                                     help_text="操作方式",
                                     default='MANUAL')
    serial_number = models.IntegerField(default=0, help_text="审批节点序列号")
    f_approval_id = models.IntegerField(help_text="FApproval ID", default=0)
    del_tag = models.SmallIntegerField(default=0, help_text="删除标识")


class FAForm(models.Model):
    """fei shu approval form"""
    type_choices = (('input', '单行文本'), ('textarea', '多行文本'), ('number', '数字'),
                    ('date', '日期'), ('radioV2', '单选项'))
    name = models.CharField(max_length=150, help_text="控件名称", default=None)
    en_name = models.CharField(max_length=250, help_text="控件英文名", default=None)
    type = models.CharField(choices=type_choices,
                            max_length=50,
                            help_text="控件类型",
                            default='input')
    serial_number = models.IntegerField(default=0, help_text="表单元素序列号")
    f_approval_id = models.IntegerField(help_text="FApproval ID", default=0)
    del_tag = models.SmallIntegerField(default=0, help_text="删除标识")


class SelectOptions(models.Model):
    """select options"""
    value = models.CharField(max_length=120, default="", help_text="选项value值")
    form_id = models.IntegerField(default=0, help_text="form id")
    del_tag = models.SmallIntegerField(default=0, help_text="删除标识")


class FInstanceValue(models.Model):
    """fei shu instance value"""
    instance_code = models.CharField(max_length=250,
                                     unique=True,
                                     help_text="实例编码",
                                     default=None)
    status_choices = (('BEFORE', '未开始'), ('PENDING', '审批中'), ('DEPLOY', '发布中'),
                      ('APPROVED', '通过'), ('REJECTED', '拒绝'),
                      ("REFUTED", "驳回"), ('DONE', '完成'))
    status = models.CharField(choices=status_choices,
                              help_text="审批实例状态",
                              max_length=100,
                              default='PENDING')
    user_id = models.CharField(max_length=150, help_text="用户ID", default=None)
    descriptions = models.CharField(max_length=250,
                                    help_text="描述信息",
                                    default=None)
    appendix = models.CharField(max_length=1200,
                                help_text="附件信息",
                                default=None)
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
    f_instance_id = models.IntegerField(help_text="FInstanceValue ID",
                                        default=0)
    del_tag = models.SmallIntegerField(default=0, help_text="删除标识")


class FIForm(models.Model):
    """fei shu instance form"""
    type_choices = (('input', '单行文本'), ('textarea', '多行文本'), ('number', '数字'),
                    ('date', '日期'), ('radioV2', '单选项'))
    name = models.CharField(max_length=250, help_text="控件名称", default=None)
    en_name = models.CharField(max_length=250, help_text="控件英文名", default=None)
    type = models.CharField(choices=type_choices,
                            max_length=50,
                            help_text="控件类型",
                            default='input')
    value = models.CharField(max_length=250, help_text="控件值", default=None)
    f_instance_id = models.IntegerField(help_text="FInstanceValue ID",
                                        default=0)
    service = models.SmallIntegerField(default=1, help_text="服务编号")
    del_tag = models.SmallIntegerField(default=0, help_text="删除标识")


class FITask(models.Model):
    """fei shu instance task list"""
    name = models.CharField(max_length=150, help_text="节点名称", default=None)
    user_id = models.CharField(max_length=150, help_text="审批人", default=None)
    need_approver = models.BooleanField(default=False, help_text="是否发起人自选节点")
    status_choices = (('BEFORE', '未开始'), ('PENDING', '审批中'),
                      ('APPROVED', '同意'), ('REJECTED', '拒绝'),
                      ('TRANSFERRED', '已转交'), ('DONE', '完成'), ('REFUTED',
                                                               '驳回'))
    status = models.CharField(choices=status_choices,
                              max_length=250,
                              help_text="任务状态",
                              default=0)
    type_choices = (('AND', '会签'), ('OR', '或签'))
    type = models.CharField(choices=type_choices,
                            max_length=250,
                            help_text="审批方式",
                            default='input')
    comment = models.CharField(max_length=250, help_text="评论内容", default=None)
    start_time = models.CharField(default="",
                                  max_length=250,
                                  help_text="task 开始时间")
    end_time = models.CharField(default="",
                                max_length=250,
                                help_text="task 完成时间")
    f_node_id = models.IntegerField(help_text="FANode ID", default=0)
    f_instance_id = models.IntegerField(help_text="FInstanceValue ID",
                                        default=0)
    serial_number = models.IntegerField(default=0, help_text="task任务序列号")
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
    f_instance_id = models.IntegerField(help_text="FInstanceValue ID",
                                        default=0)
    del_tag = models.SmallIntegerField(default=0, help_text="删除标识")


class DeployHistory(models.Model):
    """deploy history"""
    env = models.CharField(max_length=150, help_text="环境", default=None)
    project = models.CharField(max_length=150, help_text="项目", default=None)
    service = models.CharField(max_length=150, help_text="服务", default=None)
    service_type = models.CharField(max_length=150, help_text="服务类型", default=None)
    BrName = models.CharField(max_length=150, help_text="版本", default=None)
    deploy_time = models.CharField(max_length=250,
                                   help_text="发布时间",
                                   default=None)
    result = models.CharField(max_length=50, help_text="发布结果", default=None)
    del_tag = models.SmallIntegerField(default=0, help_text="删除标识")


class BusinessEnvironment(models.Model):
    name = models.CharField(null=True,
                            blank=True,
                            max_length=200,
                            verbose_name="环境名")
    remark = models.CharField(null=True,
                              blank=True,
                              max_length=500,
                              verbose_name="备注")
    job_name = models.CharField(null=True,
                              blank=True,
                              max_length=500,
                              verbose_name="job name")
    del_tag = models.SmallIntegerField(default=0, help_text="删除标识")
    c_time = models.CharField(null=True,
                              blank=True,
                              max_length=100,
                              verbose_name="创建时间")
    u_time = models.CharField(null=True,
                              blank=True,
                              max_length=100,
                              verbose_name="更新入库时间")

class BusinessServices(models.Model):
    name = models.CharField(null=True,
                            blank=True,
                            max_length=200,
                            verbose_name="服务名")
    remark = models.CharField(null=True,
                              blank=True,
                              max_length=500,
                              verbose_name="备注")
    manager = models.CharField(null=True,
                               blank=True,
                               max_length=100,
                               verbose_name="服务负责人")
    rel_project = models.IntegerField(null=True,
                                      blank=True,
                                      verbose_name="关联项目")
    del_tag = models.SmallIntegerField(default=0, help_text="删除标识")
    c_time = models.CharField(null=True,
                              blank=True,
                              max_length=100,
                              verbose_name="创建时间")
    u_time = models.CharField(null=True,
                              blank=True,
                              max_length=100,
                              verbose_name="更新入库时间")
    type_choices = ((1, 'apis'), (2, 'services'), (3, 'web'))
    service_type = models.IntegerField(choices=type_choices,
                                       verbose_name="服务类型",
                                       default=1)

class DnsRecords(models.Model):  # 域名解析记录
    domain_id = models.IntegerField(null=True, blank=True, verbose_name="域名id")
    line_choices = (('default', '默认'), ('telecom', '电信'), ('unicom', '联通'),
                    ('mobile', '移动'), ('oversea', '境外'), ('edu', '教育网'),
                    ('drpeng', '鹏博士'), ('btvn', '广电网'))
    line = models.CharField(choices=line_choices,
                            max_length=20,
                            verbose_name="解析线路")
    locked = models.CharField(null=True,
                              blank=True,
                              max_length=20,
                              verbose_name="解析记录锁定状态")
    priority = models.IntegerField(null=True,
                                   blank=True,
                                   verbose_name="MX记录优先级")  # 仅MX,SRV,URI
    RR = models.CharField(null=True,
                          blank=True,
                          max_length=100,
                          verbose_name="主机记录")
    record_id = models.CharField(null=True,
                                 blank=True,
                                 max_length=20,
                                 verbose_name="解析记录ID")
    remark = models.CharField(null=True,
                              blank=True,
                              max_length=100,
                              verbose_name="备注")
    status_id = models.IntegerField(null=True, blank=True,
                                    verbose_name="状态")  # ENABLE
    ttl = models.IntegerField(null=True, blank=True, verbose_name="生存时间")
    type = models.CharField(null=True,
                            blank=True,
                            max_length=20,
                            verbose_name="记录类型")
    value = models.CharField(null=True,
                             blank=True,
                             max_length=100,
                             verbose_name="记录值")
    weight = models.CharField(null=True,
                              blank=True,
                              max_length=20,
                              verbose_name="均衡权重")
    u_time = models.CharField(null=True,
                              blank=True,
                              max_length=100,
                              verbose_name="更新入库时间")
    del_tag = models.SmallIntegerField(default=0, help_text="删除标识")
    belong_to = models.CharField(null=True,
                                 blank=True,
                                 max_length=100,
                                 verbose_name="域名解析归属")
    proxy_choices = ((True, '已代理'), (False, '仅DNS'))
    proxied = models.BooleanField(choices=proxy_choices,
                                  null=True,
                                  blank=True,
                                  verbose_name='代理状态')  # cloudflare
