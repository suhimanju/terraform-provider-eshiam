# example_identity_profile (Resource)

Manages an identity profile and its authoritative source mapping.

## Example Usage

```terraform
# Manages an identity profile.
resource "example_identity_profile" "employees" {
  name        = "Employees"
  description = "Employee identity profile"

  owner = {
    id = "2c9180835d191a86015d28455b4b232a"
  }

  authoritative_source = {
    id = example_source.hr.id
  }
}
```

## Import

Import is supported using the resource id:

```shell
terraform import example_identity_profile.example <id>
```

> This page is a hand-written stub. In a published provider these docs are
> generated from the schema with `tfplugindocs`.
