# 1. docker-compose 应用
1. syslog
    /app/syslog -dc urumqi -conf /config/nodes.conf
    client:
        /bin/syslogd -R shepherd-master-ha.:514
            Log to master:514[udp]
    client log -> syslog -> tunnels[9092] -> kafka[kafka/100.70.81.93] 
    -> clickhouse -> grafana 
2. tunnels
    autossh -M 0 -o ExitOnForwardFailure=yes -o ServerAliveCountMax=3 -o ServerAliveInterval=10 -NT -L 127.0.0.1:9092:100.70.81.93:9092 -L 10.40.23.250:8086:100.71.38.14:8086 -L 127.0.0.1:2379:100.69.106.78:2379 svc@edgegw01-a1.alt-chain.io    
3. shepherd
    /app/shepherd -conf /config/shepherd.json
    127.0.0.1:8080
4. pxeserver
    dnsmasq -kC /config/dnsmasq.d/pxe.conf
    -> :80/api -> 127.0.0.1:8080
    client [启动网卡内置小程序，发送dhcp请求] -> master dnsmasq: dhcp 服务给出相应
    -> dnsmasq
    -> 配置子网/路由/网关信息
    -> /boot/ipxe.efi[执行ipxe小程序]
    -> /api/v1/pxe
        -> handlePXE
            -> 通过请求头 X-Forwarded-For 获取客户端地址
            -> findNetbootEntryBySubnet 通过客户端地址 获取当前子网所在 shepherd.conf 中的配置信息 
            -> PXE: 找到当前配置中 PXE 信息
                -> 当前条数据是shell 脚本
                    -> 下载 vmlinuz 镜像和 initrd-zrt 驱动
                    -> boot 引导启动
    [怎么进入下一步的？]
    -> /api/v1/get-config
        -> handleConfig
            -> 通过请求头 X-Forwarded-For 获取客户端地址
            -> findNetbootEntryBySubnet 通过客户端地址 获取当前子网所在 shepherd.conf 中的配置
            -> 从当前配置中获取 AgentConfig,AgentConfig是一个json对象
                -> 结构：
                    {
                        supervisor:
                            images:
                                uri: userland
                                mountpoint: /
                        metrics:
                            input:
                                bminer:
                                    device:
                                    solver:
                            output:
                                influxdb:
                                    url:
                                    db:
                    }
        [哪里驱动这解压下面两个东西并执行init？]
        -> 获取 A 卡或 N 卡的 userland 包并解压到 /
        -> 获取 bminer 程序包，解压到 /tmp/workeragent/bminer-current,并执行 init 脚本
    -> init 脚本
        -> /api/v1/get-initial-config: bminer 服务启动前获取相关环境变量
            -> handleInitialConfig: 通过请求头 X-Forwarded-For 获取客户端地址
                -> getInitialEnv: 通过客户端地址获取env内容
                    -> findUUIDBySubnet: 通过子网获取UUID
                        -> 通过UUID获取dcConfigKey，通过readValue，获取dcConfigKey值[etcd]
                        -> 通过getMacAddr获取MAC地址，通过readValue，获取workerEnv值[etcd]
                    -> 将dcConfigKey和workerEnv整合，返回给客户端/tmp/workeragent/worker.env
        -> 启动 超频脚本
        -> 启动 workeragent
        -> 启动 bminer -managed -> 去 bminer-backend 获取屏蔽卡信息 -> 启动 bminer
5. etcd-grpc-proxy
    etcd grpc-proxy start --endpoints=127.0.0.1:2379 --listen-addr=10.40.23.250:23790
    client etcd -> etcd-grpc-proxy(10.40.23.250:23790 -> 127.0.0.1:2379) 
    -> tunnels[2379] -> etcd[shepherd/100.69.106.78]