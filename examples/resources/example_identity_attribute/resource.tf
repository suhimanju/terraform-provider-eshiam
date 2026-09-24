# Manages an identity attribute.
resource "example_identity_attribute" "cost_center" {
  name         = "costCenter"
  display_name = "Cost Center"
  type         = "string"
  searchable   = true
  sources = [
    {
      type = "rule"
      properties = jsonencode({
        ruleType = "IdentityAttribute"
        ruleName = "Cloud Promote Identity Attribute"
      })
    }
  ]
}
