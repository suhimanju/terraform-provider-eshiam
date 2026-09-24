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
