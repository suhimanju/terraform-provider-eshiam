// Package source_native_change_detection implements the
// example_source_native_change_detection sub-resource.
package source_native_change_detection

import (
	"context"

	"terraform-provider-eshiam/internal/crud"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource              = &sourceNativeChangeDetectionResource{}
	_ resource.ResourceWithConfigure = &sourceNativeChangeDetectionResource{}
)

// NewSourceNativeChangeDetectionResource is the constructor registered with the provider.
func NewSourceNativeChangeDetectionResource() resource.Resource {
	return &sourceNativeChangeDetectionResource{crud.SubResource[sourceNativeChangeDetectionModel, *sourceNativeChangeDetectionModel]{
		TypeNameSuffix: "_source_native_change_detection",
		CollectionPath: func(m *sourceNativeChangeDetectionModel) string { return m.path() },
		ObjectPath:     func(m *sourceNativeChangeDetectionModel) string { return m.path() },
		Singleton:      true,
		SchemaFn:       sourceNativeChangeDetectionSchema,
	}}
}

type sourceNativeChangeDetectionResource struct {
	crud.SubResource[sourceNativeChangeDetectionModel, *sourceNativeChangeDetectionModel]
}

func sourceNativeChangeDetectionSchema(_ context.Context) schema.Schema {
	return schema.Schema{
		Description: "Manages the native-change-detection configuration of a source.",
		Attributes: map[string]schema.Attribute{
			"source_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"enabled":                             schema.BoolAttribute{Required: true},
			"operations":                          schema.SetAttribute{ElementType: types.StringType, Optional: true},
			"all_entitlements":                    schema.BoolAttribute{Optional: true, Computed: true},
			"all_non_entitlement_attributes":      schema.BoolAttribute{Optional: true, Computed: true},
			"selected_entitlements":               schema.SetAttribute{ElementType: types.StringType, Optional: true},
			"selected_non_entitlement_attributes": schema.SetAttribute{ElementType: types.StringType, Optional: true},
		},
	}
}
