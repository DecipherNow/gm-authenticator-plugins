package login

import (
	"encoding/gob"

	"github.com/hashicorp/go-plugin"
)

func init() {
	gob.Register(map[string]interface{}{})
	gob.Register([][]string{})
	gob.Register(Challenges{})
	gob.Register([]Response{})
	gob.Register(Subject{})
	gob.Register(plugin.BasicError{})
}
