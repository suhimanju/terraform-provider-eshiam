package connector

import (
	"context"
	"fmt"

	"terraform-provider-eshiam/internal/client"
	"terraform-provider-eshiam/internal/crud"
	"terraform-provider-eshiam/internal/util"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &connectorDataSource{}
	_ datasource.DataSourceWithConfigure = &connectorDataSource{}
)

// NewConnectorDataSource is the data source constructor registered with the provider.
func NewConnectorDataSource() datasource.DataSource {
	return &connectorDataSource{}
}

type connectorDataSource struct {
	client *client.APIClient
}

func (d *connectorDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	d.client = crud.ClientFromProviderData(req.ProviderData, &resp.Diagnostics)
}

func (d *connectorDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_connector"
}

func (d *connectorDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Looks up a connector by name.",
		Attributes: map[string]schema.Attribute{
			"name":        schema.StringAttribute{Description: "The connector name.", Required: true},
			"type":        schema.StringAttribute{Description: "The connector type.", Computed: true},
			"script_name": schema.StringAttribute{Description: "The connector script name.", Computed: true},
		},
	}
}

func (d *connectorDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config connectorDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	results, spResp, err := d.client.ListObjects(ctx, "/beta/connectors")
	if err != nil {
		resp.Diagnostics.AddError("Error Reading Connector", err.Error()+"\n"+util.GetBody(spResp))
		return
	}
	for _, item := range results {
		if name, _ := item["name"].(string); name == config.Name.ValueString() {
			config.Type = stringOr(item["type"])
			config.ScriptName = stringOr(item["scriptName"])
			resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
			return
		}
	}
	resp.Diagnostics.AddError("Connector Not Found", fmt.Sprintf("No connector found with name '%s'.", config.Name.ValueString()))
}

func stringOr(v any) types.String {
	if s, ok := v.(string); ok {
		return types.StringValue(s)
	}
	return types.StringNull()
}
