// Package org_config implements the example_org_config singleton resource.
package org_config

import (
	"context"
	"net/http"

	"terraform-provider-eshiam/internal/crud"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

var (
	_ resource.Resource              = &orgConfigResource{}
	_ resource.ResourceWithConfigure = &orgConfigResource{}
)

// NewOrgConfigResource is the constructor registered with the provider.
func NewOrgConfigResource() resource.Resource {
	return &orgConfigResource{crud.SingletonResource[orgConfigModel, *orgConfigModel]{
		TypeNameSuffix: "_org_config",
		Path:           "/v3/org-config",
		WriteMethod:    http.MethodPatch,
		SchemaFn:       orgConfigSchema,
	}}
}

type orgConfigResource struct {
	crud.SingletonResource[orgConfigModel, *orgConfigModel]
}

func orgConfigSchema(_ context.Context) schema.Schema {
	return schema.Schema{
		Description: "Manages tenant org configuration.",
		Attributes: map[string]schema.Attribute{
			"time_zone": schema.StringAttribute{Required: true},
		},
	}
}
