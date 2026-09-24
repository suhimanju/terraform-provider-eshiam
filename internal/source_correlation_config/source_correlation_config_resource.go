// Package source_correlation_config implements the
// example_source_correlation_config sub-resource.
package source_correlation_config

import (
	"context"

	"terraform-provider-eshiam/internal/crud"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var (
	_ resource.Resource              = &sourceCorrelationConfigResource{}
	_ resource.ResourceWithConfigure = &sourceCorrelationConfigResource{}
)

// NewSourceCorrelationConfigResource is the constructor registered with the provider.
func NewSourceCorrelationConfigResource() resource.Resource {
	return &sourceCorrelationConfigResource{crud.SubResource[sourceCorrelationConfigModel, *sourceCorrelationConfigModel]{
		TypeNameSuffix: "_source_correlation_config",
		CollectionPath: func(m *sourceCorrelationConfigModel) string { return m.path() },
		ObjectPath:     func(m *sourceCorrelationConfigModel) string { return m.path() },
		Singleton:      true,
		SchemaFn:       sourceCorrelationConfigSchema,
	}}
}

type sourceCorrelationConfigResource struct {
	crud.SubResource[sourceCorrelationConfigModel, *sourceCorrelationConfigModel]
}

func sourceCorrelationConfigSchema(_ context.Context) schema.Schema {
	return schema.Schema{
		Description: "Manages the account-correlation configuration of a source.",
		Attributes: map[string]schema.Attribute{
			"source_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"id":   schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"name": schema.StringAttribute{Optional: true, Computed: true},
			"attribute_assignments": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"property":      schema.StringAttribute{Required: true},
						"value":         schema.StringAttribute{Required: true},
						"operation":     schema.StringAttribute{Optional: true, Computed: true},
						"complex":       schema.BoolAttribute{Optional: true, Computed: true},
						"ignore_case":   schema.BoolAttribute{Optional: true, Computed: true},
						"match_mode":    schema.StringAttribute{Optional: true, Computed: true},
						"filter_string": schema.StringAttribute{Optional: true, Computed: true},
					},
				},
			},
		},
	}
}
