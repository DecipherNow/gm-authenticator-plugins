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
