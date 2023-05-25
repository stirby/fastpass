package main

import (
	"flag"
	"time"

	"strings"

	"github.com/s-kirby/fastpass"
	"github.com/s-kirby/fastpass/passgen"
)

func passwordFromGenerator() (password string) {
	switch config.Generator {
	case "human":
		return passgen.Human()
	case "hex":
		return passgen.Hex()
	case "base62":
		return passgen.Base62()
	default:
		fail("unknown generator: %v", config.Generator)
	}
	return ""
}

func cmdNew(fp *fastpass.FastPass) {
	name := flag.Arg(1)
	entry := fp.Entries.FindByName(name)
	if entry != nil {
		fail("an entry with name %v already exists", name)
	}
	entry = &fastpass.Entry{
		Name:      strings.ToLower(name),
		Password:  passwordFromGenerator(),
		CreatedAt: time.Now(),
	}
	fp.Entries = append(fp.Entries, entry)
	copyPassword(entry)
}
