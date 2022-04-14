# ITSM Operation System
> 工单系统，帮助开发运维人员更快捷发布服务
# 本地环境开发

```
$ mkdir build
$ python3 -m venv build
$ pip3 install -r requirements.txt
$ export DJANGO_SETTINGS_MODULE=djentry.settings.development
$ export DJANGO_SECRET_KEY='qe187p7jy6w%hrxdly6t=rq%2izaq@owxtefa=sc(gpnialsa!'
$ export ENABLE='False'
$ python3 manage.py runserver
# http://121.43.41.139:9115/api/users/v1/event
```
# v1 版本使用
## 1. 组成
> v1 版本工单审批不牵扯到前端展示，通过飞书审批创建实例，到达指定流程会触发 jenkins 构建，jenkins pipeline 反馈构建结果
## 2. 工单与job绑定与取消绑定
```
/approval/v1/subscribe
{
    "approval_name": "飞书审批名称",
    "approval_code": "飞书审批 code 码",
    "job_name": "jenkins job 名称",
    "subscribe": 1 # 1 绑定，0 取消绑定
}
```
1. 绑定后数据落库
## 3. 绑定完成后飞书发布审批工单
## 4. 审批到达经办人后触发 jenkins 构建
# v2 版本使用
## 1. 新增发布记录接口
```shell
curl --location --request POST 'http://ep.devops.epverse.net/api/approval/v2/deploy-history' \
--header 'Content-Type: application/json' \
--data-raw '{
    "env": "xx",
    "project":"xx",
    "service": "xx",
    "BrName": "xx",
    "result": "SUCCESS" # SUCCESS|FAILURE
}'
```
## 2. 发布结果飞书通知接口
```shell
curl --location --request POST 'http://ep.devops.epverse.net/api/approval/v1/send-result-to-group2' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "xx",
    "job_name": "xx",
    "number": 0,
    "result": "xx",
    "env": "xx",
    "project": "xx",
    "service": "xx"
}'
```