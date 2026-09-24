# example_identity_attribute (Resource)

Manages an identity attribute and its attribute sources.

## Example Usage

```terraform
# Manages an identity attribute.
resource "example_identity_attribute" "cost_center" {
  name         = "costCenter"
  display_name = "Cost Center"
  type         = "string"
  searchable   = true
  sources = [
    {
      type = "rule"
      properties = jsonencode({
        ruleType = "IdentityAttribute"
        ruleName = "Cloud Promote Identity Attribute"
      })
    }
  ]
}
```

## Import

Import is supported using the resource id:

```shell
terraform import example_identity_attribute.example <id>
```

> This page is a hand-written stub. In a published provider these docs are
> generated from the schema with `tfplugindocs`.
