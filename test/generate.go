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
