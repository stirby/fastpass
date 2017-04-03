package fastpass

import (
	"crypto/sha256"
	"fmt"

	"github.com/howeyc/gopass"
	"golang.org/x/crypto/pbkdf2"
)

var salt = []byte("seems like a decent salt")

const iters = 1024 * 64

//GetPassword retrieves a password from the command line.
//It returns the the stretched representation of the password using pbkdf2.
func GetPassword() (key [32]byte) {
	fmt.Printf("Enter password: ")
	rawpwd, err := gopass.GetPasswdMasked()
	if err != nil {
		panic(err)
	}
	copy(key[:], pbkdf2.Key(rawpwd, salt, iters, 32, sha256.New))
	return
}
