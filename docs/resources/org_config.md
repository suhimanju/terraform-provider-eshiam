# example_org_config (Resource)

Manages tenant-wide org configuration (singleton).

## Example Usage

```terraform
# Manages tenant-wide org configuration (singleton).
resource "example_org_config" "this" {
  time_zone = "America/New_York"
}
```

## Import

Import is supported using the resource id:

```shell
terraform import example_org_config.example <id>
```

> This page is a hand-written stub. In a published provider these docs are
> generated from the schema with `tfplugindocs`.
