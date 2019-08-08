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

package plugins

import (
	"github.com/hashicorp/go-plugin"
)

const (
	// MagicCookieKey is the value of the MagicCookieKey in the HandshakeConfig.
	MagicCookieKey = "AUTHENTICATOR_PROVIDER"

	// MagicCookieValue is the value of the MagicCookieValue in the HandshakeConfig.
	MagicCookieValue = "xrj4kMghmVY7RE7k01dq"

	// ProtocolVersion is the value of the ProtocolVersion in the HandshakeConfig.
	ProtocolVersion = 1
)

// HandshakeConfig returns the handshake configuration for plugins.
func HandshakeConfig() plugin.HandshakeConfig {
	return plugin.HandshakeConfig{
		MagicCookieKey:   MagicCookieKey,
		MagicCookieValue: MagicCookieValue,
		ProtocolVersion:  ProtocolVersion,
	}
}
