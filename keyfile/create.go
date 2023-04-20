package keyfile

import (
	"crypto/rand"
	"io"

	"os"

	"github.com/pkg/errors"
)

//Create creates a key file at path and returns the key.
func Create(path string) (key [32]byte, err error) {
	_, err = io.ReadFull(rand.Reader, key[:])
	if err != nil {
		panic(err)
	}
	fi, err := os.OpenFile(path, os.O_EXCL|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return key, errors.Wrap(err, "failed to open fi")
	}
	defer fi.Close()
	_, err = fi.Write(key[:])
	return key, errors.Wrap(err, "failed to write file")
}
