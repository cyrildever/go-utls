package model_test

import (
	"testing"

	"github.com/cyrildever/go-utls/model"
	"gotest.tools/assert"
)

// TestIsBase64String ...
func TestIsBase64String(t *testing.T) {
	str := "BLOyg4JUDKU9HwYGwnFA+3/0pcLYgwDUUSQ14Wz4biai0oHpvCd2+dqDfVSrpuRkRM4GVc/vU4fNelJDueAlLDZty2qOwcf3uqAzM6FrQOm8uHMpMNJ1qjpk6sqnx6TTiI+UmbIQfMhVsahPfxPU3zjFjKz9AFER4Y4Z1Shkwr7kRtsg9jf4BNT2vjxVGLfcx5jA2nwrW1QsC8ZXrqNfTRQ="
	assert.Assert(t, model.IsBase64String(str))

	wrongStr := "Z^$x."
	assert.Assert(t, model.IsBase64String(wrongStr) == false)
}
