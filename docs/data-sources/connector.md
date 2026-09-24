# example_connector (Data Source)

Looks up a connector by name.

## Example Usage

```terraform
# Looks up a connector by name.
data "example_connector" "active_directory" {
  name = "Active Directory - Direct"
}

output "connector_type" {
  value = data.example_connector.active_directory.type
}
```
