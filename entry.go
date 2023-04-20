package fastpass

import "time"

//Entry is a fastpass entry
type Entry struct {
	Name      string
	Password  string
	CreatedAt time.Time
	Stats     Stats
	Notes     string
}

// //HasTag returns if the entry has a certain tag
// func (e *Entry) HasTag(tag string) bool {
// 	tag = strings.ToLower(tag)
// 	for _, t := range e.Tags {
// 		if tag == t {
// 			return true
// 		}
// 	}
// 	return false
// }
