# example_connector_rule (Resource)

Manages a connector rule (BeanShell rule executed by the connector).

## Example Usage

```terraform
# Manages a connector rule.
resource "example_connector_rule" "build_map" {
  name        = "BuildMapRule"
  description = "Custom BuildMap rule"
  type        = "BuildMap"

  source_code = jsonencode({
    version = "1.0"
    script  = "return null;"
  })
}
```

## Import

Import is supported using the resource id:

```shell
terraform import example_connector_rule.example <id>
```

> This page is a hand-written stub. In a published provider these docs are
> generated from the schema with `tfplugindocs`.
