package workflow

import (
	"terraform-provider-eshiam/internal/util"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// workflowModel mirrors a workflow (flat collection resource).
type workflowModel struct {
	Id          types.String         `tfsdk:"id"`
	Name        types.String         `tfsdk:"name"`
	Owner       *util.ReferenceModel `tfsdk:"owner"`
	Description types.String         `tfsdk:"description"`
	Enabled     types.Bool           `tfsdk:"enabled"`
	Definition  *definition          `tfsdk:"definition"`
	Trigger     *trigger             `tfsdk:"trigger"`
}

type definition struct {
	Start types.String         `tfsdk:"start"`
	Steps jsontypes.Normalized `tfsdk:"steps"`
}

type trigger struct {
	Type types.String `tfsdk:"type"`
	// Attributes holds the trigger-type-specific configuration as free-form
	// JSON, keeping the model generic across trigger types.
	Attributes jsontypes.Normalized `tfsdk:"attributes"`
}

func (m *workflowModel) GetID() string { return m.Id.ValueString() }
