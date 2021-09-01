1. 开机重新挂在 /cdrom
mount -o remount,rw /cdrom
2. 将 /home/cs/bminer... 移动到 /cdrom ，并启动
mv /home/cs/bminer /cdrom
cd /cdrom/bminer