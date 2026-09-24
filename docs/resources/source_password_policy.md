# example_source_password_policy (Resource)

Associates password policies with a source.

## Example Usage

```terraform
# Associates password policies with a source.
resource "example_source_password_policy" "hr" {
  source_id = example_source.hr.id

  password_policies = jsonencode([
    { policyId = "2c91808573e5 f0e2", selectors = {} }
  ])
}
```

## Import

Import is supported using the resource id:

```shell
terraform import example_source_password_policy.example <id>
```

> This page is a hand-written stub. In a published provider these docs are
> generated from the schema with `tfplugindocs`.
