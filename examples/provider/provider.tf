terraform {
  required_providers {
    example = {
      source = "eshiam-corp/eshiam"
    }
  }
}

provider "example" {
  # These can also be supplied via EXAMPLE_HOST, EXAMPLE_CLIENT_ID,
  # and EXAMPLE_CLIENT_SECRET environment variables.
  host          = "https://api.example.com"
  client_id     = var.client_id
  client_secret = var.client_secret
}

variable "client_id" {
  type = string
}

variable "client_secret" {
  type      = string
  sensitive = true
}
