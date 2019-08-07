package plugins

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHandshakeConfig(t *testing.T) {

	Convey("When plugin.HandshakeConfig is invoked", t, func() {

		handshakeConfig := HandshakeConfig()

		Convey("it should return a HandshakeConfig with", func() {

			Convey("the correct MagicCookieKey", func() {
				So(handshakeConfig.MagicCookieKey, ShouldEqual, MagicCookieKey)
			})

			Convey("the correct MagicCookieValue", func() {
				So(handshakeConfig.MagicCookieValue, ShouldEqual, MagicCookieValue)
			})

			Convey("the correct ProtocolVersion", func() {
				So(handshakeConfig.ProtocolVersion, ShouldEqual, ProtocolVersion)
			})
		})
	})
}
