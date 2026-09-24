# example_workflow (Resource)

Manages a workflow definition and trigger.

## Example Usage

```terraform
# Manages a workflow.
resource "example_workflow" "joiner" {
  name        = "Joiner Notification"
  description = "Notifies on new identity"
  enabled     = false

  owner = {
    id = "2c9180835d191a86015d28455b4b232a"
  }

  definition = {
    start = "Send Email"
    steps = jsonencode({
      "Send Email" = { type = "action" }
    })
  }

  trigger = {
    type       = "EVENT"
    attributes = jsonencode({ id = "idn:identity-created" })
  }
}
```

## Import

Import is supported using the resource id:

```shell
terraform import example_workflow.example <id>
```

> This page is a hand-written stub. In a published provider these docs are
> generated from the schema with `tfplugindocs`.
