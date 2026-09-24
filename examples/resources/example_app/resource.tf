# Manages a source app.
resource "example_app" "salesforce" {
  name        = "Salesforce"
  description = "Salesforce source app"
  enabled     = true

  account_source = {
    id = example_source.hr.id
  }
}
