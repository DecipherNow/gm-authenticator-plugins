package login

// Challenge represents an authentication challenge.
type Challenge struct {

	// Description is a description of the challenge.
	Description string `json:"description" mapstructure:"description" yaml:"description"`

	// Name is the name of the challenge.
	Name string `json:"name" mapstructure:"name" yaml:"name"`

	// Sensitive indicates whether the response to the challenge is sensitive.
	Sensitive bool `json:"sensitive" mapstructure:"sensitive" yaml:"sensitive"`

	// Validator is a regex for validating the response to the challenge.
	Validator string `json:"validator" mapstructure:"validator" yaml:"validator"`
}
