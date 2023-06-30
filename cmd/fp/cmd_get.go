package main

import (
	"flag"

	"fmt"

	"github.com/s-kirby/fastpass"
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
		for i, r := range results[1:] {
			//show a maximum of five suggestions
			if i > 5 {
				break
			}
			fmt.Printf("%v ", r.Name)
		}
		fmt.Printf("\n")
	}

	copyPassword(e)
}
