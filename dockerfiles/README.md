# 1. 先制作镜像
```shell
]# docker build -t centos:python37 .
```
# 2. 根据镜像启动容器
```shell
# 需要先启动 mysql:8.0
]# docker run -itd --name mysql8 --hostname mysql8 -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 --restart always -v C:\data\mysql8\data:/var/lib/mysql mysql:8.0
]# docker run -itd --name cmdb_backend --hostname cmdb_backend -p 9092:8000 -v C:\data\projects\backend\Projects_dj:/app/cmdb_backend --restart always --link mysql8 centos:python37
```
# 3. 环境变量
```shell
export DJANGO_SETTINGS_MODULE=djentry.settings.production
export DJANGO_SECRET_KEY='qe187p7jy6w%hrxdly6t=rq%2izaq@owxtefa=sc(gpnialsa!'
export CLOVER_MYSQL_DB='test'
export CLOVER_MYSQL_USER='root'
export CLOVER_MYSQL_PASSWORD='YHhg9-5.*'
export CLOVER_MYSQL_HOST='mysql8'
export REDIS_PORT=16379
export REDIS_PASS="EcoPlants"
```

# 3. 测试环境部署
```
# 1. 安装配置 redis , node , docker[mysql8]
$ apt install redis
$ vim /etc/redis/redis.conf
requirepass EcoPlants
port 16379
systemctl restart redis
$ cd /data/apps
$ wget https://nodejs.org/dist/v16.13.1/node-v16.13.1-linux-x64.tar.xz
$ tar -xf node-v16.13.1-linux-x64.tar.xz
$ ln -s /data/apps/node-v16.13.1-linux-x64/bin/node /usr/local/bin/
$ ln -s /data/apps/node-v16.13.1-linux-x64/bin/npm /usr/local/bin/
$ node -v && npm -v
$ apt install docker
$ docker run -itd --name mysql8 --hostname mysql8 -p 3306:3306 -e MYSQL_ROOT_PASSWORD=YHhg9-5.* -v /data/mysql8/data:/var/lib/mysql --restart always mysql:8.0
# 2. 安装配置 nginx
$ apt install nginx
$ vim /etc/nginx/conf.d/cmdb.conf
server {
  listen 9115;
  server_name localhost;
  location / {
    root /data/projects/cmdb_frontend/dist;
    index index.html index.htm;
    try_files  $uri $uri/ /index.html;
  }
  location ^~ /api/ {
    proxy_set_header Host $proxy_host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_pass http://127.0.0.1:8081/;
  }
}
$ nginx -t 
$ nginx -s reload
# 3. 安装配置 supervisor
$ apt install supervisor
$ vim /etc/supervisor/supervisord.conf
[include]
files = conf.d/*.ini
$ vim /etc/supervisor/conf.d/cmdb_backend.ini
[program:cmdb_backend]
autorestart=True
autostart=True 
redirect_stderr=True
command=/data/projects/cmdb_backend/build/python/bin/uvicorn  djentry.asgi:application --host 0.0.0.0 --port 8081
user=root
stdout_logfile_maxbytes = 20MB
stdout_logfile_backups = 20
stdout_logfile = /tmp/cmdb_backend.log
$ vim /etc/supervisor/conf.d/cmdb_beat.ini
[program:cmdb_beat]
autorestart=True
autostart=True 
redirect_stderr=True
command=/data/projects/cmdb_backend/build/python/bin/celery -A djentry beat 
user=root
stdout_logfile_maxbytes = 20MB
stdout_logfile_backups = 20
stdout_logfile = /tmp/cmdb_beat.log
$ vim /etc/supervisor/conf.d/cmdb_worker.ini
[program:cmdb_worker]
autorestart=True
autostart=True 
redirect_stderr=True
command=/data/projects/cmdb_backend/build/python/bin/celery -A djentry worker -l info
user=root
stdout_logfile_maxbytes = 20MB
stdout_logfile_backups = 20
stdout_logfile = /tmp/cmdb_worker.log
$ systemctl restart supervisor
$ supervisorctl status
# 4. 配置 profile 环境变量
$ vim /etc/profile
export DJANGO_SETTINGS_MODULE=djentry.settings.production
export DJANGO_SECRET_KEY='qe187p7jy6w%hrxdly6t=rq%2izaq@owxtefa=sc(gpnialsa!'
export CLOVER_MYSQL_DB='test'
export CLOVER_MYSQL_USER='root'
export CLOVER_MYSQL_PASSWORD='YHhg9-5.*'
export CLOVER_MYSQL_HOST='mysql8'
export REDIS_PORT=16379
export REDIS_PASS="EcoPlants"
export CMDB_BACK_PORT=8081
$ source /etc/profile
# 5. 拉取并构建后端代码
$ git clone git@code.corp.ecoplants.tech:cmdb_operation/cmdb_backend.git
$ ROOT_DIR=/data/projects/cmdb_backend
$ mkdir -p $ROOT_DIR/build/python
$ apt install python3.8-venv
$ python3 -m venv $ROOT_DIR/build/python
$ source $ROOT_DIR/build/python/bin/activate
$ pip install -r requirements.txt
$ mkdir -p ${ROOT_DIR}/backend/proto
$ python setup.py gen_grpc -o ${ROOT_DIR}/backend/proto
$ supervisorctl restart cmdb_backend
# 6. 拉取并配置前端代码
$ git clone git@code.corp.ecoplants.tech:cmdb_operation/cmdb_frontend.git
$ npm install
$ npm run build:prod
# 7. 访问测试
$ 浏览器 http://121.43.41.139:9115
```
# 4. 接入 ldap
## 1. 安装依赖
```sh
sudo apt-get install libsasl2-dev python-dev libldap2-dev libssl-dev
sudo apt-get install -y python-dev libldap2-dev libsasl2-dev libssl-dev
sudo yum install python-devel
sudo yum install openldap-devel
pip install python-ldap
pip install django-auth-ldap
```
# 5. webhook 机器人
## 1. 机器人地址
```
https://open.feishu.cn/open-apis/bot/v2/hook/e0df4f59-e05e-4343-a5de-5305c58a1213
```
## 2. Verification TOken
```
7eAUaMHhiMi0QViW2zX7igqvTd6Pa3AM
```
# 6. HelloWorld
https://open.feishu.cn/open-apis/bot/v2/hook/57210d66-1c63-49a4-9be4-72db49f21780
# 7. sql 导入
mysql -h mysql8 -uroot -p123456 test < do.sql --default-character-set=utf8