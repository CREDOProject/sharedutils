package types

import "fmt"

func To[T interface{}](element any) (*T, error) {
	var converted *T
	switch s := element.(type) {
	case T:
		converted = &s
	case *T:
		converted = s
	default:
		return nil, fmt.Errorf("Error converting type, %v", s)
	}
	return converted, nil
}
