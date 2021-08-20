#!/bin/bash
# 1. supervisor 配置更改
Project=`hostname`
ProjectDir=/data/${Project}
sed -i "s,demo,${Project}," /etc/supervisord.conf
# 2. 清理日志和不必要数据
find ${ProjectDir} -name "__pycache__" -exec rm -rf {} \;
rm -rf ${ProjectDir}/logs/*
# 3. 修改项目目录下文件与配置
# mv /data/django_demo ${ProjectDir}
mv ${ProjectDir}/demo ${ProjectDir}/${Project}
mv ${ProjectDir}/DemoIn ${ProjectDir}/${Project}In
mv ${ProjectDir}/demo.ini ${ProjectDir}/${Project}.ini
sed -i 's,demo,'${Project}',g' `grep -rl demo ${ProjectDir}`
sed -i 's,Demo,'${Project}',g' `grep -rl Demo ${ProjectDir}`
# 4. 启动 supervisor 服务并查看状态
supervisord -c /etc/supervisord.conf
supervisorctl status

# 启动：docker run -itd --name ktz  --hostname ktz  -v /data/go/src/github.com/MX-Steve/jobs/scripts/services/demo:/data/ktz  -p 9092:9092 --privileged=true  mycentosbypy:v7