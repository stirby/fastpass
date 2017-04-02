package main

import (
	"flag"

	"fmt"

	"github.com/ammario/fastpass"
	"github.com/atotto/clipboard"
	"github.com/fatih/color"
)

func cmdGet(fp *fastpass.FastPass) {
	search := flag.Arg(0)

	if len(flag.Args()) > 1 {
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

	color.New(color.Bold).Printf("%v", e.Name)
	if config.Show {
		color.New(color.FgHiMagenta).Printf(" -> %q", e.Password)
	}
	if config.Copy {
		if err := clipboard.WriteAll(e.Password); err != nil {
			fail("cannot copy to clipboard: %v", err)
		}
		fmt.Printf(" -> Password Copied!")
	}
	fmt.Printf("\n")
}
