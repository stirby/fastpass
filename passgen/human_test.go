package passgen

import (
	"fmt"
	"testing"
)

func TestHuman(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Human Pass: %s\n", Human())
	}
}
