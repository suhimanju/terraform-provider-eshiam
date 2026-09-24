# example_launcher (Resource)

Manages a launcher.

## Example Usage

```terraform
# Manages a launcher.
resource "example_launcher" "aws_console" {
  name        = "AWS Console"
  description = "Launch AWS console session"
  type        = "INTERACTIVE_PROCESS"
  disabled    = false
}
```

## Import

Import is supported using the resource id:

```shell
terraform import example_launcher.example <id>
```

> This page is a hand-written stub. In a published provider these docs are
> generated from the schema with `tfplugindocs`.
