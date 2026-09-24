# example_source (Resource)

Manages a source (connector instance) including its connector attributes.

## Example Usage

```terraform
# Manages a source.
resource "example_source" "hr" {
  name        = "HR Source"
  description = "Human Resources delimited file source"
  type        = "DelimitedFile"
  connector   = "delimited-file-angularsc"

  owner = {
    id = "2c9180835d191a86015d28455b4b232a"
  }

  cluster = {
    id = "2c9180866166b5b0016167c32ef31a66"
  }

  connector_attributes = jsonencode({
    deleteThresholdPercentage = 10
  })
}
```

## Import

Import is supported using the resource id:

```shell
terraform import example_source.example <id>
```

> This page is a hand-written stub. In a published provider these docs are
> generated from the schema with `tfplugindocs`.
