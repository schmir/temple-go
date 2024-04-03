package greet_test

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/schmir/temple-go/greet"
)

func TestGreet(t *testing.T) {
	assert.Equal(t, greet.Greet("World"), "Hello World!")
}
