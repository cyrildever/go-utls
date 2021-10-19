package email

import (
	"fmt"

	"github.com/cyrildever/go-utls/normalizer"
)

//--- TYPES

// Email ...
type Email struct {
	// The actual e-mail address, eg. john@doe.com
	Address string

	// The name to display, eg. John Doe
	Name string
}

//--- METHODS

// Display returns the display name
func (e Email) Display() string {
	if e.Name != "" {
		// TODO Remove unwanted characters?
		return e.Name
	}
	return e.Address
}

// IsValid ...
func (e Email) IsValid() bool {
	return normalizer.IsValidEmail(e.Address)
}

// String ...
func (e Email) String() string {
	if e.Display() == e.Address {
		return e.Address
	}
	return fmt.Sprintf("%s <%s>", e.Display(), e.Address)
}

//--- FUNCTIONS

// NewEmail ...
func NewEmail(address, displayName string) *Email {
	return &Email{
		Address: address,
		Name:    displayName,
	}
}
