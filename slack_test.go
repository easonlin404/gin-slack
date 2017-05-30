package slack

import (
	"testing"
)

func TestSlack_SendMessage(t *testing.T) {
	s := New()
	s.WebhookURL = "https://hooks.slack.com/services/xxxxxxx/xxxxx/xxxxn" //your slack webhook

	message := Message{
		Text:        "just text",
		Attachments: nil,
	}
	s.SendMessage(message)
}
