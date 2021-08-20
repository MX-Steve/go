#/usr/bin/env python3
# -*- coding -*- : utf-8

import etcd

client = etcd.Client(host='192.168.226.130', port=2379)

client.write('/a', 'kitty')
ret = client.read('/a').value
print(ret)

# client.write('/jsontest2', '{"Network": "10.244.0.0/16","Backend": {"Type": "vxla111n" }}')
# print("/jsontest2")
# ret = client.read('/jsontest2').value
# print(ret)