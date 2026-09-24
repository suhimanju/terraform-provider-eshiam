package source_password_policy

import "github.com/hashicorp/terraform-plugin-framework/types"

// sourcePasswordPolicyModel mirrors the association of password policies with a
// source (a singleton sub-resource).
type sourcePasswordPolicyModel struct {
	Id               types.String                      `tfsdk:"id"`
	SourceId         types.String                      `tfsdk:"source_id"`
	PasswordPolicies []sourcePasswordPolicyHolderModel `tfsdk:"password_policies"`
}

type sourcePasswordPolicyHolderModel struct {
	PolicyId   types.String                        `tfsdk:"policy_id"`
	PolicyName types.String                        `tfsdk:"policy_name"`
	Selectors  *sourcePasswordPolicySelectorsModel `tfsdk:"selectors"`
}

type sourcePasswordPolicySelectorsModel struct {
	IdentityAttr []sourcePasswordPolicyIdentityAttrModel `tfsdk:"identity_attr"`
}

type sourcePasswordPolicyIdentityAttrModel struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

func (m *sourcePasswordPolicyModel) GetID() string { return m.SourceId.ValueString() }

func (m *sourcePasswordPolicyModel) path() string {
	return "/beta/sources/" + m.SourceId.ValueString() + "/source-config-password-policies"
}
