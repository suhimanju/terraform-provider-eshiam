# Adding a Resource

This provider uses a **generic CRUD base** so you rarely write
Create/Read/Update/Delete by hand. To add a resource you mainly describe its
*schema* and *model*, then point it at an API endpoint. Follow the steps below.

If you are new to the codebase, open `internal/launcher/` first — it is the
smallest complete example and the one these steps are based on.

## Overview

Each resource lives in its own package under `internal/<name>/` and consists of
two files:

| File | Purpose |
|------|---------|
| `<name>_model.go` | A Go struct describing the resource's attributes. |
| `<name>_resource.go` | The schema and the constructor wired to a CRUD base. |

## 1. Create the package

Create a folder `internal/widget/` (replace `widget` with your resource name).

## 2. Write the model — `widget_model.go`

The model is a plain struct whose fields map to Terraform attributes via
`tfsdk` tags. It must have a `GetID()` method so the generic base knows how to
identify the object.

```go
package widget

import "github.com/hashicorp/terraform-plugin-framework/types"

// widgetModel mirrors the widget resource.
type widgetModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// GetID lets the generic CRUD base identify this object.
func (m *widgetModel) GetID() string { return m.Id.ValueString() }
```

Field type cheat-sheet:

| Terraform data | Go type |
|----------------|---------|
| string | `types.String` |
| bool | `types.Bool` |
| number | `types.Int64` / `types.Float64` |
| free-form JSON | `jsontypes.Exact` or `jsontypes.Normalized` |
| `{ type, id, name }` reference | `*util.ReferenceModel` |

## 3. Write the resource — `widget_resource.go`

Define the schema and a constructor that embeds `crud.BaseResource`. There are
no manual CRUD methods to write — the base handles them.

```go
// Package widget implements the example_widget resource.
package widget

import (
	"context"

	"terraform-provider-eshiam/internal/crud"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var (
	_ resource.Resource                = &widgetResource{}
	_ resource.ResourceWithConfigure   = &widgetResource{}
	_ resource.ResourceWithImportState = &widgetResource{}
)

// NewWidgetResource is the constructor registered with the provider.
func NewWidgetResource() resource.Resource {
	return &widgetResource{crud.BaseResource[widgetModel, *widgetModel]{
		TypeNameSuffix: "_widget",       // becomes example_widget
		Endpoint:       "/v1/widgets",   // your REST collection path
		SchemaFn:       widgetSchema,
	}}
}

type widgetResource struct {
	crud.BaseResource[widgetModel, *widgetModel]
}

func widgetSchema(_ context.Context) schema.Schema {
	return schema.Schema{
		Description: "Manages a widget.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name": schema.StringAttribute{Required: true},
		},
	}
}
```

### Which base should I use?

| Base | Use when… | Example package |
|------|-----------|-----------------|
| `crud.BaseResource[M, PM]` | The object is a flat, top-level collection (`/v1/widgets`). | `internal/launcher` |
| `crud.SubResource` | The object lives under a parent (`/v1/sources/{id}/...`). | `internal/provisioning_policy` |
| `crud.SingletonResource` | There is exactly one per tenant (no list). | `internal/org_config` |

## 4. Register the constructor

Add your package to the `Resources` slice in
`internal/provider/provider.go`:

```go
func (p *exampleProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		// ...existing resources...
		widget.NewWidgetResource, // <-- add this
	}
}
```

Remember to add the matching `import` for the `widget` package at the top of the
file.

## 5. Add an example and docs

- `examples/resources/example_widget/resource.tf` — a small HCL snippet.
- `docs/resources/widget.md` — a short reference page.

## 6. Verify

Run these three commands; all should pass with no output:

```bash
go build ./...   # compiles
go vet ./...     # static checks
gofmt -l .       # lists badly-formatted files (empty = good)
```

## Need a PATCH-based update?

If your API updates via JSON Patch (RFC-6902), build operations with
`internal/patch` instead of sending the whole object. See
`internal/lifecycle_state/` for a working example.
