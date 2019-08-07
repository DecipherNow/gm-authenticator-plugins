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
