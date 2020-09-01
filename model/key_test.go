package model_test

import (
	"testing"

	"github.com/cyrildever/go-utls/model"
	"gotest.tools/assert"
)

// TestBytesKey ...
func TestKeyBytes(t *testing.T) {
	key := model.Key("04aa01ccc5dd17d257a8b8cbfe371f4aaa7e4f5af153bba5aa472dc9fa4f84242d0ac607aac9f8b97ff34f5df4cba7bdf202aeff3ba032e895ecf32d2443e873ce")
	bytes := []byte{4, 170, 1, 204, 197, 221, 23, 210, 87, 168, 184, 203, 254, 55, 31, 74, 170, 126, 79, 90, 241, 83, 187, 165, 170, 71, 45, 201, 250, 79, 132, 36, 45, 10, 198, 7, 170, 201, 248, 185, 127, 243, 79, 93, 244, 203, 167, 189, 242, 2, 174, 255, 59, 160, 50, 232, 149, 236, 243, 45, 36, 67, 232, 115, 206}
	assert.DeepEqual(t, bytes, key.Bytes())
}

// TestToKey ...
func TestToKey(t *testing.T) {
	ref := model.Key("04aa01ccc5dd17d257a8b8cbfe371f4aaa7e4f5af153bba5aa472dc9fa4f84242d0ac607aac9f8b97ff34f5df4cba7bdf202aeff3ba032e895ecf32d2443e873ce")
	bytes := []byte{4, 170, 1, 204, 197, 221, 23, 210, 87, 168, 184, 203, 254, 55, 31, 74, 170, 126, 79, 90, 241, 83, 187, 165, 170, 71, 45, 201, 250, 79, 132, 36, 45, 10, 198, 7, 170, 201, 248, 185, 127, 243, 79, 93, 244, 203, 167, 189, 242, 2, 174, 255, 59, 160, 50, 232, 149, 236, 243, 45, 36, 67, 232, 115, 206}
	key := model.ToKey(bytes)
	assert.DeepEqual(t, ref.Bytes(), key.Bytes())
}
