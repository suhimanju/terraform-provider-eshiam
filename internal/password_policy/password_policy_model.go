package password_policy

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// passwordPolicyModel mirrors the tenant password-policy resource.
type passwordPolicyModel struct {
	Id                                    types.String `tfsdk:"id"`
	Name                                  types.String `tfsdk:"name"`
	Description                           types.String `tfsdk:"description"`
	DefaultPolicy                         types.Bool   `tfsdk:"default_policy"`
	MinLength                             types.Int64  `tfsdk:"min_length"`
	MaxLength                             types.Int64  `tfsdk:"max_length"`
	MinAlpha                              types.Int64  `tfsdk:"min_alpha"`
	MinNumeric                            types.Int64  `tfsdk:"min_numeric"`
	MinUpper                              types.Int64  `tfsdk:"min_upper"`
	MinLower                              types.Int64  `tfsdk:"min_lower"`
	MinSpecial                            types.Int64  `tfsdk:"min_special"`
	MinCharacterTypes                     types.Int64  `tfsdk:"min_character_types"`
	MaxRepeatedChars                      types.Int64  `tfsdk:"max_repeated_chars"`
	FirstExpirationReminder               types.Int64  `tfsdk:"first_expiration_reminder"`
	AccountIdMinWordLength                types.Int64  `tfsdk:"account_id_min_word_length"`
	AccountNameMinWordLength              types.Int64  `tfsdk:"account_name_min_word_length"`
	UseAccountAttributes                  types.Bool   `tfsdk:"use_account_attributes"`
	UseIdentityAttributes                 types.Bool   `tfsdk:"use_identity_attributes"`
	UseDictionary                         types.Bool   `tfsdk:"use_dictionary"`
	EnablePasswdExpiration                types.Bool   `tfsdk:"enable_passwd_expiration"`
	PasswordExpiration                    types.Int64  `tfsdk:"password_expiration"`
	RequireStrongAuthn                    types.Bool   `tfsdk:"require_strong_authn"`
	RequireStrongAuthOffNetwork           types.Bool   `tfsdk:"require_strong_auth_off_network"`
	RequireStrongAuthUntrustedGeographies types.Bool   `tfsdk:"require_strong_auth_untrusted_geographies"`
	ValidateAgainstAccountId              types.Bool   `tfsdk:"validate_against_account_id"`
	ValidateAgainstAccountName            types.Bool   `tfsdk:"validate_against_account_name"`
}

func (m *passwordPolicyModel) GetID() string { return m.Id.ValueString() }
