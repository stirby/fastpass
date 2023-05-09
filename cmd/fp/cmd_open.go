package main

import (
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

	writeToPasswordKeyCache(fp.Key)
}
