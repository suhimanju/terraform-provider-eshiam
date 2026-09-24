// Package custom_user_level implements the example_custom_user_level resource.
package custom_user_level

import (
	"context"

	"terraform-provider-eshiam/internal/crud"
	"terraform-provider-eshiam/internal/util"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &customUserLevelResource{}
	_ resource.ResourceWithConfigure   = &customUserLevelResource{}
	_ resource.ResourceWithImportState = &customUserLevelResource{}
)

// NewCustomUserLevelResource is the constructor registered with the provider.
func NewCustomUserLevelResource() resource.Resource {
	return &customUserLevelResource{crud.BaseResource[customUserLevelModel, *customUserLevelModel]{
		TypeNameSuffix: "_custom_user_level",
		Endpoint:       "/beta/work-reassignment/custom-user-levels",
		SchemaFn:       customUserLevelSchema,
	}}
}

type customUserLevelResource struct {
	crud.BaseResource[customUserLevelModel, *customUserLevelModel]
}

func customUserLevelSchema(_ context.Context) schema.Schema {
	return schema.Schema{
		Description: "Manages a custom user level.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name":        schema.StringAttribute{Required: true},
			"description": schema.StringAttribute{Optional: true},
			"owner":       util.ResourceReferenceSchema("IDENTITY", true, "The level owner."),
			"right_sets": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"status":                      schema.StringAttribute{Optional: true, Computed: true},
			"custom":                      schema.BoolAttribute{Computed: true},
			"admin_assignable":            schema.BoolAttribute{Optional: true, Computed: true},
			"associated_identities_count": schema.Int64Attribute{Computed: true},
		},
	}
}
