package login

import (
	"errors"
	"reflect"
	"testing"

	"github.com/deciphernow/gm-authenticator-plugins/test"
	"github.com/hashicorp/go-plugin"
	. "github.com/smartystreets/goconvey/convey"
)

func TestServerConfigure(t *testing.T) {
	configurationError := plugin.NewBasicError(errors.New(test.MustGenerate(reflect.TypeOf(""), t).Interface().(string)))
	authenticateError := plugin.NewBasicError(errors.New(test.MustGenerate(reflect.TypeOf(""), t).Interface().(string)))
	challengeError := plugin.NewBasicError(errors.New(test.MustGenerate(reflect.TypeOf(""), t).Interface().(string)))
	provider := NewTestProvider(*configurationError, *authenticateError, *challengeError)
	server := &Server{provider: provider}
	configuration := map[string]interface{}{"subject": test.MustGenerate(reflect.TypeOf(""), t).Interface().(string)}

	Convey("When .Configure is invoked", t, func() {
		var response plugin.BasicError
		server.Configure(configuration, &response)

		Convey("it passes the configuration to the provider", func() {
			So(provider.configuration, ShouldEqual, configuration)
		})

		Convey("it returns the error from the provider", func() {
			So(response.Message, ShouldEqual, configurationError.Message)
		})
	})
}

func TestServerChallenges(t *testing.T) {
	challengeError := plugin.NewBasicError(errors.New(test.MustGenerate(reflect.TypeOf(""), t).Interface().(string)))
	provider := NewTestProvider(plugin.BasicError{}, plugin.BasicError{}, *challengeError)
	server := &Server{provider: provider}

	Convey("When .Challenges is invoked", t, func() {
		var response Challenges
		var args map[string]interface{}
		server.Challenges(args, &response)

		Convey("it returns the correct array of challenges", func() {
			So(response.Value, ShouldBeEmpty)
		})

		Convey("it returns the error from the provider", func() {
			So(response.Error.Message, ShouldEqual, challengeError.Message)
		})
	})
}

func TestServerAuthenticate(t *testing.T) {
	authenticateError := plugin.NewBasicError(errors.New(test.MustGenerate(reflect.TypeOf(""), t).Interface().(string)))
	provider := NewTestProvider(plugin.BasicError{}, *authenticateError, plugin.BasicError{})
	server := &Server{provider: provider}
	subject := test.MustGenerate(reflect.TypeOf(""), t).Interface().(string)
	configuration := map[string]interface{}{"subject": subject}

	var configResponse plugin.BasicError
	server.Configure(configuration, &configResponse)

	Convey("When .Authenticate is invoked", t, func() {
		var responses []Response
		var subjects [][]string
		var response Subject
		args := map[string]interface{}{"responses": responses, "subjects": subjects}

		server.Authenticate(args, &response)

		Convey("it returns the configured subject", func() {
			So(response.Value, ShouldEqual, subject)
		})

		Convey("it returns the error from the provider", func() {
			So(response.Error.Message, ShouldEqual, authenticateError.Message)
		})
	})
}
