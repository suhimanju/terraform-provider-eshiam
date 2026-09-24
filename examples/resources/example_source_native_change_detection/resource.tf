# Manages native change detection for a source.
resource "example_source_native_change_detection" "hr" {
  source_id = example_source.hr.id
  enabled   = true

  operations                     = ["ACCOUNT_UPDATED", "ACCOUNT_CREATED", "ACCOUNT_DELETED"]
  all_entitlements               = true
  all_non_entitlement_attributes = true
}
