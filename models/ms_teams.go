package models

import (
	"context"
	"fmt"

	msTeams "github.com/atc0005/go-teams-notify/v2"
	"github.com/mw3tv123/go-notify/forms"
	"github.com/pkg/errors"
)

// MSTeamsService struct holds necessary data to communicate with the MSTeams API.
type MSTeamsService struct {
	client   msTeams.API
	webHooks []string
}

// NewMSTeamsService returns a new instance of a MSTeams notification service.
// For more information about MSTeams api token:
func NewMSTeamsService() *MSTeamsService {
	client := msTeams.NewClient()

	m := &MSTeamsService{
		client:   client,
		webHooks: []string{},
	}

	return m
}

// DisableWebhookValidation disables the validation of webhook URLs, including the validation of known prefixes so that
// custom/private webhook URL endpoints can be used (e.g., testing purposes).
func (m *MSTeamsService) DisableWebhookValidation() {
	m.client.SkipWebhookURLValidationOnSend(true)
}

// AddReceivers takes MSTeams channel web-hooks and adds them to the internal web-hook list. The Send method will send
// a given message to all those chats.
func (m *MSTeamsService) AddReceivers(webHooks ...string) {
	m.webHooks = append(m.webHooks, webHooks...)
}

// AddImageCardSection add an image url to message card section
func (m MSTeamsService) AddImageCardSection(imgTitle, imgUrl string) msTeams.MessageCardSectionImage {
	img := msTeams.NewMessageCardSectionImage()

	img.Title = imgTitle
	img.Image = imgUrl

	return img
}

// CreateFactsFromList return a list of MessageCardSectionFact from the given key/value
func (m MSTeamsService) CreateFactsFromList(kv map[string]string) (facts []msTeams.MessageCardSectionFact) {
	for k, v := range kv {
		fact := msTeams.NewMessageCardSectionFact()
		fact.Name = k
		fact.Value = v

		facts = append(facts, fact)
	}
	return facts
}

// GenerateAlertCard generates a Message Card from the request alert form
func (m MSTeamsService) GenerateAlertCard(alertForm forms.CreateMSTeamsAlertForm) (msTeams.MessageCard, error) {
	alertCard := msTeams.NewMessageCard()
	alertCard.Summary = "Alert from Notify Web Service"

	// Add message title
	msgCardSection := msTeams.NewMessageCardSection()
	msgCardSection.Title = fmt.Sprintf("**%s**", alertForm.Title)

	// Add message body
	err := msgCardSection.AddFact(m.CreateFactsFromList(map[string]string{"Monitor Name": alertForm.MonitorName, "Description": alertForm.Description, "Critical Level": string(rune(alertForm.Priority)), "Submitted On": alertForm.CreateDate.String()})...)
	if err != nil {
		return alertCard, err
	}

	err = alertCard.AddSection(msgCardSection)

	return alertCard, err
}

// SendMessage accepts a subject and a message body and sends them to all previously specified channels. Message body
// supports html as markup language.
func (m MSTeamsService) SendMessage(ctx context.Context, subject, message string) error {
	msgCard := msTeams.NewMessageCard()
	msgCard.Title = subject
	msgCard.Text = message

	return m.Send(ctx, msgCard)
}

// Send accepts a subject and a message body and sends them to all previously specified channels. Message body supports
// html as markup language.
func (m MSTeamsService) Send(ctx context.Context, msgCard msTeams.MessageCard) error {
	for _, webHook := range m.webHooks {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			err := m.client.SendWithContext(ctx, webHook, msgCard)
			if err != nil {
				return errors.Wrapf(err, "Failed to send message to Microsoft Teams via webhook '%s'", webHook)
			}
		}
	}

	return nil
}
