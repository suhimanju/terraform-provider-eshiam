# example_parameter_storage (Resource)

Manages a parameter storage entry.

## Example Usage

```terraform
# Manages a parameter storage entry.
resource "example_parameter_storage" "api_scope" {
  name       = "oauth-scope"
  auth_scope = "read:all"
}
```

## Import

Import is supported using the resource id:

```shell
terraform import example_parameter_storage.example <id>
```

> This page is a hand-written stub. In a published provider these docs are
> generated from the schema with `tfplugindocs`.
