package parameter_storage

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// parameterStorageModel mirrors the parameter-storage resource.
type parameterStorageModel struct {
	Id                          types.String    `tfsdk:"id"`
	Name                        types.String    `tfsdk:"name"`
	Description                 types.String    `tfsdk:"description"`
	Type                        types.String    `tfsdk:"type"`
	OwnerId                     types.String    `tfsdk:"owner_id"`
	PrimaryField                types.String    `tfsdk:"primary_field"`
	PublicFields                jsontypes.Exact `tfsdk:"public_fields"`
	PrivateFields               jsontypes.Exact `tfsdk:"private_fields"`
	LastModifiedAt              types.String    `tfsdk:"last_modified_at"`
	LastModifiedBy              types.String    `tfsdk:"last_modified_by"`
	PrivateFieldsLastModifiedAt types.String    `tfsdk:"private_fields_last_modified_at"`
	PrivateFieldsLastModifiedBy types.String    `tfsdk:"private_fields_last_modified_by"`
}

func (m *parameterStorageModel) GetID() string { return m.Id.ValueString() }
