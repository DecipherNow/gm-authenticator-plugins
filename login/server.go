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
	plugins "github.com/deciphernow/gm-authenticator-plugins"
	"github.com/hashicorp/go-plugin"
)

// Server represents the RPC server for login provider plugins, according to the
// net/rpc requirements.
type Server struct {
	provider Provider
}

// Authenticate authenticates the responses and context then returns the subject or an error
func (s *Server) Authenticate(args map[string]interface{}, response *Subject) error {
	*response = s.provider.Authenticate(args["responses"].([]Response), args["subjects"].([][]string))
	return nil
}

// Challenges returns the challenges for this provider
func (s *Server) Challenges(args interface{}, response *Challenges) error {
	*response = s.provider.Challenges()
	return nil
}

// Configure configures the login provider or returns an error
func (s *Server) Configure(args map[string]interface{}, response *plugin.BasicError) error {
	*response = s.provider.Configure(args)
	return nil
}

// Serve starts the provider as a plugin and blocks indefinitely.
func Serve(provider Provider) {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: plugins.HandshakeConfig(),
		Plugins:         map[string]plugin.Plugin{"login": &Plugin{provider: provider}},
	})
}
