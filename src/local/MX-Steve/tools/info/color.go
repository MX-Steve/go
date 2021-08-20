package info

import (
	"fmt"
)

func Yellow(s string) {
	fmt.Printf("\n%c[5;33;40m%s%c[0m\n", 0x1B, s, 0x1B)
}

func Red(s string) {
	fmt.Printf("\n%c[5;31;40m%s%c[0m\n", 0x1B, s, 0x1B)
}

func Green(s string) {
	fmt.Printf("\n%c[5;32;40m%s%c[0m\n", 0x1B, s, 0x1B)
}