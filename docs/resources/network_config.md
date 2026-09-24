# example_network_config (Resource)

Manages tenant network access configuration (singleton).

## Example Usage

```terraform
# Manages tenant network access configuration (singleton).
resource "example_network_config" "this" {
  range     = ["1.2.3.4/32"]
  whitelist = true
}
```

## Import

Import is supported using the resource id:

```shell
terraform import example_network_config.example <id>
```

> This page is a hand-written stub. In a published provider these docs are
> generated from the schema with `tfplugindocs`.
