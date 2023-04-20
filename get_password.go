package fastpass

import "github.com/howeyc/gopass"
import "crypto/sha256"
import "fmt"

//GetPassword retrieves a password from the command line.
//It returns the sha256 representation of the password.
func GetPassword() [32]byte {
	fmt.Printf("Enter password: ")
	rawpwd, err := gopass.GetPasswdMasked()
	if err != nil {
		panic(err)
	}
	return sha256.Sum256(rawpwd)
}
