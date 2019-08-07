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
