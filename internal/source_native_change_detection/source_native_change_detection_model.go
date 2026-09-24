package source_native_change_detection

import "github.com/hashicorp/terraform-plugin-framework/types"

// sourceNativeChangeDetectionModel mirrors the native-change-detection
// configuration of a source (a singleton sub-resource).
type sourceNativeChangeDetectionModel struct {
	SourceId                         types.String `tfsdk:"source_id"`
	Enabled                          types.Bool   `tfsdk:"enabled"`
	Operations                       types.Set    `tfsdk:"operations"`
	AllEntitlements                  types.Bool   `tfsdk:"all_entitlements"`
	AllNonEntitlementAttributes      types.Bool   `tfsdk:"all_non_entitlement_attributes"`
	SelectedEntitlements             types.Set    `tfsdk:"selected_entitlements"`
	SelectedNonEntitlementAttributes types.Set    `tfsdk:"selected_non_entitlement_attributes"`
}

func (m *sourceNativeChangeDetectionModel) GetID() string { return m.SourceId.ValueString() }

func (m *sourceNativeChangeDetectionModel) path() string {
	return "/beta/sources/" + m.SourceId.ValueString() + "/native-change-detection-config"
}
