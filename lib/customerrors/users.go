package customerrors

import "fmt"

func InsufficientArgumentsSupplied(arguments string) error {
	return fmt.Errorf("insufficient arguments supplied to perform this operation! Missing: (%s)", arguments)
}

func CannotSetToUnset() error {
	return fmt.Errorf("cannot set a user's privacy to UNSET")
}
