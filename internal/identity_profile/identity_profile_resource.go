// Package identity_profile implements the example_identity_profile resource.
package identity_profile

import (
	"context"
	"net/http"

	"terraform-provider-eshiam/internal/crud"
	"terraform-provider-eshiam/internal/util"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var (
	_ resource.Resource                = &identityProfileResource{}
	_ resource.ResourceWithConfigure   = &identityProfileResource{}
	_ resource.ResourceWithImportState = &identityProfileResource{}
)

// NewIdentityProfileResource is the constructor registered with the provider.
func NewIdentityProfileResource() resource.Resource {
	return &identityProfileResource{crud.BaseResource[identityProfileModel, *identityProfileModel]{
		TypeNameSuffix: "_identity_profile",
		Endpoint:       "/v3/identity-profiles",
		UpdateMethod:   http.MethodPatch,
		SchemaFn:       identityProfileSchema,
	}}
}

type identityProfileResource struct {
	crud.BaseResource[identityProfileModel, *identityProfileModel]
}

func identityProfileSchema(_ context.Context) schema.Schema {
	return schema.Schema{
		Description: "Manages an identity profile.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name":                      schema.StringAttribute{Required: true},
			"description":               schema.StringAttribute{Optional: true},
			"priority":                  schema.Int64Attribute{Optional: true, Computed: true},
			"authoritative_source":      util.ResourceReferenceSchema("SOURCE", true, "The authoritative source."),
			"owner":                     util.ResourceReferenceSchema("IDENTITY", false, "The profile owner."),
			"identity_attribute_config": schema.StringAttribute{CustomType: jsontypes.ExactType{}, Optional: true},
		},
	}
}
