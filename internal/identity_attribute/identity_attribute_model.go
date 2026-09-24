package identity_attribute

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// identityAttributeModel mirrors the identity-attribute resource. Identity
// attributes are keyed by name rather than a generated id.
type identityAttributeModel struct {
	Name        types.String    `tfsdk:"name"`
	DisplayName types.String    `tfsdk:"display_name"`
	Type        types.String    `tfsdk:"type"`
	Standard    types.Bool      `tfsdk:"standard"`
	System      types.Bool      `tfsdk:"system"`
	Multi       types.Bool      `tfsdk:"multi"`
	Searchable  types.Bool      `tfsdk:"searchable"`
	Sources     jsontypes.Exact `tfsdk:"sources"`
}

// GetID returns the name, which is the identity attribute's identifier.
func (m *identityAttributeModel) GetID() string { return m.Name.ValueString() }
