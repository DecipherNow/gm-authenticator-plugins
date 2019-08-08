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
