# 1. 发布历史记录模块

## 1. 介绍

1. 发布历史记录模块记录了飞书审批工单，服务发布的版本，jenkins 构建记录，服务发布结果等信息
2. 方便进行回滚操作
3. 方便查询当前服务版本信息

## 2. 表单设计

1. DHistory
   - approval_name: 审批名称 <`char(150)`>
   - instance_code: 审批实例 CODE 码 <`char(250)`>
   - env_name: 环境名称 <`char(150)`>
   - project_name: 项目名 <`char(150)`>
   - service_type: 服务类型 <`smallInt()`>
   - service_name: 服务名 <`char(150)`>
   - branch_name: 分支名 <`char(150)`>
   - version: 版本号 <`char(150)`>
   - publisher: 发布人 <`char(150)`>
   - job_name: jenkins job 名称 <`char(150)`>
   - build_num: job 编号 <`int()`>
   - res_code: 发布结果 code 码 <`int()`>
   - c_time: 发布时间 <`char(150)`>
   - u_time: 更新时间 <`char(150)`>
   - del_tag: 删除标识 <`smallInt()`>

## 3. 接口设计

1. 记录新增功能
   - 提供更新函数
     - 平台内部状态更新
   - 提供更新 url 接口
     - 供 jenkins 中 pipeline 脚本使用
2. 查询功能
   - 服务基线功能
   - 服务发布详情功能 [版本，日志等]
   - 服务历史版本查询功能
