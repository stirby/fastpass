package main

import (
	"bytes"
	"fmt"

	"github.com/ammario/fastpass"
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
