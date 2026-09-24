# example_service_desk_integration_time_check_config (Resource)

Manages the SLA time-check configuration of a service desk integration (singleton).

## Example Usage

```terraform
# Manages the SLA time-check configuration of a service desk integration (singleton).
resource "example_service_desk_integration_time_check_config" "this" {
  provisioning_status_check_interval_minutes = "240"
  provisioning_max_status_check_days         = "30"
}
```

## Import

Import is supported using the resource id:

```shell
terraform import example_service_desk_integration_time_check_config.example <id>
```

> This page is a hand-written stub. In a published provider these docs are
> generated from the schema with `tfplugindocs`.
