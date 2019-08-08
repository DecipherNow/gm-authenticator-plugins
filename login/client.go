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
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

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
