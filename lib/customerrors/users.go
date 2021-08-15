package customerrors

import "fmt"

func InsufficientArgumentsSupplied(arguments string) error {
	return fmt.Errorf("insufficient arguments supplied to perform this operation! Missing: (%s)", arguments)
}
