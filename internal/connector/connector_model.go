package connector

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// connectorModel mirrors the custom-connector resource, keyed by script_name.
type connectorModel struct {
	Name                 types.String `tfsdk:"name"`
	ScriptName           types.String `tfsdk:"script_name"`
	Type                 types.String `tfsdk:"type"`
	ClassName            types.String `tfsdk:"class_name"`
	DirectConnect        types.Bool   `tfsdk:"direct_connect"`
	Status               types.String `tfsdk:"status"`
	ApplicationXml       types.String `tfsdk:"application_xml"`
	SourceConfigXml      types.String `tfsdk:"source_config_xml"`
	CorrelationConfigXml types.String `tfsdk:"correlation_config_xml"`
	ConnectorMetadata    types.String `tfsdk:"connector_metadata"`
}

// GetID returns the script name, the connector's identifier.
func (m *connectorModel) GetID() string { return m.ScriptName.ValueString() }

// connectorDataSourceModel is the connector data source state model.
type connectorDataSourceModel struct {
	Name       types.String `tfsdk:"name"`
	Type       types.String `tfsdk:"type"`
	ScriptName types.String `tfsdk:"script_name"`
}
