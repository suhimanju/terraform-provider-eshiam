// Package role implements the example_role resource.
package role

import (
	"context"

	"terraform-provider-eshiam/internal/crud"
	"terraform-provider-eshiam/internal/util"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &roleResource{}
	_ resource.ResourceWithConfigure   = &roleResource{}
	_ resource.ResourceWithImportState = &roleResource{}
)

// NewRoleResource is the constructor registered with the provider.
func NewRoleResource() resource.Resource {
	return &roleResource{crud.BaseResource[roleModel, *roleModel]{
		TypeNameSuffix: "_role",
		Endpoint:       "/v3/roles",
		SchemaFn:       roleSchema,
	}}
}

type roleResource struct {
	crud.BaseResource[roleModel, *roleModel]
}

func roleSchema(_ context.Context) schema.Schema {
	json := func(required bool) schema.StringAttribute {
		return schema.StringAttribute{CustomType: jsontypes.ExactType{}, Optional: !required, Required: required}
	}
	return schema.Schema{
		Description: "Manages a role.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name":                      schema.StringAttribute{Required: true},
			"description":               schema.StringAttribute{Optional: true},
			"enabled":                   schema.BoolAttribute{Optional: true, Computed: true},
			"requestable":               schema.BoolAttribute{Optional: true, Computed: true},
			"owner":                     util.ResourceReferenceSchema("IDENTITY", true, "The role owner."),
			"entitlements":              json(false),
			"access_profiles":           json(false),
			"membership":                json(false),
			"access_request_config":     json(false),
			"revocation_request_config": json(false),
			"segments":                  schema.ListAttribute{ElementType: types.StringType, Optional: true},
		},
	}
}
