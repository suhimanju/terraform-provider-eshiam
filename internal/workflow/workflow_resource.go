// Package workflow implements the example_workflow resource, a flat collection
// object with nested definition and trigger blocks.
package workflow

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
	_ resource.Resource                = &workflowResource{}
	_ resource.ResourceWithConfigure   = &workflowResource{}
	_ resource.ResourceWithImportState = &workflowResource{}
)

// NewWorkflowResource is the constructor registered with the provider.
func NewWorkflowResource() resource.Resource {
	return &workflowResource{crud.BaseResource[workflowModel, *workflowModel]{
		TypeNameSuffix: "_workflow",
		Endpoint:       "/beta/workflows",
		SchemaFn:       workflowSchema,
	}}
}

type workflowResource struct {
	crud.BaseResource[workflowModel, *workflowModel]
}

func workflowSchema(_ context.Context) schema.Schema {
	return schema.Schema{
		Description: "Manages a workflow.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name":        schema.StringAttribute{Required: true},
			"description": schema.StringAttribute{Optional: true},
			"enabled":     schema.BoolAttribute{Optional: true, Computed: true},
			"owner":       util.ResourceReferenceSchema("IDENTITY", true, "The workflow owner."),
			"definition": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"start": schema.StringAttribute{Required: true},
					"steps": schema.StringAttribute{CustomType: jsontypes.NormalizedType{}, Required: true},
				},
			},
			"trigger": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"type":       schema.StringAttribute{Required: true},
					"attributes": schema.StringAttribute{CustomType: jsontypes.NormalizedType{}, Optional: true},
				},
			},
		},
	}
}
