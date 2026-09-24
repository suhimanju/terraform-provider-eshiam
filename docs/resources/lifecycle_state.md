# example_lifecycle_state (Resource)

Manages a lifecycle state of an identity profile.

## Example Usage

```terraform
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
```

## Import

Import is supported using the resource id:

```shell
terraform import example_lifecycle_state.example <id>
```

> This page is a hand-written stub. In a published provider these docs are
> generated from the schema with `tfplugindocs`.
