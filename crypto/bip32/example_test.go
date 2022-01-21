// Copyright (c) 2014 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package bip32_test

import (
	"crypto/rand"
	"fmt"
	"io"

	"github.com/sammyne/base58"

	"github.com/cyrildever/go-utls/crypto/bip32"
)

// This example demonstrates how to generate a cryptographically random seed
// then use it to create a new master node (extended key).
func ExampleNewMasterKey() {
	// Generate a random seed at the recommended length.
	seed := make([]byte, bip32.RecommendedSeedLen)
	if _, err := io.ReadFull(rand.Reader, seed); nil != err {
		fmt.Println(err)
		return
	}

	// Generate a new master node using the seed.
	_, err := bip32.NewMasterKey(seed, *bip32.MainNetPrivateKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	// then key is the private key ready to use

	// Output:
	//
}

// This example demonstrates the default hierarchical deterministic wallet
// layout as described in BIP0032.
func Example_defaultWalletLayout() {
	// The default wallet layout described in BIP0032 is:
	//
	// Each account is composed of two keypair chains: an internal and an
	// external one. The external keychain is used to generate new public
	// addresses, while the internal keychain is used for all other
	// operations (change addresses, generation addresses, ..., anything
	// that doesn't need to be communicated).
	//
	//   * m/iH/0/k
	//     corresponds to the k'th keypair of the external chain of account
	//     number i of the HDW derived from master m.
	//   * m/iH/1/k
	//     corresponds to the k'th keypair of the internal chain of account
	//     number i of the HDW derived from master m.

	// Ordinarily this would either be read from some encrypted source
	// and be decrypted or generated as the NewMasterKey example shows, but
	// for the purposes of this example, the private extended key for the
	// master node is being hard coded here.
	// This is a base58-check encoded private key
	master58 := "xprv9s21ZrQH143K3QTDL4LXw2F7HEK3wJUD2nW2nRk4stbPy6cq3jP" +
		"PqjiChkVvvNKmPGJxWUtg6LnF5kejMRNNU3TGtRBeJgk33yuGBxrMPHi"

	// Start by getting an extended key instance for the master node.
	// This gives the path:
	//   m
	masterKey, err := bip32.ParsePrivateKey(master58)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Derive the extended key for account 0.  This gives the path:
	//   m/0H
	acct0, err := masterKey.Child(bip32.HardenIndex(0))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Derive the extended key for the account 0 external chain.  This
	// gives the path:
	//   m/0H/0
	acct0Ext, err := acct0.Child(0)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Derive the extended key for the account 0 internal chain.  This gives
	// the path:
	//   m/0H/1
	acct0Int, err := acct0.Child(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	// At this point, acct0Ext and acct0Int are ready to derive the keys for
	// the external and internal wallet chains.

	// Derive the 10th extended key for the account 0 external chain.  This
	// gives the path:
	//   m/0H/0/10
	acct0Ext10, err := acct0Ext.Child(10)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Derive the 0th extended key for the account 0 internal chain.  This
	// gives the path:
	//   m/0H/1/0
	acct0Int0, err := acct0Int.Child(0)
	if err != nil {
		fmt.Println(err)
		return
	}

	const mainNetID = 0x00
	// Get and show the address associated with the extended keys for the
	// main bitcoin	network.
	acct0ExtAddr := base58.CheckEncode(acct0Ext10.AddressPubKeyHash(), mainNetID)
	acct0IntAddr := base58.CheckEncode(acct0Int0.AddressPubKeyHash(), mainNetID)

	const format = `
Account %d %s Address %d
  - address: %s
  - depth: %d
  - hardened: %v
  - for main net: %v
`

	fmt.Printf(format, 0, "External", acct0Ext10.Index(), acct0ExtAddr,
		acct0Ext10.Depth(), acct0Ext10.Hardened(),
		acct0Ext10.IsForNet(*bip32.MainNetPrivateKey))
	fmt.Printf(format, 0, "Internal", acct0Int0.Index(), acct0IntAddr,
		acct0Int0.Depth(), acct0Int0.Hardened(),
		acct0Int0.IsForNet(*bip32.MainNetPrivateKey))

	// Output:
	// Account 0 External Address 10
	//   - address: 1HVccubUT8iKTapMJ5AnNA4sLRN27xzQ4F
	//   - depth: 3
	//   - hardened: false
	//   - for main net: true
	//
	// Account 0 Internal Address 0
	//   - address: 1J5rebbkQaunJTUoNVREDbeB49DqMNFFXk
	//   - depth: 3
	//   - hardened: false
	//   - for main net: true
}

// This example demonstrates the audits use case in BIP0032.
func Example_audits() {
	// The audits use case described in BIP0032 is:
	//
	// In case an auditor needs full access to the list of incoming and
	// outgoing payments, one can share all account public extended keys.
	// This will allow the auditor to see all transactions from and to the
	// wallet, in all accounts, but not a single secret key.
	//
	//   * N(m/*)
	//   corresponds to the neutered master extended key (also called
	//   the master public extended key)

	// Ordinarily this would either be read from some encrypted source
	// and be decrypted or generated as the NewMasterKey example shows, but
	// for the purposes of this example, the private extended key for the
	// master node is being hard coded here.
	master := "xprv9s21ZrQH143K3QTDL4LXw2F7HEK3wJUD2nW2nRk4stbPy6cq3jP" +
		"PqjiChkVvvNKmPGJxWUtg6LnF5kejMRNNU3TGtRBeJgk33yuGBxrMPHi"

	// Start by getting an extended key instance for the master node.
	// This gives the path:
	//   m
	masterKey, err := bip32.ParsePrivateKey(master)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Neuter the master key to generate a master public extended key.  This
	// gives the path:
	//   N(m/*)
	masterPubKey, err := masterKey.Neuter()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Share the master public extended key with the auditor.
	fmt.Println("Audit key N(m/*):", masterPubKey)

	// Output:
	// Audit key N(m/*): xpub661MyMwAqRbcFtXgS5sYJABqqG9YLmC4Q1Rdap9gSE8NqtwybGhePY2gZ29ESFjqJoCu1Rupje8YtGqsefD265TMg7usUDFdp6W1EGMcet8
}

// this example demonstrates the conversion from the extended private key to
// to its corresponding secp256k1 private key
func ExamplePrivateKey_ToECPrivate() {
	xprv, _ := bip32.ParsePrivateKey("xprv9s21ZrQH143K3QTDL4LXw2F7HEK3wJUD2nW2nRk4stbPy6cq3jPPqjiChkVvvNKmPGJxWUtg6LnF5kejMRNNU3TGtRBeJgk33yuGBxrMPHi")

	priv := xprv.ToECPrivate()

	fmt.Println("X:", priv.X)
	fmt.Println("Y:", priv.Y)
	fmt.Println("D:", priv.D)

	// Output:
	// X: 26070491525757212273923430130929023832023083399656255554599993618152067728834
	// Y: 27475340966630338619946172401610299714452031745281566289223681642426902078081
	// D: 105366245268346348601399826821003822098691517983742654654633135381666943167285
}

// this example demonstrates the derivation of child for public key
func ExamplePublicKey_walletLayout() {
	master, _ := bip32.ParsePublicKey(
		"xpub661MyMwAqRbcFW31YEwpkMuc5THy2PSt5bDMsktWQcFF8syAmRUapSCGu8ED9W6oDMS" +
			"gv6Zz8idoc4a6mr8BDzTJY47LJhkJ8UB7WEGuduB")

	const format = `
----
  path: %s
  key: %s
  depth: %d
  hardened: %v,
  for main net: %v
`

	// m
	fmt.Printf(format, "m", master.String(), master.Depth(),
		master.Hardened(), master.IsForNet(*bip32.MainNetPublicKey))

	// m/0
	c0, err := master.Child(0)
	if nil != err {
		fmt.Println(err)
		return
	}
	fmt.Printf(format, "m/0", c0.String(), c0.Depth(),
		c0.Hardened(), c0.IsForNet(*bip32.MainNetPublicKey))

	// m/0/2147483647H
	c0c2147483647H, err := c0.Child(bip32.HardenIndex(2147483647))
	if nil != err {
		fmt.Printf("\n----\n%v\n", err)
		return
	}
	fmt.Printf(format, "m/0/2147483647H", c0c2147483647H.String(),
		c0c2147483647H.Depth(), c0c2147483647H.Hardened(),
		c0c2147483647H.IsForNet(*bip32.MainNetPublicKey))

	// Output:
	// ----
	//   path: m
	//   key: xpub661MyMwAqRbcFW31YEwpkMuc5THy2PSt5bDMsktWQcFF8syAmRUapSCGu8ED9W6oDMSgv6Zz8idoc4a6mr8BDzTJY47LJhkJ8UB7WEGuduB
	//   depth: 0
	//   hardened: false,
	//   for main net: true
	//
	// ----
	//   path: m/0
	//   key: xpub69H7F5d8KSRgmmdJg2KhpAK8SR3DjMwAdkxj3ZuxV27CprR9LgpeyGmXUbC6wb7ERfvrnKZjXoUmmDznezpbZb7ap6r1D3tgFxHmwMkQTPH
	//   depth: 1
	//   hardened: false,
	//   for main net: true
	//
	// ----
	// cannot derive a hardened key from a public key
}
