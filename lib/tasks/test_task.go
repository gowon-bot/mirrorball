package tasks

import (
	"fmt"
	"time"
)

// TestTask is just a test it won't last long
func TestTask(str string) (string, error) {
	time.Sleep(10 * time.Second)

	return fmt.Sprintf("Task run with %s", str), nil
}
