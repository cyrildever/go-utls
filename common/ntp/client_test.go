package ntp_test

import (
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/cyrildever/go-utls/common/ntp"
	"gotest.tools/assert"
)

// TestNTPClient ...
func TestNTPClient(t *testing.T) {
	var limit float64 = 999

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
	elapsed := math.Abs(float64(t0.UnixNano()/1e6) - float64(t1.UnixNano()/1e6))
	fmt.Println("elapsed", elapsed)
	assert.Assert(t, elapsed < limit)
}
