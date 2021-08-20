CentOS/8.3.2011
1. 启动容器
docker run -itd --name=mycentos mycentos:v1
docker cp /data/apps/back-auth mycentos:/data/back-auth
docker cp /data/apps/Python-3.6.1.tgz mycentos:/data/Python-3.6.1.tgz
docker exec -it mycentos /bin/bash
2. 容器内操作
mkdir /data
cd /data
tar -xf Python-3.6.1.tgz 
cd Python-3.6.1
yum -y install gcc gcc-c++ make zlib* openssl openssl-devel openldap-devel python36-devel mysql-devel mariadb vim
./configure --prefix=/usr/local/python/ --enable-shared --enable-loadable-sqlite-extensions 
make && make install
echo "export PATH=$PATH:/usr/local/python/bin/" >> /etc/profile
source /etc/profile
python3 -m ensurepip
pip3 install -r requirements.txt 
3. 使用新镜像启动
docker run -itd --name=ktzauthv1 -p 9092:9092 ktzauth:v1