#!/bin/sh
#

CONFIG_=/config/dc/urc01/pxeserver/dnsmasq.d/nodes.conf

function update() {
    SYMBOL=$1
    COMMAND="wget http://shepherd-master-ha./deployments/apply.sh -O -|sh"
    remote_execute $SYMBOL $COMMAND
}

function remote_execute() {
    SYMBOL=$1
    COMMAND="${@:2}"
    [ -z "$SYMBOL" ] && echo "symbol nedded!" && exit 10

    for ip in `awk -F, '/'$SYMBOL'/ {print $NF}' $CONFIG`;do
        timeout 10s ssh -f -oStrictHostKeyChecking=no -oCheckHostIP=no root@${ip} "${COMMAND}"
    done
}

for sub in $@;do
    update $sub
    sleep 60
done