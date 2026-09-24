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
