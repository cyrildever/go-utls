package bip32

// HardenIndex maps an index in range [0,HardenedKeyStart) to
// its harhened corresponding one.
// Note: No overflow checking is implemented now.
func HardenIndex(i uint32) uint32 {
	/* if i > math.MaxUint32-HardenedKeyStart { } */
	return i + HardenedKeyStart
}

// Zero clears all data items to 0
func Zero(data []byte) {
	for i := range data {
		data[i] = 0
	}
}
