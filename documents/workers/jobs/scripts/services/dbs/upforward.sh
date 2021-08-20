#!/bin/bash
# */30 * * * * /usr/local/bin/upforward
ps aux | grep  "3306:3306" | grep -v grep
if [ $? -ne 0 ];then 
nohup kubectl port-forward --address 0.0.0.0 svc/mariadb 3306:3306 &
fi
ps aux | grep  "8123:8123" | grep -v grep
if [ $? -ne 0 ];then 
nohup kubectl port-forward  --address 0.0.0.0 svc/clickhouse-miner-stats 8123:8123 -n data &
fi