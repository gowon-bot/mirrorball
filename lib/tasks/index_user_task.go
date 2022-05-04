package tasks

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/services/indexing"
	"github.com/jivison/gowon-indexer/lib/services/webhook"
)

// IndexUserTask fully indexes a user
func IndexUserTask(userJSON string, token string) (string, error) {
	indexingService := indexing.CreateService()
	webhookService := webhook.CreateService()

	user := &db.User{}
	json.Unmarshal([]byte(userJSON), user)

	err := indexingService.NewFullIndex(user)

	var data *bytes.Buffer

	if err != nil {
		data = webhookService.BuildTaskErrorRequest(token, err.Error())
	} else {
		data = webhookService.BuildTaskCompleteRequest(token)
	}

	webhookService.Post(data)

	return fmt.Sprintf("Indexed user %s", user.Username), nil
}
