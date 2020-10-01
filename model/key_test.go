package model_test

import (
	"testing"

	"github.com/cyrildever/go-utls/model"
	"gotest.tools/assert"
)

// TestKeyBytes ...
func TestKeyBytes(t *testing.T) {
	key := model.Key("04aa01ccc5dd17d257a8b8cbfe371f4aaa7e4f5af153bba5aa472dc9fa4f84242d0ac607aac9f8b97ff34f5df4cba7bdf202aeff3ba032e895ecf32d2443e873ce")
	bytes := []byte{4, 170, 1, 204, 197, 221, 23, 210, 87, 168, 184, 203, 254, 55, 31, 74, 170, 126, 79, 90, 241, 83, 187, 165, 170, 71, 45, 201, 250, 79, 132, 36, 45, 10, 198, 7, 170, 201, 248, 185, 127, 243, 79, 93, 244, 203, 167, 189, 242, 2, 174, 255, 59, 160, 50, 232, 149, 236, 243, 45, 36, 67, 232, 115, 206}
	assert.DeepEqual(t, bytes, key.Bytes())
}

// TestKeyString ...
func TestKeyString(t *testing.T) {
	ref := "1234567890abcdef"
	key := model.Key(ref)
	assert.Equal(t, ref, key.String())

	wrong := model.Key("wrong key")
	assert.Assert(t, wrong.IsEmpty())
}

// TestToKey ...
func TestToKey(t *testing.T) {
	ref := model.Key("04aa01ccc5dd17d257a8b8cbfe371f4aaa7e4f5af153bba5aa472dc9fa4f84242d0ac607aac9f8b97ff34f5df4cba7bdf202aeff3ba032e895ecf32d2443e873ce")
	bytes := []byte{4, 170, 1, 204, 197, 221, 23, 210, 87, 168, 184, 203, 254, 55, 31, 74, 170, 126, 79, 90, 241, 83, 187, 165, 170, 71, 45, 201, 250, 79, 132, 36, 45, 10, 198, 7, 170, 201, 248, 185, 127, 243, 79, 93, 244, 203, 167, 189, 242, 2, 174, 255, 59, 160, 50, 232, 149, 236, 243, 45, 36, 67, 232, 115, 206}
	key := model.ToKey(bytes)
	assert.DeepEqual(t, ref.Bytes(), key.Bytes())
}

// TestKeysEquals ...
func TestKeysEquals(t *testing.T) {
	var keys1, keys2, keys3 model.Keys

	key1 := model.Key("04aa01ccc5dd17d257a8b8cbfe371f4aaa7e4f5af153bba5aa472dc9fa4f84242d0ac607aac9f8b97ff34f5df4cba7bdf202aeff3ba032e895ecf32d2443e873ce")
	key2 := model.Key("04e315a987bd79b9f49d3a1c8bd1ef5a401a242820d52a3f22505da81dfcd992cc5c6e2ae9bc0754856ca68652516551d46121daa37afc609036ab5754fe7a82a3")
	key3 := model.Key("04abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890")

	keys1 = append(keys1, key1)
	keys1 = append(keys1, key2)
	keys1 = append(keys1, key3)

	keys2 = append(keys2, key1)
	keys2 = append(keys2, key2)
	keys2 = append(keys2, key3)

	keys3 = append(keys3, key3)
	keys3 = append(keys3, key1)
	keys3 = append(keys3, key2)

	assert.Assert(t, keys1.Equals(&keys2))
	assert.Assert(t, !keys1.Equals(&keys3))
}
