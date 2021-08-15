package customerrors

import "fmt"

// CSVParseError occurs when parsing a csv (namely from rateyourmusic) goes wrong
func CSVParseError() error {
	return fmt.Errorf("an error ocurred parsing your CSV file, please use the correct format")
}
