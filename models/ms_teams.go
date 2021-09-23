package models

import (
	"context"
	"encoding/json"
	"strconv"

	msTeams "github.com/atc0005/go-teams-notify/v2"
	"github.com/mw3tv123/go-notify/forms"
	"github.com/mw3tv123/go-notify/utils"
	"github.com/pkg/errors"
)

// MSTeamsService struct holds necessary data to communicate with the MSTeams API.
type MSTeamsService struct {
	client        msTeams.API
	webHooks      []string
	templatePaths []string
}

// NewMSTeamsService returns a new instance of a MSTeams notification service.
// For more information about MSTeams api token:
func NewMSTeamsService() *MSTeamsService {
	client := msTeams.NewClient()

	m := &MSTeamsService{
		client:        client,
		webHooks:      []string{},
		templatePaths: []string{""},
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

// loadTemplate load template from JSON file
func (m MSTeamsService) loadTemplate(path string) (card *msTeams.MessageCard, err error) {
	tpl, err := utils.LoadJSON(path)
	if err != nil {
		return nil, errors.Wrapf(err, "Fail to load template file: ")
	}

	err = json.Unmarshal(tpl, &card)

	return card, err
}

// parseTemplate simple replace field from template with value from request
func (m MSTeamsService) parseTemplate(card *msTeams.MessageCard, form forms.CreateMSTeamsAlertForm) {
	card.Sections[0].ActivityTitle = form.Title
	for i, fact := range card.Sections[0].Facts {
		switch fact.Name {
		case "Monitor Name":
			card.Sections[0].Facts[i].Value = form.MonitorName
		case "Description":
			card.Sections[0].Facts[i].Value = form.Description
		case "Critical Level":
			card.Sections[0].Facts[i].Value = strconv.Itoa(form.Priority)
		case "Created On":
			card.Sections[0].Facts[i].Value = form.CreateDate.String()
		default:
			// Omit other fields
		}
	}
}

// GenerateAlertCard generates a Message Card from the request alert form
func (m MSTeamsService) GenerateAlertCard(alertForm forms.CreateMSTeamsAlertForm) (*msTeams.MessageCard, error) {
	alertCard, err := m.loadTemplate("templates/alert.json")
	if err != nil {
		alertCard, err = m.loadTemplate("../templates/alert.json")
		if err != nil {
			return nil, err
		}
	}

	m.parseTemplate(alertCard, alertForm)

	return alertCard, nil
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
