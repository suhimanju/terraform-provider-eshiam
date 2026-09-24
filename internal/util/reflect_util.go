package util

import "reflect"

// IsNil reports whether i is nil, including typed nil pointers, maps, slices,
// channels, and arrays. Use this for nil-safe checks on interface values.
func IsNil(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}
