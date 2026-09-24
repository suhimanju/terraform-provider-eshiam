# Manages account correlation configuration for a source (singleton).
resource "example_source_correlation_config" "hr" {
  source_id = example_source.hr.id

  attribute_assignments = jsonencode([
    { property = "email", operation = "EQ", value = "$mail" }
  ])
}
