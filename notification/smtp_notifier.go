package notification

import (
	"fmt"
	"net/smtp"

	"github.com/clems4ever/authelia/configuration/schema"
)

// SMTPNotifier a notifier to send emails to SMTP servers.
type SMTPNotifier struct {
	address  string
	sender   string
	username string
	password string
	host     string
	port     int
}

// NewSMTPNotifier create an SMTPNotifier targeting a given address.
func NewSMTPNotifier(configuration schema.SMTPNotifierConfiguration) *SMTPNotifier {
	return &SMTPNotifier{
		host:     configuration.Host,
		port:     configuration.Port,
		address:  fmt.Sprintf("%s:%d", configuration.Host, configuration.Port),
		sender:   configuration.Sender,
		username: configuration.Username,
		password: configuration.Password,
	}
}

// Send send a identity verification link to a user.
func (n *SMTPNotifier) Send(recipient string, subject string, body string) error {
	msg := "From: " + n.sender + "\n" +
		"To: " + recipient + "\n" +
		"Subject: " + subject + "\n" +
		"Content-Type: text/html\n" +
		"MIME-Version: 1.0\n" +
		"Content-Transfer-Encoding: quoted-printable\n\n" +
		body

	// Connect to the remote SMTP server.
	c, err := smtp.Dial(n.address)
	if err != nil {
		return err
	}

	// Set the sender and recipient first
	if err := c.Mail(n.sender); err != nil {
		return err
	}

	if err := c.Rcpt(recipient); err != nil {
		return err
	}

	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(wc, msg)
	if err != nil {
		return err
	}
	err = wc.Close()
	if err != nil {
		return err
	}

	// Send the QUIT command and close the connection.
	err = c.Quit()
	if err != nil {
		return err
	}
	return nil
}
