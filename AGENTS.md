# Agent Instructions — terraform-provider-eshiam (Go framework)

A **generic boilerplate** for a Terraform provider built on
**terraform-plugin-framework** (not the legacy SDK).

## Quick Reference

- **Build**: `go build ./...`
- **Unit tests**: `go test ./...`
- **Lint**: `golangci-lint run ./...`
- **Format**: `gofmt -w .`

## Architecture

```
internal/
├── provider/   # Provider definition; registers all resources/data sources
├── client/     # Generic OAuth2 client-credentials REST client
├── crud/       # Generic CRUD bases (BaseResource, SubResource, SingletonResource)
├── util/       # Shared schema builders, JSON helpers, plan modifiers
├── patch/      # RFC-6902 JSON Patch builder for PATCH updates
├── launcher/   # Simplest example resource — copy this when adding new ones
│   ├── launcher_resource.go   # Schema + constructor wired to a CRUD base
│   └── launcher_model.go      # Model struct + GetID()
└── <25 more resource packages and 4 data-source packages>
```

Most resources do **not** hand-write CRUD. They embed a generic base from
`internal/crud` and only supply a schema, an endpoint, and a model.

## Conventions

- Never import `hashicorp/terraform-plugin-sdk`; use `terraform-plugin-framework`.
- One package per resource/data source under `internal/`.
- Declare interface compliance at the top of each resource:
  ```go
  var (
      _ resource.Resource                = &widgetResource{}
      _ resource.ResourceWithConfigure   = &widgetResource{}
      _ resource.ResourceWithImportState = &widgetResource{}
  )
  ```
- Constructor: exported `NewXxxResource()` returning `resource.Resource`, which
  embeds a `crud.BaseResource` (or `SubResource` / `SingletonResource`) with the
  `TypeNameSuffix`, `Endpoint`, and `SchemaFn` set.
- Models: unexported (camelCase) structs; tags `tfsdk:"snake_case"`; each model
  implements `GetID() string`.
- Reuse helpers from `internal/util` and `internal/patch`.
- Use tabs for Go; 2-space indent for HCL/JSON/YAML/Markdown; LF endings.

## Adding a Resource

1. Create `internal/<name>/` with `<name>_model.go` and `<name>_resource.go`.
2. Define the model struct with `tfsdk` tags and a `GetID()` method.
3. Define the schema and a constructor embedding the right `crud` base.
4. Register `NewXxxResource` in `internal/provider/provider.go`.
5. Add an example under `examples/resources/example_<name>/` and docs under
   `docs/resources/<name>.md`.

See [docs/adding-a-resource.md](docs/adding-a-resource.md) for a full
walkthrough, and `internal/launcher/` for the reference implementation.
