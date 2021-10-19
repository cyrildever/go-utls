package email_test

import (
	"testing"

	"github.com/cyrildever/go-utls/common/email"
	"gotest.tools/assert"
)

// TestEmail ...
func TestEmail(t *testing.T) {
	ref := "Cyril Dever <cdever@edgewhere.fr>"

	mail := email.NewEmail("cdever@edgewhere.fr", "Cyril Dever")
	assert.Equal(t, mail.String(), ref)
}

// NB: To make it work, turn on 'Autoriser les applications moins sécurisées' in Google account, uncomment all and fill in the password field
// func TestSend(t *testing.T) {
// 	username := "cdever@edgewhere.fr"
// 	password := ""

// 	to := []email.Email{
// 		{
// 			Address: "support@edgewhere.fr",
// 			Name:    "Support Edgewhere",
// 		},
// 		{
// 	}
// 	client := email.NewClient(username, password, "smtp.gmail.com", 587)
// 	recipients, err := client.Send("Test", "Ceci est un test", email.Email{Address: username, Name: "Cyril Dever"}, to)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	assert.Equal(t, len(recipients), len(to))
// }
