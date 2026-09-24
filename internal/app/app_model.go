package app

import (
	"terraform-provider-eshiam/internal/util"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// appModel mirrors the source-app resource.
type appModel struct {
	Id            types.String         `tfsdk:"id"`
	Name          types.String         `tfsdk:"name"`
	Description   types.String         `tfsdk:"description"`
	Enabled       types.Bool           `tfsdk:"enabled"`
	Type          types.String         `tfsdk:"type"`
	AccountSource *util.ReferenceModel `tfsdk:"account_source"`
	Owner         *util.ReferenceModel `tfsdk:"owner"`
}

func (m *appModel) GetID() string { return m.Id.ValueString() }
