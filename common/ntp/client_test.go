package ntp_test

import (
	"math"
	"testing"
	"time"

	"github.com/cyrildever/go-utls/common/ntp"
	"gotest.tools/assert"
)

// TestNTPClient ...
func TestNTPClient(t *testing.T) {
	var limit float64 = 500

	_, err := ntp.Time("")
	assert.Error(t, err, "NTP wasn't initialized")

	err = ntp.Initialize("pool.ntp.org", limit)
	if err != nil {
		t.Fatal(err)
	}
	t0 := time.Now()
	t1, err := ntp.Time("")
	if err != nil {
		t.Fatal(err)
	}
	assert.Assert(t, math.Abs(float64(t0.UnixNano()/1e6)-float64(t1.UnixNano()/1e6)) < limit)
}
