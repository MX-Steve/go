#!/usr/bin/env python3
import subprocess
import datetime
import time


def getPercentOfDf(name, top):
    status, output = subprocess.getstatusoutput(
        "/usr/local/bin/kubectl exec -i %s -n data  -- df -h | awk '/clickhouse/{print $5}'" % (name))
    if status == 0:
        print(output)
        df_percent = int(output.replace('%', ''))
        if df_percent >= top:
            now = datetime.datetime.now()
            warning_msg = "%s Warning: df of %s : %d; over use default %d" % (now.strftime('%Y-%m-%d %H:%M:%S'), name, df_percent, top)
            print(warning_msg)
            with open("check_clickhouse_df.log","a") as fobj:
                fobj.write(warning_msg+"\n")
    else:
        print("can not excute kubectl, please check it")

while True:
    getPercentOfDf('clickhouse-miner-stats-0', 90)
    time.sleep(60)