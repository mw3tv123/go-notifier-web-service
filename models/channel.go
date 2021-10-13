package models

import (
	"context"

	"github.com/mw3tv123/go-notify/forms"
)

// Channel represents a channel where message/alert is sent into. This channel can be MS Teams, Telegram or even mail.
type Channel interface {
	// Send accepts forms.MessageForm and sends them to all previously specified channels.
	// Message content supports html as markup language.
	Send(ctx context.Context, form forms.MessageForm) error
}
