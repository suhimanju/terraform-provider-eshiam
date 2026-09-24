// Package connector implements the example_connector resource and data source.
package connector

import (
	"context"

	"terraform-provider-eshiam/internal/crud"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var (
	_ resource.Resource                = &connectorResource{}
	_ resource.ResourceWithConfigure   = &connectorResource{}
	_ resource.ResourceWithImportState = &connectorResource{}
)

// NewConnectorResource is the constructor registered with the provider.
func NewConnectorResource() resource.Resource {
	return &connectorResource{crud.BaseResource[connectorModel, *connectorModel]{
		TypeNameSuffix: "_connector",
		Endpoint:       "/beta/connectors",
		SchemaFn:       connectorSchema,
	}}
}

type connectorResource struct {
	crud.BaseResource[connectorModel, *connectorModel]
}

func connectorSchema(_ context.Context) schema.Schema {
	return schema.Schema{
		Description: "Manages a custom connector.",
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{Required: true},
			"script_name": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"type":                   schema.StringAttribute{Required: true},
			"class_name":             schema.StringAttribute{Optional: true},
			"direct_connect":         schema.BoolAttribute{Optional: true, Computed: true},
			"status":                 schema.StringAttribute{Optional: true, Computed: true},
			"application_xml":        schema.StringAttribute{Optional: true},
			"source_config_xml":      schema.StringAttribute{Optional: true},
			"correlation_config_xml": schema.StringAttribute{Optional: true},
			"connector_metadata":     schema.StringAttribute{Optional: true},
		},
	}
}
