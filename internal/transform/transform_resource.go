// Package transform implements the example_transform resource, a generic mirror
// of a Transform object. Copy this package as a template for scalar + JSON
// resources built on crud.BaseResource.
package transform

import (
	"context"

	"terraform-provider-eshiam/internal/crud"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// Interface compliance.
var (
	_ resource.Resource                = &transformResource{}
	_ resource.ResourceWithConfigure   = &transformResource{}
	_ resource.ResourceWithImportState = &transformResource{}
)

// NewTransformResource is the constructor registered with the provider.
func NewTransformResource() resource.Resource {
	return &transformResource{crud.BaseResource[transformModel, *transformModel]{
		TypeNameSuffix: "_transform",
		Endpoint:       "/v1/transforms",
		SchemaFn:       transformSchema,
	}}
}

type transformResource struct {
	crud.BaseResource[transformModel, *transformModel]
}

func transformSchema(_ context.Context) schema.Schema {
	return schema.Schema{
		Description: "Manages a transform.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name": schema.StringAttribute{
				Required:      true,
				Validators:    []validator.String{stringvalidator.LengthAtMost(50)},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"type": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"attributes": schema.StringAttribute{
				CustomType: jsontypes.ExactType{},
				Required:   true,
			},
		},
	}
}
