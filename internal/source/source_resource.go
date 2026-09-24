// Package source implements the example_source resource.
package source

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
	_ resource.Resource                = &sourceResource{}
	_ resource.ResourceWithConfigure   = &sourceResource{}
	_ resource.ResourceWithImportState = &sourceResource{}
)

// NewSourceResource is the constructor registered with the provider.
func NewSourceResource() resource.Resource {
	return &sourceResource{crud.BaseResource[sourceModel, *sourceModel]{
		TypeNameSuffix: "_source",
		Endpoint:       "/v3/sources",
		SchemaFn:       sourceSchema,
	}}
}

type sourceResource struct {
	crud.BaseResource[sourceModel, *sourceModel]
}

func sourceSchema(_ context.Context) schema.Schema {
	return schema.Schema{
		Description: "Manages a source.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name":                             schema.StringAttribute{Required: true},
			"description":                      schema.StringAttribute{Optional: true},
			"owner":                            util.ResourceReferenceSchema("IDENTITY", true, "The source owner."),
			"cluster":                          util.ResourceReferenceSchema("CLUSTER", false, "The source cluster."),
			"connector":                        schema.StringAttribute{Required: true},
			"type":                             schema.StringAttribute{Optional: true, Computed: true},
			"authoritative":                    schema.BoolAttribute{Optional: true, Computed: true},
			"status":                           schema.StringAttribute{Optional: true, Computed: true},
			"features":                         schema.ListAttribute{ElementType: types.StringType, Optional: true},
			"connector_attributes":             schema.StringAttribute{CustomType: jsontypes.ExactType{}, Optional: true},
			"connector_attributes_credentials": schema.StringAttribute{CustomType: jsontypes.ExactType{}, Optional: true, Sensitive: true},
			"delete_threshold":                 schema.Int64Attribute{Optional: true, Computed: true},
		},
	}
}
