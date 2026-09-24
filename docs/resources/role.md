# example_role (Resource)

Manages a role and its access model membership.

## Example Usage

```terraform
# Manages a role.
resource "example_role" "finance" {
  name        = "Finance Access"
  description = "Bundle of finance entitlements"

  owner = {
    id = "2c9180835d191a86015d28455b4b232a"
  }
}
```

## Import

Import is supported using the resource id:

```shell
terraform import example_role.example <id>
```

> This page is a hand-written stub. In a published provider these docs are
> generated from the schema with `tfplugindocs`.
