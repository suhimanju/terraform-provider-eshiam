# example_source_schema (Resource)

Manages the object schema (account/group) of a source.

## Example Usage

```terraform
# Manages the account schema of a source.
resource "example_source_schema" "account" {
  source_id           = example_source.hr.id
  name                = "account"
  native_object_type  = "User"
  identity_attribute  = "id"
  display_attribute   = "name"
  include_permissions = false

  attributes = jsonencode([
    { name = "id", type = "STRING", isMultiValued = false },
    { name = "name", type = "STRING", isMultiValued = false }
  ])
}
```

## Import

Import is supported using the resource id:

```shell
terraform import example_source_schema.example <id>
```

> This page is a hand-written stub. In a published provider these docs are
> generated from the schema with `tfplugindocs`.
