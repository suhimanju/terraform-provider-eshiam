# Manages a connector rule.
resource "example_connector_rule" "build_map" {
  name        = "BuildMapRule"
  description = "Custom BuildMap rule"
  type        = "BuildMap"

  source_code = jsonencode({
    version = "1.0"
    script  = "return null;"
  })
}
