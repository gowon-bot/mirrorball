package tasks

import (
	"fmt"
	"time"

	"github.com/jivison/gowon-indexer/lib/services/webhook"
)

// TestTask is just a test it won't last long
func TestTask(str string, token string) (string, error) {
	webhookService := webhook.CreateService()

	time.Sleep(10 * time.Second)

	data := webhookService.BuildTaskCompleteRequest(token)

	webhookService.PostTo("https://webhook.site/ad3f1cff-a496-4cc5-b1f4-85acb6bb1bb8", data)

	return fmt.Sprintf("Task run with %s", str), nil
}
