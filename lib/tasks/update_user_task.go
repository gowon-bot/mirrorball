package tasks

import (
	"fmt"
)

// UpdateUserTask updates a user with the latest data
func UpdateUserTask(username string, token string) (string, error) {
	return fmt.Sprintf("Updated user %s", username), nil
}
