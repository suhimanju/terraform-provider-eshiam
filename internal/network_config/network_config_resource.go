// Package network_config implements the example_network_config singleton resource.
package network_config

import (
	"context"

	"terraform-provider-eshiam/internal/crud"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource              = &networkConfigResource{}
	_ resource.ResourceWithConfigure = &networkConfigResource{}
)

// NewNetworkConfigResource is the constructor registered with the provider.
func NewNetworkConfigResource() resource.Resource {
	return &networkConfigResource{crud.SingletonResource[networkConfigModel, *networkConfigModel]{
		TypeNameSuffix: "_network_config",
		Path:           "/beta/network-config",
		SchemaFn:       networkConfigSchema,
	}}
}

type networkConfigResource struct {
	crud.SingletonResource[networkConfigModel, *networkConfigModel]
}

func networkConfigSchema(_ context.Context) schema.Schema {
	return schema.Schema{
		Description: "Manages tenant network access configuration.",
		Attributes: map[string]schema.Attribute{
			"whitelisted": schema.BoolAttribute{Optional: true, Computed: true},
			"range":       schema.ListAttribute{ElementType: types.StringType, Optional: true},
			"geolocation": schema.ListAttribute{ElementType: types.StringType, Optional: true},
		},
	}
}
