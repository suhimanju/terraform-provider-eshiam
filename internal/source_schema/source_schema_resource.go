// Package source_schema implements the example_source_schema sub-resource of a
// source. It demonstrates crud.SubResource for parent-scoped objects.
package source_schema

import (
	"context"

	"terraform-provider-eshiam/internal/crud"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource              = &sourceSchemaResource{}
	_ resource.ResourceWithConfigure = &sourceSchemaResource{}
)

// NewSourceSchemaResource is the constructor registered with the provider.
func NewSourceSchemaResource() resource.Resource {
	return &sourceSchemaResource{crud.SubResource[sourceSchemaModel, *sourceSchemaModel]{
		TypeNameSuffix: "_source_schema",
		CollectionPath: func(m *sourceSchemaModel) string { return m.collectionPath() },
		ObjectPath:     func(m *sourceSchemaModel) string { return m.objectPath() },
		SchemaFn:       sourceSchemaSchema,
	}}
}

type sourceSchemaResource struct {
	crud.SubResource[sourceSchemaModel, *sourceSchemaModel]
}

func sourceSchemaSchema(_ context.Context) schema.Schema {
	return schema.Schema{
		Description: "Manages the object schema (account/group) of a source.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"source_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name":                schema.StringAttribute{Required: true},
			"native_object_type":  schema.StringAttribute{Optional: true, Computed: true},
			"identity_attribute":  schema.StringAttribute{Optional: true, Computed: true},
			"display_attribute":   schema.StringAttribute{Optional: true, Computed: true},
			"hierarchy_attribute": schema.StringAttribute{Optional: true, Computed: true},
			"include_permissions": schema.BoolAttribute{Optional: true, Computed: true},
			"features":            schema.ListAttribute{ElementType: types.StringType, Optional: true},
			"configuration":       schema.StringAttribute{CustomType: jsontypes.ExactType{}, Optional: true},
			"attributes":          schema.StringAttribute{CustomType: jsontypes.ExactType{}, Optional: true},
		},
	}
}
