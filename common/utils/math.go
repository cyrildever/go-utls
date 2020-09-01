package utils

import (
	"errors"
	"math/big"
	"sort"
)

// EuclideanDivision returns the integer tuple (quotient, remainder) from the division of the past integers
func EuclideanDivision(numerator, denominator int) (quotient, remainder int, err error) {
	if denominator == 0 {
		err = errors.New("division by zero") // TODO typed error
		return
	}
	quotient = numerator / denominator
	remainder = numerator % denominator
	return
}

// FindClosest returns the closest number to the target in the passed array of numbers
// see https://www.geeksforgeeks.org/find-closest-number-array/
func FindClosest(arr []*big.Int, target *big.Int) (*big.Int, bool) {
	n := len(arr)
	if target == nil || len(arr) == 0 {
		return nil, false
	}

	sort.Slice(arr, func(i int, j int) bool {
		res := arr[i].Cmp(arr[j])
		return res == -1
	})

	// Corner cases
	if res := target.Cmp(arr[0]); res == -1 {
		return arr[0], true
	}
	if res := target.Cmp(arr[n-1]); res == 1 {
		return arr[n-1], true
	}

	// Doing binary search
	i := 0
	j := n
	mid := 0
	for i < j {
		mid := i + j/2
		if res := target.Cmp(arr[mid]); res == 0 {
			return arr[mid], true
		} else if res == -1 {
			// If target is less than array element, then search in left
			if r := target.Cmp(arr[mid-1]); mid > 0 && r == 1 {
				return getClosest(arr[mid-1], arr[mid], target), true
			}
			// Repeat for left half
			j = mid
		} else {
			// If target is greater than mid
			if r := target.Cmp(arr[mid+1]); mid < n-1 && r == -1 {
				return getClosest(arr[mid], arr[mid+1], target), true
			}
			i = mid + 1
		}
	}
	return arr[mid], true
}

func getClosest(val1 *big.Int, val2 *big.Int, target *big.Int) *big.Int {
	sub1 := new(big.Int)
	sub1.Sub(target, val1)
	sub2 := new(big.Int)
	sub2.Sub(val2, target)
	if r := sub1.Cmp(sub2); r >= 0 {
		return val2
	}
	return val1
}
