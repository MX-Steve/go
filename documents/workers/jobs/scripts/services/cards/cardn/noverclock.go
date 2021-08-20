package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
	"os/exec"
	"net"
	"flag"
	"time"
	"context"
	"regexp"
	"github.com/coreos/etcd/clientv3"
	log "github.com/sirupsen/logrus"
)

// 设置 N 卡超频
func setOverclocking(n1 string,n2 string) {
	log.Infof(`export ROOT=/tmp/workeragent
	&& chroot $ROOT nvidia-xconfig -a --cool-bits=28 --allow-empty-initial-configuration
	&& chroot $ROOT Xorg &
	&& chroot $ROOT nvidia-settings -c :0 -a "[gpu]/GPUGraphicsClockOffset[1]=`+n1+`"
	&& chroot $ROOT nvidia-settings -c :0 -a "[gpu]/GPUMemoryTransferRateOffset[1]=`+n2+`"`)
	exec.Command("/bin/sh", "-c", `export ROOT=/tmp/workeragent
		&& chroot $ROOT nvidia-xconfig -a --cool-bits=28 --allow-empty-initial-configuration
		&& chroot $ROOT Xorg &
		&& chroot $ROOT nvidia-settings -c :0 -a "[gpu]/GPUGraphicsClockOffset[1]=`+n1+`"
		&& chroot $ROOT nvidia-settings -c :0 -a "[gpu]/GPUMemoryTransferRateOffset[1]=`+n2+`"`).Output()
}

func getMacByEth(ethName string) string {
	mac := ""
	interfaces, err1 := net.Interfaces()
	if err1 != nil {
		log.Warn(err1)
		return mac 
	}
	for _, inter := range interfaces {
		if inter.Name == ethName {
			mac = fmt.Sprintf("%s", inter.HardwareAddr)
		}
	}
	if len(mac) != 17 {
		log.Warn("mac not found")
	}
	return mac
}

func getMacByHostname(hostname string) string {
	counts := make(map[string]int)
	file := "nodes.conf"
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(os.Stderr,err)
		return ""
	}
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	f.Close()
	mac := ""
	for line := range counts {
		if strings.Contains(line, hostname) {
			mac = fmt.Sprintf("%s", line[10:27])
		}
	}
	return mac
}

func dataFormat(data string) bool {
	x := "/"
	result := false
	xORb := regexp.MustCompile(x)
	xmatches := xORb.FindAllStringIndex(data, -1)
	if len(xmatches) == 1 {
		result = true
	}else{
		log.Warn("Key's format error!")
	}
	return result
}

func getDataFromEtcd(etcd string,DcUUID string,eth string) (string,int) {
	result := ""
	errCode := 404
	var (
		config clientv3.Config
		err    error
		client *clientv3.Client
		kv     clientv3.KV
		getResp *clientv3.GetResponse
	)
	config = clientv3.Config{
		Endpoints:   []string{etcd},
		DialTimeout: time.Second * 5,
	}
	if client, err = clientv3.New(config); err != nil {
		log.Warn(err)
	}else{
		globalKey := fmt.Sprintf("/shepherd/%s/worker/occonfig", DcUUID)
		kv = clientv3.NewKV(client)
		getResp, err = kv.Get(context.TODO(), globalKey)
		if err != nil {
			log.Fatal(err)
		}else {
			if len(getResp.Kvs) == 1 {
				data := string(getResp.Kvs[0].Value)
				if dataFormat(data){
					errCode = 0 
					mac := getMacByEth(eth)
					macKey := fmt.Sprintf("/shepherd/%s/worker/%s/occonfig",DcUUID,mac)
					kv = clientv3.NewKV(client)
					getResp, err = kv.Get(context.TODO(), macKey)
					if err != nil {
						log.Fatal(err)
					}else {
						if len(getResp.Kvs) == 1 {
							data2 := string(getResp.Kvs[0].Value)
							if !dataFormat(data2){
								log.Warn("Mac key '",macKey,"' format error,use Global key")
								result = data
							}else{
								result = data2
							}
						}else{
							log.Infof("Mac key '",macKey,"' has no value,use Global key")
						}
					}
				}else {
					log.Warn("Global key '",globalKey,"' format error")
				}
			}else {
				log.Warn("Global Key '",globalKey,"' has no value!")
			}
		}
	}
	return result,errCode
}

func setDataToEtcd(etcd string,DcUUID string,mac string,data string) int{
	errCode := 404
	var (
		config clientv3.Config
		err    error
		client *clientv3.Client
		kv     clientv3.KV
	)
	config = clientv3.Config{
		Endpoints:   []string{etcd},
		DialTimeout: time.Second * 5,
	}
	//连接 创建一个客户端
	if client, err = clientv3.New(config); err != nil {
		log.Warn(err)
	}else{
		kv = clientv3.NewKV(client)
		macKey := fmt.Sprintf("/shepherd/%s/worker/%s/occonfig", DcUUID,mac)
		_, err = kv.Put(context.TODO(), macKey, data, clientv3.WithPrevKV())
		if err != nil {
			log.Warn(err)
		}else{
			errCode=0
		}
	}
	return errCode
}

func main(){
	var etcdIp string
    var etcdPort string
    var method string
    var eth string
	var DcUUID string
	var hostname string
	var macValue string
	flag.StringVar(&etcdIp, "etcdIp", "127.0.0.1", "etcd ip addr")
	flag.StringVar(&etcdPort, "etcdPort", "2379", "etcd port")
	flag.StringVar(&method, "method", "get", "get | put")
	flag.StringVar(&eth, "eth", "ens33", "eth name")
	flag.StringVar(&DcUUID, "DcUUID", "adbf6827-b962-4dcd-b396-f1b5dc704926", "DC UUID")
	flag.StringVar(&hostname, "hostname", "r0614m5", "hostname | IP")
	flag.StringVar(&macValue, "macValue", "0/0", "macValue")
	flag.Parse()
	if method == "get" {
		data,code:=getDataFromEtcd(etcdIp+":"+etcdPort,DcUUID,eth)
		if code == 0{
			datas:=strings.Split(data,"/")
			setOverclocking(datas[0],datas[1])
		}else {
			log.Warn("set error")
		}
	}else if method == "put"{
		log.Infof("Client: configure mac")
		mac := getMacByHostname(hostname)
		code := setDataToEtcd(etcdIp+":"+etcdPort,DcUUID,mac,macValue)
		if code == 0 {
			log.Infof("Set mac key/value successful")
		}else{
			log.Warn("Set mac key/value error")
		}

	}
}