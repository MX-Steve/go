#!/bin/bash
ifnames=`ifconfig | awk -F'[ :]+' '!NF{if(eth!=""&&ip=="")print eth;eth=ip4=""}/^[^ ]/{eth=$1}/inet addr:/{ip=$4}'`
# echo $ifnames 
for i in $ifnames; do
    if [ "$i" != "lo" ];then
        ifname=$i 
        break 
    fi
done
echo $ifname 
mac=`ifconfig $ifname | awk '/ether/{print $2}'`
echo $mac
sleep 1
curl -X POST http://47.108.224.218:9103/machine/${mac} > temp.log
sleep 2
port=`cat temp.log`
echo $port 
sleep 2
mport=`expr ${port} + 1`
autossh -M $mport -fCNR $port:localhost:22 root@47.108.224.218
#ssh -fCNR $port:localhost:22 root@47.108.224.218