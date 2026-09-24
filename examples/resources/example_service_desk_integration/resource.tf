# Manages a service desk integration.
resource "example_service_desk_integration" "snow" {
  name        = "ServiceNow SDIM"
  description = "ServiceNow ticketing integration"
  type        = "ServiceNowSDIM"

  owner_ref = {
    id = "2c9180835d191a86015d28455b4b232a"
  }

  cluster_ref = {
    id = "2c9180866166b5b0016167c32ef31a66"
  }

  attributes = jsonencode({
    url            = "https://example.service-now.com"
    authentication = "Basic"
  })
}
