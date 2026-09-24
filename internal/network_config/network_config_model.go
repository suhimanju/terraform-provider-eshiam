package network_config

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// networkConfigModel mirrors the tenant network-config singleton.
type networkConfigModel struct {
	Whitelisted types.Bool `tfsdk:"whitelisted"`
	Range       types.List `tfsdk:"range"`
	Geolocation types.List `tfsdk:"geolocation"`
}

// GetID returns an empty identifier; network config is a singleton.
func (m *networkConfigModel) GetID() string { return "" }
