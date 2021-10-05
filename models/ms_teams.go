package models

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	msTeams "github.com/atc0005/go-teams-notify/v2"
	"github.com/mw3tv123/go-notify/config"
	"github.com/mw3tv123/go-notify/forms"
	"github.com/mw3tv123/go-notify/utils"
	"github.com/pkg/errors"
)

// MSTeamsService struct holds necessary data to communicate with the MSTeams API.
type MSTeamsService struct {
	client   msTeams.API
	webHooks []string
}

// NewMSTeamsService returns a new instance of a MSTeams notification service.
// For more information about MSTeams api token:
func NewMSTeamsService(wh ...string) *MSTeamsService {
	client := msTeams.NewClient()

	m := &MSTeamsService{
		client:   client,
		webHooks: wh,
	}

	return m
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
func (m MSTeamsService) parseTemplate(card *msTeams.MessageCard, form forms.RequestAlertForm) {
	card.Sections[0].ActivityTitle = form.Title
	for i, fact := range card.Sections[0].Facts {
		switch fact.Name {
		case "Service Name":
			card.Sections[0].Facts[i].Value = form.ServiceName
		case "Description":
			card.Sections[0].Facts[i].Value = form.Description
		case "Critical Level":
			card.Sections[0].Facts[i].Value = strconv.Itoa(form.Priority)
		case "Created On":
			if form.CreateDate.IsZero() {
				location, _ := time.LoadLocation(config.GetConfig("TZ"))
				card.Sections[0].Facts[i].Value = time.Now().In(location).Format(time.RFC1123Z)
			} else {
				card.Sections[0].Facts[i].Value = form.CreateDate.String()
			}
		default:
			// Omit other fields
		}
	}
}

// generateAlertCard generates a Message Card from the request alert form
func (m MSTeamsService) generateAlertCard(alertForm forms.RequestAlertForm) (*msTeams.MessageCard, error) {
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

// send accepts a subject and a message body and sends them to all previously specified channels. Message body supports
// html as markup language.
func (m MSTeamsService) send(ctx context.Context, msgCard msTeams.MessageCard) error {
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

// SendMessage sends a message (in this case is a MessageCard) into MS Teams channel with the given webhook.
// This message only contains basic part like title and contents.
func (m MSTeamsService) SendMessage(ctx context.Context, form forms.RequestMessageForm) error {
	msgCard := msTeams.NewMessageCard()
	msgCard.Title = form.Title
	msgCard.Text = form.Content

	return m.send(ctx, msgCard)
}

// SendAlert sends a MessageCard with more information field than a normal message. Message is load from a template
// json file from template directory.
func (m MSTeamsService) SendAlert(ctx context.Context, alertForm forms.RequestAlertForm) error {
	msg, err := m.generateAlertCard(alertForm)
	if err != nil {
		return err
	}

	return m.send(ctx, *msg)
}
