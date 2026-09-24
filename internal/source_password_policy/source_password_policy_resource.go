// Package source_password_policy implements the example_source_password_policy
// sub-resource.
package source_password_policy

import (
	"context"

	"terraform-provider-eshiam/internal/crud"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var (
	_ resource.Resource              = &sourcePasswordPolicyResource{}
	_ resource.ResourceWithConfigure = &sourcePasswordPolicyResource{}
)

// NewSourcePasswordPolicyResource is the constructor registered with the provider.
func NewSourcePasswordPolicyResource() resource.Resource {
	return &sourcePasswordPolicyResource{crud.SubResource[sourcePasswordPolicyModel, *sourcePasswordPolicyModel]{
		TypeNameSuffix: "_source_password_policy",
		CollectionPath: func(m *sourcePasswordPolicyModel) string { return m.path() },
		ObjectPath:     func(m *sourcePasswordPolicyModel) string { return m.path() },
		Singleton:      true,
		SchemaFn:       sourcePasswordPolicySchema,
	}}
}

type sourcePasswordPolicyResource struct {
	crud.SubResource[sourcePasswordPolicyModel, *sourcePasswordPolicyModel]
}

func sourcePasswordPolicySchema(_ context.Context) schema.Schema {
	return schema.Schema{
		Description: "Associates password policies with a source.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"source_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"password_policies": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"policy_id":   schema.StringAttribute{Required: true},
						"policy_name": schema.StringAttribute{Optional: true, Computed: true},
						"selectors": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"identity_attr": schema.ListNestedAttribute{
									Optional: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name":  schema.StringAttribute{Required: true},
											"value": schema.StringAttribute{Required: true},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
