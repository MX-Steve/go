export ROOT=/tmp/workeragent
chroot $ROOT nvidia-xconfig -a --cool-bits=28 --allow-empty-initial-configuration
chroot $ROOT Xorg &
chroot $ROOT nvidia-settings -c :0 -a "[gpu]/GPUGraphicsClockOffset[1]=150"
chroot $ROOT nvidia-settings -c :0 -a "[gpu]/GPUMemoryTransferRateOffset[1]=200"