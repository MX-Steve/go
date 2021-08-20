1. 安装 npm
cd /data/apps
wget https://npm.taobao.org/mirrors/node/v15.8.0/node-v15.8.0-linux-x64.tar.gz
tar -zxvf node-v15.8.0-linux-x64.tar.gz 
ln -s /data/apps/node-v15.8.0-linux-x64/bin/npm /usr/local/bin/npm
ln -s /data/apps/node-v15.8.0-linux-x64/bin/node /usr/local/bin/node
echo "ln -s /data/apps/node-v15.8.0-linux-x64/bin/npm /usr/local/bin/npm" >> /etc/profile
echo "ln -s /data/apps/node-v15.8.0-linux-x64/bin/node /usr/local/bin/node" >> /etc/profile
2. npm install
cd /data/apps/vue-project
rm -rf node_modules/*
rm -f package-lock.json
npm config set registry https://registry.npm.taobao.org 
npm install  --unsafe-perm
3. 测试运行
npm run dev
4. 构建包
npm run build:prod 
5. 打包上传到镜像仓库
zip -r dist.zip dist/