package tasks

import (
	"encoding/json"
	"fmt"

	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/services/indexing"
)

// IndexUserTask fully indexes a user
func IndexUserTask(userJSON string, token string) (string, error) {
	indexingService := indexing.CreateService()

	user := &db.User{}
	json.Unmarshal([]byte(userJSON), user)

	indexingService.FullIndex(user)

	return fmt.Sprintf("Indexed user %s (%s)", user.Username, user.UserType), nil
}
