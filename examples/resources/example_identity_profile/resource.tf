# Manages an identity profile.
resource "example_identity_profile" "employees" {
  name        = "Employees"
  description = "Employee identity profile"

  owner = {
    id = "2c9180835d191a86015d28455b4b232a"
  }

  authoritative_source = {
    id = example_source.hr.id
  }
}
