systemctl disable snapd
service snapd stop
rm -rf /var/cache/apt/*
rm -rf /var/lib/snapd/*
rm -rf /var/log/boot.log*
service rsyslog stop
systemctl disable rsyslog
journalctl --vacuum-size=10M
rm -rf /var/log/journal/*
rm -rf /home/cs/bminer-v16.4.6-d77cc9b-amd64.tar.xz 
rm -rf /home/cs/NVIDIA-Linux-x86_64-465.31.run
rm -rf /var/cache/apt/*
# 00:e0:4c:36:12:09 10x16x2x205
# 00:e0:4c:36:11:a1 10.16.3.136 dhispeth1.rack310x16x3x136
# 00:e0:4c:36:0d:fa 10.16.2.219 dhispeth1.rack210x16x2x219