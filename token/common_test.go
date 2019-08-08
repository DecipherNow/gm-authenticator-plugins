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
