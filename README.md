# Terraform Provider Framework (Go)

A **generic, vendor-neutral boilerplate** for building a Terraform provider with
the [terraform-plugin-framework](https://developer.hashicorp.com/terraform/plugin/framework).
It captures a clean, reusable architecture — provider setup, an OAuth2 API
client, shared utilities, a JSON Patch builder, and a fully worked example
resource and data source — with no customer- or vendor-specific code.

Use it as the starting point for a new provider: rename `example`, plug in your
API client, and add resources following the established pattern.

## Features

- **terraform-plugin-framework** (Protocol 6), not the legacy SDK.
- **OAuth2 client-credentials** API client with automatic token caching/refresh.
- **Generic CRUD bases** in `internal/crud` — build a resource by filling in a
  schema and an endpoint instead of hand-writing Create/Read/Update/Delete.
- **Reusable helpers** in `internal/util` (JSON, reference schemas, plan
  modifiers, nil-safe checks).
- **RFC-6902 JSON Patch builder** in `internal/patch` for PATCH-based updates.
- **26 worked resources + 4 data sources** in `internal/` that all follow the
  same pattern — copy the closest one when adding your own.
- Linting config (`.golangci.yml`) and registry manifest included.

## Project Layout

```
main.go                     # Provider entry point (set your registry address)
internal/
  provider/provider.go      # Provider definition, schema, resource registration
  client/client.go          # Generic OAuth2 REST client
  crud/                     # Generic CRUD base types (BaseResource, SubResource, ...)
  util/                     # Shared helpers (JSON, schema, plan modifiers)
  patch/                    # RFC-6902 JSON Patch builder
  launcher/                 # A simple example resource to copy from
  <26 resources + 4 data sources, one package each>
examples/                   # HCL usage examples
docs/                       # Documentation
mock/                       # Mockoon mock API for offline tests
pipelines/                  # CI/CD templates (see pipelines/README.md)
```


## Prerequisites

Install these first (all free):

- **[Go](https://go.dev/dl/) 1.23 or newer** — the language this provider is
  written in. Verify with `go version`.
- **[Terraform](https://developer.hashicorp.com/terraform/install) 1.7 or newer**
  — used to run the provider. Verify with `terraform version`.
- **git** — to clone this repository.

No paid accounts or cloud services are required to build and test.

## Quick Start

1. **Clone the repository:**
   ```bash
   git clone <your-fork-url> terraform-provider-eshiam
   cd terraform-provider-eshiam
   ```

2. **Rename the module and provider.** Replace `terraform-provider-eshiam`,
   `example`, and the registry address `registry.terraform.io/eshiam-corp/eshiam`
   with your own names throughout the project.

3. **Build** (compiles all packages; produces no output on success):
   ```bash
   go build ./...
   ```

4. **Run the tests:**
   ```bash
   go test ./...
   ```

5. **Try it locally** with a dev override so Terraform uses your freshly built
   binary. Step-by-step instructions are in
   [docs/local-development.md](docs/local-development.md).

## Configuration

The provider reads credentials from config or environment variables:

| Setting        | Env var                 | Description            |
|----------------|-------------------------|------------------------|
| `host`         | `EXAMPLE_HOST`          | API base URL           |
| `client_id`    | `EXAMPLE_CLIENT_ID`     | OAuth2 client ID       |
| `client_secret`| `EXAMPLE_CLIENT_SECRET` | OAuth2 client secret   |

## Adding a Resource

See [docs/adding-a-resource.md](docs/adding-a-resource.md) for a full
walkthrough. In short: copy the closest existing package (for a simple one,
`internal/launcher`), adjust the schema and model, point it at your API
endpoint, and register the constructor in `internal/provider/provider.go`.

## Publishing & CI/CD

Ready to ship? See
[docs/publishing-to-the-registry.md](docs/publishing-to-the-registry.md) for a
step-by-step guide to signing keys, the Terraform Registry, and automating
releases with Jenkins or any other CI platform. Pipeline templates for every
major CI system live in [pipelines/README.md](pipelines/README.md).

## Documentation

- **[docs/setup.md](docs/setup.md) — start here: full first-time setup &
  installation (Go + Terraform, dev overrides, credentials, Docker, registry).**
- [docs/local-development.md](docs/local-development.md) — build and run locally.
- [docs/docker.md](docs/docker.md) — build & test in Docker (no Go install) and
  link to a consumer framework.
- [docs/adding-a-resource.md](docs/adding-a-resource.md) — add a new resource.
- [docs/publishing-to-the-registry.md](docs/publishing-to-the-registry.md) —
  publish and automate releases.
- [pipelines/README.md](pipelines/README.md) — CI/CD templates and secrets.

## License

See [LICENSE](LICENSE).
