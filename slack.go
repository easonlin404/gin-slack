package slack

import (
	"github.com/easonlin404/esrest"
	"net/http"
)

type SlackClient struct {
	WebhookURL string
}

func New() *SlackClient {
	return &SlackClient{}
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

func (s *SlackClient) SendMessage(message Message) (*http.Response, error) {
	return esrest.New().
		Header("Content-Type", "text/plain").
		Debug(true).
		Post(s.WebhookURL).
		Body(message).
		Do()
}
