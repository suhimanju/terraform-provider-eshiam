package cluster

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
	_ datasource.DataSource              = &clusterDataSource{}
	_ datasource.DataSourceWithConfigure = &clusterDataSource{}
)

// NewClusterDataSource is the constructor registered with the provider.
func NewClusterDataSource() datasource.DataSource {
	return &clusterDataSource{}
}

type clusterDataSource struct {
	client *client.APIClient
}

type clusterDataSourceModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Pod  types.String `tfsdk:"pod"`
	Org  types.String `tfsdk:"org"`
}

func (d *clusterDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	d.client = crud.ClientFromProviderData(req.ProviderData, &resp.Diagnostics)
}

func (d *clusterDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cluster"
}

func (d *clusterDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Looks up a cluster (managed cluster) by name.",
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{Description: "The cluster name.", Required: true},
			"id":   schema.StringAttribute{Description: "The cluster id.", Computed: true},
			"pod":  schema.StringAttribute{Description: "The cluster pod.", Computed: true},
			"org":  schema.StringAttribute{Description: "The cluster org.", Computed: true},
		},
	}
}

func (d *clusterDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config clusterDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	results, spResp, err := d.client.ListObjects(ctx, "/v3/managed-clusters")
	if err != nil {
		resp.Diagnostics.AddError("Error Reading Cluster", err.Error()+"\n"+util.GetBody(spResp))
		return
	}
	for _, item := range results {
		if name, _ := item["name"].(string); name == config.Name.ValueString() {
			config.Id = stringOr(item["id"])
			config.Pod = stringOr(item["pod"])
			config.Org = stringOr(item["org"])
			resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
			return
		}
	}
	resp.Diagnostics.AddError("Cluster Not Found", fmt.Sprintf("No cluster found with name '%s'.", config.Name.ValueString()))
}

func stringOr(v any) types.String {
	if s, ok := v.(string); ok {
		return types.StringValue(s)
	}
	return types.StringNull()
}
