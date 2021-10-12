package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// Config struct
type Config struct {
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

var c Config

func initConf() {
	c.Bminer = os.Getenv("bminer")
	c.Miningops = os.Getenv("miningops")
	c.Shepherd = os.Getenv("shepherd")
	c.Pool = os.Getenv("pool")
	c.Worker = os.Getenv("worker")
	c.Host.Address = "mariadb.default.svc.cluster.local"
	c.Host.Port = "3306"
}

func main() {
	dataList := []string{"altchain_pool", "bminer", "conflux_pool", "miningops", "operation_portal", "shepherd"}
	initConf()
	user := ""
	password := ""
	for _, db := range dataList {
		if db == "operation_portal" {
			user = "worker"
			password = c.Worker
		} else if db == "shepherd" {
			user = "shepherd"
			password = c.Shepherd
		} else if db == "bminer" {
			user = "bminer"
			password = c.Bminer
		} else if db == "miningops" {
			user = "miningops"
			password = c.Miningops
		} else {
			user = "pool"
			password = c.Pool
		}
		_, backupPath := BackupMySqlDb(c.Host.Address, c.Host.Port, user, password, db, "", "/opt/dbback")
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
