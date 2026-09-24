// Package provider defines the example Terraform provider, its configuration
// schema, and the registration of resources and data sources.
//
// The authentication pattern (host + client_id + client_secret with env-var
// fallback) is a generic OAuth2 client-credentials flow. Adapt the Configure
// method to your API's auth scheme.
package provider

import (
	"context"
	"os"

	"terraform-provider-eshiam/internal/app"
	"terraform-provider-eshiam/internal/client"
	"terraform-provider-eshiam/internal/cluster"
	"terraform-provider-eshiam/internal/connector"
	"terraform-provider-eshiam/internal/connector_rule"
	"terraform-provider-eshiam/internal/custom_form"
	"terraform-provider-eshiam/internal/custom_user_level"
	"terraform-provider-eshiam/internal/entitlement"
	"terraform-provider-eshiam/internal/identity"
	"terraform-provider-eshiam/internal/identity_attribute"
	"terraform-provider-eshiam/internal/identity_profile"
	"terraform-provider-eshiam/internal/launcher"
	"terraform-provider-eshiam/internal/lifecycle_state"
	"terraform-provider-eshiam/internal/network_config"
	"terraform-provider-eshiam/internal/org_config"
	"terraform-provider-eshiam/internal/parameter_storage"
	"terraform-provider-eshiam/internal/password_policy"
	"terraform-provider-eshiam/internal/provisioning_policy"
	"terraform-provider-eshiam/internal/role"
	"terraform-provider-eshiam/internal/service_desk_integration"
	"terraform-provider-eshiam/internal/service_desk_integration_time_config"
	"terraform-provider-eshiam/internal/source"
	"terraform-provider-eshiam/internal/source_aggregation_schedule"
	"terraform-provider-eshiam/internal/source_attribute_sync_config"
	"terraform-provider-eshiam/internal/source_correlation_config"
	"terraform-provider-eshiam/internal/source_native_change_detection"
	"terraform-provider-eshiam/internal/source_password_policy"
	"terraform-provider-eshiam/internal/source_schema"
	"terraform-provider-eshiam/internal/transform"
	"terraform-provider-eshiam/internal/workflow"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure exampleProvider satisfies the provider.Provider interface.
var _ provider.Provider = &exampleProvider{}

// New returns a factory that constructs the provider with the given version.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &exampleProvider{version: version}
	}
}

// exampleProvider is the provider implementation.
type exampleProvider struct {
	// version is "dev" when run locally, "test" during acceptance tests, and
	// the release version on published builds.
	version string
}

// exampleProviderModel maps provider configuration to a Go type.
type exampleProviderModel struct {
	Host         types.String `tfsdk:"host"`
	ClientId     types.String `tfsdk:"client_id"`
	ClientSecret types.String `tfsdk:"client_secret"`
}

func (p *exampleProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "example"
	resp.Version = p.version
}

func (p *exampleProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Interact with the example API.",
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				Description: "Base URL for the API. May also be set via the EXAMPLE_HOST environment variable.",
				Optional:    true,
			},
			"client_id": schema.StringAttribute{
				Description: "OAuth2 client ID. May also be set via the EXAMPLE_CLIENT_ID environment variable.",
				Optional:    true,
			},
			"client_secret": schema.StringAttribute{
				Description: "OAuth2 client secret. May also be set via the EXAMPLE_CLIENT_SECRET environment variable.",
				Optional:    true,
				Sensitive:   true,
			},
		},
	}
}

func (p *exampleProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring example API client")

	var config exampleProviderModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.Host.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Unknown API Host",
			"Cannot create the API client with an unknown host value. Set it statically or via the EXAMPLE_HOST environment variable.",
		)
	}
	if config.ClientId.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("client_id"),
			"Unknown API client_id",
			"Cannot create the API client with an unknown client_id value. Set it statically or via the EXAMPLE_CLIENT_ID environment variable.",
		)
	}
	if config.ClientSecret.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("client_secret"),
			"Unknown API client_secret",
			"Cannot create the API client with an unknown client_secret value. Set it statically or via the EXAMPLE_CLIENT_SECRET environment variable.",
		)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	// Environment variables provide defaults; explicit config wins.
	host := os.Getenv("EXAMPLE_HOST")
	clientId := os.Getenv("EXAMPLE_CLIENT_ID")
	clientSecret := os.Getenv("EXAMPLE_CLIENT_SECRET")

	if !config.Host.IsNull() {
		host = config.Host.ValueString()
	}
	if !config.ClientId.IsNull() {
		clientId = config.ClientId.ValueString()
	}
	if !config.ClientSecret.IsNull() {
		clientSecret = config.ClientSecret.ValueString()
	}

	if host == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Missing API Host",
			"Set the host in the provider configuration or via the EXAMPLE_HOST environment variable.",
		)
	}
	if clientId == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("client_id"),
			"Missing API client_id",
			"Set the client_id in the provider configuration or via the EXAMPLE_CLIENT_ID environment variable.",
		)
	}
	if clientSecret == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("client_secret"),
			"Missing API client_secret",
			"Set the client_secret in the provider configuration or via the EXAMPLE_CLIENT_SECRET environment variable.",
		)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	apiClient := client.NewAPIClient(client.Config{
		BaseURL:      host,
		ClientId:     clientId,
		ClientSecret: clientSecret,
		TokenURL:     host + "/oauth/token",
	})

	// Make the client available to resources and data sources.
	resp.DataSourceData = apiClient
	resp.ResourceData = apiClient
}

// Resources registers the provider's managed resources.
func (p *exampleProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		identity_attribute.NewIdentityAttributeResource,
		transform.NewTransformResource,
		source.NewSourceResource,
		identity_profile.NewIdentityProfileResource,
		source_schema.NewSourceSchemaResource,
		source_aggregation_schedule.NewSourceAggregationScheduleResource,
		source_attribute_sync_config.NewSourceAttributeSyncConfigResource,
		lifecycle_state.NewLifecycleStateResource,
		connector_rule.NewConnectorRuleResource,
		workflow.NewWorkflowResource,
		role.NewRoleResource,
		org_config.NewOrgConfigResource,
		source_native_change_detection.NewSourceNativeChangeDetectionResource,
		source_password_policy.NewSourcePasswordPolicyResource,
		network_config.NewNetworkConfigResource,
		provisioning_policy.NewProvisioningPolicyResource,
		connector.NewConnectorResource,
		source_correlation_config.NewSourceCorrelationConfigResource,
		service_desk_integration.NewServiceDeskIntegrationResource,
		custom_form.NewCustomFormResource,
		custom_user_level.NewCustomUserLevelResource,
		launcher.NewLauncherResource,
		app.NewAppResource,
		service_desk_integration_time_config.NewServiceDeskIntegrationTimeCheckConfigResource,
		parameter_storage.NewParameterStorageResource,
		password_policy.NewPasswordPolicyResource,
	}
}

// DataSources registers the provider's data sources.
func (p *exampleProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		identity.NewIdentityDataSource,
		cluster.NewClusterDataSource,
		connector.NewConnectorDataSource,
		entitlement.NewEntitlementDataSource,
	}
}
