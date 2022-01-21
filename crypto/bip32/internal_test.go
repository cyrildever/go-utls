package bip32

import (
	"bytes"
	"encoding/hex"
	"reflect"
	"testing"

	"github.com/sammyne/base58"
)

func Test_decodePublicKey(t *testing.T) {
	type expect struct {
		pub *PublicKey
		err error
	}
	testCases := []struct {
		data   string
		expect expect
	}{
		{ // no error
			"xpub661MyMwAqRbcFtXgS5sYJABqqG9YLmC4Q1Rdap9gSE8NqtwybGhePY2gZ29ESFjqJoCu1Rupje8YtGqsefD265TMg7usUDFdp6W1EGMcet8",
			expect{
				&PublicKey{
					ChainCode: []byte{
						0x87, 0x3d, 0xff, 0x81, 0xc0, 0x2f, 0x52, 0x56,
						0x23, 0xfd, 0x1f, 0xe5, 0x16, 0x7e, 0xac, 0x3a,
						0x55, 0xa0, 0x49, 0xde, 0x3d, 0x31, 0x4b, 0xb4,
						0x2e, 0xe2, 0x27, 0xff, 0xed, 0x37, 0xd5, 0x8,
					},
					ChildIndex: 0x0,
					Data: []byte{
						0x3, 0x39, 0xa3, 0x60, 0x13, 0x30, 0x15, 0x97,
						0xda, 0xef, 0x41, 0xfb, 0xe5, 0x93, 0xa0, 0x2c,
						0xc5, 0x13, 0xd0, 0xb5, 0x55, 0x27, 0xec, 0x2d,
						0xf1, 0x5, 0xe, 0x2e, 0x8f, 0xf4, 0x9c, 0x85,
						0xc2,
					},
					Level:    0x0,
					ParentFP: []byte{0x0, 0x0, 0x0, 0x0},
					Version:  []byte{0x4, 0x88, 0xb2, 0x1e},
				},
				nil,
			},
		},
		{ // base58 decoding failure: last 8->9
			"xpub661MyMwAqRbcFtXgS5sYJABqqG9YLmC4Q1Rdap9gSE8NqtwybGhePY2gZ29ESFjqJoCu1Rupje8YtGqsefD265TMg7usUDFdp6W1EGMcet9",
			expect{nil, base58.ErrChecksum},
		},
		{ // invalid key length: append a 0xff the public key data
			"5FQT7TT6bZmQ6QjZkciSR3iW58jYrY1rhLE3ozYsiUF7K4LwZQpHenGJQ2TxRaL3LJU44DYwWYtx9hCtKjJviZDe3oQfLFfWMm75bUsH21iGB5AmT",
			expect{nil, ErrInvalidKeyLen},
		},
	}

	for i, c := range testCases {
		pub, err := decodePublicKey(c.data)

		if err != c.expect.err {
			t.Fatalf("#%d unexpected error: got %v, expect %v", i, err, c.expect.err)
		}

		/*
			if 0 == i {
				pub.Data = append(pub.Data, 0xff)
				t.Log(pub.String())
			}*/

		if nil == err && !reflect.DeepEqual(pub, c.expect.pub) {
			t.Fatalf("#%d invalid public key: got %#v, expect %#v", i, pub,
				c.expect.pub)
		}
	}
}

func Test_derivePublicKey(t *testing.T) {
	// test vector decoded from the golden as specified in testdata/bip32.golden
	testCases := []struct {
		priv   string
		expect string
	}{
		{
			"e8f32e723decf4051aefac8e2c93c9c5b214313817cdb01a1494b917c8436b35",
			"0339a36013301597daef41fbe593a02cc513d0b55527ec2df1050e2e8ff49c85c2",
		},
		{
			"edb2e14f9ee77d26dd93b4ecede8d16ed408ce149b6cd80b0715a2d911a0afea",
			"035a784662a4a20a65bf6aab9ae98a6c068a81c52e4b032c0fb5400c706cfccc56",
		},
		{
			"3c6cb8d0f6a264c91ea8b5030fadaa8e538b020f0a387421a12de9319dc93368",
			"03501e454bf00751f24b1b489aa925215d66af2234e3891c3b21a52bedb3cd711c",
		},
		{
			"cbce0d719ecf7431d88e6a89fa1483e02e35092af60c042b1df2ff59fa424dca",
			"0357bfe1e341d01c69fe5654309956cbea516822fba8a601743a012a7896ee8dc2",
		},
		{
			"0f479245fb19a38a1954c5c7c0ebab2f9bdfd96a17563ef28a6a4b1a2a764ef4",
			"02e8445082a72f29b75ca48748a914df60622a609cacfce8ed0e35804560741d29",
		},
		{
			"471b76e389e528d6de6d816857e012c5455051cad6660850e58372a6c3e6e7c8",
			"022a471424da5e657499d1ff51cb43c47481a03b1e77f951fe64cec9f5a48f7011",
		},
		{
			"4b03d6fc340455b363f51020ad3ecca4f0850280cf436c70c727923f6db46c3e",
			"03cbcaa9c98c877a26977d00825c956a238e8dddfbd322cce4f74b0b5bd6ace4a7",
		},
		{
			"abe74a98f6c7eabee0428f53798f0ab8aa1bd37873999041703c742f15ac7e1e",
			"02fc9e5af0ac8d9b3cecfe2a888e2117ba3d089d8585886c9c826b6b22a98d12ea",
		},
		{
			"877c779ad9687164e9c2f4f0f4ff0340814392330693ce95a58fe18fd52e6e93",
			"03c01e7425647bdefa82b12d9bad5e3e6865bee0502694b94ca58b666abc0a5c3b",
		},
		{
			"704addf544a06e5ee4bea37098463c23613da32020d604506da8c0518e1da4b7",
			"03a7d1d856deb74c508e05031f9895dab54626251b3806e16b4bd12e781a7df5b9",
		},
		{
			"f1c7c871a54a804afe328b4c83a1c33b8e5ff48f5087273f04efa83b247d6a2d",
			"02d2b36900396c9282fa14628566582f206a5dd0bcc8d5e892611806cafb0301f0",
		},
		{
			"bb7d39bdb83ecf58f2fd82b6d918341cbef428661ef01ab97c28a4842125ac23",
			"024d902e1a2fc7a8755ab5b694c575fce742c48d9ff192e63df5193e4c7afe1f9c",
		},
		{
			"00ddb80b067e0d4993197fe10f2657a844a384589847602d56f0c629c81aae32",
			"03683af1ba5743bdfc798cf814efeeab2735ec52d95eced528e692b8e34c4e5669",
		},
		{
			"491f7a2eebc7b57028e0d3faa0acda02e75c33b03c48fb288c41e2ea44e1daef",
			"026557fdda1d5d43d79611f784780471f086d58e8126b8c40acb82272a7712e7f2",
		},
	}

	for i, c := range testCases {
		priv, _ := hex.DecodeString(c.priv)
		expect, _ := hex.DecodeString(c.expect)

		if got := derivePublicKey(priv); !bytes.Equal(got, expect) {
			t.Fatalf("#%d failed: got %x, expect %x", i, got, expect)
		}
	}
}
