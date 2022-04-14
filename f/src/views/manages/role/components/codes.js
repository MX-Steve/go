export default [
  {
    key: "dashboard",
    label: "Dashboard",
    pages: [
      {
        key: "dashboard",
        label: "Dashboard",
        perms: [{ key: "view", label: "查看Dashboard" }],
      },
    ],
  },
  {
    key: "assets",
    label: "资产管理",
    pages: [
      {
        key: "basic",
        label: "基础资产",
        perms: [
          { key: "view", label: "查看基础资产" },
          { key: "edit", label: "编辑基础资产" },
          { key: "add", label: "添加基础资产" },
          { key: "del", label: "删除基础资产" },
        ],
      },
      {
        key: "network-list",
        label: "网络列表",
        perms: [
          { key: "view", label: "查看网络设备" },
          { key: "add", label: "添加网络设备" },
          { key: "edit", label: "编辑网络设备" },
          { key: "del", label: "删除网络设备" },
          { key: "down", label: "下载网络设备" },
          { key: "conn", label: "域名解析操作" },
        ],
      },
      {
        key: "store-list",
        label: "存储列表",
        perms: [
          { key: "view", label: "查看存储设备" },
          { key: "add", label: "添加存储设备" },
          { key: "edit", label: "编辑存储设备" },
          { key: "del", label: "删除存储设备" },
          { key: "down", label: "下载存储设备" },
        ],
      },
      {
        key: "machine-list",
        label: "机器列表",
        perms: [
          { key: "view", label: "查看机器设备" },
          { key: "edit", label: "编辑机器设备" },
          { key: "del", label: "删除机器设备" },
          { key: "down", label: "下载机器设备" },
          { key: "conn", label: "连接机器" },
        ],
      },
      {
        key: "idc-input",
        label: "区域机房录入",
        perms: [
          { key: "view", label: "查看资产" },
          { key: "add", label: "添加区域机房信息" },
        ],
      },
      {
        key: "machine-input",
        label: "设备信息录入",
        perms: [
          { key: "view", label: "查看资产" },
          { key: "add", label: "添加设备信息" },
        ],
      },
      {
        key: "domain-record",
        label: "域名解析页",
        perms: [
          { key: "view", label: "查看解析" },
          { key: "add", label: "添加解析" },
          { key: "del", label: "删除解析" },
          { key: "edit", label: "编辑解析" },
        ],
      },
    ],
  },
  {
    key: "projects",
    label: "项目管理",
    pages: [
      {
        key: "edit",
        label: "配置中心",
        perms: [
          { key: "view", label: "展示配置中心" },
          { key: "edit", label: "编辑配置" },
          { key: "add", label: "添加配置" },
          { key: "del", label: "删除配置" },
        ],
      },
      {
        key: "tree",
        label: "业务树展示",
        perms: [
          { key: "view", label: "展示业务树" },
          { key: "edit", label: "编辑业务树" },
          { key: "add", label: "添加业务树" },
          { key: "del", label: "删除业务树" },
          { key: "conn", label: "连接机器" },
        ],
      },
      {
        key: "ssh",
        label: "ssh终端",
        perms: [{ key: "view", label: "ssh终端展示" }],
      },
    ],
  },
  {
    key: "jobs",
    label: "任务管理",
    pages: [
      {
        key: "list",
        label: "任务列表",
        perms: [
          { key: "view", label: "展示任务列表" },
          { key: "add", label: "添加任务" },
          { key: "edit", label: "编辑任务" },
          { key: "del", label: "删除任务" },
          { key: "conn", label: "运行一次权限" },
        ],
      },
    ],
  },
  {
    key: "approvals",
    label: "审批管理",
    pages: [
      {
        key: "manage",
        label: "任务创建",
        perms: [
          { key: "view", label: "展示任务信息" },
          { key: "add", label: "添加任务" },
          { key: "edit", label: "编辑任务" },
          { key: "del", label: "删除任务" },
        ],
      },
      {
        key: "list",
        label: "任务列表",
        perms: [
          { key: "view", label: "展示任务列表" },
          { key: "add", label: "添加任务" },
          { key: "edit", label: "编辑任务" },
          { key: "del", label: "删除任务" },
          { key: "conn", label: "任务详情" },
        ],
      },
      {
        key: "start",
        label: "创建工单",
        perms: [
          { key: "view", label: "展示工单信息" },
          { key: "add", label: "添加工单" },
          { key: "edit", label: "编辑工单" },
          { key: "del", label: "删除工单" },
        ],
      },
      {
        key: "fi-list",
        label: "工单列表",
        perms: [
          { key: "view", label: "展示工单列表" },
          { key: "add", label: "添加工单" },
          { key: "conn", label: "工单详情" },
          { key: "del", label: "删除工单" },
        ],
      },
      {
        key: "detail",
        label: "工单详情",
        perms: [
          { key: "view", label: "展示工单信息" },
          { key: "conn", label: "触发jenkins按钮" },
        ],
      },
      {
        key: "deploy-history",
        label: "发布历史",
        perms: [
          { key: "view", label: "展示发布历史" },
          { key: "del", label: "删除发布记录" },
          { key: "down", label: "发布记录导出" },
        ],
      },
    ],
  },
  {
    key: "audits",
    label: "审计管理",
    pages: [
      {
        key: "list",
        label: "审计日志",
        perms: [{ key: "view", label: "展示审计日志" }],
      },
    ],
  },
  {
    key: "manages",
    label: "系统管理",
    pages: [
      {
        key: "user",
        label: "用户管理",
        perms: [
          { key: "view", label: "展示用户列表" },
          { key: "status_change", label: "用户状态更新" },
          { key: "roles_change", label: "用户角色更新" },
        ],
      },
      {
        key: "role",
        label: "角色管理",
        perms: [
          { key: "view", label: "展示角色列表" },
          { key: "add", label: "新增角色" },
          { key: "edit", label: "编辑角色" },
          { key: "del", label: "删除角色" },
        ],
      },
    ],
  },
  {
    key: "personal",
    label: "个人中心",
    pages: [
      {
        key: "user_center",
        label: "个人中心",
        perms: [
          { key: "view", label: "展示个人中心" },
          { key: "config_robot", label: "机器人配置" },
        ],
      },
    ],
  },
];
