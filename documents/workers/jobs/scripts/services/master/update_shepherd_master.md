  # **注意点**
    1. 通过 AWS 平台 EC2 ipsec_tunnel-gateway 的防火墙安全策略可知：
        - 对 22 端口原地址没有做限制
        - 迁移场地访问 ipsec 理论上没有问题，如果访问不同，需要 麦爷 授权
    2. 新地址 eno2 网卡口
    3. 让矿场给我们分配 8 个网段
  # **网络配置变更和配置网络规则**
    1. 配置变更
        1. /etc/netplan/xxx.yaml
            0. 10.40.23.250/21 -> 新地址
            1. 10.40.23.250 -> 新地址
            2. 10.40.23.254 -> 新路由
            3. Vlan30/40 已注释
    2. 应用网络配置并配置网络规则
        1. netplan apply [root/或者使用 svc 但要知道 svc 的密码]
        2. iptables-restore < /config/dc/urc01/master/iptables-rules.v4-shepherd-master01 [root/或者使用 svc 但要知道 svc 的密码]
  # **容器配置文件更新和启动**
    1. 容器配置文件更新
        1. docker-compose
            0. 出入口端口不做限制
                至少：9092/8086/2379/23790/...
            1. pxeserver,shepherd,syslog 不用动
            2. tunnels
                10.40.23.250:8086:100.71.38.14:8086 ->
                新地址:8086:100.71.38.14:8086 
            3. etcd-grpc-proxy 
                --listen-addr=10.40.23.250:23790 ->
                --listen-addr=新地址:23790
        2. shepherd
            0. sed -i "s,10.40.23.250,新地址,g" /home/svc/config/dc/urc01/shepherd/shepherd.json
            1. 对应的子网根据场地需求划分
        3. master 无须更改   
            0. /home/svc/config/dc/urc01/master/iptables-rules.v4-shepherd-master01
        4. pxeserver
            0. nginx.d/pxe.conf 无须动
            1. dnsmasq.d/nodes.conf 
                0. ip 地址更换，是否需要清空    
            2. dnsmasq.d/pxe.conf
                0. listen-address=新地址 
                1. dhcp-range=起点,终点,子网,12h
                2. dhcp-option=option:router,路由地址
                3. dhcp-option=option:dns-server,新地址
                4. dhcp-range=tag:RPI,起点,终点,子网,12h
                5. dhcp-option=tag:RPI,option:router,路由地址
                6. dhcp-boot=tag:!ipxe,tag:HTTP,http://新地址/boot/ipxe.efi
                7. dhcp-boot=tag:ipxe,http://新地址/api/v1/pxe
                8. address=/shepherd-master-ha/新地址
                9. address=/shepherd-master01/新地址
    2. 使用 docker-compose 启动相应服务
        1. docker-compose up -d -f /config/dc/urc01/docker-compose/docker-compose.yaml
  # **周期性计划任务更新设置**
    1. master设置 [svc] 
        1. */3 * * * * /home/svc/crontab/master-monitor -dc urumqi -endpoint http://本机地址:8086 
    2. k8s集群验证
        1. kubectl exec -it influxdb-66c8874495-w66k5 /bin/bash
        2. influx
        3. > use monitor
        4. > select count(*) from (select mean(status) from monitor where time > now() - 10m and dc='urumqi' group by ops)
  # **检查测试**
    1. 容器日志查看 
        1. docker logs Cid
    2. 找一个客户端，启动装机测试 PXE
        1. client 是否可以正常登录
        2. 验证 etcdctl
            1. cd /tmp/workeragent/bin/
            2. wget http://shepherd-master-ha./deployments/etcdctl
            3. chmod +x etcdctl
            4. export ETCDCTL_API=3
            5. chroot /tmp/workeragent etcdctl --endpoints=10.40.23.250:23790 get /a
        3. bminer 是否正常提供服务
            1. tail -100 /tmp/workeragent/bminer.log
        4. 监控平台是否可以正常获取监控数据
        5. bminer 平台是否可以正常获取到 client 数据
  # **其他服务**
    1. socat -v TCP-LISTEN:7000,fork,reuseaddr TCP:eth-asia1.nanopool.org:9999
    2. socat -v TCP-LISTEN:13333,bind=10.0.0.200,fork,reuseaddr TCP:cn.sparkpool.com:13333
  # **破解服务器密码**
    1. 重启 linux 系统，出现 GRUB 启动菜单
    2. 按 e 键进入编辑状态
    3. 按向下的方向健，划到 linux16 所在行，把光标停在行末尾
    4. 在步骤 3 的行末尾，添加 rd.break console=tty0 
    5. 按下 Ctrl + x 键进入恢复模式[单用户模式]
    6. 以可写方式挂载硬盘中的根目录，并重设 root 密码
        1. mount -o remount,rw /sysroot
        2. chroot /sysroot
        3. passwd root
        4. touch /.autorelabel
        5. exit
        6. reboot
    7. 完成以上步骤则可以进入 linux 系统，（已经更改的密码）


1. 注意点
    1. ipsec_tunnel-gateway 这台机器做端口转发，看防火墙策略没有限制，
        如果 master 到时候无法转发获取，则需要 麦爷 授权登录
    2. master 机器牵扯到网络配置 netplan, 这个服务需要改配置再激活策略，
        需要 root 用户或使用 svc 用户提权操作
        你那边是否知道 root 或 svc 账号密码
        如不知，则需要破解密码使用 root 登录操作