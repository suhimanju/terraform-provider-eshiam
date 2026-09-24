# Manages a transform.
resource "example_transform" "demo" {
  name = "demo-transform"
  type = "lower"

  attributes = jsonencode({
    input = {
      type = "accountAttribute"
      attributes = {
        sourceName    = "HR Source"
        attributeName = "firstName"
      }
    }
  })
}
