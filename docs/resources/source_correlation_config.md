# example_source_correlation_config (Resource)

Manages account correlation configuration for a source (singleton).

## Example Usage

```terraform
# Manages account correlation configuration for a source (singleton).
resource "example_source_correlation_config" "hr" {
  source_id = example_source.hr.id

  attribute_assignments = jsonencode([
    { property = "email", operation = "EQ", value = "$mail" }
  ])
}
```

## Import

Import is supported using the resource id:

```shell
terraform import example_source_correlation_config.example <id>
```

> This page is a hand-written stub. In a published provider these docs are
> generated from the schema with `tfplugindocs`.
