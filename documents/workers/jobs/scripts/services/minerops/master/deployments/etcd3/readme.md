## 1. DC_UUID 设置全局配置
master]#  ./etcdctl --endpoints=10.0.0.200:23790 put /shepherd/adbf6827-b962-4dcd-b396-f1b5dc704926/worker/occonfig "1100/875,2000/900;1100/875,2000/900;1100/875,2000/900;1100/875,2000/900;1100/875,2000/900;1100/875,2000/900;1100/875,2000/900;1100/875,2000/900" 
## 2. DC_UUID 查看全局配置
master]#  ./etcdctl --endpoints=10.0.0.200:23790 get /shepherd/adbf6827-b962-4dcd-b396-f1b5dc704926/worker/occonfig
## 3. 查看某台机器配置
master]#  ./etcdctl --endpoints=10.0.0.200:23790 get /shepherd/adbf6827-b962-4dcd-b396-f1b5dc704926/worker/00:e2:69:13:96:6d/occonfig
## 4. 测试
localhost]# scp -P 2020  -i /root/.ssh/id_ed25519_altchain goEtcd  altchainops@220.77.208.215:/home/altchainops
master]# sudo -su svc
master]# cp goEtcd /home/svc/services/pxeserver/data/httproot/deployments/goEtcd
master]# exit
master]# ssh root@10.0.1.210
client]# rm -f /tmp/workeragent/bin/goEtcd
client]# cd /tmp/workeragent/bin/
client]# wget http://shepherd-master-ha./deployments/goEtcd
client]# chmod +x goEtcd
client]# chroot /tmp/workeragent goEtcd 10.0.0.200 23790 get eth0 adbf6827-b962-4dcd-b396-f1b5dc704926
## 5. 修改 crater-deploy.sh [韩国]
```shell
# overclocking-begin
DC_UUID="adbf6827-b962-4dcd-b396-f1b5dc704926"
cd /bin/ && wget http://shepherd-master-ha./deployments/goEtcd && chmod +x goEtcd
goEtcd $MASTER 23790 get eth0 $DC_UUID
# overclocking-end
```
## 6. master 上进行指定机器指定卡超频
master (/home/svc/config/dc/icn01/pxeserver/dnsmasq.d)# ./goEtcd   10.0.0.200 23790 put 10.0.1.210 adbf6827-b962-4dcd-b396-f1b5dc704926 0:1100/890,2000/900