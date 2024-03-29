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

package test

import (
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
	"time"
)

var (
	random = rand.New(rand.NewSource(time.Now().Unix()))
)

// MustGenerate generates a random value for testing or fails the test.
func MustGenerate(tipe reflect.Type, test *testing.T) reflect.Value {
	value, ok := quick.Value(tipe, random)
	if !ok {
		test.Errorf("failed to create random value of type [%s]", tipe)
	}
	return value
}
