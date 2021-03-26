package customerrors

import "fmt"

func InsufficientArgumentsSupplied(arguments string) error {
	return fmt.Errorf("Insufficient arguments supplied to perform this operation! Missing: (%s)", arguments)
}
