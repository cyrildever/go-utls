package email_test

import (
	"testing"

	"github.com/cyrildever/go-utls/common/email"
	"gotest.tools/assert"
)

// TestEmail ...
func TestEmail(t *testing.T) {
	ref := "John Doe <john@doe.com>"
	adrs := "john@doe.com"

	mail := email.NewEmail(adrs, "John Doe")
	assert.Equal(t, mail.String(), ref)

	noDisplay := email.NewEmail(adrs)
	assert.Equal(t, noDisplay.String(), adrs)
}

// TestSendAWS ...
//
// NB: To make it work, provided you are in the 'eu-west-3' region, fill in the appropriate AWS Simple Email Service SMTP credentials
// (@see https://eu-west-3.console.aws.amazon.com/ses/home?region=eu-west-3#smtp-settings:)
//
// Furthermore, while still in AWS Sandbox mode, each email must have been verified beforehand via the AWS SES console
// (@see https://eu-west-3.console.aws.amazon.com/ses/home?region=eu-west-3#verified-senders-email:)
func TestSendAWS(t *testing.T) {
	username := "" // Your AWS SES username
	password := "" // Your AWS SES password

	if username == "" || password == "" {
		assert.Assert(t, true) // To avoid breaking the tests
		t.Log("Empty credentials in TestSendAWS")
		return
	}

	to := []email.Email{
		{
			Address: "cdever@edgewhere.fr", // AWS verified e-mail
			Name:    "Cyril Dever",
		},
	}
	client := email.NewClient(username, password, "email-smtp.eu-west-3.amazonaws.com", 587) // Adapt to own region
	sender := email.NewEmail("support@edgewhere.fr", "Support Edgewhere")                    // AWS verified e-mail
	msg := "Ceci est un test en UTF-8\n" +
		"sur plusieurs lignes.\n\n" +
		"Votre équipe Edgewhere"
	recipients, err := client.Send("Test via AWS", msg, *sender, to)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, len(recipients), len(to))
}

// TestSendGmail ...
//
// NB: To make it work, turn on 'Authorize less secure applications' in your Google account, uncomment all and fill in the credentials
func TestSendGmail(t *testing.T) {
	username := "" // Your Gmail e-mail address
	password := "" // Your Gmail password

	if username == "" || password == "" {
		assert.Assert(t, true) // To avoid breaking the tests
		t.Log("Empty credentials in TestSendGmail")
		return
	}

	to := []email.Email{
		{
			Address: "support@edgewhere.fr",
			Name:    "Support Edgewhere",
		},
	}
	client := email.NewClient(username, password, "smtp.gmail.com", 587)
	recipients, err := client.Send("Test via Gmail", "Ceci est un test", email.Email{Address: username, Name: "Your Test Name"}, to)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, len(recipients), len(to))
}
