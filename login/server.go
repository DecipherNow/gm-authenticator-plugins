package login

import (
	"encoding/gob"

	plugins "github.com/deciphernow/gm-authenticator-plugins"
	"github.com/hashicorp/go-plugin"
)

func init() {
	gob.Register(map[string]interface{}{})
	gob.Register([][]string{})
	gob.Register(Challenges{})
	gob.Register([]Response{})
	gob.Register(Subject{})
	gob.Register(plugin.BasicError{})
}

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
