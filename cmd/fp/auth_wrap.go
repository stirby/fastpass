package main

import (
	"github.com/ammario/fastpass"
)

//authWraps provides a valid fp to a handler
func authWrap(f func(fp *fastpass.FastPass)) {
	fp := fastpass.New()
	fp.Key = getKey()
	if err := fp.Open(config.DB); err != nil {
		fail("failed to open db: %v", err)
	}
	f(fp)
	if err := fp.Close(); err != nil {
		fail("failed to close fp: %v", err)
	}
}
