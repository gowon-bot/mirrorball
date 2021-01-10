package webhook

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Webhook holds methods for posting to webhooks
type Webhook struct{}

// PostTo posts to a given url with response data
func (w Webhook) PostTo(url string, data *bytes.Buffer) {

	http.Post(url, "application/json", data)
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
	service := &Webhook{}

	return service
}
