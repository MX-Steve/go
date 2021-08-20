package main

import (
	"encoding/hex"
	"os"

	"github.com/btcsuite/btcutil/base58"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("")

	bin := keyStrToBinary(os.Args[1])
	log.Info(hex.EncodeToString(bin))
}

func keyStrToBinary(s string) []byte {
	decoded, version, _ := base58.CheckDecode(s)

	b := []byte{version}
	b = append(b, decoded...)
	return b
}
