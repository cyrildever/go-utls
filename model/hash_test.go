package model_test

import (
	"testing"

	"github.com/cyrildever/go-utls/model"
	"gotest.tools/assert"
)

// TestBytesHash ...
func TestHashBytes(t *testing.T) {
	hash := model.Hash("123e3b30d0b76c2349a0229fc70128f5709c9101b67de347216b7506efecb310")
	bytes := []byte{18, 62, 59, 48, 208, 183, 108, 35, 73, 160, 34, 159, 199, 1, 40, 245, 112, 156, 145, 1, 182, 125, 227, 71, 33, 107, 117, 6, 239, 236, 179, 16}
	assert.DeepEqual(t, bytes, hash.Bytes())
}

// TestLongString ...
func TestLongString(t *testing.T) {
	ref := "045206550e06560c0c54040f5c010151010857555b59515e58075d0a0c550f070e5004525504000f550855015d545a545a5709060504575f565f0206055a5100"
	hash := model.Hash(ref)
	assert.Equal(t, hash.String(), ref)
}

// TestToHash ...
func TestToHash(t *testing.T) {
	ref := model.Hash("123e3b30d0b76c2349a0229fc70128f5709c9101b67de347216b7506efecb310")
	bytes := []byte{18, 62, 59, 48, 208, 183, 108, 35, 73, 160, 34, 159, 199, 1, 40, 245, 112, 156, 145, 1, 182, 125, 227, 71, 33, 107, 117, 6, 239, 236, 179, 16}
	hash := model.ToHash(bytes)
	assert.DeepEqual(t, ref.Bytes(), hash.Bytes())
}

// TestEquals ...
func TestEquals(t *testing.T) {
	var hashes1, hashes2, hashes3 model.Hashes

	hash1 := model.Hash("cb0518922221d45f0a1d2bdbe73a2c1d814e59e72d58a959e39ab57ad174a0dd")
	hash2 := model.Hash("936177c069b16a8bc026edfc4cfb633943692f2fa0323d87ff772fe86fb916fe")
	hash3 := model.Hash("b02156e935c543acb42467e8b9f636a14d477635ed309f3aa58c8fd5bb3060ac")

	hashes1 = append(hashes1, hash1)
	hashes1 = append(hashes1, hash2)
	hashes1 = append(hashes1, hash3)

	hashes2 = append(hashes2, hash1)
	hashes2 = append(hashes2, hash2)
	hashes2 = append(hashes2, hash3)

	hashes3 = append(hashes3, hash3)
	hashes3 = append(hashes3, hash1)
	hashes3 = append(hashes3, hash2)

	assert.Assert(t, hashes1.Equals(&hashes2))
	assert.Assert(t, !hashes1.Equals(&hashes3))
}

// TestNonEmpty ...
func TestNonEmpty(t *testing.T) {
	empty := model.Hash("")
	assert.Assert(t, empty.NonEmpty() == false)

	hash := model.Hash("cb0518922221d45f0a1d2bdbe73a2c1d814e59e72d58a959e39ab57ad174a0dd")
	assert.Assert(t, hash.NonEmpty())
}
