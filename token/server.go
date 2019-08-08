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
