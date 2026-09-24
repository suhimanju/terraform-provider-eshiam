// Package provisioning_policy implements the example_provisioning_policy
// sub-resource of a source, keyed by usage type.
package provisioning_policy

import (
	"context"

	"terraform-provider-eshiam/internal/crud"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var (
	_ resource.Resource              = &provisioningPolicyResource{}
	_ resource.ResourceWithConfigure = &provisioningPolicyResource{}
)

// NewProvisioningPolicyResource is the constructor registered with the provider.
func NewProvisioningPolicyResource() resource.Resource {
	return &provisioningPolicyResource{crud.SubResource[provisioningPolicyModel, *provisioningPolicyModel]{
		TypeNameSuffix: "_provisioning_policy",
		CollectionPath: func(m *provisioningPolicyModel) string { return m.collectionPath() },
		ObjectPath:     func(m *provisioningPolicyModel) string { return m.objectPath() },
		SchemaFn:       provisioningPolicySchema,
	}}
}

type provisioningPolicyResource struct {
	crud.SubResource[provisioningPolicyModel, *provisioningPolicyModel]
}

func provisioningPolicySchema(_ context.Context) schema.Schema {
	return schema.Schema{
		Description: "Manages a provisioning policy (account template) of a source.",
		Attributes: map[string]schema.Attribute{
			"source_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"usage_type": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name":        schema.StringAttribute{Required: true},
			"description": schema.StringAttribute{Optional: true},
			"fields": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name":            schema.StringAttribute{Required: true},
						"type":            schema.StringAttribute{Optional: true, Computed: true},
						"is_required":     schema.BoolAttribute{Optional: true},
						"is_multi_valued": schema.BoolAttribute{Optional: true},
						"transform":       schema.StringAttribute{CustomType: jsontypes.ExactType{}, Optional: true},
						"attributes":      schema.StringAttribute{CustomType: jsontypes.ExactType{}, Optional: true},
					},
				},
			},
		},
	}
}
