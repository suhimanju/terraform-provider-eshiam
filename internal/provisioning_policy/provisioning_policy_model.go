package provisioning_policy

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// provisioningPolicyModel mirrors a provisioning policy, a sub-resource of a
// source keyed by usage_type (for example CREATE).
type provisioningPolicyModel struct {
	SourceId    types.String `tfsdk:"source_id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	UsageType   types.String `tfsdk:"usage_type"`
	Fields      []fieldModel `tfsdk:"fields"`
}

type fieldModel struct {
	Name          types.String    `tfsdk:"name"`
	Transform     jsontypes.Exact `tfsdk:"transform"`
	Attributes    jsontypes.Exact `tfsdk:"attributes"`
	IsRequired    types.Bool      `tfsdk:"is_required"`
	Type          types.String    `tfsdk:"type"`
	IsMultiValued types.Bool      `tfsdk:"is_multi_valued"`
}

// GetID returns the usage type, which keys the policy within its source.
func (m *provisioningPolicyModel) GetID() string { return m.UsageType.ValueString() }

func (m *provisioningPolicyModel) collectionPath() string {
	return "/v3/sources/" + m.SourceId.ValueString() + "/provisioning-policies"
}

func (m *provisioningPolicyModel) objectPath() string {
	return m.collectionPath() + "/" + m.UsageType.ValueString()
}
