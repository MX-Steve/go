#!/bin/sh
#
ROOT=/tmp/workeragent

ENV_FILE=$ROOT/worker.env
wget -q -O $ENV_FILE http://shepherd-master-ha./api/v1/get-initial-config
if [ "$?" -ne 0 ];then
    echo "get initial config error" > /tmp/error.log
    exit 10
fi
source $ENV_FILE

for h in $(ls /sys/class/hwmon);do
    for i in 1 2 3; do
        echo $FANSPEED > /sys/class/hwmon/$h/pwm$i
    done
done

TOTAL_GPU=$(chroot $ROOT /usr/bin/nvidia-smi -L|grep "^GPU "|wc -l)
i=0
while [ $i -lt $TOTAL_GPU ]; do
    chroot $ROOT /usr/bin/nvidia-smi -i $i -pl $POWER_BUDGET
    i=$((i+1))
done

killall bminer
sleep 5
killall -9 bminer
cd $ROOT

exec chroot /tmp/workeragent sh /bminer-current/start.sh