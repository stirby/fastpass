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
	pass := flag.Arg(2)
	entry := fp.Entries.FindByName(name)
	if entry != nil {
		fail("an entry with name %v already exists", name)
	}
	if pass == "" {
		pass = passwordFromGenerator()
	}
	entry = &fastpass.Entry{
		Name:      strings.ToLower(name),
		Password:  pass,
		CreatedAt: time.Now(),
	}
	fp.Entries = append(fp.Entries, entry)
	copyPassword(entry)
}
