package token

import "github.com/hashicorp/go-plugin"

// Provider defines the interface for token providers.
//
// Token providers convert a subject and a set of scopes into a token (i.e., a set of claims).  The token will be used
// to identify the user or system making the request and any associated attributes.
type Provider interface {

	// Configure configures the provider or returns an error.
	Configure(config map[string]interface{}) plugin.BasicError

	// Token returns a token (i.e., a set of claims) for the provided subject and scopes.
	Token(subject string, scopes []string) Token
}
