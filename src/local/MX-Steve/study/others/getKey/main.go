package main 

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main(){
	hasher := md5.New()
	hasher.Write([]byte("updateWarning"))
	hasher.Write([]byte("clickhouse"))
	fmt.Println(hex.EncodeToString(hasher.Sum(nil)))
}