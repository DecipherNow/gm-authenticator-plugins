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

import (
	plugins "github.com/deciphernow/gm-authenticator-plugins"
	"github.com/hashicorp/go-plugin"
)

// Server implements the RPC server for token plugins.
type Server struct {
	provider Provider
}

// Configure configures the provider or returns an error.
func (s *Server) Configure(args map[string]interface{}, response *plugin.BasicError) error {
	*response = s.provider.Configure(args)
	return nil
}

// Token returns a token (i.e., a set of claims) for the provided args (i.e., a subject and scopes) or an error.
func (s *Server) Token(args map[string]interface{}, response *Token) error {
	*response = s.provider.Token(args["subject"].(string), args["scopes"].([]string))
	return nil
}

// Serve starts the provider as a plugin and blocks indefinitely.
func Serve(provider Provider) {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: plugins.HandshakeConfig(),
		Plugins:         map[string]plugin.Plugin{"token": &Plugin{provider: provider}},
	})
}
