package token

import (
	"errors"
	"reflect"
	"testing"

	"github.com/deciphernow/gm-authenticator-plugins/test"
	"github.com/hashicorp/go-plugin"
	. "github.com/smartystreets/goconvey/convey"
)

// TestServerConfigure tests the Server.Configure method.
func TestServerConfigure(t *testing.T) {

	configErr := test.MustGenerate(reflect.TypeOf(plugin.BasicError{}), t).Interface().(plugin.BasicError)
	tokenErr := test.MustGenerate(reflect.TypeOf(plugin.BasicError{}), t).Interface().(plugin.BasicError)
	provider := NewTestProvider(configErr, tokenErr)
	server := &Server{provider: provider}

	Convey("When .Configure is invoked", t, func() {

		var response plugin.BasicError

		configuration := make(map[string]interface{})
		err := server.Configure(configuration, &response)

		Convey("it passes the configuration to the provider", func() {
			So(provider.configuration, ShouldEqual, configuration)
		})

		Convey("it returns the nil", func() {
			So(err, ShouldBeNil)
		})

		Convey("it returns the error from the provider", func() {
			So(response, ShouldResemble, configErr)
		})
	})
}

// TestToken tests the Server.Token method.
func TestServerToken(t *testing.T) {

	configurationError := errors.New(test.MustGenerate(reflect.TypeOf(""), t).Interface().(string))
	tokenError := errors.New(test.MustGenerate(reflect.TypeOf(""), t).Interface().(string))

	configErr := plugin.NewBasicError(configurationError)
	tokenErr := plugin.NewBasicError(tokenError)
	provider := NewTestProvider(*configErr, *tokenErr)
	server := &Server{provider: provider}

	Convey("When .Token is invoked", t, func() {

		var response Token

		scopes := test.MustGenerate(reflect.TypeOf([]string{}), t).Interface().([]string)
		subject := test.MustGenerate(reflect.TypeOf(""), t).Interface().(string)
		server.Token(map[string]interface{}{"subject": subject, "scopes": scopes}, &response)

		Convey("it passes the subject to the provider", func() {
			So(provider.subject, ShouldEqual, subject)
		})

		Convey("it passes the scopes to the provider", func() {
			So(provider.scopes, ShouldResemble, scopes)
		})

		Convey("it returns the error from the provider", func() {
			So(response.Error, ShouldResemble, provider.tokenError)
		})

		Convey("the response should be the token", func() {
			So(response.Value, ShouldEqual, provider.token)
		})
	})
}
