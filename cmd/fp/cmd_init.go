package main

import (
	"os"

	"github.com/ammario/fastpass"
	"github.com/ammario/fastpass/keyfile"
	"github.com/pkg/errors"
)

func cmdInit() {
	fp := fastpass.New()
	if _, err := os.Stat(config.DB); err == nil {
		fail("db @ %v exists", config.DB)
	}

	var key [32]byte
	var err error

	if config.KeyFile == "" {
		key = confirmPassword()
	} else {
		if key, err = keyfile.Load(config.KeyFile); os.IsNotExist(errors.Cause(err)) {
			if key, err = keyfile.Create(config.KeyFile); err != nil {
				fail("failed to create key file @ %v: %v", config.KeyFile, err)
			}
		} else if err != nil {
			fail("unexpected error while loading %v: %v", config.KeyFile, err)
		}
		info("using key file @ %v", config.KeyFile)
	}
	fp.Key = key

	if err := fp.Create(config.DB); err != nil {
		fail("failed to create db: %v", err)
	}

	fp.Close()

	writeToPasswordKeyCache(fp.Key)

	success("created db @ %v", config.DB)
}
