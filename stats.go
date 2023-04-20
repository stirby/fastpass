package fastpass

import "time"

//halving parameters
const (
	NextHalfInc = time.Hour * 24 * 7
)

func maxNextHalf() time.Time {
	return time.Now().AddDate(0, 2, 0)
}

//Stats contains stats about an entry's access.
type Stats struct {
	//Hits should not be incremented directly
	Hits int
	//NextHalfInc is added to NextHalf every hit, up to maxNextHalf
	//It's goal is to prevent entries that are no longer frequently accessed from
	//being actively selected
	NextHalf time.Time
}

//Hit adds a hit to s
func (s *Stats) Hit() {
	if time.Now().After(s.NextHalf) {
		s.NextHalf = time.Now()
		s.Hits = s.Hits / 2
	}
	s.Hits++
	if s.NextHalf.Before(maxNextHalf()) {
		s.NextHalf = s.NextHalf.Add(NextHalfInc)
	}
}
