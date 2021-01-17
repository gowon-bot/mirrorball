package helpers

import (
	"strconv"
	"time"
)

// ParseUnix parse a unix timestamp
func ParseUnix(timestamp string) (time.Time, error) {
	uts, err := strconv.ParseInt(timestamp, 10, 64)

	if err != nil {
		return time.Now(), nil
	}

	return time.Unix(uts, 0), nil
}
