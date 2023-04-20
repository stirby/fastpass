package main

import (
	"os"

	"io/ioutil"

	"github.com/s-kirby/fastpass/keyfile"
)

//getKey retrieves the active key
func getKey() (key [32]byte) {
	var err error
	if config.KeyFile != "" {
		key, err = keyfile.Load(config.KeyFile)
		if err != nil {
			fail("failed to load key file @ %v: %v", config.KeyFile, err)
		}
		return key
	}

	fi, err := os.OpenFile(passwordKeyCache, os.O_RDONLY, 0600)
	if err != nil {
		fail("failed to open password cache (maybe you forgot to run `fp open`)", passwordKeyCache)
	}
	defer fi.Close()

	dat, err := ioutil.ReadAll(fi)
	if err != nil {
		fail("failed to read %v: %v", fi.Name(), err)
	}

	copy(key[:], dat)
	return
}
