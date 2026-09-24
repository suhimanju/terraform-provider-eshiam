package util

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// MarshalToJsonType marshals an arbitrary value into a jsontypes.Exact value.
// A nil input yields a null value.
func MarshalToJsonType(v any, diagnostics *diag.Diagnostics) jsontypes.Exact {
	if v == nil {
		return jsontypes.NewExactNull()
	}
	jsonBytes, err := json.Marshal(&v)
	if err != nil {
		diagnostics.AddError(
			"Error Processing JSON value",
			fmt.Sprintf("Cannot marshal value to JSON '%v': %s", v, err.Error()),
		)
		return jsontypes.NewExactValue("")
	}
	return jsontypes.NewExactValue(string(jsonBytes))
}

// MarshalToJsonTypeNormalized is like MarshalToJsonType but yields a
// jsontypes.Normalized value (semantic JSON equality).
func MarshalToJsonTypeNormalized(v any, diagnostics *diag.Diagnostics) jsontypes.Normalized {
	if v == nil {
		return jsontypes.NewNormalizedNull()
	}
	jsonBytes, err := json.Marshal(&v)
	if err != nil {
		diagnostics.AddError(
			"Error Processing JSON value",
			fmt.Sprintf("Cannot marshal value to JSON '%v': %s", v, err.Error()),
		)
		return jsontypes.NewNormalizedValue("")
	}
	return jsontypes.NewNormalizedValue(string(jsonBytes))
}

// UnmarshalJsonType decodes a jsontypes.Exact JSON string into a generic map.
// Null or unknown inputs return nil.
func UnmarshalJsonType(jsonObj jsontypes.Exact, diagnostics *diag.Diagnostics) map[string]interface{} {
	if jsonObj.IsNull() || jsonObj.IsUnknown() {
		return nil
	}
	var obj map[string]interface{}
	diagnostic := jsonObj.Unmarshal(&obj)
	diagnostics.Append(diagnostic...)
	if diagnostic.HasError() {
		return make(map[string]interface{})
	}
	return obj
}

// UnmarshalJsonTypeNormalized decodes a jsontypes.Normalized JSON string into a
// generic map. Null or unknown inputs return nil.
func UnmarshalJsonTypeNormalized(jsonObj jsontypes.Normalized, diagnostics *diag.Diagnostics) map[string]interface{} {
	if jsonObj.IsNull() || jsonObj.IsUnknown() {
		return nil
	}
	var obj map[string]interface{}
	diagnostic := jsonObj.Unmarshal(&obj)
	diagnostics.Append(diagnostic...)
	if diagnostic.HasError() {
		return make(map[string]interface{})
	}
	return obj
}
