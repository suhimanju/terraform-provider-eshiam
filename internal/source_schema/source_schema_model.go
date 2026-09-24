package source_schema

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// sourceSchemaModel mirrors the source-schema sub-resource of a source.
type sourceSchemaModel struct {
	Id                 types.String    `tfsdk:"id"`
	SourceId           types.String    `tfsdk:"source_id"`
	Name               types.String    `tfsdk:"name"`
	NativeObjectType   types.String    `tfsdk:"native_object_type"`
	IdentityAttribute  types.String    `tfsdk:"identity_attribute"`
	DisplayAttribute   types.String    `tfsdk:"display_attribute"`
	HierarchyAttribute types.String    `tfsdk:"hierarchy_attribute"`
	IncludePermissions types.Bool      `tfsdk:"include_permissions"`
	Features           types.List      `tfsdk:"features"`
	Configuration      jsontypes.Exact `tfsdk:"configuration"`
	Attributes         jsontypes.Exact `tfsdk:"attributes"`
}

func (m *sourceSchemaModel) GetID() string { return m.Id.ValueString() }

func (m *sourceSchemaModel) collectionPath() string {
	return "/v3/sources/" + m.SourceId.ValueString() + "/schemas"
}

func (m *sourceSchemaModel) objectPath() string {
	return m.collectionPath() + "/" + m.GetID()
}
