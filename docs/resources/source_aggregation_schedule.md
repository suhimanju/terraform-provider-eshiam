# example_source_aggregation_schedule (Resource)

Manages the account-aggregation cron schedule of a source.

## Example Usage

```terraform
# Manages the account-aggregation schedule of a source.
resource "example_source_aggregation_schedule" "daily" {
  source_id       = example_source.hr.id
  type            = "ACCOUNT_AGGREGATION"
  cron_expression = "0 0 6 * * ?"
}
```

## Import

Import is supported using the resource id:

```shell
terraform import example_source_aggregation_schedule.example <id>
```

> This page is a hand-written stub. In a published provider these docs are
> generated from the schema with `tfplugindocs`.
