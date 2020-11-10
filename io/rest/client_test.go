package rest_test

import (
	"testing"

	"github.com/cyrildever/go-utls/io/rest"
	"gotest.tools/assert"
)

// TestAPIClient ...
func TestAPIClient(t *testing.T) {
	c := &rest.Client{}
	assert.Assert(t, isAPIClient(c))
}

func isAPIClient(t interface{}) bool {
	switch t.(type) {
	case rest.APIClient:
		return true
	default:
		return false
	}
}
