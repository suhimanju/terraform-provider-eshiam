# Manages a lifecycle state of an identity profile.
resource "example_lifecycle_state" "active" {
  identity_profile_id = example_identity_profile.employees.id
  name                = "Active"
  technical_name      = "active"
  enabled             = true

  email_notification_option = {
    notify_managers   = true
    notify_all_admins = false
  }

  account_actions = [
    {
      action      = "ENABLE"
      all_sources = true
    }
  ]
}
