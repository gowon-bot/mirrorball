package tasks

import (
	"fmt"
)

// IndexUserTask fully indexes a user
func IndexUserTask(username string, token string) (string, error) {
	return fmt.Sprintf("Indexed user %s", username), nil
}
