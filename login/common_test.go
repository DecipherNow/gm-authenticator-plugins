package login

import (
	"crypto/tls"

	"github.com/hashicorp/go-plugin"
)

type TestProvider struct {
	configuration      map[string]interface{}
	configurationError plugin.BasicError
	challengeError     plugin.BasicError
	authenticateError  plugin.BasicError
	subject            string
	subjects           [][]string
	state              tls.ConnectionState
}

func NewTestProvider(configurationError plugin.BasicError, authenticateError plugin.BasicError, challengeError plugin.BasicError) *TestProvider {
	return &TestProvider{
		configuration:      nil,
		configurationError: configurationError,
		authenticateError:  authenticateError,
		challengeError:     challengeError,
		subject:            "",
	}
}

func (p *TestProvider) Authenticate(responses []Response, subjects [][]string) Subject {
	p.subjects = subjects
	return Subject{
		Value: p.subject,
		Error: p.authenticateError,
	}
}

func (p *TestProvider) Challenges() Challenges {
	return Challenges{
		Value: *new([]Challenge),
		Error: p.challengeError,
	}
}

func (p *TestProvider) Configure(args map[string]interface{}) plugin.BasicError {
	p.configuration = args
	p.subject = args["subject"].(string)
	return p.configurationError
}
