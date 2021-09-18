#!/bin/bash

echo "2021-09"
total=0
for d in `seq 14`
do 
  # echo "2021-09-$d"
  dayCount=`clickhouse-client --database=miner_stats --query="select time,hex(site_key),coin,estimation_reward from site_estimation_reward where site_key=unhex('007ACB68F8A4FFDFC7741C413E4BFC94FF4D7D85B1') and coin='ETH' and time > '2021-09-$d 00:00:00' and time < '2021-09-`expr $d+1` 00:00:00'" | wc -l`
  dayTotal=`clickhouse-client --database=miner_stats --query="select sum(estimation_reward) from site_estimation_reward where site_key=unhex('007ACB68F8A4FFDFC7741C413E4BFC94FF4D7D85B1') and coin='ETH' and time > '2021-09-$d 00:00:00' and time < '2021-09-`expr $d+1` 00:00:00'"`
  dayAvg=`awk 'BEGIN{printf "%0.6f",'${dayTotal}'/'${dayCount}'}'`
  # echo $dayAvg
  total=`expr ${dayAvg}+${total}`
done 
echo $total