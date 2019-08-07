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
