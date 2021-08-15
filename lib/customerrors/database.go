package customerrors

import (
	"fmt"
)

// DatabaseUnknownError occurs when a database operation unexpectedly goes wrong
func DatabaseUnknownError() error {
	return fmt.Errorf("an unexpected error occurred while trying to execute this operation")
}

// EntityAlreadyExistsError occurs when an entity is created, but violates a unique constraint
func EntityAlreadyExistsError(entityName string) error {
	return fmt.Errorf("that %s already exists", entityName)
}

// EntityDoesntExistError occurs when a database entity is operated on, but doesn't exist
func EntityDoesntExistError(entityName string) error {
	return fmt.Errorf("that %s doesn't exist", entityName)
}
