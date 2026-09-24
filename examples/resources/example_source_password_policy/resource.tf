# Associates password policies with a source.
resource "example_source_password_policy" "hr" {
  source_id = example_source.hr.id

  password_policies = jsonencode([
    { policyId = "2c91808573e5 f0e2", selectors = {} }
  ])
}
