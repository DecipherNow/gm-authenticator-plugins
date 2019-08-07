package token

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

// Plugin implements the plugin interface for token plugins.
type Plugin struct {
	provider Provider
}

// Server returns an RPC server for token plugins.
func (p *Plugin) Server(broker *plugin.MuxBroker) (interface{}, error) {
	return &Server{provider: p.provider}, nil
}

// Client returns an RPC client for token plugins.
func (p *Plugin) Client(broker *plugin.MuxBroker, client *rpc.Client) (interface{}, error) {
	return &Client{client}, nil
}
