package main

import (
	"context"
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"regexp"
	"strings"
	"time"
	"strconv"
	"github.com/coreos/etcd/clientv3"
	"os/exec"
)

var cardsDatas = map[string]map[string]string{}

const LogFile = "/tmp/overclocking.log"

func IfDataOk(str string) bool {
	f := ";" // 7
	d := "," // 8
	x := "/" // 16
	fORb := regexp.MustCompile(f)
	fmatches := fORb.FindAllStringIndex(str, -1)
	dORb := regexp.MustCompile(d)
	dmatches := dORb.FindAllStringIndex(str, -1)
	xORb := regexp.MustCompile(x)
	xmatches := xORb.FindAllStringIndex(str, -1)
	if len(fmatches) != 7 || len(dmatches) != 8 || len(xmatches) != 16 {
		return false
	} else {
		return true
	}
}

func getDataFormat(str string) {
	cards := strings.Split(str, ";")
	// var cardDatas []string
	for i, card := range cards {
		m1 := make(map[string]string)
		sms := strings.Split(card, ",")
		m1["s"] = strings.Replace(sms[0], "/", " ", 1)
		m1["m"] = strings.Replace(sms[1], "/", " ", 1)
		c := fmt.Sprintf("card%d", i)
		cardsDatas[c] = m1
	}
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

/**
* 往文件中写入操作
 */

func writeToFile(file string, msg string) {
	var d1 = []byte(msg)
	ioutil.WriteFile(file, d1, 0666)
}

func setOverclocking() {
	for i, card := range cardsDatas {
		fmt.Println(i, " setting")
		file := fmt.Sprintf("/sys/class/drm/%s/device/power_dpm_force_performance_level", i)
		result := checkFileIsExist(file)
		if result != true {
			fmt.Printf("%s is not exists，skip\n", i)
			continue
		}
		fmt.Printf("echo manual > /sys/class/drm/%s/device/power_dpm_force_performance_level\n", i)
		writeToFile(file,"manual")
		for j, item := range card {
			if j == "s" {
				fmt.Printf("echo s 0 %s > /sys/class/drm/%s/device/pp_od_clk_voltage\n", item, i)
				exec.Command("/bin/sh", "-c", `echo 's 0 '`+item+` > /sys/class/drm/`+i+`/device/pp_od_clk_voltage`).Output()
				fmt.Printf("echo s 1 %s > /sys/class/drm/%s/device/pp_od_clk_voltage\n", item, i)
				exec.Command("/bin/sh", "-c", `echo 's 1 '`+item+` > /sys/class/drm/`+i+`/device/pp_od_clk_voltage`).Output()
				fmt.Printf("echo s 2 %s > /sys/class/drm/%s/device/pp_od_clk_voltage\n", item, i)
				exec.Command("/bin/sh", "-c", `echo 's 2 '`+item+` > /sys/class/drm/`+i+`/device/pp_od_clk_voltage`).Output()
				fmt.Printf("echo s 3 %s > /sys/class/drm/%s/device/pp_od_clk_voltage\n", item, i)
				exec.Command("/bin/sh", "-c", `echo 's 3 '`+item+` > /sys/class/drm/`+i+`/device/pp_od_clk_voltage`).Output()
				fmt.Printf("echo s 4 %s > /sys/class/drm/%s/device/pp_od_clk_voltage\n", item, i)
				exec.Command("/bin/sh", "-c", `echo 's 4 '`+item+` > /sys/class/drm/`+i+`/device/pp_od_clk_voltage`).Output()
				fmt.Printf("echo s 5 %s > /sys/class/drm/%s/device/pp_od_clk_voltage\n", item, i)
				exec.Command("/bin/sh", "-c", `echo 's 5 '`+item+` > /sys/class/drm/`+i+`/device/pp_od_clk_voltage`).Output()
				fmt.Printf("echo s 6 %s > /sys/class/drm/%s/device/pp_od_clk_voltage\n", item, i)
				exec.Command("/bin/sh", "-c", `echo 's 6 '`+item+` > /sys/class/drm/`+i+`/device/pp_od_clk_voltage`).Output()
				fmt.Printf("echo s 7 %s > /sys/class/drm/%s/device/pp_od_clk_voltage\n", item, i)
				exec.Command("/bin/sh", "-c", `echo 's 7 '`+item+` > /sys/class/drm/`+i+`/device/pp_od_clk_voltage`).Output()
			} else if j == "m" {
				fmt.Printf("echo m 2 %s > /sys/class/drm/%s/device/pp_od_clk_voltage\n", item, i)
				exec.Command("/bin/sh", "-c", `echo 'm 2 '`+item+` > /sys/class/drm/`+i+`/device/pp_od_clk_voltage`).Output()
			}
		}
		fmt.Printf("echo c > /sys/class/drm/%s/device/pp_od_clk_voltage\n", i)
		fmt.Printf("echo 3 > /sys/class/drm/%s/device/pp_dpm_sclk\n", i)
		fmt.Printf("echo 2 > /sys/class/drm/%s/device/pp_dpm_mclk\n", i)
		exec.Command("/bin/sh", "-c", `echo 'c' > /sys/class/drm/`+i+`/device/pp_od_clk_voltage`).Output()
		exec.Command("/bin/sh", "-c", `echo '3' > /sys/class/drm/`+i+`/device/pp_dpm_sclk`).Output()
		exec.Command("/bin/sh", "-c", `echo '3' > /sys/class/drm/`+i+`/device/pp_dpm_mclk`).Output()
	}
}

func getMacByEth(ethName string) string{
	MAC := ""
	// 获取本机 mac
	interfaces, err1 := net.Interfaces()
	if err1 != nil {
		fmt.Println(err1)
		os.Exit(1)
	}
	for _, inter := range interfaces {
		if inter.Name == ethName {
			MAC = fmt.Sprintf("%s",inter.HardwareAddr)
		}
	}
	if len(MAC) != 17 {
		fmt.Println("MAC not found")
		os.Exit(1)
	}
	return MAC
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
			counts[input.Text()]++
	}
}
func getMacByHostname(hostname string) string{
	counts := make(map[string]int)
	file := "nodes.conf"

	f, err := os.Open(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
		return ""
	}
	countLines(f, counts)
	f.Close()
	mac := ""
	for line, _ := range counts {
		if strings.Contains(line,hostname){
			mac = fmt.Sprintf("%s",line[10:27])
		}
	}
	return mac
}

func main() {
	if len(os.Args) < 6 {
		fmt.Println("Get Data: ")
		fmt.Println("  goEtcd EtcdIp EtcdPort get IFname DC_UUID")
		fmt.Println("PUT Data: ")
		fmt.Println("  goEtcd EtcdIp EtcdPort put hostname DC_UUID 0:1000/800,2000/900")
		fmt.Println("  cardnum:s_first_value/s_second_value,m_first_value/m_second_value")
		return
	}
	EtcdIp := os.Args[1]
	EtcdPort := os.Args[2]
	dataType := os.Args[3]
	if !(dataType == "get" || dataType == "put") {
		fmt.Println("arg3 must be : get | put")
		os.Exit(1)
	}
	var (
		config  clientv3.Config
		err     error
		client  *clientv3.Client
		kv      clientv3.KV
		// putResp *clientv3.PutResponse
		getResp *clientv3.GetResponse
	)
	//配置
	config = clientv3.Config{
		Endpoints:   []string{EtcdIp + ":" + EtcdPort},
		DialTimeout: time.Second * 5,
	}
	//连接 创建一个客户端
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}
	if dataType == "get" {
		argLens := len(os.Args)
		if argLens != 6 {
			fmt.Println("goEtcd EtcdIp EtcdPort get IFname DC_UUID")
			os.Exit(1)
		}
		IFname := os.Args[4]
		DC_UUID := os.Args[5]
		global_key := fmt.Sprintf("/shepherd/%s/worker/occonfig", DC_UUID)
		MAC := getMacByEth(IFname)
		mac_key := fmt.Sprintf("/shepherd/%s/worker/%s/occonfig", DC_UUID, MAC)
		kv = clientv3.NewKV(client)
		// 判断局部配置
		getResp, err = kv.Get(context.TODO(), mac_key)
		if err != nil {
			fmt.Println(err)
			return
		}
		data := ""
		for _, ev := range getResp.Kvs {
			data = string(ev.Value)
		}
		result := IfDataOk(data)
		if result {
			fmt.Println("read local configuration")
			getDataFormat(data)
			setOverclocking()
		} else {
			fmt.Println("read global configuration")
			// 判断全局配置
			getResp, err = kv.Get(context.TODO(), global_key)
			if err != nil {
				fmt.Println(err)
				return
			}
			data = ""
			for _, ev := range getResp.Kvs {
				data = string(ev.Value)
			}
			result2 := IfDataOk(data)
			if result2 {
				getDataFormat(data)
				setOverclocking()
			} else {
				fmt.Println("the Stored data format error")
			}
		}
	}
	// 写
	if dataType == "put" {
		if len(os.Args) != 7 {
			fmt.Println("goEtcd EtcdIp EtcdPort put hostname DC_UUID 0:1000/800,2000/900")
			fmt.Println("cardnum:s_first_value/s_second_value,m_first_value/m_second_value")
			return
		}
		hostname := os.Args[4]
		DC_UUID := os.Args[5]
		kv = clientv3.NewKV(client)
		// 1. 获取全局配置，如果没有或者格式错误，则退出
		global_key := fmt.Sprintf("/shepherd/%s/worker/occonfig", DC_UUID)
		getResp, err = kv.Get(context.TODO(), global_key)
		if err != nil {
			fmt.Println(err)
			return
		}
		data := ""
		for _, ev := range getResp.Kvs {
			data = string(ev.Value)
		}
		result := IfDataOk(data)
		if ! result{
			fmt.Println("This DC_UUID hasn't configured values or the Stored data format error: ",global_key)
			return
		}
		// 如果全局配置有且格式没有问题，则获取局部配置
		// 如果局部配置没有，则拉取全局配置作为局部配置，然后修改局部配置
		// 如果局部配置存在，则直接修改局部配置
		MAC := getMacByHostname(hostname)
		if MAC == ""{
			return
		}
		mac_key := fmt.Sprintf("/shepherd/%s/worker/%s/occonfig", DC_UUID, MAC)
		getResp, err = kv.Get(context.TODO(), mac_key)
		if err != nil {
			fmt.Println(err)
			return
		}
		data2 := ""
		for _, ev := range getResp.Kvs {
			data2 = string(ev.Value)
		}
		result2 := IfDataOk(data2)
		if ! result2{
			fmt.Println("pull global to local")
			_, err = kv.Put(context.TODO(), mac_key, data, clientv3.WithPrevKV())
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("set ok!reboot the client then read the new configuration.\n")
		}
		fmt.Println("edit data of mac key")
		card := string(os.Args[6])
		cardNum,_ := strconv.Atoi(card[0:1])
		cardData := card[2:]
		s2 := strings.Split(string(data2),";")
		s2[cardNum] = cardData
		data3 := strings.Join(s2,";")
		_, err = kv.Put(context.TODO(), mac_key, data3, clientv3.WithPrevKV())
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s set ok!reboot the client then read the new configuration.\n",card)
	}
}
