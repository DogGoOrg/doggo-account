package helpers

import (
	"errors"
)

type checkConstraint interface {
	string | int | float32 | float64 | byte | rune | uint16 | uint32 | uint64 | uint | int16 | int64 | int8
}

func CheckForNullValues[T checkConstraint](vals ...T) error {
	if len(vals) == 0 {
		return errors.New("no values provided")
	}

	for _, value := range vals {
		switch inferedVal := any(value).(type) {
		case string:
			if inferedVal == "" {
				return errors.New("arg is null value")
			}
		case int:
		case int8:
		case int16:
		case rune:
		case int64:
		case byte:
		case uint:
		case uint16:
		case uint32:
		case uint64:
			if inferedVal == 0 {
				return errors.New("arg is null value")
			}
		default:
			return errors.New("unsupported value type")
		}
	}

	return nil
}
