package login

import "github.com/hashicorp/go-plugin"

// Subject represents an authentication response.
type Subject struct {

	// Value is the authenticated subject if error is not nil.
	Value string

	// Error captures the error raised if any.
	Error plugin.BasicError
}
