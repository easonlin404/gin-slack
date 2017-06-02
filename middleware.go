package slack

import (


	"github.com/gin-gonic/gin"

	"runtime/debug"
)

func Recovery(client *SlackClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				debug.PrintStack()


				message := Message{
				}
				client.SendMessage(message)
			}
		}()
		c.Next()
	}
}
