package tasks

import (
	"encoding/json"
	"fmt"

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

	indexingService.Update(user)

	data := webhookService.BuildTaskCompleteRequest(token)

	webhookService.Post(data)

	return fmt.Sprintf("Updated user %s (%s)", user.Username, user.UserType), nil
}
