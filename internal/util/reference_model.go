package util

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ReferenceModel is a reusable model for API object references of the common
// shape { type, id, name }. Reference it from resource models with either a
// single value (*ReferenceModel) or a list ([]ReferenceModel).
type ReferenceModel struct {
	Type types.String `tfsdk:"type"`
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// ToMap converts the reference into a generic map, omitting null/unknown
// members so they are excluded from a JSON request body.
func (r ReferenceModel) ToMap() map[string]any {
	out := map[string]any{}
	if !r.Type.IsNull() && !r.Type.IsUnknown() {
		out["type"] = r.Type.ValueString()
	}
	if !r.Id.IsNull() && !r.Id.IsUnknown() {
		out["id"] = r.Id.ValueString()
	}
	if !r.Name.IsNull() && !r.Name.IsUnknown() {
		out["name"] = r.Name.ValueString()
	}
	return out
}

// ReferenceFromMap builds a ReferenceModel from a generic map (e.g. an API
// response fragment). Missing members become null.
func ReferenceFromMap(data map[string]any) *ReferenceModel {
	if data == nil {
		return nil
	}
	ref := &ReferenceModel{
		Type: types.StringNull(),
		Id:   types.StringNull(),
		Name: types.StringNull(),
	}
	if v, ok := data["type"].(string); ok {
		ref.Type = types.StringValue(v)
	}
	if v, ok := data["id"].(string); ok {
		ref.Id = types.StringValue(v)
	}
	if v, ok := data["name"].(string); ok {
		ref.Name = types.StringValue(v)
	}
	return ref
}
