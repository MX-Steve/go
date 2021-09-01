#!/bin/bash
# sshpass -p'123456'  ssh -p 10002 -o StrictHostKeyChecking=no root@localhost "/home/cs/bminer-v16.4.6-d77cc9b/bminer -version"
ports=`awk -F',' '{print $2}' machines`
new_version="v16.4.6-d77cc9b"

update(){
    port=$1
    version=$2
    echo "$port:$version"
    echo "kill bminer..."
    sshpass -p'123456'  timeout 10  ssh -p $port -o StrictHostKeyChecking=no root@localhost "killall -9 bminer"
    sleep 2
    echo "backup now bminer..."
    sshpass -p'123456'  timeout 10  ssh -p $port -o StrictHostKeyChecking=no root@localhost "mv /home/cs/bminer-v16.4.6-d77cc9b/bminer /home/cs/bminer-v16.4.6-d77cc9b/bminer-${version}"
    sleep 2
    echo "scp new bminer to client..."
    sshpass -p'123456'  timeout 10 scp -P $port -o StrictHostKeyChecking=no /data/bminer-${new_version}/bminer root@localhost:/home/cs/bminer-v16.4.6-d77cc9b/bminer 
    sleep 2
    echo "start new bminer..."
    sshpass -p'123456'  timeout 10  ssh -p $port -o StrictHostKeyChecking=no root@localhost "/home/cs/bminer-v16.4.6-d77cc9b/bminer -managed"
    echo "show new version..."
    sshpass -p'123456'  timeout 10  ssh -p $port -o StrictHostKeyChecking=no root@localhost "/home/cs/bminer-v16.4.6-d77cc9b/bminer -version"
}
> versions
for port in $ports
do
    version=`sshpass -p'123456'  timeout 10  ssh -p $port -o StrictHostKeyChecking=no root@localhost "/home/cs/bminer-v16.4.6-d77cc9b/bminer -version"`
    if [ $? -eq 0 ];then
        version=`echo $version | awk -F"[()]" '{print $2}'`
        echo "$port,$version,ok" >> versions
        if [ $version != ${new_version} ];then
            echo "$port need update"
            update $port $version
        else
            echo "$port don't need update"
        fi
    else 
        echo "$port,,no" >> versions
        echo "$port connect refused"
    fi 
done
