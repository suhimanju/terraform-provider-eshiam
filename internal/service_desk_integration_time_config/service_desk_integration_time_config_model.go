package service_desk_integration_time_config

import "github.com/hashicorp/terraform-plugin-framework/types"

// serviceDeskIntegrationTimeConfigModel mirrors the tenant-wide service desk
// integration time-check configuration (a singleton object).
type serviceDeskIntegrationTimeConfigModel struct {
	ProvisioningStatusCheckIntervalMinutes types.String `tfsdk:"provisioning_status_check_interval_minutes"`
	ProvisioningMaxStatusCheckDays         types.String `tfsdk:"provisioning_max_status_check_days"`
}

// GetID satisfies crud.Identifiable; this singleton has no id.
func (m *serviceDeskIntegrationTimeConfigModel) GetID() string { return "" }
