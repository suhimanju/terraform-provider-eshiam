package entitlement

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
	_ datasource.DataSource              = &entitlementDataSource{}
	_ datasource.DataSourceWithConfigure = &entitlementDataSource{}
)

// NewEntitlementDataSource is the constructor registered with the provider.
func NewEntitlementDataSource() datasource.DataSource {
	return &entitlementDataSource{}
}

type entitlementDataSource struct {
	client *client.APIClient
}

type entitlementDataSourceModel struct {
	Id        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	Attribute types.String `tfsdk:"attribute"`
	Value     types.String `tfsdk:"value"`
	SourceId  types.String `tfsdk:"source_id"`
}

func (d *entitlementDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	d.client = crud.ClientFromProviderData(req.ProviderData, &resp.Diagnostics)
}

func (d *entitlementDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_entitlement"
}

func (d *entitlementDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Looks up an entitlement on a source by attribute and value.",
		Attributes: map[string]schema.Attribute{
			"source_id": schema.StringAttribute{Description: "The source id owning the entitlement.", Required: true},
			"attribute": schema.StringAttribute{Description: "The entitlement attribute name.", Required: true},
			"value":     schema.StringAttribute{Description: "The entitlement value.", Required: true},
			"id":        schema.StringAttribute{Description: "The entitlement id.", Computed: true},
			"name":      schema.StringAttribute{Description: "The entitlement display name.", Computed: true},
		},
	}
}

func (d *entitlementDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config entitlementDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	uri := fmt.Sprintf("/beta/entitlements?filters=source.id eq \"%s\" and attribute eq \"%s\" and value eq \"%s\"",
		config.SourceId.ValueString(), config.Attribute.ValueString(), config.Value.ValueString())
	results, spResp, err := d.client.ListObjects(ctx, uri)
	if err != nil {
		resp.Diagnostics.AddError("Error Reading Entitlement", err.Error()+"\n"+util.GetBody(spResp))
		return
	}
	for _, item := range results {
		attr, _ := item["attribute"].(string)
		value, _ := item["value"].(string)
		if attr == config.Attribute.ValueString() && value == config.Value.ValueString() {
			config.Id = stringOr(item["id"])
			config.Name = stringOr(item["name"])
			resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
			return
		}
	}
	resp.Diagnostics.AddError("Entitlement Not Found",
		fmt.Sprintf("No entitlement found for attribute '%s' value '%s'.", config.Attribute.ValueString(), config.Value.ValueString()))
}

func stringOr(v any) types.String {
	if s, ok := v.(string); ok {
		return types.StringValue(s)
	}
	return types.StringNull()
}
