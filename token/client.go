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
