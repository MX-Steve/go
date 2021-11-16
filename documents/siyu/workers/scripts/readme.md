## 实施步骤
0. 制作AMI  （如果增加region, AMI里面的代码不需要做改动）  
a. 进入/root/app/scripts路径  
b. 上传miner.py代码和bminer.conf/xmr.conf(更新ami时修改哪一个更新哪一个)  
c. 确认路径下有以下6个文件  
```sh
bminer.conf  xmr.conf miner.py start.sh workeragent.conf  workeragent.py
```
zx负责: bminer.conf  xmr.conf miner.py start.sh    
zy团队负责: workeragent.conf  workeragent.py  

1. 检查  
两个账号，4个region，分别检查  
是否有VPC与subnet，subnet应该包含当前region的所有AZ  (每个AZ必须有一个单独的subnet)
IAM role名为miner（中国账户需要再创建一个，手工创建），应有EC2和S3的权限（参照已有IAM role权限可以创建新的）
2. 创建sg  
每个region创建一个名为miner的sg，ssh只对监控跳板机公网IP开放登录（范思宇购买一台跳板机，t2.micro），root密码登录，所有spot实例使用相同root密码。监控服务器及网页登录必须严格管理Access。

3. 创建launch template（autoscaling中）
命名为miner，实例类型选择g4dn.xlarge
在详情里面需要指定应该使用哪一个AMI，如果更新为新的AMI，需要在这个更改为新的。
(参考录屏)
EC2页面下左侧，Auto Scaling Group，里面配置了上述Launch template，选择Spot allocation strategy - Lowest price - diversified across the 20 lowest priced pools，另外还需要选择AZ，Subnet ID.

1. 创建launch template（autoscaling中）
命名为miner，实例类型选择g4dn.xlarge，在元数据中运行python代码
在详情里面需要指定应该使用哪一个AMI，如果更新为新的AMI，需要在这个更改为新的。



### Security Group
所有region中sg命名为miner  
设置了入站流量限制，只能从3.142.141.226进行跳转  

### VPC和subnet
就使用默认的，足够使用

### 实例类型
默认都使用g4dn.xlarge

### AMI
1. 源AMI  
Name: Deep Learning AMI (Ubuntu 18.04) Version 36.0  
id: ami-063585f0e06d22308  


### 应用下载和启动
1. Bminer
https://www.bminer.me/zh/releases/
https://www.bminercontent.com/releases/bminer-v16.4.6-d77cc9b-amd64.tar.xz
```sh
wget https://www.bminercontent.com/releases/bminer-v16.4.6-d77cc9b-amd64.tar.xz
tar xvf bminer-v16.4.6-d77cc9b-amd64.tar.xz
cd bminer-v16.4.6-d77cc9b/
/root/app/bminer-v16.4.6-d77cc9b/bminer -uri ethproxy://0xF7332BA2dA27c495b1Fb0cb56A0D6d22a99C6644.%s@us1.ethermine.org:4444 -api 0.0.0.0:1880
```
2. XMR
https://github.com/xmrig/xmrig/releases/
https://github.com/xmrig/xmrig/releases/download/v6.12.1/xmrig-6.12.1-linux-x64.tar.gz
```sh
wget https://github.com/xmrig/xmrig/releases/download/v6.12.1/xmrig-6.12.1-linux-x64.tar.gz
tar xvf xmrig-6.12.1-linux-x64.tar.gz
cd xmrig-6.12.1-linux-x64/
/root/app/xmrig-6.12.1/xmrig  -o pool.supportxmr.com:3333 -u 49stwWL5s4jH6tdJeixg59bYNA2pACbfXdXC5zFFRrGs94E6EuJeHXQiXSmSzP2jioNdd5yU7yyeHHG132JW3BNJ93TAr9A.%s -t 4 --http-port 1889
```

### 账号密码
1. 测试服务器（俄勒冈）  
ip：52.43.177.248  
用户名：root  
密码： UBUNTUminer  

### 通知
EC2新建的时候需要邮件通知:fsy@beavotech.com  

### DataBase Info（俄亥俄）  
operation: https://blog.csdn.net/weixx3/article/details/80782479  
IP:3.142.141.226  
username: root  
passwd: Q1W2E3R4T5#$%qwe  

### Monitor
ip: http://52.83.203.39:3000  
username:admin  
password:Biwo2021!  

### Auto Start start.sh
修改/etc/rc.local为以下内容
```sh
#!/bin/sh
bash /root/app/scripts/start.sh
```
