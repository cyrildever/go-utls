package utils_test

import (
	"math/big"
	"testing"

	"github.com/cyrildever/go-utls/common/utils"
	"gotest.tools/assert"
)

// TestEuclideanDivision ...
func TestEuclideanDivision(t *testing.T) {
	numerator := 13
	denominator := 2
	quotient, remainder, _ := utils.EuclideanDivision(numerator, denominator)
	assert.Equal(t, quotient, 6)
	assert.Equal(t, remainder, 1)

	denominator = 0
	_, _, err := utils.EuclideanDivision(numerator, denominator)
	assert.Error(t, err, "division by zero")
}

// TestFindClosest ...
func TestFindClosest(t *testing.T) {
	bi1 := big.NewInt(25)
	bi2 := big.NewInt(30)
	bi3 := big.NewInt(20)
	arr := []*big.Int{bi1, bi2, bi3}

	under := big.NewInt(10)
	closest, _ := utils.FindClosest(arr, under)
	assert.Assert(t, closest.Cmp(bi3) == 0)

	over := big.NewInt(50)
	closest, _ = utils.FindClosest(arr, over)
	assert.Assert(t, closest.Cmp(bi2) == 0)

	middle1 := big.NewInt(27)
	closest, _ = utils.FindClosest(arr, middle1)
	assert.Assert(t, closest.Cmp(bi1) == 0)

	middle2 := big.NewInt(28)
	closest, _ = utils.FindClosest(arr, middle2)
	assert.Assert(t, closest.Cmp(bi2) == 0)

	middle3 := big.NewInt(22)
	closest, _ = utils.FindClosest(arr, middle3)
	assert.Assert(t, closest.Cmp(bi3) == 0)

	bi4 := big.NewInt(23)
	arr = append(arr, bi4)
	closest, _ = utils.FindClosest(arr, middle3)
	assert.Assert(t, closest.Cmp(bi4) == 0)
}
