package bip32_test

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/cyrildever/go-utls/crypto/bip32"
)

type NewMasterKeyGoldie struct {
	Seed    []byte
	KeyID   bip32.Magic
	PrivKey *bip32.PrivateKey
}

func (goldie *NewMasterKeyGoldie) UnmarshalJSON(data []byte) error {
	var elem bip32.Goldie
	if err := json.Unmarshal(data, &elem); nil != err {
		return err
	}

	goldie.Seed, _ = hex.DecodeString(elem.Seed)
	goldie.KeyID = *bip32.MainNetPrivateKey
	for _, c := range elem.Chains {
		if "m" == c.Path {
			var err error
			goldie.PrivKey, err = bip32.ParsePrivateKey(c.ExtendedPrivateKey)

			if nil != err {
				return err
			}
		}
	}

	if nil == goldie.PrivKey {
		return errors.New("missing private key")
	}

	return nil
}

func ReadGoldenJSON(t *testing.T, name string, golden interface{}) {
	fd, err := os.Open(filepath.Join(bip32.GoldenBase, name))
	if nil != err {
		t.Fatal(err)
	}
	defer fd.Close()

	if err := json.NewDecoder(fd).Decode(golden); nil != err {
		t.Fatal(err)
	}
}
