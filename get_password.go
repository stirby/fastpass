package fastpass

import (
	"crypto/sha256"
	"fmt"

	"github.com/howeyc/gopass"
	"golang.org/x/crypto/pbkdf2"
)

var salt = []byte("seems like a decent salt")

//GetPassword retrieves a password from the command line.
//It returns the sha256 representation of the password.
func GetPassword() (key [32]byte) {
	fmt.Printf("Enter password: ")
	rawpwd, err := gopass.GetPasswdMasked()
	if err != nil {
		panic(err)
	}
	copy(key[:], pbkdf2.Key(rawpwd, salt, 4096, 32, sha256.New))
	return
}
