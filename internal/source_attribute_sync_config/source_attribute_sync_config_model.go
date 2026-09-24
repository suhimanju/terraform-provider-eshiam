package source_attribute_sync_config

import (
	"terraform-provider-eshiam/internal/util"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// sourceAttributeSyncConfigModel mirrors the attribute-synchronization
// configuration of a source (a singleton sub-resource).
type sourceAttributeSyncConfigModel struct {
	Source         util.ReferenceModel `tfsdk:"source"`
	AttributesSync []attributesSync    `tfsdk:"attributes_sync"`
}

type attributesSync struct {
	Name        types.String `tfsdk:"name"`
	DisplayName types.String `tfsdk:"display_name"`
	Enabled     types.Bool   `tfsdk:"enabled"`
	Target      types.String `tfsdk:"target"`
}

func (m *sourceAttributeSyncConfigModel) GetID() string { return m.Source.Id.ValueString() }

func (m *sourceAttributeSyncConfigModel) path() string {
	return "/beta/sources/" + m.Source.Id.ValueString() + "/attribute-sync-config"
}
