// Copyright 2019 Decipher Technology Studios
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
