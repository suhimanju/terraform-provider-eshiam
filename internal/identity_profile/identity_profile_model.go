package identity_profile

import (
	"terraform-provider-eshiam/internal/util"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// identityProfileModel mirrors the identity-profile resource.
type identityProfileModel struct {
	Id                      types.String         `tfsdk:"id"`
	Name                    types.String         `tfsdk:"name"`
	Description             types.String         `tfsdk:"description"`
	Priority                types.Int64          `tfsdk:"priority"`
	AuthoritativeSource     *util.ReferenceModel `tfsdk:"authoritative_source"`
	Owner                   *util.ReferenceModel `tfsdk:"owner"`
	IdentityAttributeConfig jsontypes.Exact      `tfsdk:"identity_attribute_config"`
}

func (m *identityProfileModel) GetID() string { return m.Id.ValueString() }
