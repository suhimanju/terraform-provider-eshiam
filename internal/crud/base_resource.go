package crud

import (
	"context"
	"net/http"

	"terraform-provider-eshiam/internal/client"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Identifiable is implemented by resource models so the generic base resource
// can build per-object URLs. Return the server identifier (usually the id, or
// the name for name-keyed objects).
type Identifiable interface {
	GetID() string
}

// BaseResource is a generic terraform-plugin-framework resource implementation
// that handles Configure, Metadata, CRUD, and ImportState for models that map
// cleanly to a REST collection endpoint.
//
// Type parameters:
//   - M  is the model struct type.
//   - PM is *M and must implement Identifiable.
//
// Concrete resources embed a *BaseResource and only provide the schema, the
// type-name suffix, and the endpoint.
type BaseResource[M any, PM interface {
	*M
	Identifiable
}] struct {
	client *client.APIClient

	// TypeNameSuffix is appended to the provider type name, e.g. "_transform".
	TypeNameSuffix string
	// Endpoint is the REST collection path, e.g. "/v1/transforms".
	Endpoint string
	// UpdateMethod is the HTTP method used for updates (PUT or PATCH). Defaults
	// to PUT when empty.
	UpdateMethod string
	// SchemaFn returns the resource schema.
	SchemaFn func(ctx context.Context) schema.Schema
}

func (b *BaseResource[M, PM]) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	b.client = ClientFromProviderData(req.ProviderData, &resp.Diagnostics)
}

func (b *BaseResource[M, PM]) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + b.TypeNameSuffix
}

func (b *BaseResource[M, PM]) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = b.SchemaFn(ctx)
}

func (b *BaseResource[M, PM]) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan M
	resp.Diagnostics.Append(req.Plan.Get(ctx, PM(&plan))...)
	if resp.Diagnostics.HasError() {
		return
	}
	Create(ctx, b.client, b.Endpoint, PM(&plan), &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, PM(&plan))...)
}

func (b *BaseResource[M, PM]) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state M
	resp.Diagnostics.Append(req.State.Get(ctx, PM(&state))...)
	if resp.Diagnostics.HasError() {
		return
	}
	if !Read(ctx, b.client, b.objectPath(PM(&state)), PM(&state), &resp.Diagnostics) {
		resp.State.RemoveResource(ctx)
		return
	}
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, PM(&state))...)
}

func (b *BaseResource[M, PM]) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan M
	resp.Diagnostics.Append(req.Plan.Get(ctx, PM(&plan))...)
	if resp.Diagnostics.HasError() {
		return
	}
	method := b.UpdateMethod
	if method == "" {
		method = http.MethodPut
	}
	Update(ctx, b.client, method, b.objectPath(PM(&plan)), PM(&plan), &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, PM(&plan))...)
}

func (b *BaseResource[M, PM]) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state M
	resp.Diagnostics.Append(req.State.Get(ctx, PM(&state))...)
	if resp.Diagnostics.HasError() {
		return
	}
	Delete(ctx, b.client, b.objectPath(PM(&state)), &resp.Diagnostics)
}

func (b *BaseResource[M, PM]) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// objectPath builds the per-object URL: "<endpoint>/<id>".
func (b *BaseResource[M, PM]) objectPath(m PM) string {
	return b.Endpoint + "/" + m.GetID()
}
