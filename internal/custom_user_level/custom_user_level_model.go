package custom_user_level

import (
	"terraform-provider-eshiam/internal/util"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// customUserLevelModel mirrors the custom-user-level resource.
type customUserLevelModel struct {
	Id                        types.String        `tfsdk:"id"`
	Name                      types.String        `tfsdk:"name"`
	Description               types.String        `tfsdk:"description"`
	Owner                     util.ReferenceModel `tfsdk:"owner"`
	RightSets                 types.Set           `tfsdk:"right_sets"`
	Status                    types.String        `tfsdk:"status"`
	Custom                    types.Bool          `tfsdk:"custom"`
	AdminAssignable           types.Bool          `tfsdk:"admin_assignable"`
	AssociatedIdentitiesCount types.Int64         `tfsdk:"associated_identities_count"`
}

func (m *customUserLevelModel) GetID() string { return m.Id.ValueString() }
