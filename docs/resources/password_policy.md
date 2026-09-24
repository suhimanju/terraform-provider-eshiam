# example_password_policy (Resource)

Manages a tenant password policy.

## Example Usage

```terraform
# Manages a tenant password policy.
resource "example_password_policy" "standard" {
  name        = "Standard Policy"
  description = "Baseline password policy"

  min_length = 8
  max_length = 64
}
```

## Import

Import is supported using the resource id:

```shell
terraform import example_password_policy.example <id>
```

> This page is a hand-written stub. In a published provider these docs are
> generated from the schema with `tfplugindocs`.
