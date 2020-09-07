package crypto_test

import (
	"testing"
	"time"

	"github.com/cyrildever/go-utls/common/utils"
	"github.com/cyrildever/go-utls/crypto"
	"github.com/cyrildever/go-utls/model"
	"gotest.tools/assert"
)

// TestHash ...
func TestHash(t *testing.T) {
	ref1 := model.Hash("0ac95c9bd0a75c835a03ae3a03f1d23726916e72645507cfb6b6d74a253c69a7")
	hash1, _ := crypto.Hash([]byte("Edgewhere"))
	assert.DeepEqual(t, ref1.Bytes(), hash1)

	ref2 := model.Hash("68403eeafc1431eb85aabd250480752b3950f50463023a42f4d4d3b301544451")
	hash2, _ := crypto.Hash(model.Hash("e5dcc2bdf7310ea3abc82a3580621bd0a487c181113d889d0775d3f8eba21e84").Bytes())
	assert.DeepEqual(t, ref2.Bytes(), hash2)
}

// TestIsHashedValue ...
func TestIsHashedValue(t *testing.T) {
	ok := model.Hash("0ac95c9bd0a75c835a03ae3a03f1d23726916e72645507cfb6b6d74a253c69a7")
	assert.Assert(t, crypto.IsHashedValue(ok.String()))

	wrong := "not_a_hash"
	assert.Assert(t, crypto.IsHashedValue(wrong) == false)
}

// TestPerformance ...
func TestPerformance(t *testing.T) {
	nbOfTimes := 1000000
	t0 := time.Now().UnixNano()
	hash, _ := crypto.Hash([]byte("Edgewhere"))
	for i := 1; i < nbOfTimes; i++ {
		hash, _ = crypto.Hash(hash)
	}
	t1 := time.Now().UnixNano()
	if nbOfTimes == 1000000 {
		assert.Assert(t, (t1-t0)/int64(time.Millisecond) < 1100)
		assert.Equal(t, utils.ToHex(hash), "aa019937660afa59ec3b302dc2970fd7a34be3eb851ed1d477b073797a3c0335")
	}
	// fmt.Printf("final hash: '%s' computed in %d ms\n", utils.ToHex(hash), (t1-t0)/int64(time.Millisecond))
	// assert.Assert(t, false)
}
