package patch

import (
	"encoding/json"
	"reflect"
)

// isEmpty reports whether v is nil or a zero/empty value (empty string, empty
// slice/map, or nil pointer). Used to decide between add/remove/replace.
func isEmpty(v any) bool {
	if v == nil {
		return true
	}
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Ptr, reflect.Interface:
		return rv.IsNil()
	case reflect.String:
		return rv.Len() == 0
	case reflect.Slice, reflect.Map, reflect.Array:
		return rv.Len() == 0
	}
	return false
}

// valuesEqual performs a deep, JSON-normalized equality check. Comparing via
// JSON avoids false negatives from differing but equivalent representations.
func valuesEqual(a, b any) bool {
	if reflect.DeepEqual(a, b) {
		return true
	}
	aj, errA := json.Marshal(a)
	bj, errB := json.Marshal(b)
	if errA != nil || errB != nil {
		return false
	}
	return string(aj) == string(bj)
}
