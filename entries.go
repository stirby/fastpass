package fastpass

import (
	"sort"

	"strings"

	"github.com/renstrom/fuzzysearch/fuzzy"
)

//Entries is a set of entries
type Entries []*Entry

//FuzzyMatch returns all entries with names fuzzy matching search
func (es Entries) FuzzyMatch(search string) (ret Entries) {
	for _, entry := range es {
		if fuzzy.Match(search, entry.Name) {
			ret = append(ret, entry)
		}
	}
	return
}

//DeleteByName deletes an entry from es and returns the new slice
func (es Entries) DeleteByName(name string) (cleaned Entries) {
	for i, e := range es {
		if e.Name == name {
			cleaned = append(cleaned, es[:i]...)
			if i < (len(es) - 1) {
				cleaned = append(cleaned, es[i+1:]...)
			}
			return
		}
	}
	return es
}

//FindByName finds an entry by it's name.
//It returns nil if no entry was found.
func (es Entries) FindByName(name string) *Entry {
	for _, e := range es {
		if e.Name == name {
			return e
		}
	}
	return nil
}

//SortByName sorts es by name
func (es Entries) SortByName() Entries {
	sort.Slice(es, func(i, j int) bool {
		return es[i].Name < es[j].Name
	})
	return es
}

//SortByHits sorts es by hits
func (es Entries) SortByHits() Entries {
	sort.Slice(es, func(i, j int) bool {
		return es[i].Stats.Activity > es[j].Stats.Activity
	})
	return es
}

//SortByBestMatch tries to sort entries by best match
func (es Entries) SortByBestMatch(search string) Entries {
	distances := make([]int, len(es))
	for i, e := range es {
		distances[i] = fuzzy.RankMatch(search, e.Name)
	}
	sort.Slice(es, func(i, j int) bool {
		//if i contains the substring but j doesn't i is much more likely to be correct.
		if strings.Contains(es[i].Name, search) && !strings.Contains(es[j].Name, search) {
			return true
		}

		iDistance, jDistance := distances[i], distances[j]

		if iDistance < 0 {
			return false
		}
		if iDistance == 0 {
			return true
		}
		if jDistance == 0 {
			return false
		}

		iScore := (float64(es[i].Stats.Activity) / (float64(iDistance) / float64(len(es[i].Name))))
		jScore := (float64(es[j].Stats.Activity) / (float64(jDistance) / float64(len(es[j].Name))))
		return iScore > jScore
	})
	return es
}

// //FilterByTag returns all entries with a certain tag
// func (entries Entries) FilterByTag(tag string) Entries {
// 	var matches Entries
// 	for _, e := range entries {
// 		if e.HasTag(tag) {
// 			matches = append(matches, e)
// 		}
// 	}
// 	return matches
// }
