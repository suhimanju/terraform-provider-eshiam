# example_custom_form (Resource)

Manages a custom form definition.

## Example Usage

```terraform
# Manages a custom form definition.
resource "example_custom_form" "access_request" {
  name        = "Access Request Form"
  description = "Collects justification"

  owner = {
    id = "2c9180835d191a86015d28455b4b232a"
  }
}
```

## Import

Import is supported using the resource id:

```shell
terraform import example_custom_form.example <id>
```

> This page is a hand-written stub. In a published provider these docs are
> generated from the schema with `tfplugindocs`.
