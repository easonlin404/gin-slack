package slack

import (
	"github.com/easonlin404/esrest"
)

//TODO: extract slack token to config
const Url = "https://hooks.slack.com/services/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"

type Slack struct {
	WebhookURL string
}

func New() *Slack {
	return &Slack{}
}

type Message struct {
	//TODO: other
	Text        string       `json:"text"`
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	Title    string   `json:"title"`
	Text     string   `json:"text"`
	MrkdwnIn []string `json:"mrkdwn_in"`
	Fields   []Field  `json:"fields"`
}

type Field struct {
	Title string `json:"title"`
	Value int    `json:"value"`
	Short bool   `json:"short"`
}

func (s *Slack) SendMessage(message Message) {
	esrest.New().
		Header("Content-Type", "text/plain").
		Debug(true).
		Post(s.WebhookURL).
		Body(message).
		Do()
}
