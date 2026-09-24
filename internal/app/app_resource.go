// Package app implements the example_app resource.
package app

import (
	"context"
	"net/http"

	"terraform-provider-eshiam/internal/crud"
	"terraform-provider-eshiam/internal/util"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var (
	_ resource.Resource                = &appResource{}
	_ resource.ResourceWithConfigure   = &appResource{}
	_ resource.ResourceWithImportState = &appResource{}
)

// NewAppResource is the constructor registered with the provider.
func NewAppResource() resource.Resource {
	return &appResource{crud.BaseResource[appModel, *appModel]{
		TypeNameSuffix: "_app",
		Endpoint:       "/v1/source-apps",
		UpdateMethod:   http.MethodPatch,
		SchemaFn:       appSchema,
	}}
}

type appResource struct {
	crud.BaseResource[appModel, *appModel]
}

func appSchema(_ context.Context) schema.Schema {
	return schema.Schema{
		Description: "Manages a source app.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name": schema.StringAttribute{Required: true},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"enabled": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(true),
			},
			"type": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"account_source": util.ResourceReferenceSchema("SOURCE", false, "The account source backing this app."),
			"owner":          util.ResourceReferenceSchema("IDENTITY", false, "The app owner."),
		},
	}
}
