package custom_form

import (
	"terraform-provider-eshiam/internal/util"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// customFormModel mirrors the custom-form resource. Complex nested collections
// (input, conditions, elements) are represented as JSON blobs to keep the
// generic mirror flat and portable.
type customFormModel struct {
	Id             types.String         `tfsdk:"id"`
	Name           types.String         `tfsdk:"name"`
	Description    types.String         `tfsdk:"description"`
	Owner          *util.ReferenceModel `tfsdk:"owner"`
	UsedBy         jsontypes.Exact      `tfsdk:"used_by"`
	FormInput      jsontypes.Exact      `tfsdk:"form_input"`
	FormConditions jsontypes.Exact      `tfsdk:"form_conditions"`
	FormElements   jsontypes.Exact      `tfsdk:"form_elements"`
}

func (m *customFormModel) GetID() string { return m.Id.ValueString() }
