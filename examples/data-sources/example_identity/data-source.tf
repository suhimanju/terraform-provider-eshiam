# Looks up an identity by alias (account name).
data "example_identity" "jdoe" {
  alias = "jdoe"
}

output "identity_id" {
  value = data.example_identity.jdoe.id
}
