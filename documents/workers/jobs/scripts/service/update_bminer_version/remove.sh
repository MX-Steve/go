#!/bin/bash
# sshpass -p'123456'  ssh -p 10002 -o StrictHostKeyChecking=no root@localhost "/home/cs/bminer-v16.4.6-d77cc9b/bminer -version"
ports=`awk -F',' '{print $2}' machines`

remove(){
    port=$1
    echo "$port"
    sshpass -p'123456'  timeout 10  ssh -p $port -o StrictHostKeyChecking=no root@localhost "rm -f /home/cs/bminer*.tar.xz"
    sshpass -p'123456'  timeout 10  ssh -p $port -o StrictHostKeyChecking=no root@localhost "rm -f /var/log/boot.log*"
    sshpass -p'123456'  timeout 10  ssh -p $port -o StrictHostKeyChecking=no root@localhost "rm -f /var/log/syslog*"
    sshpass -p'123456'  timeout 10  ssh -p $port -o StrictHostKeyChecking=no root@localhost "rm -f /boot/*42*"
    sshpass -p'123456'  timeout 10  ssh -p $port -o StrictHostKeyChecking=no root@localhost "service stop rsyslog"
    sshpass -p'123456'  timeout 10  ssh -p $port -o StrictHostKeyChecking=no root@localhost "journalctl --vacuum-size=10M"
    sshpass -p'123456'  timeout 10  ssh -p $port -o StrictHostKeyChecking=no root@localhost "rm -rf /var/log/journal/*"
    sshpass -p'123456'  timeout 10  ssh -p $port -o StrictHostKeyChecking=no root@localhost "systemctl restart systemd-journald.service"
}
> versions
for port in $ports
do
    version=`sshpass -p'123456'  timeout 10  ssh -p $port -o StrictHostKeyChecking=no root@localhost "/home/cs/bminer-v16.4.6-d77cc9b/bminer -version"`
    if [ $? -eq 0 ];then
        remove $port
    fi 
done
