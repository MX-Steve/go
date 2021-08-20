package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//定义config结构体
type Config struct {
	UserName  string
	Password  string
	Host      Host
	Worker    string
	Shepherd  string
	Bminer    string
	Miningops string
	Pool      string
}

//json中的嵌套对应结构体的嵌套
type Host struct {
	Address string
	Port    string
}

func getConf() (user string, password string, address string, port string, worker string, shepherd string, bminer string, miningops string, pool string) {
	config := viper.New()
	config.AddConfigPath("./")
	config.SetConfigName("config")
	config.SetConfigType("json")
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}
	config.WatchConfig()
	config.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := config.ReadInConfig(); err != nil {
			panic(err)
		}
	})

	//直接反序列化为Struct
	var configjson Config
	if err := config.Unmarshal(&configjson); err != nil {
		fmt.Println(err)
	}
	WorkerBytes, err := base64.StdEncoding.DecodeString(configjson.Worker)
	if err != nil {
		log.Fatalln(err)
	}
	Worker := string(WorkerBytes)
	ShepherdBytes, err := base64.StdEncoding.DecodeString(configjson.Shepherd)
	if err != nil {
		log.Fatalln(err)
	}
	Shepherd := string(ShepherdBytes)
	BminerBytes, err := base64.StdEncoding.DecodeString(configjson.Bminer)
	if err != nil {
		log.Fatalln(err)
	}
	Bminer := string(BminerBytes)
	MiningopsBytes, err := base64.StdEncoding.DecodeString(configjson.Miningops)
	if err != nil {
		log.Fatalln(err)
	}
	Miningops := string(MiningopsBytes)
	PoolBytes, err := base64.StdEncoding.DecodeString(configjson.Pool)
	if err != nil {
		log.Fatalln(err)
	}
	Pool := string(PoolBytes)
	return configjson.UserName, configjson.Password, configjson.Host.Address, configjson.Host.Port, Worker, Shepherd, Bminer, Miningops, Pool
}

func main() {
	dataList := []string{"altchain_pool", "bminer", "conflux_pool", "miningops", "operation_portal", "shepherd"}
	user, password, address, port, worker, shepherd, bminer, miningops, pool := getConf()
	for _, db := range dataList {
		if db == "operation_portal" {
			user = "worker"
			password = worker
		} else if db == "shepherd" {
			user = "shepherd"
			password = shepherd
		} else if db == "bminer" {
			user = "bminer"
			password = bminer
		} else if db == "miningops" {
			user = "miningops"
			password = miningops
		} else {
			user = "pool"
			password = pool
		}
		_, backupPath := BackupMySqlDb(address, port, user, password, db, "", "/opt/dbback")
		fmt.Printf("%s backup finished: %s\n", db, backupPath)
		time.Sleep(time.Second * 2)
	}
}

/**
 *
 * 备份MySql数据库
 * @param 	host: 			数据库地址: localhost
 * @param 	port:			端口: 3306
 * @param 	user:			用户名: root
 * @param 	password:		密码: root
 * @param 	databaseName:	需要被分的数据库名: test
 * @param 	tableName:		需要备份的表名: user
 * @param 	sqlPath:		备份SQL存储路径: D:/backup/test/
 * @return 	backupPath
 *
 */

func BackupMySqlDb(host, port, user, password, databaseName, tableName, sqlPath string) (error, string) {
	// 删除3天前的备份
	currentTime := time.Now()
	oldTime := currentTime.AddDate(0, 0, -3).Format("2006010215")
	err := os.RemoveAll(sqlPath + "/" + oldTime)
	if err != nil {
		fmt.Println("目录不存在，无法删除")
	}
	// 实现备份
	var cmd *exec.Cmd

	if tableName == "" {
		cmd = exec.Command("mysqldump", "--opt", "-h"+host, "-P"+port, "-u"+user, "-p"+password, databaseName)
	} else {
		cmd = exec.Command("mysqldump", "--opt", "-h"+host, "-P"+port, "-u"+user, "-p"+password, databaseName, tableName)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Println(err)
		return err, ""
	}

	if err := cmd.Start(); err != nil {
		log.Println(err)
		return err, ""
	}

	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Println(err)
		return err, ""
	}
	now := time.Now().Format("2006010215")
	var backupPath string
	err2 := os.MkdirAll(sqlPath+"/"+now, 0766)
	if err2 != nil {
		fmt.Println(err2)
	}
	sqlPath = fmt.Sprintf("%s/%s/", sqlPath, now)
	if tableName == "" {
		backupPath = sqlPath + databaseName + ".sql"
	} else {
		backupPath = sqlPath + databaseName + "_" + tableName + ".sql"
	}
	err = ioutil.WriteFile(backupPath, bytes, 0644)

	if err != nil {
		log.Println(err)
		return err, ""
	}
	return nil, backupPath
}
