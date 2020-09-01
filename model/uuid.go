package model

import (
	"regexp"

	"github.com/gofrs/uuid"
)

var uuidRegex = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)

//--- TYPES

// UUID is the string representation of a UUID, eg. 'e05572b3-230a-45fd-a779-604c2b8ceb24'
type UUID string

//--- METHODS

// Bytes ...
func (u UUID) Bytes() []byte {
	if u.String() == "" {
		return nil
	}
	return []byte(u)
}

// String ...
func (u UUID) String() string {
	if string(u) == "" || !uuidRegex.MatchString(string(u)) {
		return ""
	}
	return string(u)
}

// NonEmpty ...
func (u UUID) NonEmpty() bool {
	return u.String() != ""
}

// ToUUID ...
func ToUUID(bytes []byte) UUID {
	if bytes == nil {
		return UUID("")
	}
	return UUID(string(bytes))
}

// GenerateUUID creates a new UUID
func GenerateUUID() (id UUID, err error) {
	u, err := uuid.NewV4()
	if err != nil {
		return
	}
	id = ToUUID([]byte(u.String()))
	return
}
