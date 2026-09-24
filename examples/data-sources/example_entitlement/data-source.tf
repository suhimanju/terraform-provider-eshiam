# Looks up an entitlement on a source by attribute and value.
data "example_entitlement" "domain_admins" {
  source_id = example_source.hr.id
  attribute = "memberOf"
  value     = "CN=Domain Admins,CN=Users,DC=example,DC=com"
}

output "entitlement_id" {
  value = data.example_entitlement.domain_admins.id
}
