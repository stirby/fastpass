package keyfile

import (
	"io/ioutil"

	"github.com/pkg/errors"
)

//Load loads a key file
func Load(path string) (key [32]byte, err error) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return key, errors.Wrap(err, "failed to read file")
	}
	if len(dat) != 32 {
		return key, errors.Errorf("keyfile must be 32 bytes")
	}
	copy(key[:], dat)
	return
}
