#!/usr/bin/env bash
function check_df() {
    output=`kubectl exec -i clickhouse-miner-stats-0 -n data  -- df -h | awk '/clickhouse/{print $5}'`
    result=${output%%%}
    if [ $result -gt 96 ];then
      date >> check_df.log
      echo "Over user df, now is $result" >> check_df.log
      /usr/bin/python3  pagerduty_incident_trigger.py
      echo "========================================" >> check_df.log
    fi
}
while [ 1 == 1 ]; do
    check_df
    sleep 600
done

