package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/julienschmidt/httprouter"
)

var mutex sync.Mutex

func main() {
	router := httprouter.New()
	router.GET("/health", func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		w.Write([]byte("OK"))
	})

	router.POST("/machine/:mac/:ip", handleSetMachine)
	router.POST("/machine/:mac/:ip/:sitename", handleSetMachine)
	router.GET("/allmachines", handleAllmachines)

	server := &http.Server{
		Addr:    "0.0.0.0:9103",
		Handler: router,
	}
	server.ListenAndServe()
}

const filePath = "./machines"

type machineInfo struct {
	MAC        string
	IP         string
	TunnelPort string
	SiteName   string
}

func readFile() (map[string]machineInfo, map[string]bool, error) {
	mutex.Lock()
	defer mutex.Unlock()
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return nil, nil, err
	}

	machineInfos := map[string]machineInfo{}
	ports := map[string]bool{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		raw := string(scanner.Bytes())
		if raw == "" {
			continue
		}
		if strings.HasPrefix(raw, "#") {
			continue
		}

		r := strings.Split(raw, ",")
		if len(r) != 3 && len(r) != 4 {
			return nil, nil, errors.New("column number error")
		}
		macAddr := r[0]
		port := r[1]
		ipAddr := r[2]
		sitename := "DEFAULT"
		if len(r) == 4 {
			sitename = r[3]
		}
		machineInfos[macAddr] = machineInfo{
			MAC:        macAddr,
			TunnelPort: port,
			IP:         ipAddr,
			SiteName:   sitename,
		}
		ports[port] = true
	}
	return machineInfos, ports, nil
}

func handleAllmachines(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	log.Info("handle all machines")
	machineInfos, _, err := readFile()
	if err != nil {
		log.Fatal(err)
	}
	machines := []machineInfo{}
	for _, v := range machineInfos {
		machines = append(machines, v)
	}
	log.Info(machines)
	machinesBody, err := json.Marshal(machines)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(machinesBody)
}

func handleSetMachine(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	log.Info("handleSetMachine")
	machines, ports, err := readFile()
	if err != nil {
		log.Fatal(err)
	}

	mac := ps.ByName("mac")
	ip := ps.ByName("ip")
	sitename := ps.ByName("sitename")
	if sitename == "" {
		sitename = "DEFAULT"
	}
	log.Infof("One machine connected, mac address:%s, ip address:%s, sitename:%s", mac, ip, sitename)
	_, err = net.ParseMAC(mac)
	if err != nil {
		http.Error(w, "Bad mac", 400)
		return
	}
	if net.ParseIP(ip) == nil {
		http.Error(w, "Bad ip", 400)
		return
	}

	machine, ok := machines[mac]
	if !ok {
		newPort := getRandomPort(ports)
		log.Infof("This is a new machine, port:%s", newPort)
		err := writeFile(mac, newPort, ip, sitename, true)
		if err != nil {
			log.Fatal(err)
		}
		w.Write([]byte(newPort))
	} else {
		log.Infof("This is a existed machine, port:%s", machine.TunnelPort)
		if ip != machine.IP || sitename != machine.SiteName {
			log.Infof("Update machine %s ip from:%s to :%s, sitename from:%s to:%s", mac, machine.IP, ip, machine.SiteName, sitename)
			err := writeFile(mac, machine.TunnelPort, ip, sitename, false)
			if err != nil {
				log.Fatal(err)
			}
		}

		w.Write([]byte(machine.TunnelPort))
	}
	return
}

func startReverse(newPort string) error {
	initStr := "-fCNL *:%s:localhost:%s localhost"
	str := fmt.Sprintf(initStr, newPort, newPort)
	cmd := exec.Command("ssh", str)
	if err := cmd.Start(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func writeFile(mac, port, ip, sitename string, isNew bool) error {
	mutex.Lock()
	defer mutex.Unlock()

	if isNew {
		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer file.Close()
		str := mac + "," + port + "," + ip + "," + sitename + "\n"
		if _, err := file.WriteString(str); err != nil {
			return err
		}
	} else {
		input, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Fatal(err)
		}

		lines := strings.Split(string(input), "\n")

		for i, line := range lines {
			if strings.Contains(line, mac) {
				lines[i] = mac + "," + port + "," + ip + "," + sitename
			}
		}
		output := strings.Join(lines, "\n")
		err = ioutil.WriteFile(filePath, []byte(output), 0644)
		if err != nil {
			log.Fatal("writeFileErr:", err)
		}
	}

	return nil
}

func getRandomPort(ports map[string]bool) string {
	baseNum := 10000
	for {
		num := rand.Intn(2000)
		port := baseNum + num
		if _, ok := ports[strconv.Itoa(port)]; ok {
			continue
		}
		return strconv.Itoa(port)
	}
}
