# example_source_native_change_detection (Resource)

Manages native change detection configuration for a source.

## Example Usage

```terraform
# Manages native change detection for a source.
resource "example_source_native_change_detection" "hr" {
  source_id = example_source.hr.id
  enabled   = true

  operations                     = ["ACCOUNT_UPDATED", "ACCOUNT_CREATED", "ACCOUNT_DELETED"]
  all_entitlements               = true
  all_non_entitlement_attributes = true
}
```

## Import

Import is supported using the resource id:

```shell
terraform import example_source_native_change_detection.example <id>
```

> This page is a hand-written stub. In a published provider these docs are
> generated from the schema with `tfplugindocs`.
