// Package identity_attribute implements the example_identity_attribute resource.
package identity_attribute

import (
	"context"
	"net/http"

	"terraform-provider-eshiam/internal/crud"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var (
	_ resource.Resource              = &identityAttributeResource{}
	_ resource.ResourceWithConfigure = &identityAttributeResource{}
)

// NewIdentityAttributeResource is the constructor registered with the provider.
func NewIdentityAttributeResource() resource.Resource {
	return &identityAttributeResource{crud.BaseResource[identityAttributeModel, *identityAttributeModel]{
		TypeNameSuffix: "_identity_attribute",
		Endpoint:       "/v3/identity-attributes",
		UpdateMethod:   http.MethodPut,
		SchemaFn:       identityAttributeSchema,
	}}
}

type identityAttributeResource struct {
	crud.BaseResource[identityAttributeModel, *identityAttributeModel]
}

func identityAttributeSchema(_ context.Context) schema.Schema {
	return schema.Schema{
		Description: "Manages an identity attribute.",
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"display_name": schema.StringAttribute{Required: true},
			"type":         schema.StringAttribute{Optional: true, Computed: true},
			"standard":     schema.BoolAttribute{Optional: true, Computed: true},
			"system":       schema.BoolAttribute{Optional: true, Computed: true},
			"multi":        schema.BoolAttribute{Optional: true, Computed: true},
			"searchable":   schema.BoolAttribute{Optional: true, Computed: true},
			"sources":      schema.StringAttribute{CustomType: jsontypes.ExactType{}, Optional: true},
		},
	}
}
