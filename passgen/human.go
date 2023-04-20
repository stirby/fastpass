package passgen

import (
	"bytes"
	"fmt"

	"github.com/ammario/crand"
)

var nouns [][]byte
var adjectives [][]byte

func init() {
	clean := func(a [][]byte) {
		for i := range a {
			a[i] = bytes.TrimSpace(a[i])
		}
	}

	nouns = bytes.Split(MustAsset("word_list/nouns.txt"), []byte("\n"))
	adjectives = bytes.Split(MustAsset("word_list/adjectives.txt"), []byte("\n"))

	clean(nouns)
	clean(adjectives)

	if len(nouns) == 0 {
		panic("no nouns")
	}

	if len(adjectives) == 0 {
		panic("no adjectives")
	}
}

//Human generates a password easy for a human to read.
func Human() string {
	buf := &bytes.Buffer{}

	// log.Fatalf("%v\n", crand.Uint(uint(len(adjectives))))
	//write both adjectives
	buf.Write(bytes.Title(adjectives[crand.Uint(uint(len(adjectives)))]))
	buf.Write(bytes.Title(adjectives[crand.Uint(uint(len(adjectives)))]))

	//write both nouns
	buf.Write(bytes.Title(nouns[crand.Uint(uint(len(nouns)))]))
	buf.Write(bytes.Title(nouns[crand.Uint(uint(len(nouns)))]))

	//write random number
	buf.WriteString(fmt.Sprintf("%.3v", crand.Uint(1000)))

	return buf.String()
}
