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
