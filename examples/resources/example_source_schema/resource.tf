# Manages the account schema of a source.
resource "example_source_schema" "account" {
  source_id           = example_source.hr.id
  name                = "account"
  native_object_type  = "User"
  identity_attribute  = "id"
  display_attribute   = "name"
  include_permissions = false

  attributes = jsonencode([
    { name = "id", type = "STRING", isMultiValued = false },
    { name = "name", type = "STRING", isMultiValued = false }
  ])
}
