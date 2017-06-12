package ginslack

import (
	"bytes"
	"fmt"
	"github.com/easonlin404/go-slack"
	"gopkg.in/gin-gonic/gin.v1"
	"log"
	"net/http/httputil"
	"runtime/debug"
)

func Recovery(s *slack.Slack) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				//debug.PrintStack()
				httprequest, _ := httputil.DumpRequest(c.Request, false)
				stack := debug.Stack()
				log.Printf("[Recovery] panic recovered:\n%s\n%s\n%s", string(httprequest), err, stack)

				message := genSlackMessage(httprequest, err, stack)

				s.SendMessage(message)
				c.AbortWithStatus(500)

			}
		}()
		c.Next()
	}
}

func genSlackMessage(httprequest []byte, err interface{}, stack []byte) slack.Message {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%s%s ```%s```", httprequest, err, stack)

	message := slack.Message{
		Attachments: []slack.Attachment{
			{Title: "notification",
				Text:     string(buf.Bytes()),
				MrkdwnIn: []string{"text"},
			},
		},
	}

	return message
}
