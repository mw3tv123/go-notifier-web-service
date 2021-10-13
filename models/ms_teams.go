package models

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	msTeams "github.com/atc0005/go-teams-notify/v2"
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
func (m MSTeamsService) parseTemplate(card *msTeams.MessageCard, form forms.MessageForm) {
	if form.Contents["title"] != nil {
		card.Title = form.Contents["title"].(string)
	}

	for i := range card.Sections[0].Facts {
		field := strings.ToLower(strings.ReplaceAll(card.Sections[0].Facts[i].Name, " ", "_"))
		if form.Contents[field] != "" {
			card.Sections[0].Facts[i].Value = fmt.Sprintf("%v", form.Contents[field])
		}
	}
}

// generateMessage generates a Message Card from the request form.
// Message Card maybe generated from normal MS Teams Message Card or parse into a pre define template.
// Template is a template file.
// TODO: template store in memory, database,...
func (m MSTeamsService) generateMessage(form forms.MessageForm) (msgCard *msTeams.MessageCard, err error) {
	// From pre defined template file
	if form.Template != "" {
		msgCard, err = m.loadTemplate(fmt.Sprintf("templates/msteams_%s.json", form.Template))
		if err != nil {
			return nil, err
		}

		m.parseTemplate(msgCard, form)

	} else { // Normal Message Card
		msgCard = &msTeams.MessageCard{
			Type:    "MessageCard",
			Context: "https://schema.org/extensions",
			Title:   "Message from Notify system",
		}

		if form.Contents["title"] != nil {
			msgCard.Title = form.Contents["title"].(string)
		}

		for k, v := range form.Contents {
			msgCard.Text += fmt.Sprintf("%s: %v\n", k, v)
		}
	}

	return msgCard, nil
}

// Send accepts a subject and a message body and sends them to all previously specified channels. Message body supports
// html as markup language.
func (m MSTeamsService) Send(ctx context.Context, form forms.MessageForm) error {
	msgCard, err := m.generateMessage(form)
	if err != nil {
		return err
	}

	for _, webHook := range m.webHooks {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			err = m.client.SendWithContext(ctx, webHook, *msgCard)
			if err != nil {
				return errors.Wrapf(err, "Failed to send message to Microsoft Teams via webhook '%s'", webHook)
			}
		}
	}

	return nil
}
