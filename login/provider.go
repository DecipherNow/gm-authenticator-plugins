package login

import "github.com/hashicorp/go-plugin"

// Provider defines the interface for login providers.
type Provider interface {

	// Authenticate authenticates the responses and verified subjects then returns the subject.
	Authenticate([]Response, [][]string) Subject

	// Challenges returns the challenges for this provider.
	Challenges() Challenges

	// Configure configures the provider or returns an error.
	Configure(map[string]interface{}) plugin.BasicError
}
