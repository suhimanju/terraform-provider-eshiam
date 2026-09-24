// Package password_policy implements the example_password_policy resource.
package password_policy

import (
	"context"

	"terraform-provider-eshiam/internal/crud"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var (
	_ resource.Resource                = &passwordPolicyResource{}
	_ resource.ResourceWithConfigure   = &passwordPolicyResource{}
	_ resource.ResourceWithImportState = &passwordPolicyResource{}
)

// NewPasswordPolicyResource is the constructor registered with the provider.
func NewPasswordPolicyResource() resource.Resource {
	return &passwordPolicyResource{crud.BaseResource[passwordPolicyModel, *passwordPolicyModel]{
		TypeNameSuffix: "_password_policy",
		Endpoint:       "/v3/password-policies",
		SchemaFn:       passwordPolicySchema,
	}}
}

type passwordPolicyResource struct {
	crud.BaseResource[passwordPolicyModel, *passwordPolicyModel]
}

func passwordPolicySchema(_ context.Context) schema.Schema {
	opt := func() schema.BoolAttribute { return schema.BoolAttribute{Optional: true, Computed: true} }
	num := func() schema.Int64Attribute { return schema.Int64Attribute{Optional: true, Computed: true} }
	return schema.Schema{
		Description: "Manages a tenant password policy.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name":                            schema.StringAttribute{Required: true},
			"description":                     schema.StringAttribute{Optional: true},
			"default_policy":                  opt(),
			"min_length":                      num(),
			"max_length":                      num(),
			"min_alpha":                       num(),
			"min_numeric":                     num(),
			"min_upper":                       num(),
			"min_lower":                       num(),
			"min_special":                     num(),
			"min_character_types":             num(),
			"max_repeated_chars":              num(),
			"first_expiration_reminder":       num(),
			"account_id_min_word_length":      num(),
			"account_name_min_word_length":    num(),
			"use_account_attributes":          opt(),
			"use_identity_attributes":         opt(),
			"use_dictionary":                  opt(),
			"enable_passwd_expiration":        opt(),
			"password_expiration":             num(),
			"require_strong_authn":            opt(),
			"require_strong_auth_off_network": opt(),
			"require_strong_auth_untrusted_geographies": opt(),
			"validate_against_account_id":               opt(),
			"validate_against_account_name":             opt(),
		},
	}
}
