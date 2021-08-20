# 1. 数据库与表的操作
```
#创建数据库
create database "db_name"
 
#显示所有的数据库
show databases
 
#删除数据库
drop database "db_name"
 
#使用数据库
use db_name
 
#显示该数据库中所有的表
show measurements
 
#创建表，直接在插入数据的时候指定表名
insert test,host=127.0.0.1,monitor_name=test count=1
 
#删除表
drop measurement "measurement_name"
```
# 2. 增
```
> use metrics
Using database metrics
> insert test,host=127.0.0.1,monitor_name=test count=1
```
# 3. 查
```
> use metrics
Using database metrics
> select * from test order by time desc
```
# 4. 其他查询语句
```
SHOW FIELD KEYS --查看当前数据库所有表的字段
SHOW series from pay --查看key数据
SHOW TAG KEYS FROM "pay" --查看key中tag key值
SHOW TAG VALUES FROM "pay" WITH KEY = "merId" --查看key中tag 指定key值对应的值
SHOW TAG VALUES FROM cpu WITH KEY IN ("region", "host") WHERE service = 'redis'
DROP SERIES FROM <measurement_name[,measurement_name]> WHERE <tag_key>='<tag_value>' --删除key
SHOW CONTINUOUS QUERIES   --查看连续执行命令
SHOW QUERIES  --查看最后执行命令
KILL QUERY <qid> --结束命令
SHOW RETENTION POLICIES ON mydb  --查看保留数据
# 查询数据
SELECT * FROM /.*/ LIMIT 1  --查询当前数据库下所有表的第一行记录
select * from pay  order by time desc limit 2
select * from  db_name."POLICIES name".measurement_name --指定查询数据库下数据保留中的表数据 POLICIES name数据保留
# 删除数据
delete from "query" --删除表所有数据，则表就不存在了
drop MEASUREMENT "query"   --删除表（注意会把数据保留删除使用delete不会）
DELETE FROM cpu
DELETE FROM cpu WHERE time < '2000-01-01T00:00:00Z'
DELETE WHERE time < '2000-01-01T00:00:00Z'
DROP DATABASE “testDB” --删除数据库
DROP RETENTION POLICY "dbbak" ON mydb --删除保留数据为dbbak数据
DROP SERIES from pay where tag_key='' --删除key中的tag

SHOW SHARDS  --查看数据存储文件
DROP SHARD 1
SHOW SHARD GROUPS
SHOW SUBSCRIPTIONS
```
# 5. 数据保存策略
```
show retention policies on "db_name"
create retention policy "rp_name" on "db_name" duration 3w replication 1 default
    # rp_name：策略名
    # db_name：具体的数据库名
    # 3w：保存3周，3周之前的数据将被删除，influxdb具有各种事件参数，比如：h（小时），d（天），w（星期）
    # replication 1：副本个数，一般为1就可以了
    # default：设置为默认策略
# 修改Retention Policies
alter retention policy "rp_name" on "db_name" duration 30d default
# 删除Retention Policies
drop retention policy "rp_name" on "db_name"     
```
# 6. 用户管理
```
#显示用户
show users
 
#创建用户
create user "username" with password 'password'
 
#创建管理员权限用户
create user "username" with password 'password' with all privileges
 
#删除用户
drop user "username"
```
# 7. 连续查询
1. 当数据超过保存策略里指定的时间之后就会被删除，但是这时候可能并不想数据被完全删掉，influxdb提供了连续查询，可以做数据统计采样。
2. 查看数据库的 Continous Queries
```
show continuous queries
```
3. 创建新的 Continous Queries
```
create continous query cq_name on db_name begin select sum(count) into new_table_name from table_name group by time(30m) end
# cq_name：连续查询名字
# db_name：数据库名字
# sum(count)：计算总和
# table_name：当前表名
# new_table_name：存新的数据的表名
# 30m：时间间隔为30分钟
```
4. 删除 Continous Queries
```
drop continous query cq_name on db_name
```
