# Manages a parameter storage entry.
resource "example_parameter_storage" "api_scope" {
  name       = "oauth-scope"
  auth_scope = "read:all"
}
