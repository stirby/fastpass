package main

import (
	"flag"

	"fmt"

	"github.com/ammario/fastpass"
)

func cmdGet(fp *fastpass.FastPass) {
	search := flag.Arg(0)

	if len(flag.Args()) != 1 {
		usage()
	}

	results := fp.Entries.SortByName()

	if search != "" {
		results = fp.Entries.SortByBestMatch(search)
	}

	if len(results) == 0 {
		fail("no results found")
	}

	e := results[0]
	e.Stats.Hit()

	if len(results) > 1 {
		fmt.Printf("similar: ")
		for _, r := range results[1:] {
			fmt.Printf("%v ", r.Name)
		}
		fmt.Printf("\n")
	}

	copyPassword(e)
}
