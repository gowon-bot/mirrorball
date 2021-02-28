package customerrors

import (
	"fmt"
)

// WavyNotSupportedError occurs when someone with a Wavy user type attempts to do an action that is Last.fm only
func WavyNotSupportedError() error {
	return fmt.Errorf("This feature hasn't been implemented for Wavy.fm users yet")
}

// LastFMError occurs when Last.fm returns an error from its api
func LastFMError(errorMessage string, errorCode int) error {
	return fmt.Errorf("%s (error %d)", errorMessage, errorCode)
}
