# Manages a custom connector.
resource "example_connector" "custom" {
  name           = "My Custom Connector"
  type           = "custom-connector"
  class_name     = "sailpoint.connector.OpenConnectorAdapter"
  direct_connect = true
  status         = "RELEASED"
}
