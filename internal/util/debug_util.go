package util

import (
	"encoding/json"
	"io"
	"net/http"
)

// PrettyPrint returns an indented JSON representation of i, useful for logging.
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

// ToJsonString returns a compact JSON representation of i.
func ToJsonString(i interface{}) string {
	s, _ := json.Marshal(i)
	return string(s)
}

// ConvertToMap marshals v and unmarshals it into a generic map.
func ConvertToMap(v any) map[string]interface{} {
	var m map[string]interface{}
	b, _ := json.Marshal(v)
	_ = json.Unmarshal(b, &m)
	return m
}

// GetBody reads and returns the response body as a string. The body is fully
// consumed; use only for error reporting.
func GetBody(resp *http.Response) string {
	if resp == nil {
		return ""
	}
	all, _ := io.ReadAll(resp.Body)
	return string(all)
}
