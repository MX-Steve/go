**1. DC_UUID 设置全局配置**

```
master]# export ETCDCTL_API=3
master]# DC_UUID="adbf6827-b962-4dcd-b396-f1b5dc704926"
master]# ETCD="10.0.0.200:23790"
master]#  ./etcdctl --endpoints=$ETCD put /shepherd/${DC_UUID}/worker/occonfig "1100/875,2000/900;1100/875,2000/900;1100/875,2000/900;1100/875,2000/900;1100/875,2000/900;1100/875,2000/900;1100/875,2000/900;1100/875,2000/900" 
```
**2. DC_UUID 查看全局配置**

```
master]#  ./etcdctl --endpoints=$ETCD get /shepherd/${DC_UUID}/worker/occonfig
```
**3. 查看某台机器配置**

```
master]# MAC="00:e2:69:13:96:6d"
master]#  ./etcdctl --endpoints=$ETCD get /shepherd/${DC_UUID}/worker/$MAC/occonfig
```
**4. 测试**

```
localhost]# scp -P 2020  -i /root/.ssh/id_ed25519_altchain overclock  altchainops@220.77.208.215:/home/altchainops
master]# sudo -su svc
master]# cp overclock /home/svc/services/pxeserver/data/httproot/deployments/overclock
master]# cp overclock /home/svc/config/dc/yul02/pxeserver/dnsmasq.d/overclock
master]# exit
master]# ssh root@10.0.1.210
client]# rm -f /tmp/workeragent/bin/overclock
client]# cd /tmp/workeragent/bin/
client]# wget http://shepherd-master-ha./deployments/overclock
client]# chmod +x overclock
client]# chroot /tmp/workeragent overclock 10.0.0.200 23790 get eth0 adbf6827-b962-4dcd-b396-f1b5dc704926
```
**5. 修改 crater-deploy.sh [韩国]**

```
# overclocking-begin
DC_UUID="adbf6827-b962-4dcd-b396-f1b5dc704926"
cd /bin/ && wget http://shepherd-master-ha./deployments/overclock && chmod +x overclock
overclock $MASTER 23790 get eth0 $DC_UUID
# overclocking-end
```
**6. master 上进行指定机器指定卡超频**

```
master (/home/svc/config/dc/icn01/pxeserver/dnsmasq.d)# ./overclock   10.0.0.200 23790 put 10.0.1.210 adbf6827-b962-4dcd-b396-f1b5dc704926 0:1100/890,2000/900
```
**7. 加拿大 A 卡 vega 修改 init,并重新打包**

```
master] vim init
while [ $i -lt $TOTAL_GPU ]; do
    base_dir=/sys/class/drm/card${i}/device
    if [ -f "/tmp/pp_table" ]; then
	    cat /tmp/pp_table > ${base_dir}/pp_table
        #[ -f "${base_dir}/power_dpm_force_performance_level" ] && echo "manual" > ${base_dir}/power_dpm_force_performance_level
        #[ -f "${base_dir}/pp_od_clk_voltage" ] && echo "m 3 ${CLOCK} ${VOLTAGE}" > ${base_dir}/pp_od_clk_voltage
        #echo "c" >  ${base_dir}/pp_od_clk_voltage
        #echo 7 > ${base_dir}/pp_dpm_sclk
        #echo 3 > ${base_dir}/pp_dpm_mclk
    fi
    HWMON_DIR=$(ls -1d /sys/class/drm/card$i/device/hwmon/hwmon*)
    if [ "x${POWER_BUDGET}" != "x" ]; then 
        [ -f "$HWMON_DIR/power1_cap" ] && echo ${POWER_BUDGET} > $HWMON_DIR/power1_cap
    fi
    i=$((i+1))
done
wget -q -O $ROOT/bin/overclock http://shepherd-master-ha./deployments/overclock
chmod +x $ROOT/bin/overclock
chroot $ROOT /bin/overclock 10.3.0.1 23790 get eth0 f02e0814-05b8-4ffd-9cce-b1c8c5346024
master] tar Jcvf bminer-crater.tar.xz  init  libcratercl.so  start.sh bminer
```
**8. 各 master 对应 UUID**
| master | UUID | 网段 | 国家| masterIP | 卡名 |
| ---- | ---- | ---- | ---- | ---- | ---- |
| shepherd-yul02 | f02e0814-05b8-4ffd-9cce-b1c8c5346024 | "10.3.107.0/24", "10.3.108.0/24","10.3.116.0/22","10.3.120.0/21" | 加拿大 | 38.122.102.122 | A卡 Vega |
| shepherd-yul02 | 3f154da2-4532-4435-abe6-758fc4f7fc85 | "10.3.109.0/24","10.3.96.0/20","10.3.112.0/21" | 加拿大 | 38.122.102.122 | N 卡 |
| shepherd-master01 | 7086bef2-71db-46b3-8c16-85fed38dae5e | "10.0.0.0/16" | 韩国 | 59.0.87.158 | A卡580 |
| shepherd-master01 | adbf6827-b962-4dcd-b396-f1b5dc704926 | "10.0.0.0/16" | 韩国 | 220.77.208.215 | A卡580 |
