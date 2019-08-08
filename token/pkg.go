package token

import (
	"encoding/gob"

	"github.com/hashicorp/go-plugin"
)

func init() {
	gob.Register(map[string]interface{}{})
	gob.Register([]string{})
	gob.Register(Token{})
	gob.Register(plugin.BasicError{})
}
