# Manages a role.
resource "example_role" "finance" {
  name        = "Finance Access"
  description = "Bundle of finance entitlements"

  owner = {
    id = "2c9180835d191a86015d28455b4b232a"
  }
}
