# example_connector (Resource)

Manages a custom connector definition.

## Example Usage

```terraform
# Manages a custom connector.
resource "example_connector" "custom" {
  name           = "My Custom Connector"
  type           = "custom-connector"
  class_name     = "sailpoint.connector.OpenConnectorAdapter"
  direct_connect = true
  status         = "RELEASED"
}
```

## Import

Import is supported using the resource id:

```shell
terraform import example_connector.example <id>
```

> This page is a hand-written stub. In a published provider these docs are
> generated from the schema with `tfplugindocs`.
