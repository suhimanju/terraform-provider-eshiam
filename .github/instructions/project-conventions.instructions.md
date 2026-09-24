---
description: "Conventions for the generic terraform-plugin-framework Go boilerplate. Use when writing Go resources, data sources, the client, utilities, or the patch builder."
applyTo: "**"
---
# Project Conventions — terraform-provider-eshiam (Go framework)

## Framework & Architecture

- Uses **terraform-plugin-framework** (NOT the legacy SDK). Never import
  `hashicorp/terraform-plugin-sdk`.
- Each resource/data source lives in its own package under `internal/`.
- Shared helpers live in `internal/util`; PATCH helpers in `internal/patch`.
  Reuse them instead of duplicating.
- The API client is `internal/client.APIClient` (generic OAuth2 REST client).

## Resource Pattern

1. **Files** per package: `<name>_resource.go`, `<name>_model.go`,
   `<name>_resource_test.go` (build tag `!integration`), and optionally
   `<name>_datasource.go`.
2. **Interface compliance** at the top:
   ```go
   var _ resource.Resource = &myResource{}
   var _ resource.ResourceWithConfigure = &myResource{}
   ```
3. **Struct naming**: unexported camelCase (`widgetResource`).
4. **Constructor**: exported `NewWidgetResource()` returning `resource.Resource`.
5. **Metadata**: `resp.TypeName = req.ProviderTypeName + "_widget"`.
6. **Configure**: extract `*client.APIClient` from `req.ProviderData`.
7. **PATCH updates**: build ops with `internal/patch.AbstractBuilder`.

## Model Conventions

- Package-private (lowercase) model structs.
- Use `types.String/Bool/Int64/Set/Object` from terraform-plugin-framework.
- JSON blobs use `jsontypes.Exact` or `jsontypes.Normalized`.
- References use `*util.ReferenceModel` / `[]util.ReferenceModel`.
- Struct tags: `tfsdk:"snake_case_attribute_name"`.

## Testing

- Unit tests use build tag `//go:build !integration` and avoid the network.
- Prefer testing pure model-mapping and patch-building helpers.

## Style

- Tabs for Go; 2-space indent for HCL/JSON/YAML/Markdown; LF line endings;
  final newline always.
- Keep everything vendor-neutral — no customer names or private endpoints.
