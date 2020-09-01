package crypto

import (
	"errors"
	"strconv"
	"strings"
)

// Path is the string representation of a BIP32 path, eg. "m/0'/0/0"
type Path string

// Indices ...
type Indices struct {
	Account  Account
	Scope    uint32
	KeyIndex uint32
}

// Account ...
type Account struct {
	Number   uint32
	Hardened bool
}

// Parse ...
func (p Path) Parse() (i Indices, err error) {
	if string(p) == "" {
		err = errors.New("invalid empty path")
		return
	}
	parts := strings.Split(string(p), "/")
	if len(parts) != 4 || parts[0] != "m" {
		err = errors.New("not a valid path")
		return
	}
	isHardened := false
	if strings.Contains(parts[1], "'") {
		isHardened = true
	}
	number := strings.Replace(parts[1], "'", "", 1)
	nb, err := strconv.ParseInt(number, 10, 32)
	if err != nil {
		return
	}
	account := Account{
		Number:   uint32(nb),
		Hardened: isHardened,
	}
	scope, err := strconv.ParseInt(parts[2], 10, 32)
	if err != nil {
		return
	}
	keyIndex, err := strconv.ParseInt(parts[2], 10, 32)
	if err != nil {
		return
	}
	i = Indices{
		Account:  account,
		Scope:    uint32(scope),
		KeyIndex: uint32(keyIndex),
	}
	return
}
