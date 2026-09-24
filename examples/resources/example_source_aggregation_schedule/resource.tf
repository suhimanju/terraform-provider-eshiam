# Manages the account-aggregation schedule of a source.
resource "example_source_aggregation_schedule" "daily" {
  source_id       = example_source.hr.id
  type            = "ACCOUNT_AGGREGATION"
  cron_expression = "0 0 6 * * ?"
}
