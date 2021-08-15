package customerrors

import "fmt"

func RankNotFound() error {
	return fmt.Errorf("you don't have any scrobbles of that artist")
}
