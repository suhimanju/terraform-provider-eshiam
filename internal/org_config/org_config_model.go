package org_config

import "github.com/hashicorp/terraform-plugin-framework/types"

// orgConfigModel mirrors the tenant org-config singleton.
type orgConfigModel struct {
	TimeZone types.String `tfsdk:"time_zone"`
}

// GetID returns an empty identifier; org config is a singleton.
func (m *orgConfigModel) GetID() string { return "" }
