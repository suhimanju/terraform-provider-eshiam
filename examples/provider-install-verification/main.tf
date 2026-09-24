terraform {
  required_version = ">=1.7.1"
  required_providers {
    example = {
      source  = "eshiam-corp/eshiam"
      version = ">= 0.1.0"
    }
  }
}

provider "example" {
  # Host and credentials can also be supplied via EXAMPLE_HOST,
  # EXAMPLE_CLIENT_ID, and EXAMPLE_CLIENT_SECRET environment variables.
  host = "https://api.example.com"
}

# A minimal data source lookup used to verify the provider is installed and
# able to authenticate against the target tenant.
data "example_identity" "default_owner" {
  alias = "example.user"
}

# A minimal resource to verify create/read/update/delete round-trips.
resource "example_identity_attribute" "demo" {
  name         = "test"
  display_name = "test"
  type         = "string"
  sources = [
    {
      type = "rule"
      properties = jsonencode({
        ruleType = "IdentityAttribute"
        ruleName = "Cloud Promote Identity Attribute"
      })
    }
  ]
}
