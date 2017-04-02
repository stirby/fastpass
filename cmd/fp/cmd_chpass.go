package main

import (
	"github.com/ammario/fastpass"
)

func cmdChpass(fp *fastpass.FastPass) {
	fp.Key = confirmPassword()
	writeToPasswordKeyCache(fp.Key)
	fp.Flush()
}
