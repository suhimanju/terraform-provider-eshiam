# example_app (Resource)

Manages a source app.

## Example Usage

```terraform
# Manages a source app.
resource "example_app" "salesforce" {
  name        = "Salesforce"
  description = "Salesforce source app"
  enabled     = true

  account_source = {
    id = example_source.hr.id
  }
}
```

## Import

Import is supported using the resource id:

```shell
terraform import example_app.example <id>
```

> This page is a hand-written stub. In a published provider these docs are
> generated from the schema with `tfplugindocs`.
