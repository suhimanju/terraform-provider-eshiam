// Package connector_rule implements the example_connector_rule resource.
package connector_rule

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
	_ resource.Resource                = &connectorRuleResource{}
	_ resource.ResourceWithConfigure   = &connectorRuleResource{}
	_ resource.ResourceWithImportState = &connectorRuleResource{}
)

// NewConnectorRuleResource is the constructor registered with the provider.
func NewConnectorRuleResource() resource.Resource {
	return &connectorRuleResource{crud.BaseResource[connectorRuleModel, *connectorRuleModel]{
		TypeNameSuffix: "_connector_rule",
		Endpoint:       "/beta/connector-rules",
		SchemaFn:       connectorRuleSchema,
	}}
}

type connectorRuleResource struct {
	crud.BaseResource[connectorRuleModel, *connectorRuleModel]
}

func connectorRuleSchema(_ context.Context) schema.Schema {
	return schema.Schema{
		Description: "Manages a connector rule.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name":        schema.StringAttribute{Required: true},
			"description": schema.StringAttribute{Optional: true},
			"type":        schema.StringAttribute{Required: true},
			"signature":   schema.StringAttribute{CustomType: jsontypes.ExactType{}, Optional: true},
			"source_code": schema.StringAttribute{CustomType: jsontypes.ExactType{}, Required: true},
			"attributes":  schema.StringAttribute{CustomType: jsontypes.ExactType{}, Optional: true},
		},
	}
}
