package customerrors

import (
	"fmt"
)

// DatabaseUnknownError occurs when a database operation unexpectedly goes wrong
func DatabaseUnknownError() error {
	return fmt.Errorf("Un unexpected error occurred while trying to execute this operation")
}

// EntityAlreadyExists occurs when an entity is created, but violates a unique constraint
func EntityAlreadyExists(entityName string) error {
	return fmt.Errorf("That %s already exists", entityName)
}

// EntityDoesntExist occurs when a database entity is operated on, but doesn't exist
func EntityDoesntExist(entityName string) error {
	return fmt.Errorf("That %s doesn't exist", entityName)
}
