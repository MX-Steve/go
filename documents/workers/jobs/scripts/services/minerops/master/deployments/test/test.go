package main
 
import (
	// "fmt"
	// "io/ioutil"
	"os/exec"
)

// func writeToFile(file string, msg string) {
// 	var d1 = []byte(msg)
// 	ioutil.WriteFile(file, d1, 0666)
// }

func main() {
	// out, err  := exec.Command(echo "s 0 1100 876" > /sys/class/drm/card4/device/pp_od_clk_voltage).Output()
	item := "1100 889"
	i := "card4"
	exec.Command("/bin/sh", "-c", `echo 's 0 '`+item+` > /sys/class/drm/`+i+`/device/pp_od_clk_voltage`).Output()
	// writeToFile("/sys/class/drm/card4/device/pp_od_clk_voltage",fmt.Sprintf("s 0 %s","2000 300"))
}

// ----------------------------------------
// package main

// import(
// 	"fmt"
// 	"os/exec"
// 	"log"
// )

// func main(){
// 	fmt.Println("hello")
// 	cmd := exec.Command("sleep", "5")
// 	err := cmd.Start()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Printf("Waiting for command to finish...")
// 	err = cmd.Wait()
// 	log.Printf("Command finished with error: %v", err)
// }