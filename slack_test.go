package slack

import (
	"testing"
)

func TestSlack_SendMessage(t *testing.T) {
	s := New()
	s.WebhookURL = "https://hooks.slack.com/services/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"

	message := Message{
		Attachments: []Attachment{
			{Title: "Order Status", MrkdwnIn: []string{"text"},
				Fields: []Field{
					{Title: "Active", Value: 0, Short: true},
					{Title: "Failed", Value: 0, Short: true},
				},
			},
		},
	}
	s.SendMessage(message)
}
