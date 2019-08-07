package login

import (
	"encoding/gob"
	"net/rpc"

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

// Client implements the RPC client for login plugins.
type Client struct {
	client *rpc.Client
}

// Authenticate calls the plugin's authenticate method and returns authentication response.
func (c *Client) Authenticate(responses []Response, subjects [][]string) Subject {
	var subject Subject
	err := c.client.Call("Plugin.Authenticate", map[string]interface{}{"responses": responses, "subjects": subjects}, &subject)
	if err != nil {
		subject.Error = *plugin.NewBasicError(err)
	}
	return subject
}

// Challenges calls the plugin's challenge method and returns the challenges response.
func (c *Client) Challenges() Challenges {
	var challenges Challenges
	err := c.client.Call("Plugin.Challenges", new(interface{}), &challenges)
	if err != nil {
		challenges.Error = *plugin.NewBasicError(err)
	}
	return challenges
}

// Configure calls the plugin's configure method and configures the plugin returns an error.
func (c *Client) Configure(configuration map[string]interface{}) plugin.BasicError {
	var response plugin.BasicError
	err := c.client.Call("Plugin.Configure", configuration, &response)
	if err != nil {
		response = *plugin.NewBasicError(err)
	}
	return response
}
