# Manages a launcher.
resource "example_launcher" "aws_console" {
  name        = "AWS Console"
  description = "Launch AWS console session"
  type        = "INTERACTIVE_PROCESS"
  disabled    = false
}
