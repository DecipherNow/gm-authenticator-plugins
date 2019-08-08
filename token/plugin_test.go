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
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// TestServer tests the token.Plugin.Server method.
func TestPluginServer(t *testing.T) {

	Convey("When Plugin.Server is invoked", t, func() {

		provider := &TestProvider{}
		plugin := &Plugin{provider: provider}
		server, err := plugin.Server(nil)

		Convey("it should return a server with the expected provider", func() {
			So(server.(*Server).provider, ShouldEqual, provider)
		})

		Convey("it should not return an error", func() {
			So(err, ShouldBeNil)
		})
	})
}

// TestClient test the Plugin.Client method.
func TestPluginClient(t *testing.T) {

	internal := &rpc.Client{}
	plugin := &Plugin{provider: &TestProvider{}}
	client, err := plugin.Client(nil, internal)

	Convey("When Plugin.Client is invoked", t, func() {

		Convey("it should return a client with the expected client", func() {
			So(client.(*Client).client, ShouldEqual, internal)
		})

		Convey("it should not return an error", func() {
			So(err, ShouldBeNil)
		})
	})
}
