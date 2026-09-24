package util

import (
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ModelToMap converts a terraform-plugin-framework model struct into a generic
// map suitable for a JSON request body. Fields are keyed by their `tfsdk` tag.
//
// Supported field types: types.String, types.Bool, types.Int64, types.Float64,
// types.List (of strings), types.Set (of strings), jsontypes.Exact,
// jsontypes.Normalized, ReferenceModel / *ReferenceModel, nested model structs
// (and pointers to them), and slices of any supported element type. Null and
// unknown values are skipped so they are omitted from the request body.
//
// The "id" field is always skipped as it is server-assigned.
func ModelToMap(model any, diagnostics *diag.Diagnostics) map[string]any {
	return structToMap(reflect.ValueOf(model), diagnostics, true)
}

// structToMap reflects over a struct value and converts tagged fields. When
// skipID is true the "id" field is omitted (used for top-level request bodies).
func structToMap(v reflect.Value, diagnostics *diag.Diagnostics, skipID bool) map[string]any {
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
	}
	out := map[string]any{}
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("tfsdk")
		if tag == "" || tag == "-" || (skipID && tag == "id") {
			continue
		}
		if val, ok := extractValue(v.Field(i).Interface(), diagnostics); ok {
			out[tag] = val
		}
	}
	return out
}

// extractValue converts a single framework attribute value to a Go value. The
// bool return reports whether the value should be included (false for
// null/unknown/absent values).
func extractValue(fieldVal any, diagnostics *diag.Diagnostics) (any, bool) {
	switch val := fieldVal.(type) {
	case types.String:
		if val.IsNull() || val.IsUnknown() {
			return nil, false
		}
		return val.ValueString(), true
	case types.Bool:
		if val.IsNull() || val.IsUnknown() {
			return nil, false
		}
		return val.ValueBool(), true
	case types.Int64:
		if val.IsNull() || val.IsUnknown() {
			return nil, false
		}
		return val.ValueInt64(), true
	case types.Int32:
		if val.IsNull() || val.IsUnknown() {
			return nil, false
		}
		return val.ValueInt32(), true
	case types.Float64:
		if val.IsNull() || val.IsUnknown() {
			return nil, false
		}
		return val.ValueFloat64(), true
	case jsontypes.Exact:
		if val.IsNull() || val.IsUnknown() {
			return nil, false
		}
		return UnmarshalJsonType(val, diagnostics), true
	case jsontypes.Normalized:
		if val.IsNull() || val.IsUnknown() {
			return nil, false
		}
		return UnmarshalJsonTypeNormalized(val, diagnostics), true
	case types.List:
		return extractCollection(val.Elements())
	case types.Set:
		return extractCollection(val.Elements())
	case *ReferenceModel:
		if val == nil {
			return nil, false
		}
		return val.ToMap(), true
	case ReferenceModel:
		return val.ToMap(), true
	}

	// Fall back to reflection for nested structs and slices.
	rv := reflect.ValueOf(fieldVal)
	switch rv.Kind() {
	case reflect.Ptr:
		if rv.IsNil() {
			return nil, false
		}
		return extractValue(rv.Elem().Interface(), diagnostics)
	case reflect.Struct:
		return structToMap(rv, diagnostics, false), true
	case reflect.Slice:
		if rv.IsNil() {
			return nil, false
		}
		out := make([]any, 0, rv.Len())
		for i := 0; i < rv.Len(); i++ {
			if el, ok := extractValue(rv.Index(i).Interface(), diagnostics); ok {
				out = append(out, el)
			}
		}
		return out, true
	}
	return nil, false
}

// extractCollection converts a list/set of string elements to a []string.
func extractCollection(elements []attr.Value) (any, bool) {
	if len(elements) == 0 {
		return []string{}, true
	}
	out := make([]string, 0, len(elements))
	for _, e := range elements {
		if s, ok := e.(types.String); ok && !s.IsNull() {
			out = append(out, s.ValueString())
		}
	}
	return out, true
}

// MapToModel populates a terraform-plugin-framework model struct from a generic
// map decoded from a JSON API response. It is the inverse of ModelToMap and
// supports the same field types.
func MapToModel(data map[string]any, model any, diagnostics *diag.Diagnostics) {
	v := reflect.ValueOf(model)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("tfsdk")
		if tag == "" || tag == "-" {
			continue
		}
		raw, present := data[tag]
		fv := v.Field(i)
		if !fv.CanSet() {
			continue
		}
		setFieldValue(fv, raw, present, diagnostics)
	}
}

// setFieldValue assigns a decoded JSON value to a model field, converting to the
// appropriate framework type. Missing values become null/zero.
func setFieldValue(fv reflect.Value, raw any, present bool, diagnostics *diag.Diagnostics) {
	switch fv.Interface().(type) {
	case types.String:
		if !present || raw == nil {
			fv.Set(reflect.ValueOf(types.StringNull()))
			return
		}
		fv.Set(reflect.ValueOf(types.StringValue(toString(raw))))
		return
	case types.Bool:
		if b, ok := raw.(bool); ok {
			fv.Set(reflect.ValueOf(types.BoolValue(b)))
		} else {
			fv.Set(reflect.ValueOf(types.BoolNull()))
		}
		return
	case types.Int64:
		if !present || raw == nil {
			fv.Set(reflect.ValueOf(types.Int64Null()))
			return
		}
		fv.Set(reflect.ValueOf(types.Int64Value(toInt64(raw))))
		return
	case types.Int32:
		if !present || raw == nil {
			fv.Set(reflect.ValueOf(types.Int32Null()))
			return
		}
		fv.Set(reflect.ValueOf(types.Int32Value(int32(toInt64(raw)))))
		return
	case types.Float64:
		if f, ok := raw.(float64); ok {
			fv.Set(reflect.ValueOf(types.Float64Value(f)))
		} else {
			fv.Set(reflect.ValueOf(types.Float64Null()))
		}
		return
	case jsontypes.Exact:
		if !present || raw == nil {
			fv.Set(reflect.ValueOf(jsontypes.NewExactNull()))
			return
		}
		fv.Set(reflect.ValueOf(MarshalToJsonType(raw, diagnostics)))
		return
	case jsontypes.Normalized:
		if !present || raw == nil {
			fv.Set(reflect.ValueOf(jsontypes.NewNormalizedNull()))
			return
		}
		fv.Set(reflect.ValueOf(MarshalToJsonTypeNormalized(raw, diagnostics)))
		return
	case types.List:
		fv.Set(reflect.ValueOf(toStringList(raw)))
		return
	case types.Set:
		fv.Set(reflect.ValueOf(toStringSet(raw)))
		return
	case ReferenceModel:
		if m, ok := raw.(map[string]any); ok {
			fv.Set(reflect.ValueOf(*ReferenceFromMap(m)))
		}
		return
	case *ReferenceModel:
		if m, ok := raw.(map[string]any); ok {
			fv.Set(reflect.ValueOf(ReferenceFromMap(m)))
		} else {
			fv.Set(reflect.Zero(fv.Type()))
		}
		return
	}

	// Reflection fallback for nested struct pointers and slices.
	switch fv.Kind() {
	case reflect.Ptr:
		if !present || raw == nil {
			fv.Set(reflect.Zero(fv.Type()))
			return
		}
		m, ok := raw.(map[string]any)
		if !ok {
			return
		}
		nested := reflect.New(fv.Type().Elem())
		MapToModel(m, nested.Interface(), diagnostics)
		fv.Set(nested)
	case reflect.Struct:
		if m, ok := raw.(map[string]any); ok {
			nested := reflect.New(fv.Type())
			MapToModel(m, nested.Interface(), diagnostics)
			fv.Set(nested.Elem())
		}
	case reflect.Slice:
		items, ok := raw.([]any)
		if !ok || items == nil {
			fv.Set(reflect.Zero(fv.Type()))
			return
		}
		elemType := fv.Type().Elem()

		// Slice of scalar framework values, e.g. []types.String.
		if elemType == reflect.TypeOf(types.String{}) {
			slice := reflect.MakeSlice(fv.Type(), 0, len(items))
			for _, it := range items {
				slice = reflect.Append(slice, reflect.ValueOf(types.StringValue(toString(it))))
			}
			fv.Set(slice)
			return
		}

		// Slice of nested model structs.
		slice := reflect.MakeSlice(fv.Type(), 0, len(items))
		for _, it := range items {
			m, ok := it.(map[string]any)
			if !ok {
				continue
			}
			elem := reflect.New(elemType)
			MapToModel(m, elem.Interface(), diagnostics)
			slice = reflect.Append(slice, elem.Elem())
		}
		fv.Set(slice)
	}
}

func toStringList(raw any) types.List {
	elems := toStringElements(raw)
	l, _ := types.ListValue(types.StringType, elems)
	return l
}

func toStringSet(raw any) types.Set {
	elems := toStringElements(raw)
	s, _ := types.SetValue(types.StringType, elems)
	return s
}

func toStringElements(raw any) []attr.Value {
	items, ok := raw.([]any)
	if !ok {
		return []attr.Value{}
	}
	elems := make([]attr.Value, 0, len(items))
	for _, it := range items {
		elems = append(elems, types.StringValue(toString(it)))
	}
	return elems
}

func toString(v any) string {
	switch val := v.(type) {
	case string:
		return val
	default:
		b, _ := json.Marshal(v)
		return string(b)
	}
}

func toInt64(v any) int64 {
	switch val := v.(type) {
	case float64:
		return int64(val)
	case int64:
		return val
	case int:
		return int64(val)
	}
	return 0
}
