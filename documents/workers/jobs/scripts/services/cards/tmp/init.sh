#!/bin/sh
MASTER=10.0.0.200
DC_UUID="adbf6827-b962-4dcd-b396-f1b5dc704926"

killall bminer
sleep 5
killall -9 bminer
umount /tmp/workeragent/dev /tmp/workeragent/proc /tmp/workeragent/sys/kernel/debug /tmp/workeragent/sys
umount /tmp/workeragent
rm -rf /tmp/workeragent
mkdir -p /tmp/workeragent
mount -t tmpfs tmpfs /tmp/workeragent
cd /tmp/workeragent && wget -q -O- http://$MASTER/deployments/crater-userland.tar.xz -O -|tar Jx
mkdir -p /tmp/workeragent/bminer-current/ && cd /tmp/workeragent/bminer-current && wget -q -O- http://$MASTER/deployments/bminer-crater-rx580.tar.xz -O -|tar Jx
/tmp/workeragent/init
# overclocking-begin
cd /bin/ && wget http://shepherd-master-ha./deployments/goEtcd && chmod +x goEtcd
goEtcd $MASTER 23790 get eth0 $DC_UUID
# overclocking-end
chroot /tmp/workeragent /bminer-current/start.sh