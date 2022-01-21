package bip32_test

import (
	"reflect"
	"testing"

	"github.com/cyrildever/go-utls/crypto/bip32"
)

// dummy test for coverage report
func TestPath_ChildIndices(t *testing.T) {
	type expect struct {
		indices []*bip32.ChildIndex
		hasErr  bool
	}

	testCases := []struct {
		path   bip32.Path
		expect expect
	}{
		{ // good path
			"m/0/123H",
			expect{
				[]*bip32.ChildIndex{
					{Index: 0, Hardened: false},
					{Index: 123, Hardened: true},
				},
				false,
			},
		},
		{ // path contains invalid child index
			"m/0 /123H",
			expect{nil, true},
		},
	}

	for i, c := range testCases {
		got, err := c.path.ChildIndices()

		if nil != err && !c.expect.hasErr {
			t.Fatalf("#%d unexpected error: %v", i, err)
		} else if nil == err && c.expect.hasErr {
			t.Fatalf("#%d expect error but got none", i)
		}

		if nil == err && !reflect.DeepEqual(got, c.expect.indices) {
			t.Fatalf("#%d invalid child indices: got %#v, expect %#v", i, got,
				c.expect.indices)
		}
	}
}
