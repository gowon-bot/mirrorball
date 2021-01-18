package tasks

import (
	"fmt"

	"github.com/jivison/gowon-indexer/lib/services/indexing"
	"github.com/jivison/gowon-indexer/lib/services/user"
	"github.com/jivison/gowon-indexer/lib/services/webhook"
)

// UpdateUserTask updates a user with the latest data
func UpdateUserTask(username string, token string) (string, error) {
	webhookService := webhook.CreateService()
	userService := user.CreateService()
	indexingService := indexing.CreateService()

	user, _ := userService.GetUser(username)

	indexingService.Update(user)

	data := webhookService.BuildTaskCompleteRequest(token)

	webhookService.PostTo("https://webhook.site/ad3f1cff-a496-4cc5-b1f4-85acb6bb1bb8", data)

	return fmt.Sprintf("Indexed user %s", username), nil
}