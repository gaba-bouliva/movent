package data

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

type Runtime struct {
	Minutes int32
}

func (r Runtime) MarshalJSON() ([]byte, error) {
	value := strconv.Itoa(int(r.Minutes))
	jsonStr := fmt.Sprintf("%q", value+" mins")
	return []byte(jsonStr), nil
}

func (r *Runtime) UnmarshalJSON(data []byte) error {
	inputStr := strings.Trim(string(data), "\"")
	splits := strings.Split(inputStr, " ")
	if len(splits) != 2 || splits[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}
	newValue, err := strconv.Atoi(splits[0])
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	r.Minutes = int32(newValue)

	return nil
}

func (r *Runtime) Scan(value interface{}) error {
	if v, ok := value.(int64); ok {
		r.Minutes = int32(v)
		return nil
	}
	return fmt.Errorf("cannot scan %T into Runtime", value)
}

func (r Runtime) Value() (driver.Value, error) {
	return int64(r.Minutes), nil
}
