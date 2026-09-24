# Manages the attribute-sync configuration of a source.
resource "example_source_attribute_sync_config" "hr" {
  source_id = example_source.hr.id

  attributes = jsonencode([
    { name = "email", displayName = "Email", enabled = true, target = "mail" }
  ])
}
