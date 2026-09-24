# example Provider

The **example** provider is a generic, vendor-neutral Terraform provider
boilerplate built on the
[terraform-plugin-framework](https://developer.hashicorp.com/terraform/plugin/framework).
It demonstrates a complete, production-shaped provider: an OAuth2
client-credentials API client, reusable CRUD/patch/util packages, and 26
resources plus 4 data sources that follow one consistent pattern.

Replace `example` and the API endpoints with your own to build a real provider.

## Example Usage

```terraform
terraform {
  required_providers {
    example = {
      source = "eshiam-corp/eshiam"
    }
  }
}

provider "example" {
  host          = "https://api.example.com"
  client_id     = var.client_id
  client_secret = var.client_secret
}
```

## Authentication

Credentials can be supplied in the provider block or via environment variables:

| Setting         | Environment variable    |
|-----------------|-------------------------|
| `host`          | `EXAMPLE_HOST`          |
| `client_id`     | `EXAMPLE_CLIENT_ID`     |
| `client_secret` | `EXAMPLE_CLIENT_SECRET` |

## Schema

### Optional

- `host` (String) Base URL for the API.
- `client_id` (String) OAuth2 client ID.
- `client_secret` (String, Sensitive) OAuth2 client secret.

## Resources

| Resource | Description |
|----------|-------------|
| `example_app` | Manages a source app. |
| `example_connector` | Manages a custom connector. |
| `example_connector_rule` | Manages a connector rule. |
| `example_custom_form` | Manages a custom form. |
| `example_custom_user_level` | Manages a custom user level. |
| `example_identity_attribute` | Manages an identity attribute. |
| `example_identity_profile` | Manages an identity profile. |
| `example_launcher` | Manages a launcher. |
| `example_lifecycle_state` | Manages a lifecycle state. |
| `example_network_config` | Manages tenant network config. |
| `example_org_config` | Manages tenant org config. |
| `example_parameter_storage` | Manages a parameter storage entry. |
| `example_password_policy` | Manages a password policy. |
| `example_provisioning_policy` | Manages a provisioning policy. |
| `example_role` | Manages a role. |
| `example_service_desk_integration` | Manages a service desk integration. |
| `example_service_desk_integration_time_check_config` | Manages SDIM SLA time-check config. |
| `example_source` | Manages a source. |
| `example_source_aggregation_schedule` | Manages a source aggregation schedule. |
| `example_source_attribute_sync_config` | Manages source attribute-sync config. |
| `example_source_correlation_config` | Manages source correlation config. |
| `example_source_native_change_detection` | Manages source native change detection. |
| `example_source_password_policy` | Manages source password policies. |
| `example_source_schema` | Manages a source schema. |
| `example_transform` | Manages a transform. |
| `example_workflow` | Manages a workflow. |

## Data Sources

| Data Source | Description |
|-------------|-------------|
| `example_cluster` | Looks up a managed cluster by name. |
| `example_connector` | Looks up a connector by name. |
| `example_entitlement` | Looks up an entitlement by attribute/value. |
| `example_identity` | Looks up an identity by alias. |
