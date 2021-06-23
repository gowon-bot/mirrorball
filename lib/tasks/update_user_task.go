package tasks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/services/indexing"
	"github.com/jivison/gowon-indexer/lib/services/webhook"
)

// UpdateUserTask updates a user with the latest data
func UpdateUserTask(userJSON string, token string) (string, error) {
	indexingService := indexing.CreateService()
	webhookService := webhook.CreateService()

	user := &db.User{}
	json.Unmarshal([]byte(userJSON), user)

	err := indexingService.Update(user)

	var data *bytes.Buffer

	if err != nil {
		data = webhookService.BuildTaskErrorRequest(token, err.Error())
	} else {
		data = webhookService.BuildTaskCompleteRequest(token)
	}

	// Give the client a moment to start waiting for the webhook
	time.Sleep(200 * time.Millisecond)

	webhookService.Post(data)

	return fmt.Sprintf("Updated user %s (%s)", user.Username, user.UserType), nil
}
