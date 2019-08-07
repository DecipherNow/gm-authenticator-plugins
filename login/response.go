package login

// Response represents an authentication response.
type Response struct {

	// Name is the name of the challenge.
	Name string `json:"name" mapstructure:"name" yaml:"name"`

	// Value is the value of the response.
	Value string `json:"value" mapstructure:"value" yaml:"value"`
}
