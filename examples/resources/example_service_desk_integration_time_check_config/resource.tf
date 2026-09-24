# Manages the SLA time-check configuration of a service desk integration (singleton).
resource "example_service_desk_integration_time_check_config" "this" {
  provisioning_status_check_interval_minutes = "240"
  provisioning_max_status_check_days         = "30"
}
