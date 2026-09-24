// Package crud provides generic create/read/update/delete helpers that operate
// on terraform-plugin-framework model structs via reflection. Resources use
// these to avoid repeating identical HTTP + mapping boilerplate.
//
// Models must use `tfsdk` struct tags and the field types supported by
// util.ModelToMap / util.MapToModel.
package crud

import (
	"context"
	"net/http"

	"terraform-provider-eshiam/internal/client"
	"terraform-provider-eshiam/internal/util"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// Create POSTs the model to collectionPath and maps the response back into the
// model (including the server-assigned id).
func Create(ctx context.Context, c *client.APIClient, collectionPath string, model any, diagnostics *diag.Diagnostics) {
	body := util.ModelToMap(model, diagnostics)
	if diagnostics.HasError() {
		return
	}
	out, resp, err := c.CreateObject(ctx, collectionPath, body)
	if err != nil {
		diagnostics.AddError("Error Creating Resource", err.Error()+"\n"+util.GetBody(resp))
		return
	}
	util.MapToModel(out, model, diagnostics)
}

// Read GETs objectPath and maps the response into the model. It returns false
// (and removes nothing) when the object is gone (HTTP 404) so callers can drop
// it from state.
func Read(ctx context.Context, c *client.APIClient, objectPath string, model any, diagnostics *diag.Diagnostics) (found bool) {
	out, resp, err := c.GetObject(ctx, objectPath)
	if resp != nil && resp.StatusCode == http.StatusNotFound {
		return false
	}
	if err != nil {
		diagnostics.AddError("Error Reading Resource", err.Error()+"\n"+util.GetBody(resp))
		return false
	}
	util.MapToModel(out, model, diagnostics)
	return true
}

// Update sends the model to objectPath using method (PUT or PATCH) and maps the
// response back into the model.
func Update(ctx context.Context, c *client.APIClient, method, objectPath string, model any, diagnostics *diag.Diagnostics) {
	body := util.ModelToMap(model, diagnostics)
	if diagnostics.HasError() {
		return
	}
	out, resp, err := c.UpdateObject(ctx, method, objectPath, body)
	if err != nil {
		diagnostics.AddError("Error Updating Resource", err.Error()+"\n"+util.GetBody(resp))
		return
	}
	util.MapToModel(out, model, diagnostics)
}

// Delete DELETEs objectPath, tolerating a 404.
func Delete(ctx context.Context, c *client.APIClient, objectPath string, diagnostics *diag.Diagnostics) {
	resp, err := c.DeleteObject(ctx, objectPath)
	if resp != nil && resp.StatusCode == http.StatusNotFound {
		return
	}
	if err != nil {
		diagnostics.AddError("Error Deleting Resource", err.Error()+"\n"+util.GetBody(resp))
	}
}
