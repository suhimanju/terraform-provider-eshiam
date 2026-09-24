// Package lifecycle_state implements the example_lifecycle_state sub-resource of
// an identity profile, demonstrating crud.SubResource with nested blocks.
package lifecycle_state

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
	_ resource.Resource              = &lifecycleStateResource{}
	_ resource.ResourceWithConfigure = &lifecycleStateResource{}
)

// NewLifecycleStateResource is the constructor registered with the provider.
func NewLifecycleStateResource() resource.Resource {
	return &lifecycleStateResource{crud.SubResource[lifecycleStateModel, *lifecycleStateModel]{
		TypeNameSuffix: "_lifecycle_state",
		CollectionPath: func(m *lifecycleStateModel) string { return m.collectionPath() },
		ObjectPath:     func(m *lifecycleStateModel) string { return m.objectPath() },
		UpdateMethod:   "PATCH",
		SchemaFn:       lifecycleStateSchema,
	}}
}

type lifecycleStateResource struct {
	crud.SubResource[lifecycleStateModel, *lifecycleStateModel]
}

func lifecycleStateSchema(_ context.Context) schema.Schema {
	return schema.Schema{
		Description: "Manages a lifecycle state of an identity profile.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"identity_profile_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name":           schema.StringAttribute{Required: true},
			"technical_name": schema.StringAttribute{Required: true},
			"enabled":        schema.BoolAttribute{Optional: true, Computed: true},
			"description":    schema.StringAttribute{Optional: true},
			"identity_state": schema.StringAttribute{Optional: true, Computed: true},
			"access_profile_ids": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"email_notification_option": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"notify_managers":       schema.BoolAttribute{Optional: true},
					"notify_all_admins":     schema.BoolAttribute{Optional: true},
					"notify_specific_users": schema.BoolAttribute{Optional: true},
					"email_address_list": schema.ListAttribute{
						ElementType: types.StringType,
						Optional:    true,
					},
				},
			},
			"account_actions": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"action":             schema.StringAttribute{Required: true},
						"all_sources":        schema.BoolAttribute{Optional: true},
						"source_ids":         schema.ListAttribute{ElementType: types.StringType, Optional: true},
						"exclude_source_ids": schema.ListAttribute{ElementType: types.StringType, Optional: true},
					},
				},
			},
			"access_action_configuration": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"remove_all_access_enabled": schema.BoolAttribute{Optional: true},
				},
			},
		},
	}
}
