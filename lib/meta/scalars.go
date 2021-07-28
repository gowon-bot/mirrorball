package meta

import (
	"fmt"
	"io"
	"strconv"
)

type Date int

func (d *Date) UnmarshalGQL(v interface{}) error {
	date, ok := v.(string)

	if !ok {
		return fmt.Errorf("Date must be a string")
	}

	dateNumber, err := strconv.Atoi(date)

	if err != nil {
		return fmt.Errorf("Date must be a valid timestamp")
	}

	*d = Date(dateNumber)

	return nil
}

func (d Date) MarshalGQL(w io.Writer) {
	w.Write([]byte(fmt.Sprintf("%d", d)))
}
