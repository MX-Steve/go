package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var log = logrus.New()

// Config struct
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

// Host struct
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

	var configjson Config
	if err := config.Unmarshal(&configjson); err != nil {
		log.Error(err)
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
		log.Info(user, password)
		_, backupPath := BackupMySqlDb(address, port, user, password, db, "", "/opt/dbback")
		log.Info(db, " backup finished.", backupPath)
		time.Sleep(time.Second * 2)
	}
}

/**
 *
 * Backup MySql
 * @param 	host: 			localhost
 * @param 	port:			3306
 * @param 	user:			root
 * @param 	password:		root
 * @param 	databaseName:	test
 * @param 	tableName:		user
 * @param 	sqlPath:
 * @return 	backupPath
 *
 */

func BackupMySqlDb(host, port, user, password, databaseName, tableName, sqlPath string) (error, string) {
	// Delete backup 3 days ago
	currentTime := time.Now()
	oldTime := currentTime.AddDate(0, 0, -3).Format("2006010215")
	err := os.RemoveAll(sqlPath + "/" + oldTime)
	if err != nil {
		log.Error(err)
	}
	// backup
	var cmd *exec.Cmd

	if tableName == "" {
		cmd = exec.Command("mysqldump", "--opt", "-h"+host, "-P"+port, "-u"+user, "-p"+password, databaseName)
	} else {
		cmd = exec.Command("mysqldump", "--opt", "-h"+host, "-P"+port, "-u"+user, "-p"+password, databaseName, tableName)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Warn(err)
		return err, ""
	}

	if err := cmd.Start(); err != nil {
		log.Warn(err)
		return err, ""
	}

	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Warn(err)
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
		log.Warn(err)
		return err, ""
	}
	return nil, backupPath
}
