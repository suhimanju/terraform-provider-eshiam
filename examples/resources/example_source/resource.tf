# Manages a source.
resource "example_source" "hr" {
  name        = "HR Source"
  description = "Human Resources delimited file source"
  type        = "DelimitedFile"
  connector   = "delimited-file-angularsc"

  owner = {
    id = "2c9180835d191a86015d28455b4b232a"
  }

  cluster = {
    id = "2c9180866166b5b0016167c32ef31a66"
  }

  connector_attributes = jsonencode({
    deleteThresholdPercentage = 10
  })
}
