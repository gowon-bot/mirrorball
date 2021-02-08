package webhook

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// Webhook holds methods for posting to webhooks
type Webhook struct {
	webhookURL string
}

// Post posts to a given url with response data
func (w Webhook) Post(data *bytes.Buffer) {

	http.Post(w.webhookURL, "application/json", data)
}

// TaskCompleteRequest is what's sent to a webhook upon task completion
type TaskCompleteRequest struct {
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}

// BuildTaskCompleteRequest builds a task complete request
func (w Webhook) BuildTaskCompleteRequest(token string) *bytes.Buffer {
	data := TaskCompleteRequest{
		Data: struct {
			Token string `json:"token"`
		}{
			Token: token,
		},
	}

	jsonStr, _ := json.Marshal(&data)

	return bytes.NewBuffer(jsonStr)
}

// CreateService creates an instance of the webhook service object
func CreateService() *Webhook {
	godotenv.Load()

	webhookURL := os.Getenv("WEBHOOK_POST_URL")

	service := &Webhook{webhookURL: webhookURL}

	return service
}
