# example_source_attribute_sync_config (Resource)

Manages the attribute-synchronization configuration of a source.

## Example Usage

```terraform
# Manages the attribute-sync configuration of a source.
resource "example_source_attribute_sync_config" "hr" {
  source_id = example_source.hr.id

  attributes = jsonencode([
    { name = "email", displayName = "Email", enabled = true, target = "mail" }
  ])
}
```

## Import

Import is supported using the resource id:

```shell
terraform import example_source_attribute_sync_config.example <id>
```

> This page is a hand-written stub. In a published provider these docs are
> generated from the schema with `tfplugindocs`.
