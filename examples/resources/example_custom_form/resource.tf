# Manages a custom form definition.
resource "example_custom_form" "access_request" {
  name        = "Access Request Form"
  description = "Collects justification"

  owner = {
    id = "2c9180835d191a86015d28455b4b232a"
  }
}
