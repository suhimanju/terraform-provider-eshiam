# Manages a tenant password policy.
resource "example_password_policy" "standard" {
  name        = "Standard Policy"
  description = "Baseline password policy"

  min_length = 8
  max_length = 64
}
