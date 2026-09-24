// Package service_desk_integration_time_config implements the
// example_service_desk_integration_time_check_config singleton resource.
package service_desk_integration_time_config

import (
	"context"

	"terraform-provider-eshiam/internal/crud"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

var (
	_ resource.Resource              = &serviceDeskIntegrationTimeConfigResource{}
	_ resource.ResourceWithConfigure = &serviceDeskIntegrationTimeConfigResource{}
)

// NewServiceDeskIntegrationTimeCheckConfigResource is the constructor registered with the provider.
func NewServiceDeskIntegrationTimeCheckConfigResource() resource.Resource {
	return &serviceDeskIntegrationTimeConfigResource{crud.SingletonResource[serviceDeskIntegrationTimeConfigModel, *serviceDeskIntegrationTimeConfigModel]{
		TypeNameSuffix: "_service_desk_integration_time_check_config",
		Path:           "/v3/service-desk-integrations/status-check-configuration",
		SchemaFn:       serviceDeskIntegrationTimeConfigSchema,
	}}
}

type serviceDeskIntegrationTimeConfigResource struct {
	crud.SingletonResource[serviceDeskIntegrationTimeConfigModel, *serviceDeskIntegrationTimeConfigModel]
}

func serviceDeskIntegrationTimeConfigSchema(_ context.Context) schema.Schema {
	return schema.Schema{
		Description: "Manages the tenant-wide service desk integration status-check timing configuration.",
		Attributes: map[string]schema.Attribute{
			"provisioning_status_check_interval_minutes": schema.StringAttribute{Required: true},
			"provisioning_max_status_check_days":         schema.StringAttribute{Required: true},
		},
	}
}
