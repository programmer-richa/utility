package forms

import (
	"strconv"
	"time"
)

// LoaderFunc specifies the datatype of data is stored in a field.
type LoaderFunc func(string) (interface{}, error)

// StringLoader specifies form data as string.
var StringLoader = LoaderFunc(func(rawValue string) (interface{}, error) {
	return rawValue, nil
})

// IntLoader specifies form data as int.
var IntLoader = LoaderFunc(func(rawValue string) (interface{}, error) {
	val, err := strconv.ParseInt(rawValue, 0, 0)
	if err != nil {
		return nil, err
	}
	return int(val), nil
})

// TimeLoader specifies form data as time.
var TimeLoader = NewTimeLoader(time.RFC3339)

func NewTimeLoader(layout string) LoaderFunc {
	return LoaderFunc(func(rawValue string) (interface{}, error) {
		val, err := time.Parse(layout, rawValue)
		if err != nil {
			return nil, err
		}
		return val, nil
	})
}
