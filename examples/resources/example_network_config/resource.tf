# Manages tenant network access configuration (singleton).
resource "example_network_config" "this" {
  range     = ["1.2.3.4/32"]
  whitelist = true
}
