** 1. 介绍 **
```
export ROOT=/tmp/workeragent
chroot $ROOT nvidia-xconfig -a --cool-bits=28 --allow-empty-initial-configuration
chroot $ROOT Xorg &
chroot $ROOT nvidia-settings -c :0 -a "[gpu]/GPUGraphicsClockOffset[1]=150"
chroot $ROOT nvidia-settings -c :0 -a "[gpu]/GPUMemoryTransferRateOffset[1]=200"
```
** 2. 规划 **
globalKey="/shepherd/DcUUID/worker/occonfig"
macKey="/shepherd/DcUUID/worker/MAC/occonfig"
** 3. 新疆 **
22604be5-05b7-40eb-afbc-fe3eac0dd77e
8c6a7fb2-e9bd-40a6-aedb-ad977983c9a5
981504c6-54ce-4ec0-ac99-85c00f6c3572

export ETCDCTL_API=3
DC_UUID="22604be5-05b7-40eb-afbc-fe3eac0dd77e"
ETCD="10.3.0.1:23790"
etcdctl --endpoints=$ETCD put /shepherd/${DC_UUID}/worker/occonfig "150/200"