# Set up shepherd master
## 1. 本文适用于：所有数据都准备完全。 数据包括但不限于以下内容：
  - config
  - deployment
  - OS: Ubuntu 18.04
  - IMPORTANT: 在安装之前，先声明变量
    - export DC=yul02
    - export SHEPHERD_MASTER=10.0.0.1 
## 2. 安装基础服务
1. 安装 docker   
```
sudo apt-get update
sudo apt-get -yq install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common

curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

sudo add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"

sudo apt-get update
sudo apt-get -yq install docker-ce docker-ce-cli containerd.io docker-compose
```
2. 安装 etcd-client
```
sudo apt-get -yq install etcd-client
```
  - IMPORTANT: etcd-client 用于手动对 etcd 进行配置,在使用前需要声明
  - export ETCDCTL_API=3
## 3. 网络配置
1. 配置文件: /etc/netplan/xxx.yaml
```
network:
  version: 2
  renderer: networkd
  ethernets:
    eno2:
      addresses: [ 74.121.162.14/26 ]  # update 
      nameservers:
        search: [ yul01.farm.alt-chain.io ]  # update
        addresses: [ 8.8.8.8, 8.8.4.4 ]
      routes:
        - to: 0.0.0.0/0
          via : 74.121.162.1  # update
          on-link: true
    eno1:
      dhcp4: no
      dhcp6: no

  vlans:
    vlan30:
      dhcp4: no
      dhcp6: no
      id: 30
      link: eno1
      addresses:
        - 10.0.0.1/16
      routes:
        - to: 10.0.0.0/16
          via: 0.0.0.0
          on-link: true
    vlan40:
      dhcp4: no
      dhcp6: no
      id: 40
      link: eno1
      addresses:
        - 10.10.1.100/24
      routes:
        - to: 10.10.1.0/24
          via: 0.0.0.0
          on-link: true
```
2. 应用配置
```
sudo netplan apply
```
3. 配置网络规则
```
sudo iptables-restore < /config/dc/$DC/master/iptables-rules.v4-shepherd-master01
```

