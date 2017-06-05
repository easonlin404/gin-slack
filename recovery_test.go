package ginslack

import (
	"github.com/easonlin404/go-slack"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http/httptest"
	"os"
	"testing"
)

func TestRecovery(t *testing.T) {
	// You need get your slack WebhookURL from https://api.slack.com/incoming-webhooks
	webhookURL := os.Getenv("webhookURL")
	gin.SetMode(gin.TestMode)

	s := slack.New().WebhookURL(webhookURL)

	r := gin.New()
	r.Use(Recovery(s))

	r.GET("/", func(c *gin.Context) {
		panic("oh, it's panic")
	})
	performRequest("GET", "/", r)

}

func performRequest(method, target string, router *gin.Engine) *httptest.ResponseRecorder {
	r := httptest.NewRequest(method, target, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	return w
}
