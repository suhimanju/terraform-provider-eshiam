# example_transform (Resource)

Manages a transform used to manipulate attribute values during aggregation and provisioning.

## Example Usage

```terraform
# Manages a transform.
resource "example_transform" "demo" {
  name = "demo-transform"
  type = "lower"

  attributes = jsonencode({
    input = {
      type = "accountAttribute"
      attributes = {
        sourceName    = "HR Source"
        attributeName = "firstName"
      }
    }
  })
}
```

## Import

Import is supported using the resource id:

```shell
terraform import example_transform.example <id>
```

> This page is a hand-written stub. In a published provider these docs are
> generated from the schema with `tfplugindocs`.
