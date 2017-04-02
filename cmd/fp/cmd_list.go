package main

import (
	"fmt"
	"strconv"

	"flag"

	"github.com/ammario/fastpass"
)

func cmdList(fp *fastpass.FastPass) {
	search := flag.Arg(1)

	entries := fp.Entries.FuzzyMatch(search)

	fmt.Println("------------")
	fmt.Printf("%v: %v entries\n", config.DB, len(entries))
	fmt.Println("------------")
	largestName := 0
	for _, e := range entries {
		if runes := []rune(e.Name); len(runes) > largestName {
			largestName = len(runes)
		}
	}
	for _, e := range entries.SortByName() {
		fmt.Printf("%-"+strconv.Itoa(largestName)+"v [hits:%v created:%v]", e.Name, e.Stats.Hits, e.CreatedAt.Format("01/02/2006 15:04:05"))
		if e.Notes != "" {
			fmt.Printf(" Notes: %v", e.Notes)
		}
		fmt.Printf("\n")
	}
}
