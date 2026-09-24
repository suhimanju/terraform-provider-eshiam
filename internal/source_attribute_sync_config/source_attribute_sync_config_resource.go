// Package source_attribute_sync_config implements the
// example_source_attribute_sync_config sub-resource.
package source_attribute_sync_config

import (
	"context"

	"terraform-provider-eshiam/internal/crud"
	"terraform-provider-eshiam/internal/util"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

var (
	_ resource.Resource              = &sourceAttributeSyncConfigResource{}
	_ resource.ResourceWithConfigure = &sourceAttributeSyncConfigResource{}
)

// NewSourceAttributeSyncConfigResource is the constructor registered with the provider.
func NewSourceAttributeSyncConfigResource() resource.Resource {
	return &sourceAttributeSyncConfigResource{crud.SubResource[sourceAttributeSyncConfigModel, *sourceAttributeSyncConfigModel]{
		TypeNameSuffix: "_source_attribute_sync_config",
		CollectionPath: func(m *sourceAttributeSyncConfigModel) string { return m.path() },
		ObjectPath:     func(m *sourceAttributeSyncConfigModel) string { return m.path() },
		Singleton:      true,
		SchemaFn:       sourceAttributeSyncConfigSchema,
	}}
}

type sourceAttributeSyncConfigResource struct {
	crud.SubResource[sourceAttributeSyncConfigModel, *sourceAttributeSyncConfigModel]
}

func sourceAttributeSyncConfigSchema(_ context.Context) schema.Schema {
	return schema.Schema{
		Description: "Manages the attribute-synchronization configuration of a source.",
		Attributes: map[string]schema.Attribute{
			"source": util.ResourceReferenceSchema("SOURCE", true, "The source to configure."),
			"attributes_sync": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name":         schema.StringAttribute{Required: true},
						"display_name": schema.StringAttribute{Optional: true, Computed: true},
						"enabled":      schema.BoolAttribute{Optional: true, Computed: true},
						"target":       schema.StringAttribute{Optional: true, Computed: true},
					},
				},
			},
		},
	}
}
