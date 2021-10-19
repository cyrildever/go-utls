package email

import (
	"fmt"
	"net/smtp"
	"strings"
)

//--- TYPES

// Client ...
type Client struct {
	auth smtp.Auth
	host string
	port int
}

//--- METHODS

// Send sends an e-mail and returns the actual list of potential recipients
func (c Client) Send(subject, body string, from Email, to []Email) (recipients []string, err error) {
	for _, recipient := range to {
		if recipient.IsValid() {
			recipients = append(recipients, recipient.Address)
		}
	}

	msg := fmt.Sprintf("From: %s\r\n", from.String()) +
		fmt.Sprintf("To: %s\r\n", strings.Join(recipients, ",")) +
		fmt.Sprintf("Subject: %s\r\n", subject) +
		"Content-Type: text/plain; charset=UTF-8\r\n" +
		"\r\n" +
		fmt.Sprintf("%s\r\n", body)

	err = smtp.SendMail(fmt.Sprintf("%s:%d", c.host, c.port), c.auth, from.Address, recipients, []byte(msg))
	return
}

//--- FUNCTIONS

// NewClient ...
func NewClient(username, password, hostname string, portNumber int) *Client {
	auth := smtp.PlainAuth("", username, password, hostname)
	return &Client{
		auth: auth,
		host: hostname,
		port: portNumber,
	}
}
