package main

import (
	"bytes"
	"fmt"

	"github.com/s-kirby/fastpass"
)

func confirmPassword() [32]byte {
	pwd := fastpass.GetPassword()
	fmt.Printf("(confirm) ")
	cpwd := fastpass.GetPassword()

	if bytes.Compare(pwd[:], cpwd[:]) != 0 {
		fail("password mismatch")
	}
	return pwd
}
