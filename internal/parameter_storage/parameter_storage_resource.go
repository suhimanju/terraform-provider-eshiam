// Package parameter_storage implements the example_parameter_storage resource.
package parameter_storage

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
	_ resource.Resource                = &parameterStorageResource{}
	_ resource.ResourceWithConfigure   = &parameterStorageResource{}
	_ resource.ResourceWithImportState = &parameterStorageResource{}
)

// NewParameterStorageResource is the constructor registered with the provider.
func NewParameterStorageResource() resource.Resource {
	return &parameterStorageResource{crud.BaseResource[parameterStorageModel, *parameterStorageModel]{
		TypeNameSuffix: "_parameter_storage",
		Endpoint:       "/beta/parameter-storage",
		SchemaFn:       parameterStorageSchema,
	}}
}

type parameterStorageResource struct {
	crud.BaseResource[parameterStorageModel, *parameterStorageModel]
}

func parameterStorageSchema(_ context.Context) schema.Schema {
	computed := func() schema.StringAttribute { return schema.StringAttribute{Computed: true} }
	return schema.Schema{
		Description: "Manages a parameter-storage entry.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name":                            schema.StringAttribute{Required: true},
			"description":                     schema.StringAttribute{Optional: true},
			"type":                            schema.StringAttribute{Optional: true, Computed: true},
			"owner_id":                        schema.StringAttribute{Optional: true},
			"primary_field":                   schema.StringAttribute{Optional: true},
			"public_fields":                   schema.StringAttribute{CustomType: jsontypes.ExactType{}, Optional: true},
			"private_fields":                  schema.StringAttribute{CustomType: jsontypes.ExactType{}, Optional: true, Sensitive: true},
			"last_modified_at":                computed(),
			"last_modified_by":                computed(),
			"private_fields_last_modified_at": computed(),
			"private_fields_last_modified_by": computed(),
		},
	}
}
