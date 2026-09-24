# example_provisioning_policy (Resource)

Manages a provisioning policy (account create/update template) of a source.

## Example Usage

```terraform
# Manages a provisioning policy (account template) of a source.
resource "example_provisioning_policy" "create" {
  source_id  = example_source.hr.id
  usage_type = "CREATE"
  name       = "Account Create Policy"

  fields = [
    {
      name        = "userName"
      type        = "string"
      is_required = true
      attributes  = jsonencode({ template = "$${firstname}.$${lastname}" })
    }
  ]
}
```

## Import

Import is supported using the resource id:

```shell
terraform import example_provisioning_policy.example <id>
```

> This page is a hand-written stub. In a published provider these docs are
> generated from the schema with `tfplugindocs`.
