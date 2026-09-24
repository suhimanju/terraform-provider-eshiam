package launcher

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// launcherModel mirrors the launcher resource.
type launcherModel struct {
	Id          types.String    `tfsdk:"id"`
	Name        types.String    `tfsdk:"name"`
	Description types.String    `tfsdk:"description"`
	Type        types.String    `tfsdk:"type"`
	Disabled    types.Bool      `tfsdk:"disabled"`
	Config      jsontypes.Exact `tfsdk:"config"`
	Reference   jsontypes.Exact `tfsdk:"reference"`
}

func (m *launcherModel) GetID() string { return m.Id.ValueString() }
