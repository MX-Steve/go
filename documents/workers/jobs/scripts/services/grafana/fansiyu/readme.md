# 1. 镜像下载
```
]# docker pull mariadb:10.4
]# docker pull influxdb:1.6.6
]# docker pull grafana/grafana:7.1.1
]# docker pull 313557088125.dkr.ecr.us-west-2.amazonaws.com/alt-chain/telemetry-deploy:temp-aws
```
# 2. 运行容器
```
]# docker run -itd --name=grafana -p 3000:3000 grafana/grafana:7.1.1
]# docker run -itd --name influxdb -p 8086:8086 -v /data/influxdb:/data/influxdb influxdb:1.6.6
]# vim /root/app/telemetry/config/yul_dc.toml 
[[rules.aws-hashrate]]
dialect = "influxdb:farm"
query = '''SELECT sum("hashrate") FROM ( SELECT mean("hashrate") as hashrate FROM "algorithm" WHERE dc='aws' AND time >= now() - 15m AND algorithm='ethash' GROUP BY hostname,pci_bus_id ORDER BY time LIMIT 15)'''
warning_threshold = 15000000.0
critical_threshold = 12000000.0
service = "urumqi"
check_interval = 120
counter_threshold = 0.0
dc = "aws"
disable_auto_resolve = true

#[[rules.aws-livenode]]
#dialect = "influxdb:farm"
#query = '''select count(*) from (select mean(hashrate) from algorithm where time > now() - 10m and dc='aws' group by hostname)'''
#warning_threshold = 3.0
#critical_threshold = 2.0
#service = "urumqi"
#check_interval = 60
#counter_threshold = 0.0
#dc = "aws"
#disable_auto_resolve = true

[services]

[[services.urumqi]]
description = "DC operation in Urumqi"
route = "pagerduty"
pagerduty_key = "bc4d0ced80c34cac82f4611d5727bee4"

[alerters]

[datasources]

[[datasources.influxdb]]

[[datasources.influxdb.farm]]
endpoint="http://influx:8086"
db="farm"
]# docker run -itd -v /data/telemetry/config:/config --name telemetry --link influxdb:influx 313557088125.dkr.ecr.us-west-2.amazonaws.com/alt-chain/telemetry-deploy:temp-aws
```
# 3. 本地安装 influxdb 客户端调试
```
]# wget https://dl.influxdata.com/influxdb/releases/influxdb-1.8.0.x86_64.rpm
]# wget https://dl.influxdata.com/influxdb/releases/influxdb-1.8.0.x86_64.rpm
]# influx -version
]# influx -host 52.43.177.248  -port 8086
Connected to http://52.43.177.248:8086 version 1.6.6
InfluxDB shell version: 1.8.0
> create database test;
> show databases;
name: databases
name
----
_internal
test
```

# 4. Grafana 配置
1. 浏览器访问 http://52.43.177.248:3000/ admin/admin

# 5. workeragent
]# vim  worker.config
{
    "metrics":{
        "inputs":{
            "bminer":{
                "device":"http://127.0.0.1:1880/api/v1/status/device",
                "solver":"http://127.0.0.1:1880/api/v1/status/solver"
            }
        },
        "outputs":{
            "influxdb":{
                "url":"http://52.43.177.248:8086", 
                "database":"farm"
            }
        }
    }
}
# 6. dashboard
```
SELECT sum("hashrate") FROM ( SELECT mean("hashrate") as hashrate FROM "algorithm" WHERE dc='aws' AND $timeFilter AND algorithm='ethash' GROUP BY hostname,pci_bus_id) GROUP BY time(5m) fill(null)
```

52.43.177.248: 
    1. 除了现有的 xmr / bminer 服务启动，添加 workeragent 启动
    2. test_new.py 文件修改
    3. 新增 workeragent 配置
    4. 明天移除 influxdb/telemetry等镜像容器