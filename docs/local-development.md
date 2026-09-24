# Local Development

This guide shows how to build the provider and run it against your own machine
without publishing it to a registry. It assumes you have Go and Terraform
installed (see the "Prerequisites" section in the top-level `README.md`).

## Build and install

```bash
go build ./...   # compile everything (no output means success)
go install .     # install the provider binary into your Go bin folder
```

`go install` places the binary in your Go bin directory. Find its exact path
with:

```bash
go env GOBIN   # if empty, the location is: $(go env GOPATH)/bin
```

Note that path — you need it in the next step.

## Dev overrides

A "dev override" tells Terraform to use your locally built binary instead of
downloading the provider. Create a Terraform CLI config file:

- **Linux/macOS:** `~/.terraformrc`
- **Windows:** `%APPDATA%\terraform.rc`

Put this inside it, replacing the path with your Go bin folder from above:

```hcl
provider_installation {
  dev_overrides {
    # Linux/macOS example:
    "eshiam-corp/eshiam" = "/home/you/go/bin"
    # Windows example (use double backslashes):
    # "eshiam-corp/eshiam" = "C:\\Users\\you\\go\\bin"
  }
  # Install all other providers normally.
  direct {}
}
```

With the override in place, run `terraform plan`/`apply` as usual — and **do not**
run `terraform init` (dev overrides skip it). Terraform prints a warning
reminding you an override is active; that is expected.

> A ready-to-edit template lives at
> `examples/provider-install-verification/.terraformrc`.

## Environment variables

Instead of putting credentials in `.tf` files, you can export them:

```bash
export EXAMPLE_HOST="https://api.example.com"
export EXAMPLE_CLIENT_ID="your-client-id"
export EXAMPLE_CLIENT_SECRET="your-client-secret"
```

On Windows PowerShell:

```powershell
$env:EXAMPLE_HOST="https://api.example.com"
$env:EXAMPLE_CLIENT_ID="your-client-id"
$env:EXAMPLE_CLIENT_SECRET="your-client-secret"
```

## Testing

```bash
# Unit tests (no network access needed)
go test ./...

# Format your code and run the linter before committing
gofmt -w .
golangci-lint run ./...
```

To run tests against the offline mock API instead of a live tenant, see
`mock/README.md`.

