package main

import (
	"fmt"
	"strconv"

	"flag"

	"github.com/s-kirby/fastpass"
)

func cmdLs(fp *fastpass.FastPass) {
	search := flag.Arg(1)

	var entries fastpass.Entries
	if search == "" {
		entries = fp.Entries.SortByName()
	} else {
		entries = fp.Entries.FuzzyMatch(search)
	}

	if config.Bash {
		for _, e := range entries {
			fmt.Printf("%v\n", e.Name)
		}
		return
	}

	fmt.Println("------------")
	fmt.Printf("%v: %v entries\n", config.DB, len(entries))
	fmt.Println("------------")
	largestName := 0
	for _, e := range entries {
		if runes := []rune(e.Name); len(runes) > largestName {
			largestName = len(runes)
		}
	}
	for _, e := range entries {
		fmt.Printf("%-"+strconv.Itoa(largestName)+"v [activity:%03v created:%v]", e.Name, e.Stats.Activity, e.CreatedAt.Format("01/02/2006"))
		if e.Notes != "" {
			fmt.Printf(" %v", e.Notes)
		}
		fmt.Printf("\n")
	}
}
