package main

import (
	"io/ioutil"
	"os"

	"github.com/s-kirby/fastpass"
)

func cmdOpen() {
	fp := fastpass.New()

	if config.KeyFile != "" {
		info("key file set.. doing nothing")
		os.Exit(0)
	}

	fp.Key = fastpass.GetPassword()
	if err := fp.Open(config.DB); err != nil {
		fail("failed to open db %v: %v", config.DB, err)
	}

	if err := ioutil.WriteFile(passwordKeyCache, fp.Key[:], 0600); err != nil {
		fail("failed to write password key cache: %v", err)
	}
}
