package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Runtime struct {
	value int32
}

func (r Runtime) MarshalJSON() ([]byte, error) {
	value := strconv.Itoa(int(r.value))
	jsonStr := fmt.Sprintf("%q", value+" mins")
	fmt.Println("Marshaled json: ", jsonStr)
	return []byte(jsonStr), nil
}

func (r *Runtime) UnmarshalJSON(data []byte) error {
	splits := strings.Split(string(data), " ")
	if len(splits) != 2 {
		return errors.New("error parsing json invalid runtime")
	}

	newValue, err := strconv.Atoi(splits[0])
	if err != nil {
		return fmt.Errorf("error marshaling invalid runtime %+v", *r)
	}

	r.value = int32(newValue)

	return nil
}

func (r *Runtime) Scan(value any) error {
	if value == nil {
		return nil
	}
	switch v := value.(type) {
	case int32:
		r.value = v
	case int64:
		r.value = int32(v)
	case int:
		r.value = int32(v)
	default:
		return fmt.Errorf("failed to scan invalid runtime %v from db", value)
	}
	return nil
}

func (r Runtime) Value() (any, error) {
	return r.value, nil
}
