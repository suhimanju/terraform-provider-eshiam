package identity

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
	_ datasource.DataSource              = &identityDataSource{}
	_ datasource.DataSourceWithConfigure = &identityDataSource{}
)

// NewIdentityDataSource is the constructor registered with the provider.
func NewIdentityDataSource() datasource.DataSource {
	return &identityDataSource{}
}

type identityDataSource struct {
	client *client.APIClient
}

type identityDataSourceModel struct {
	Id    types.String `tfsdk:"id"`
	Alias types.String `tfsdk:"alias"`
	Name  types.String `tfsdk:"name"`
	Email types.String `tfsdk:"email"`
}

func (d *identityDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	d.client = crud.ClientFromProviderData(req.ProviderData, &resp.Diagnostics)
}

func (d *identityDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_identity"
}

func (d *identityDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Looks up an identity by alias (account name).",
		Attributes: map[string]schema.Attribute{
			"alias": schema.StringAttribute{Description: "The identity alias / account name.", Required: true},
			"id":    schema.StringAttribute{Description: "The identity id.", Computed: true},
			"name":  schema.StringAttribute{Description: "The identity display name.", Computed: true},
			"email": schema.StringAttribute{Description: "The identity email.", Computed: true},
		},
	}
}

func (d *identityDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config identityDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	uri := fmt.Sprintf("/v3/public-identities?filters=alias eq \"%s\"", config.Alias.ValueString())
	results, spResp, err := d.client.ListObjects(ctx, uri)
	if err != nil {
		resp.Diagnostics.AddError("Error Reading Identity", err.Error()+"\n"+util.GetBody(spResp))
		return
	}
	for _, item := range results {
		if alias, _ := item["alias"].(string); alias == config.Alias.ValueString() {
			config.Id = stringOr(item["id"])
			config.Name = stringOr(item["name"])
			config.Email = stringOr(item["email"])
			resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
			return
		}
	}
	resp.Diagnostics.AddError("Identity Not Found", fmt.Sprintf("No identity found with alias '%s'.", config.Alias.ValueString()))
}

func stringOr(v any) types.String {
	if s, ok := v.(string); ok {
		return types.StringValue(s)
	}
	return types.StringNull()
}
