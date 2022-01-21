package bip32

// masterHMACKey is the key used along with a random seed used to generate
// the master key in the hierarchical tree.
var masterHMACKey = []byte("Bitcoin seed")
