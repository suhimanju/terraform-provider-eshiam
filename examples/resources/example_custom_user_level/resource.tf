# Manages a custom user (admin) level.
resource "example_custom_user_level" "helpdesk" {
  name        = "Helpdesk"
  description = "Limited helpdesk permissions"

  owner = {
    id = "2c9180835d191a86015d28455b4b232a"
  }

  right_sets = ["idn:password-reset", "idn:account-unlock"]
}
