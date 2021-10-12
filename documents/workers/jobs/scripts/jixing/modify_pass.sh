#!/bin/bash
"""
/home/han/modify_mysql/modify_pass.sh
文档控制流程：
1. 更改mysql用户密码
1.5. 用新用户密码测试连接，如无法连接，则退出
2. 更改clickhouse表结构
2.5. 查看clickhouse是否同步了mysql，如无数据，则退出
3. 修改 secret 配置
4. 更改 configmap 配置
4.5. 查看configmap更改情况
5. 重启相关 deployment pods 
6. 查看 pods 状态
"""
# 1. 更改mysql用户密码
echo "1. 更改mysql用户密码 pool,shepherd,bminer,miningops,worker"
sleep 0.5
echo "kubectl exec -it mariadb-5ccc59b477-xzvwt -- mysql -uroot -p'wl4Y5FJMLd' < mod1.sql "
kubectl exec -it mariadb-5ccc59b477-xzvwt -- mysql -uroot -p'wl4Y5FJMLd' < mod1.sql
sleep 0.2
# 1.5. 测试连接
echo "测试连接"
kubectl exec -it mariadb-5ccc59b477-xzvwt -- mysql -ubminer -p'aNpalSEO7BsSwhEQ' -e 'select 1;'
if [ $? -ne 0 ];then
    echo "新密码无法连接"
    exit 1
fi 
# 2. 更改clickhouse表结构
echo "2. 更改clickhouse表结构"
sleep 0.5
echo "kubectl exec -it clickhouse-miner-stats-0 -n data -- clickhouse-client --host localhost --echo  -d miner_stats --multiquery < clickhouse_miner_stats.sql"
kubectl exec -it clickhouse-miner-stats-0 -n data -- clickhouse-client --host localhost --echo  -d miner_stats --multiquery < clickhouse_miner_stats.sql
echo "miner_stats modify finished"
echo "kubectl exec -it clickhouse-miner-stats-0 -n data -- clickhouse-client --host localhost --echo  -d conflux_pool --multiquery < clickhouse_conflux_pool.sql"
kubectl exec -it clickhouse-miner-stats-0 -n data -- clickhouse-client --host localhost --echo  -d conflux_pool --multiquery < clickhouse_conflux_pool.sql
echo "conflux_pool modify finished"
sleep 0.2
# 查看clickhouse是否同步了mysql
echo "查看clickhouse是否同步了mysql"
kubectl exec -it clickhouse-miner-stats-0 -n data -- clickhouse-client --host localhost --echo  -d miner_stats --multiquery -q "select name from user limit 5;";
if [ $? -ne 0 ];then
    echo "没有同步mysql数据"
    exit 1
fi 
# 3. 修改 secret 配置
echo "3. 修改 secret 配置"
sleep 0.5
for item in `ls ./secret/`
do 
    echo "kubectl apply -f secret/$item"
    kubectl apply -f secret/$item
    sleep 0.2
done 
# 4. 更改 configmap 配置
echo "4. 更改 configmap 配置"
sleep 0.5
for item in `ls ./cm`
do 
    echo "kubectl apply -f cm/$item" 
    kubectl apply -f cm/$item
    sleep 0.2
done 
# 4.5. 查看configmap更改情况
echo "查看configmap更改情况"
for item in `kubectl get cm -n data | awk '{print $1}'`; do echo $item; kubectl get cm $item -n data -o yaml|grep mariadb; done
sleep 5
for item in `kubectl get cm -n bminer | awk '{print $1}'`; do echo $item; kubectl get cm $item -n bminer -o yaml|grep mariadb; done
sleep 5
# 5. 重启相关 deployment pods 
echo "5. 重启相关 deployment pods "
sleep 0.5
echo "kubectl delete pod madminer-machine-update-devqinglin-fd44f6656-p8gx5 -n data "
kubectl delete pod madminer-machine-update-devqinglin-fd44f6656-p8gx5 -n data
sleep 0.2
echo "kubectl delete pod mining-pool-8477649f65-kfdwh -n bminer "
kubectl delete pod mining-pool-8477649f65-kfdwh -n bminer
sleep 0.2
echo "kubectl delete pod bminer-operation-staging-59687c869c-ltp7v -n bminer "
kubectl delete pod bminer-operation-staging-59687c869c-ltp7v -n bminer 
sleep 0.2
echo "kubectl delete pod bminer-operation-madminer-5c7d99ddc6-rr29w -n bminer "
kubectl delete pod bminer-operation-madminer-5c7d99ddc6-rr29w -n bminer
sleep 0.2
echo "kubectl delete pod bminer-operation-webserver-madminer-747c8999-nwrnb  -n bminer "
kubectl delete pod bminer-operation-webserver-madminer-747c8999-nwrnb  -n bminer 
sleep 0.2
echo "kubectl delete pod bminer-operation-webserver-staging-6957d59b89-7x42x -n bminer "
kubectl delete pod bminer-operation-webserver-staging-6957d59b89-7x42x -n bminer
sleep 0.2
echo "kubectl delete pod bminer-backend-prod-58994d6784-8kp56 -n bminer "
kubectl delete pod bminer-backend-prod-58994d6784-8kp56 -n bminer
sleep 0.2
echo "kubectl delete pod bminer-backend-prod-58994d6784-cp5np -n bminer "
kubectl delete pod bminer-backend-prod-58994d6784-cp5np -n bminer
sleep 0.2
# 6. 查看 pods 状态
echo "6. 查看 pods 状态"
kubectl get pods -n data | grep madminer-machine-update-devqinglin
sleep 0.2
kubectl get pods -n bminer | grep mining-pool
sleep 0.2
kubectl get pods -n bminer | grep bminer-operation-staging
sleep 0.2
kubectl get pods -n bminer | grep bminer-operation-madminer
sleep 0.2
kubectl get pods -n bminer | grep bminer-operation-webserver
sleep 0.2
kubectl get pods -n bminer | grep bminer-backend-prod
