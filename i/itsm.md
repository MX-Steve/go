# 1. ITSM设计流程
- 需要安装插件名称: Markdown Preview Mermaid Support
## 0. 进度状态
1. 进行中
    - style B fill:#bbf,stroke:#f66,stroke-width:2px,color:#fff,stroke-dasharray: 5 5
2. 已完成
    - style A fill:#090,stroke:099,stroke-width:2px
## 1. 使用流程
```mermaid
flowchart TD
START[开始] --> A
A(创建工单审批流) -->|顺序| B(创建表单项目)
B --> C(创建审批节点)
C --> D(绑定job并激活)
D -->|审批流创建完成| E[发布人员创建工单]
E --> F(填写表单)
F --> G(选择审批人员)
G -->|点击完成则审批开始| H[发布人员同意发布]
H --> I(触发jenkins开始构建)
I --> J{pipeline流程\n构建是否正常}
J -->|是| END[流程结束]
J -->|否| M(机器人通知)
M -.发送失败日志\n发布人员检查原因.-> H
H -->|代码原因,流程打回| E
M -.发送失败日志.-> E
style A fill:#bbf,stroke:#f66,stroke-width:2px,color:#fff,stroke-dasharray: 5 5
```
## 2. 创建工单审批流开发
### 1. 基础信息
1. req
```json
{
    "approval_name": "test",
    "approval_code": "xxx",
    "job_name": "dev",
    "descriptions": "describe ...",
}
```
2. 表单设计
3. 流程设计
4. 激活启用
