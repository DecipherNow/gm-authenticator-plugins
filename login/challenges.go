package login

import "github.com/hashicorp/go-plugin"

// Challenges represents a challenges response.
type Challenges struct {

	// Value defines the challenge values if error is not nil.
	Value []Challenge

	// Error captures the error raised if any.
	Error plugin.BasicError
}
