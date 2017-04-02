package passgen

import (
	"crypto/rand"
	"encoding/hex"
	"io"
)

//Hex returns a 16 byte hex password
func Hex() string {
	var buf [8]byte
	if _, err := io.ReadFull(rand.Reader, buf[:]); err != nil {
		panic(err)
	}
	return hex.EncodeToString(buf[:])
}
