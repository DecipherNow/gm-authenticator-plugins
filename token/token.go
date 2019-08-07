package token

import "github.com/hashicorp/go-plugin"

// Token represents an token response.
type Token struct {

	// Value is the token or nil if error is not nil.
	Value map[string]interface{}

	// Error captures the error raised if any.
	Error plugin.BasicError
}
