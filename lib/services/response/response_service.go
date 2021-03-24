package response

import (
	"crypto/rand"
	"fmt"

	"github.com/jivison/gowon-indexer/lib/graph/model"
)

// Response holds methods for generating API responses
type Response struct{}

// BuildTaskStartResponse builds a task start response
func (r Response) BuildTaskStartResponse(taskName, token string, success bool) *model.TaskStartResponse {
	return &model.TaskStartResponse{
		Success:  success,
		Token:    token,
		TaskName: taskName,
	}
}

// GenerateToken generates a token used to mark tasks
func (r Response) GenerateToken() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

// CreateService creates an instance of the response service object
func CreateService() *Response {
	service := &Response{}

	return service
}
