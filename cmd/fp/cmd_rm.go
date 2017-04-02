package main

import (
	"flag"

	"github.com/ammario/fastpass"
)

func cmdRm(fp *fastpass.FastPass) {
	toDelete := flag.Arg(1)
	if toDelete == "" {
		usage()
	}
	cleaned := fp.Entries.DeleteByName(toDelete)

	if len(cleaned) == len(fp.Entries) {
		fail("could not find entry %q", toDelete)
	}

	fp.Entries = cleaned
}
