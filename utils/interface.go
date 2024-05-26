package utils

import "fmt"

func InterfaceToString(value interface{}) (*string, error) {

	switch v := value.(type) {

	case string:
		return &v, nil

	default:
		return nil, fmt.Errorf("value isn't string")
	}

}
