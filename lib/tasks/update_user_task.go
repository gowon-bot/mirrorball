package tasks

import (
	"encoding/json"
	"fmt"

	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/services/indexing"
)

// UpdateUserTask updates a user with the latest data
func UpdateUserTask(userJSON string, token string) (string, error) {
	indexingService := indexing.CreateService()

	user := &db.User{}
	json.Unmarshal([]byte(userJSON), user)

	indexingService.Update(user)

	return fmt.Sprintf("Updated user %s (%s)", user.Username, user.UserType), nil
}
