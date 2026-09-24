// Package custom_form implements the example_custom_form resource.
package custom_form

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
	_ resource.Resource                = &customFormResource{}
	_ resource.ResourceWithConfigure   = &customFormResource{}
	_ resource.ResourceWithImportState = &customFormResource{}
)

// NewCustomFormResource is the constructor registered with the provider.
func NewCustomFormResource() resource.Resource {
	return &customFormResource{crud.BaseResource[customFormModel, *customFormModel]{
		TypeNameSuffix: "_custom_form",
		Endpoint:       "/beta/form-definitions",
		SchemaFn:       customFormSchema,
	}}
}

type customFormResource struct {
	crud.BaseResource[customFormModel, *customFormModel]
}

func customFormSchema(_ context.Context) schema.Schema {
	return schema.Schema{
		Description: "Manages a custom form definition.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name":            schema.StringAttribute{Required: true},
			"description":     schema.StringAttribute{Optional: true},
			"owner":           util.ResourceReferenceSchema("IDENTITY", false, "The form owner."),
			"used_by":         schema.StringAttribute{CustomType: jsontypes.ExactType{}, Optional: true, Computed: true},
			"form_input":      schema.StringAttribute{CustomType: jsontypes.ExactType{}, Optional: true},
			"form_conditions": schema.StringAttribute{CustomType: jsontypes.ExactType{}, Optional: true},
			"form_elements":   schema.StringAttribute{CustomType: jsontypes.ExactType{}, Optional: true},
		},
	}
}
