// Package urlconv is a small helper that maps url.Values into a given struct.
package urlconv

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"time"
)

// Unmarshal maps the given url.Values into a struct
func Unmarshal(values url.Values, s any) {
	// get the type and values from the struct
	structType := reflect.TypeOf(s).Elem()
	structVal := reflect.ValueOf(s)

	// check whether s really is a struct
	if structType.Kind() != reflect.Struct {
		// because this is an error that doesn't depend on user input and
		// should never really happen, we panic
		panic(fmt.Errorf("urlconv: given interface must be a struct"))
	}

	for i := 0; i < structType.NumField(); i++ {
		// get the field and check whether it has the form tag
		field := structType.Field(i)
		key, ok := field.Tag.Lookup("url")
		if !ok {
			continue
		}

		// get the field value and check if we can interface with it
		// this is false when the field or struct is private
		result := structVal.Elem().Field(i)
		if result.CanInterface() {
			// set value based on field type
			switch result.Interface().(type) {
			case string:
				result.SetString(values.Get(key))
			case int:
				if i, err := strconv.ParseInt(values.Get(key), 10, 64); err == nil {
					result.SetInt(i)
				}
			case float64:
				if f, err := strconv.ParseFloat(values.Get(key), 64); err == nil {
					result.SetFloat(f)
				}
			case []string:
				result.Set(reflect.ValueOf(values[key]))
			case bool:
				if b, err := strconv.ParseBool(values.Get(key)); err == nil {
					result.SetBool(b)
				}
			case time.Time:
				if t, err := time.Parse(time.RFC3339, values.Get(key)); err == nil {
					result.Set(reflect.ValueOf(t))
				}
			}
		}
	}
}
