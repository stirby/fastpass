package main

import "github.com/ammario/fastpass"
import "fmt"

func printPossibleMatches(es fastpass.Entries, search string) {
	es = es.FuzzyMatch(search).SortByHits()
	if len(es) == 0 {
		return
	}
	fmt.Printf("Possible matches: [")
	for _, e := range es {
		fmt.Printf("%q ", e.Name)
	}
	fmt.Printf("]\n")
}
