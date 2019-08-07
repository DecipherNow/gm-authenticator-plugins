package token

import "github.com/hashicorp/go-plugin"

// TestProvider implements the Provider interface for testing purposes.
type TestProvider struct {
	configuration      map[string]interface{}
	configurationError plugin.BasicError
	scopes             []string
	subject            string
	token              map[string]interface{}
	tokenError         plugin.BasicError
}

// NewTestProvider initializes a new instance of a TestProvider.
func NewTestProvider(configurationError plugin.BasicError, tokenError plugin.BasicError) *TestProvider {
	return &TestProvider{
		configuration:      nil,
		configurationError: configurationError,
		scopes:             nil,
		subject:            "",
		token:              nil,
		tokenError:         tokenError,
	}
}

// Configure implements the Provider.Configure method of testing purposes.
func (p *TestProvider) Configure(config map[string]interface{}) plugin.BasicError {
	p.configuration = config
	return p.configurationError
}

// Configure implements the Provider.Token method of testing purposes.
func (p *TestProvider) Token(subject string, scopes []string) Token {
	p.scopes = scopes
	p.subject = subject
	return Token{
		Value: p.token,
		Error: p.tokenError,
	}
}
