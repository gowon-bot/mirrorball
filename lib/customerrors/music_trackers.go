package customerrors

import (
	"fmt"
)

// LastFMError occurs when Last.fm returns an error from its api
func LastFMError(errorMessage string, errorCode int) error {
	return fmt.Errorf("%s (error %d)", errorMessage, errorCode)
}
