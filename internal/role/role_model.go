package role

import (
	"terraform-provider-eshiam/internal/util"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// roleModel mirrors the role resource.
type roleModel struct {
	Id                      types.String         `tfsdk:"id"`
	Name                    types.String         `tfsdk:"name"`
	Description             types.String         `tfsdk:"description"`
	Enabled                 types.Bool           `tfsdk:"enabled"`
	Requestable             types.Bool           `tfsdk:"requestable"`
	Owner                   *util.ReferenceModel `tfsdk:"owner"`
	Entitlements            jsontypes.Exact      `tfsdk:"entitlements"`
	AccessProfiles          jsontypes.Exact      `tfsdk:"access_profiles"`
	Membership              jsontypes.Exact      `tfsdk:"membership"`
	AccessRequestConfig     jsontypes.Exact      `tfsdk:"access_request_config"`
	RevocationRequestConfig jsontypes.Exact      `tfsdk:"revocation_request_config"`
	Segments                types.List           `tfsdk:"segments"`
}

func (m *roleModel) GetID() string { return m.Id.ValueString() }
