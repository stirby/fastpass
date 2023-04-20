package passgen

import (
	"github.com/ammario/crand"
)

//Base62 generates a base62 password
func Base62() string {
	return crand.String(16)
}
