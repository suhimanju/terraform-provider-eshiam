package crud

import (
	"context"
	"net/http"

	"terraform-provider-eshiam/internal/client"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SingletonResource is a generic resource for tenant-wide singleton objects that
// live at a fixed path with no collection or id (for example org config or
// network config). Create and Update both write the full object to Path; Delete
// is a no-op because the object always exists.
type SingletonResource[M any, PM interface {
	*M
	Identifiable
}] struct {
	client *client.APIClient

	TypeNameSuffix string
	// Path is the fixed object path, e.g. "/v3/org-config".
	Path string
	// WriteMethod is the method used for Create/Update (PUT or PATCH). Defaults
	// to PUT.
	WriteMethod string
	SchemaFn    func(ctx context.Context) schema.Schema
}

func (b *SingletonResource[M, PM]) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	b.client = ClientFromProviderData(req.ProviderData, &resp.Diagnostics)
}

func (b *SingletonResource[M, PM]) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + b.TypeNameSuffix
}

func (b *SingletonResource[M, PM]) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = b.SchemaFn(ctx)
}

func (b *SingletonResource[M, PM]) method() string {
	if b.WriteMethod == "" {
		return http.MethodPut
	}
	return b.WriteMethod
}

func (b *SingletonResource[M, PM]) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan M
	resp.Diagnostics.Append(req.Plan.Get(ctx, PM(&plan))...)
	if resp.Diagnostics.HasError() {
		return
	}
	Update(ctx, b.client, b.method(), b.Path, PM(&plan), &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, PM(&plan))...)
}

func (b *SingletonResource[M, PM]) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state M
	resp.Diagnostics.Append(req.State.Get(ctx, PM(&state))...)
	if resp.Diagnostics.HasError() {
		return
	}
	if !Read(ctx, b.client, b.Path, PM(&state), &resp.Diagnostics) {
		resp.State.RemoveResource(ctx)
		return
	}
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, PM(&state))...)
}

func (b *SingletonResource[M, PM]) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan M
	resp.Diagnostics.Append(req.Plan.Get(ctx, PM(&plan))...)
	if resp.Diagnostics.HasError() {
		return
	}
	Update(ctx, b.client, b.method(), b.Path, PM(&plan), &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, PM(&plan))...)
}

func (b *SingletonResource[M, PM]) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton objects cannot be deleted; removing from state is sufficient.
}
