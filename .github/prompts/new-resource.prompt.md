---
description: "Scaffold a new Terraform resource for this provider — creates the resource file, model file, and registers it in the provider."
mode: "agent"
---
# New Resource: ${input:resourceName}

Create a new Terraform resource `example_${input:resourceName}` following the
project conventions in `.github/instructions/project-conventions.instructions.md`.

## Steps

1. **Create package** `internal/${input:resourceName}/`

2. **Create `${input:resourceName}_model.go`** with:
   - Unexported model struct using `types.String`, `types.Bool`, etc.
   - `tfsdk:"snake_case"` struct tags
   - Reference fields as `*util.ReferenceModel`
   - A `GetID() string` method so the struct satisfies `crud.Identifiable`

3. **Create `${input:resourceName}_resource.go`** with:
   - Package declaration matching the directory name
   - Exported constructor `NewXxxResource() resource.Resource`
   - Build the resource on top of one of the generic CRUD bases:
     - `crud.BaseResource[M, PM]` for flat, top-level collections
     - `crud.SubResource` for parent-scoped collections
     - `crud.SingletonResource` for tenant-wide singletons
   - Provide the `Schema`, the API path, and the model mapping
   - Reference an existing resource such as
     `internal/launcher/launcher_resource.go`

4. **Register** the resource in `internal/provider/provider.go`:
   - Add import for the new package
   - Add `${input:resourceName}.NewXxxResource` to the `Resources()` slice

5. **Create example** in
   `examples/resources/example_${input:resourceName}/resource.tf`

6. **Create docs** in `docs/resources/${input:resourceName}.md`

7. **Verify** with `go build -v ./...`, `go vet ./...`, and `gofmt -l .`

## Reference Files

- Resource pattern: `internal/launcher/launcher_resource.go`
- Model pattern: `internal/launcher/launcher_model.go`
- CRUD bases: `internal/crud/`
- Utilities: `internal/util/`
- Patch builder: `internal/patch/`
