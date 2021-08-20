package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/coreos/etcd/clientv3"
)

const DC_UUID = "3f154da2-4532-4435-abe6-758fc4f7fc85"
const IFname = "ens33"

func getData(global_key) {
	var (
		config  clientv3.Config
		err     error
		client  *clientv3.Client
		kv      clientv3.KV
		putResp *clientv3.PutResponse
	)
	//配置
	config = clientv3.Config{
		Endpoints:   []string{os.Args[1] + ":" + os.Args[2]},
		DialTimeout: time.Second * 5,
	}
	//连接 床见一个客户端
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}
	//用于读写etcd的键值对
	kv = clientv3.NewKV(client)
	putResp, err = kv.Put(context.TODO(), global_key, "1100/875,2000/875;1100/875,2000/875;1100/875,2000/875;1100/875,2000/875;1100/875,2000/875;1100/875,2000/875;1100/875,2000/875;1100/875,2000/875", clientv3.WithPrevKV())
	if err != nil {
		fmt.Println(err)
	} else {
		//获取版本信息
		fmt.Println("Revision:", putResp.Header.Revision)
		if putResp.PrevKv != nil {
			fmt.Println("key:", string(putResp.PrevKv.Key))
			fmt.Println("Value:", string(putResp.PrevKv.Value))
			fmt.Println("Version:", string(putResp.PrevKv.Version))
		}
	}
}

// func setData(){}

func main() {
	EtcdIp := os.Args[1]
	EtcdPort := os.Args[2]
	fmt.Printf("%s:%s\n", EtcdIp, EtcdPort)
	global_key := fmt.Sprintf("/shepherd/%s/worker/occonfig", DC_UUID)
	fmt.Println("global_key")
	fmt.Println(global_key)
	MAC := ""
	// 获取本机 mac
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, inter := range interfaces {
		if inter.Name == IFname {
			MAC = fmt.Sprintf("%s", inter.HardwareAddr)
		}
	}
	mac_key := fmt.Sprintf("/shepherd/%s/worker/%s/occonfig", DC_UUID, MAC)
	fmt.Println("mac_key")
	fmt.Println(mac_key)
	getData(global_key)
}
