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
	"testing"

	plugins "github.com/deciphernow/gm-authenticator-plugins"
	"github.com/hashicorp/go-plugin"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPluginHandshakeConfig(t *testing.T) {
	Convey("When HandshakeConfig is invoked", t, func() {
		handshakeConfig := plugins.HandshakeConfig()

		Convey("it should return a HandshakeConfig with", func() {
			Convey("the correct MagicCookieKey", func() {
				So(handshakeConfig.MagicCookieKey, ShouldEqual, plugins.MagicCookieKey)
			})

			Convey("the correct MagicCookieValue", func() {
				So(handshakeConfig.MagicCookieValue, ShouldEqual, plugins.MagicCookieValue)
			})

			Convey("the correct ProtocolVersion", func() {
				So(handshakeConfig.ProtocolVersion, ShouldEqual, plugins.ProtocolVersion)
			})
		})
	})
}

func TestPluginServer(t *testing.T) {
	provider := NewTestProvider(plugin.BasicError{}, plugin.BasicError{}, plugin.BasicError{})
	plugin := &Plugin{provider: provider}

	Convey("When .Server is invoked", t, func() {
		server, err := plugin.Server(nil)

		Convey("it should return a server with the expected provider", func() {
			So(server.(*Server).provider, ShouldEqual, provider)
		})

		Convey("it should return a nil error", func() {
			So(err, ShouldBeNil)
		})
	})
}

func TestPluginClient(t *testing.T) {
	rpcClient := &rpc.Client{}
	provider := NewTestProvider(plugin.BasicError{}, plugin.BasicError{}, plugin.BasicError{})
	plugin := &Plugin{provider: provider}

	Convey("When .Client is invoked", t, func() {
		client, err := plugin.Client(nil, rpcClient)

		Convey("it should return a client with the expected client", func() {
			So(client.(*Client).client, ShouldEqual, rpcClient)
		})

		Convey("it should not return an error", func() {
			So(err, ShouldBeNil)
		})
	})
}
