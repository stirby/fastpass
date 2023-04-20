package main

import (
	"flag"

	"github.com/s-kirby/fastpass"
)

func cmdDelete(fp *fastpass.FastPass) {
	toDelete := flag.Arg(1)
	if toDelete == "" {
		usage()
	}
	cleaned := fp.Entries.DeleteByName(toDelete)

	if len(cleaned) == len(fp.Entries) {
		printPossibleMatches(cleaned, toDelete)
		fail("could not find entry %q", toDelete)
	}

	fp.Entries = cleaned
}
