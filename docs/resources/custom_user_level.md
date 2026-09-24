# example_custom_user_level (Resource)

Manages a custom user (admin) level and its right sets.

## Example Usage

```terraform
# Manages a custom user (admin) level.
resource "example_custom_user_level" "helpdesk" {
  name        = "Helpdesk"
  description = "Limited helpdesk permissions"

  owner = {
    id = "2c9180835d191a86015d28455b4b232a"
  }

  right_sets = ["idn:password-reset", "idn:account-unlock"]
}
```

## Import

Import is supported using the resource id:

```shell
terraform import example_custom_user_level.example <id>
```

> This page is a hand-written stub. In a published provider these docs are
> generated from the schema with `tfplugindocs`.
