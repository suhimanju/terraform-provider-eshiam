package connector_rule

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// connectorRuleModel mirrors the connector-rule resource.
type connectorRuleModel struct {
	Id          types.String    `tfsdk:"id"`
	Name        types.String    `tfsdk:"name"`
	Description types.String    `tfsdk:"description"`
	Type        types.String    `tfsdk:"type"`
	Signature   jsontypes.Exact `tfsdk:"signature"`
	SourceCode  jsontypes.Exact `tfsdk:"source_code"`
	Attributes  jsontypes.Exact `tfsdk:"attributes"`
}

func (m *connectorRuleModel) GetID() string { return m.Id.ValueString() }
