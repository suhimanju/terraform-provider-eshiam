package crud

import (
	"context"
	"net/http"

	"terraform-provider-eshiam/internal/client"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SubResource is a generic resource for objects nested under a parent (for
// example source schemas, provisioning policies, or lifecycle states). Paths are
// built from the model via CollectionPath (used on create) and ObjectPath (used
// on read/update/delete). This supports parent ids and sub-keys held as model
// fields.
type SubResource[M any, PM interface {
	*M
	Identifiable
}] struct {
	client *client.APIClient

	TypeNameSuffix string
	// CollectionPath returns the create endpoint, e.g.
	// "/v3/sources/<sourceId>/schemas".
	CollectionPath func(PM) string
	// ObjectPath returns the per-object endpoint. For singleton sub-resources it
	// may be equal to CollectionPath.
	ObjectPath func(PM) string
	// UpdateMethod is the update HTTP method (PUT or PATCH). Defaults to PUT.
	UpdateMethod string
	// Singleton indicates the sub-resource has no delete endpoint.
	Singleton bool
	SchemaFn  func(ctx context.Context) schema.Schema
}

func (b *SubResource[M, PM]) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	b.client = ClientFromProviderData(req.ProviderData, &resp.Diagnostics)
}

func (b *SubResource[M, PM]) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + b.TypeNameSuffix
}

func (b *SubResource[M, PM]) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = b.SchemaFn(ctx)
}

func (b *SubResource[M, PM]) method() string {
	if b.UpdateMethod == "" {
		return http.MethodPut
	}
	return b.UpdateMethod
}

func (b *SubResource[M, PM]) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan M
	resp.Diagnostics.Append(req.Plan.Get(ctx, PM(&plan))...)
	if resp.Diagnostics.HasError() {
		return
	}
	Create(ctx, b.client, b.CollectionPath(PM(&plan)), PM(&plan), &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, PM(&plan))...)
}

func (b *SubResource[M, PM]) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state M
	resp.Diagnostics.Append(req.State.Get(ctx, PM(&state))...)
	if resp.Diagnostics.HasError() {
		return
	}
	if !Read(ctx, b.client, b.ObjectPath(PM(&state)), PM(&state), &resp.Diagnostics) {
		resp.State.RemoveResource(ctx)
		return
	}
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, PM(&state))...)
}

func (b *SubResource[M, PM]) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan M
	resp.Diagnostics.Append(req.Plan.Get(ctx, PM(&plan))...)
	if resp.Diagnostics.HasError() {
		return
	}
	Update(ctx, b.client, b.method(), b.ObjectPath(PM(&plan)), PM(&plan), &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, PM(&plan))...)
}

func (b *SubResource[M, PM]) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	if b.Singleton {
		return
	}
	var state M
	resp.Diagnostics.Append(req.State.Get(ctx, PM(&state))...)
	if resp.Diagnostics.HasError() {
		return
	}
	Delete(ctx, b.client, b.ObjectPath(PM(&state)), &resp.Diagnostics)
}
