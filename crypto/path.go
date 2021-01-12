package crypto

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//--- TYPES

// Path is the string representation of a BIP32 path, eg. `m/0'/0/0`.
// NB: For compatibility reasons with browsers, the account number and the scope mustn't be greater than 2^16 - 1 and the key index not greater than 2^21 - 1.
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

//-- METHODS

// Bytes ...
func (p Path) Bytes() []byte {
	if p.IsEmpty() {
		return nil
	}
	return []byte(p)
}

// IsEmpty ...
func (p Path) IsEmpty() bool {
	return p.String() == ""
}

// IsEmpty ...
func (p Path) NonEmpty() bool {
	return !p.IsEmpty()
}

// String ...
func (p Path) String() string {
	i, err := p.Parse()
	if err != nil {
		return ""
	}
	acc := i.Account
	scop := i.Scope
	idx := i.KeyIndex
	if acc.Number > uint32(math.Pow(2, 16))-1 || scop > uint32(math.Pow(2, 16))-1 || idx > uint32(math.Pow(2, 21))-1 {
		return ""
	}
	return string(p)
}

// Next builds the next path, eventually moving up the account number and/or the scope, eg.
//	Path("m/0'/65535/2097151").Next() = Path("m/1'/0/0")
func (p Path) Next() (newPath Path, err error) {
	i, err := p.Parse()
	if err != nil {
		return
	}
	acc := i.Account
	scop := i.Scope
	idx := i.KeyIndex
	if idx < uint32(math.Pow(2, 21))-1 {
		idx++
	} else {
		idx = 0
		if scop < uint32(math.Pow(2, 16))-1 {
			scop++
		} else {
			scop = 0
			if acc.Number == uint32(math.Pow(2, 16))-1 {
				err = errors.New("account overflow: impossible to build next")
				return
			}
			acc.Number++
		}
	}
	hardened := "'"
	if !acc.Hardened {
		hardened = ""
	}
	newPath = Path(fmt.Sprintf("m/%d%s/%d/%d", acc.Number, hardened, scop, idx))

	return
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
	keyIndex, err := strconv.ParseInt(parts[3], 10, 32)
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
