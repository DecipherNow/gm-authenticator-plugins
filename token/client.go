package token

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

// Client implements the RPC client for token plugins.
type Client struct {
	client *rpc.Client
}

// Configure configures the provider or returns an error.
func (c *Client) Configure(configuration map[string]interface{}) plugin.BasicError {
	var response plugin.BasicError
	err := c.client.Call("Plugin.Configure", configuration, &response)
	if err != nil {
		response = *plugin.NewBasicError(err)
	}
	return response
}

// Token returns a token (i.e., a set of claims) for the provided subject and scopes or an error.
func (c *Client) Token(subject string, scopes []string) Token {
	var token Token
	err := c.client.Call("Plugin.Token", map[string]interface{}{"subject": subject, "scopes": scopes}, &token)
	if err != nil {
		token.Error = *plugin.NewBasicError(err)
	}
	return token
}
