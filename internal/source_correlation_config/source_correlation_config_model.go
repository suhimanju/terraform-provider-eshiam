package source_correlation_config

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// sourceCorrelationConfigModel mirrors the account-correlation configuration of
// a source (a singleton sub-resource).
type sourceCorrelationConfigModel struct {
	SourceId             types.String           `tfsdk:"source_id"`
	Id                   types.String           `tfsdk:"id"`
	Name                 types.String           `tfsdk:"name"`
	AttributeAssignments []attributeAssignments `tfsdk:"attribute_assignments"`
}

type attributeAssignments struct {
	Property     types.String `tfsdk:"property"`
	Value        types.String `tfsdk:"value"`
	Operation    types.String `tfsdk:"operation"`
	Complex      types.Bool   `tfsdk:"complex"`
	IgnoreCase   types.Bool   `tfsdk:"ignore_case"`
	MatchMode    types.String `tfsdk:"match_mode"`
	FilterString types.String `tfsdk:"filter_string"`
}

func (m *sourceCorrelationConfigModel) GetID() string { return m.SourceId.ValueString() }

func (m *sourceCorrelationConfigModel) path() string {
	return "/v3/sources/" + m.SourceId.ValueString() + "/correlation-config"
}
