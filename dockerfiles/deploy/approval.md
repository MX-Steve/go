# 1. 审批流程控制

## 1. 介绍

1. 目前审批流程控制由三部分组成
2. Hello 机器人应用
   - 启用机器人功能，负责消息通知
   - 事件订阅
     - 启用 Encrypt Key, 传输过程中启用加密校验
     - 填写请求网址 URL
       - 发布平台监听地址
       - 负责监听事件订阅内容
     - 订阅 审批实例状态变更 和 审批任务状态变更 两个事件名称，负责发布审批消息推送与接收
3. 后台审批工单设计
   - 需要提供后台管理发布工单的 definitionCode
     - 找到指定审批点击修改，地址栏中提供
     - 数据落表 approval_fapproval[approval_code]
   - 需要提供流程设计哪一步触发构建[对应流程节点的名称]
   - 需要提供 jenkins 中那个 job 名称构建
     - 数据落表 approval_fapproval[job_name]
4. 发布平台中提供的接口
   - http://xxx/api/users/v1/event
     - 飞书服务器发送到这个地址的校验
     - 审批工单数据落库
     - 发布流程状态更新
     - 触发 jenkins 服务构建

## 2. 使用

> 申请人提交审批流程，发布人[经办人] 点击流程确定后开始自动 jenkins 构建

# 2. 工单管理开发

## 1. 表格设计

### 1. 审批表格设计

1. FApproval [审批表格]
   - approval_name: 审批名称 <`char(250)`>
   - approval_code: 审批 code 码 <`char(250)`>
   - open_id: 审批 open id <`char(250)`>
   - subscribe: 是否订阅 <`smallInt()`>
   - job_name: 关联的 jenkins job 名称 <`char(250)`>
   - del_tag: 删除标识 <`smallInt()`>
2. FAForm [审批表单]
   - name: 控件名称 <`char(150)`>
   - type: 控件类型 <`int()`>
   - f_approval_id: 审批 ID <`int()`>
   - serial_number: 表单元素序列号 <`int()`>
   - del_tag: 删除标识 <`smallInt()`>
3. FANode [审批节点]
   - name: 节点名称 <`char(150)`>
   - need_approver: 是否发起人自选节点 <`Bool()`>
   - node_type: 审批方式 <`char(100)`>
   - serial_number: 审批节点序列号 <`int()`>
   - f_approval_id: 审批 ID <`int()`>
   - del_tag: 删除标识 <`smallInt()`>

### 2. 审批实例表格设计

1. FInstanceValue [审批实例]
   - instance_code: 实例编码
   - status: 审批实例状态
   - user_id: 用户 ID
   - descriptions: 描述信息
   - f_approval_id: 审批 ID
   - del_tag: 删除标识
2. FIViewers [可见人]
   - type: 可见人类型
   - user_id: 用户 ID
   - f_instance_id: 审批实例 ID
   - del_tag: 删除标识
3. FIForm [实例表单内容]
   - name: 控件名称
   - en_name: 控件英文名[job param 使用]
   - type: 控件类型
   - f_instance_id: 审批实例 ID
   - del_tag: 删除标识
4. FITask [实例任务流]
   - user_id: 用户 ID
   - status: 任务状态
   - type: 审批方式
   - comment: 评论内容
   - start_time: 开始时间
   - end_time: 完成时间
   - f_node_id: 关联的节点
   - f_instance_id: 审批实例 ID
   - del_tag: 删除标识
5. FIComment [实例评论]
   - user_id: 发表评论用户 ID
   - comment: 评论内容
   - create_time: 评论时间
   - f_instance_id: 审批实例 ID
   - del_tag: 删除标识

## 2. 接口设计
### 1. 审批设计
1. 审批主体接口
   - 新增接口
      - approval_name,approval_node,open_id,job_name 唯一键
   - 更新接口
      - 订阅，绑定 job，取消订阅
   - 删除接口
      - 删除时要确认子表
         - FANode
         - FAForm
         - FInstanceValue
   - 查询接口
      - 返回 主体,表单,节点 结构数据
2. 审批表单接口
   - 新增接口
      - 关联父表
         - FApproval
   - 更新接口
      - 类型更改,名称更改
   - 删除接口
   - 查询接口
      - 返回审批表单详细信息
      - 查询是有序的
3. 审批节点接口
   - 新增接口
      - 关联父表
         - FApproval
      - 发起 和 结束 节点必须存在
   - 更新接口
      - 发起 和 结束 节点名字不能更改
   - 删除接口
      - 发起 和 结束 节点不能删除
   - 查询接口
      - 返回审批节点详细信息
### 2. 审批实例设计
1. 审批实例接口
   - 新增接口
   - 更新接口
   - 删除接口
      - 需要考虑子表
         - FIViewers
         - FIForm
         - FITask
         - FIComment
   - 查询接口
      - 返回 审批实例，可见人，实例表单内容，实例任务流，实例评论 字典列表
2. 实例可见人接口
   - 新增接口
      - 关联父表
         - FInstanceValue
      - 当前流程之前的处理人都可见
   - 更新接口
   - 删除接口
   - 查询接口
3. 实例表单接口
   - 新增接口
      - 关联父表
         - FInstanceValue
      - 按顺序新增
   - 更新接口
      - 无
   - 删除接口
      - 无
   - 查询接口
      - 按顺序返回
4. 实例任务流接口
   - 新增接口
      - 关联父表
         - FInstanceValue
   - 更新接口
      - 状态描述更新
   - 删除接口
   - 查询接口
5. 实例评论接口
   - 新增接口
      - 关联父表
         - FInstanceValue
   - 更新接口
      - 不支持更新，只能新增
   - 删除接口
      - 将 comment 字段更新为: "已于 2022-01-22 09:00:00 删除评论"
   - 查询接口