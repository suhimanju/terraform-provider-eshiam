// Package launcher implements the example_launcher resource.
package launcher

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
	_ resource.Resource                = &launcherResource{}
	_ resource.ResourceWithConfigure   = &launcherResource{}
	_ resource.ResourceWithImportState = &launcherResource{}
)

// NewLauncherResource is the constructor registered with the provider.
func NewLauncherResource() resource.Resource {
	return &launcherResource{crud.BaseResource[launcherModel, *launcherModel]{
		TypeNameSuffix: "_launcher",
		Endpoint:       "/beta/launchers",
		SchemaFn:       launcherSchema,
	}}
}

type launcherResource struct {
	crud.BaseResource[launcherModel, *launcherModel]
}

func launcherSchema(_ context.Context) schema.Schema {
	return schema.Schema{
		Description: "Manages a launcher.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name":        schema.StringAttribute{Required: true},
			"description": schema.StringAttribute{Optional: true},
			"type":        schema.StringAttribute{Required: true},
			"disabled":    schema.BoolAttribute{Optional: true, Computed: true},
			"config":      schema.StringAttribute{CustomType: jsontypes.ExactType{}, Optional: true},
			"reference":   schema.StringAttribute{CustomType: jsontypes.ExactType{}, Optional: true},
		},
	}
}
