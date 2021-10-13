package models

import (
	"context"
	"net/smtp"
	"net/textproto"

	"github.com/jordan-wright/email"
	"github.com/mw3tv123/go-notify/forms"
	"github.com/pkg/errors"
)

// EmailService struct holds necessary data to communicate with the MSTeams API.
type EmailService struct {
	senderAddress     string
	smtpHostAddr      string
	smtpAuth          smtp.Auth
	receiverAddresses []string
}

// NewEmailService returns a new instance of a Mail notification service.
func NewEmailService(senderAddress, smtpHostAddress string) *EmailService {
	return &EmailService{
		senderAddress:     senderAddress,
		smtpHostAddr:      smtpHostAddress,
		receiverAddresses: []string{},
	}
}

// AuthenticateSMTP authenticates you to send emails via smtp.
// Example values: "", "test@gmail.com", "password123", "smtp.gmail.com"
// For more information about smtp authentication, see here:
//    -> https://pkg.go.dev/net/smtp#PlainAuth
func (m *EmailService) AuthenticateSMTP(identity, userName, password, host string) {
	m.smtpAuth = smtp.PlainAuth(identity, userName, password, host)
}

// AddReceivers takes email addresses and adds them to the internal address list. The Send method will send
// a given message to all those addresses.
func (m *EmailService) AddReceivers(addresses ...string) {
	m.receiverAddresses = append(m.receiverAddresses, addresses...)
}

// Send takes a message subject and a message body and sends them to all previously set chats. Message body supports
// html as markup language.
func (m EmailService) Send(ctx context.Context, form forms.MessageForm) error {
	var err error
	mail := &email.Email{
		To:   m.receiverAddresses,
		From: m.senderAddress,
		// Subject: form.Contents["title"].(string),
		// Text:    []byte("Text Body is, of course, supported!"),
		// HTML:    []byte(form.Content),
		Headers: textproto.MIMEHeader{},
	}
	select {
	case <-ctx.Done():
		err = ctx.Err()
	default:
		err = mail.Send(m.smtpHostAddr, m.smtpAuth)
		if err != nil {
			err = errors.Wrap(err, "Failed to send mail: ")
		}
	}

	return err
}
