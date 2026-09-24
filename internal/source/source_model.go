package source

import (
	"terraform-provider-eshiam/internal/util"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// sourceModel mirrors the source resource.
type sourceModel struct {
	Id                             types.String         `tfsdk:"id"`
	Name                           types.String         `tfsdk:"name"`
	Description                    types.String         `tfsdk:"description"`
	Owner                          *util.ReferenceModel `tfsdk:"owner"`
	Cluster                        *util.ReferenceModel `tfsdk:"cluster"`
	Connector                      types.String         `tfsdk:"connector"`
	Type                           types.String         `tfsdk:"type"`
	Authoritative                  types.Bool           `tfsdk:"authoritative"`
	Status                         types.String         `tfsdk:"status"`
	Features                       types.List           `tfsdk:"features"`
	ConnectorAttributes            jsontypes.Exact      `tfsdk:"connector_attributes"`
	ConnectorAttributesCredentials jsontypes.Exact      `tfsdk:"connector_attributes_credentials"`
	DeleteThreshold                types.Int64          `tfsdk:"delete_threshold"`
}

func (m *sourceModel) GetID() string { return m.Id.ValueString() }
