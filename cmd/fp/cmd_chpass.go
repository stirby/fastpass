package main

import (
	"github.com/s-kirby/fastpass"
)

func cmdChpass(fp *fastpass.FastPass) {
	fp.Key = confirmPassword()
	writeToPasswordKeyCache(fp.Key)
	fp.Flush()
}
