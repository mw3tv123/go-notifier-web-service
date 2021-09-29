package models

import (
	"context"

	"github.com/mw3tv123/go-notify/forms"
)

type Channel interface {
	// SendMessage accepts forms.RequestMessageForm and sends them to all previously specified channels.
	// Message content supports html as markup language.
	SendMessage(ctx context.Context, form forms.RequestMessageForm) error

	// SendAlert accepts forms.RequestAlertForm and sends them to all previously specified channels.
	// Message body supports html as markup language and other components.
	SendAlert(ctx context.Context, form forms.RequestAlertForm) error
}
