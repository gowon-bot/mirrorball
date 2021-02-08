package customerrors

import (
	"fmt"
)

func NoOneKnows(entityName string) error {
	return fmt.Errorf("No one knows %s", entityName)
}
