1. 进入容器创建用户
m]# docker exec -it ktzauthv1 /bin/bash
c]# cd /data/back-auth/
c]# python3 manage.py createsuperuser
用户名: xxx
电子邮件地址: xxx@163.com
Password: xxx
Password (again): xxx
2. 登录数据库，修改数据
m]# mysql -uroot -p
> use ktz
> update backauth_accounts set is_superuser=0,first_name="lihan",auth_code="",del_flag=0,buids="[2]", centerids="[56,54,19]"  where username="xxx";
3. 使用新用户密码登录开拓者平台