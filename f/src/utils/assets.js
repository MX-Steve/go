const kvs = {
  optional: "可选项",
  u_time: "更新时间",
  sn_id: "设备sn编号",
  username: "用户名",
  authentication_type: "认证方式",
  port: "端口号",
  password: "密码",
  sudo_password: "提权密码",
  cabinet_id: "机柜编号",
  instance_name: "实例名称",
  charge_type: "付费类型",
  create_time: "创建时间",
  expiration_time: "到期时间",
  machine_id: "机器id",
  zone: "区域名称",
  provider: "运营商",
  manager: "联系人",
  phone: "联系电话",
  address: "机房地址",
  ip_range: "ip地址段",
  bandwidth: "机房带宽",
  idc: "机房",
  description: "描述",
  count: "数量",
  email: "注册邮箱",
  registration_date: "注册日期",
  domain_owner: "域名持有者",
  domain_id: "域名ID",
  data_redundancy_type: "数据容灾类型",
  gmt_created: "创建时间",
  gmt_modified: "修改时间",
  source: "源",
  resource_groupId: "资源组ID",
  vpc_id: "所挂VpcID",
  tags: "标签列表",
  secondary_cidr_blocks: "第二网段",
  creation_time: "创建时间",
  vswitch_ids: "交换机ID列表",
  vrouter_id: "路由器ID",
  router_table_ids: "路由表ID列表",
  resource_group_id: "资源组ID",
  ssl_protocol: "https开关",
  create_time: "创建时间",
  pay_type: "付款类型",
  network_type: "私网负载均衡实例网络类型",
  internet_charge_type: "公网类型实例付费方式",
  ipv6_cidr_block: "IPV6网段",
  network_acl_id: "acl规则ID",
  is_default: "是否是默认交换机",
  os_name: "操作系统",
  instance_id: "实例ID",
  hostname: "主机名",
  family: "family",
  public_ip: "公网IP",
  mac_address: "MAC地址",
  cpu: "CPU",
  memory: "内存",
  net_charge_type: "net_charge_type",
  specs: "附加选项",
};
const server_types = [
  { id: 1, name: "物理机" },
  { id: 2, name: "云服务器" },
  { id: 3, name: "虚拟机" },
];
const os_types = [
  { id: "1", name: "Linux" },
  { id: "2", name: "Windows" },
];
const authentication_types = [
  { id: 1, name: "password" },
  { id: 2, name: "file" },
];
const machine_rules = {
  hostname: [{ required: true, message: "请输入主机名", trigger: "blur" }],
  instance_name: [{ required: true, message: "请输入主机名", trigger: "blur" }],
  username: [{ required: true, message: "请输入用户名", trigger: "blur" }],
  ip_address: [{ required: true, message: "请输入ip地址", trigger: "blur" }],
  password: [{ required: true, message: "请输入密码", trigger: "blur" }],
  sudo_password: [
    { required: true, message: "请输入提取密码", trigger: "blur" },
  ],
  select_require: [
    { required: true, message: "请至少选择一个", trigger: "change" },
  ],
};
const registrant_type_choices = [
  { id: 1, name: "个人" },
  { id: 2, name: "企业" },
];
const domain_status_choices = [
  { id: 1, name: "急需付费" },
  { id: 2, name: "急需赎回" },
  { id: 3, name: "正常" },
];
export function get_item_name(id, list) {
  var name = "";
  for (var i = 0; i < list.length; i++) {
    var item = list[i];
    if (item.id === id) {
      name = item.name;
      break;
    }
  }
  return name;
}
export function get_date_time(index) {
  var date = new Date();
  var newDate = new Date();
  newDate.setDate(date.getDate() + index);
  var year = newDate.getFullYear();
  var month = newDate.getMonth() + 1;
  month = month < 10 ? "0" + month : month;
  var day = newDate.getDate();
  day = day < 10 ? "0" + day : day;
  var hour = newDate.getHours();
  hour = hour < 10 ? "0" + hour : hour;
  var minute = newDate.getMinutes();
  minute = minute < 10 ? "0" + minute : minute;
  var second = newDate.getSeconds();
  second = second < 10 ? "0" + second : second;
  return (
    year + "-" + month + "-" + day + " " + hour + ":" + minute + ":" + second
  );
}
export {
  kvs,
  server_types,
  os_types,
  authentication_types,
  machine_rules,
  registrant_type_choices,
  domain_status_choices,
};
