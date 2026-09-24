// Package service_desk_integration implements the example_service_desk_integration
// resource, demonstrating nested objects and reference lists on crud.BaseResource.
package service_desk_integration

import (
	"context"

	"terraform-provider-eshiam/internal/crud"
	"terraform-provider-eshiam/internal/util"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var (
	_ resource.Resource                = &serviceDeskIntegrationResource{}
	_ resource.ResourceWithConfigure   = &serviceDeskIntegrationResource{}
	_ resource.ResourceWithImportState = &serviceDeskIntegrationResource{}
)

// NewServiceDeskIntegrationResource is the constructor registered with the provider.
func NewServiceDeskIntegrationResource() resource.Resource {
	return &serviceDeskIntegrationResource{crud.BaseResource[serviceDeskIntegrationModel, *serviceDeskIntegrationModel]{
		TypeNameSuffix: "_service_desk_integration",
		Endpoint:       "/v3/service-desk-integrations",
		SchemaFn:       serviceDeskIntegrationSchema,
	}}
}

type serviceDeskIntegrationResource struct {
	crud.BaseResource[serviceDeskIntegrationModel, *serviceDeskIntegrationModel]
}

func serviceDeskIntegrationSchema(_ context.Context) schema.Schema {
	return schema.Schema{
		Description: "Manages a service desk integration.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name":        schema.StringAttribute{Required: true},
			"description": schema.StringAttribute{Optional: true},
			"type":        schema.StringAttribute{Required: true},
			"owner_ref":   util.ResourceReferenceSchema("IDENTITY", false, "The integration owner."),
			"cluster_ref": util.ResourceReferenceSchema("CLUSTER", false, "The integration cluster."),
			"provisioning_config": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"managed_resource_refs": util.ResourceReferenceListSchema("SOURCE", false, "Managed resource references."),
					"plan_initializer_script": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"source": schema.StringAttribute{Optional: true},
						},
					},
					"no_provisioning_requests":        schema.BoolAttribute{Optional: true, Computed: true},
					"provisioning_request_expiration": schema.Int64Attribute{Optional: true, Computed: true},
				},
			},
			"attributes":               schema.StringAttribute{CustomType: jsontypes.NormalizedType{}, Optional: true},
			"attributes_credentials":   schema.StringAttribute{CustomType: jsontypes.ExactType{}, Optional: true, Sensitive: true},
			"before_provisioning_rule": util.ResourceReferenceSchema("RULE", false, "Before-provisioning rule reference."),
		},
	}
}
