package util

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// ResourceReferenceSchema returns a nested-attribute schema for a single object
// reference of the shape { type, id, name }. The id and name are computed when
// not supplied so the API can populate them.
func ResourceReferenceSchema(allowedType string, required bool, description string) schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Description: description,
		Required:    required,
		Optional:    !required,
		Attributes:  referenceAttributes(allowedType, false),
	}
}

// ResourceReferenceListSchema returns a list of object references.
func ResourceReferenceListSchema(allowedType string, required bool, description string) schema.ListNestedAttribute {
	return schema.ListNestedAttribute{
		Description:  description,
		Required:     required,
		Optional:     !required,
		NestedObject: schema.NestedAttributeObject{Attributes: referenceAttributes(allowedType, true)},
	}
}

// ResourceReferenceSetSchema returns a set of object references.
func ResourceReferenceSetSchema(allowedType string, required bool, description string) schema.SetNestedAttribute {
	return schema.SetNestedAttribute{
		Description:  description,
		Required:     required,
		Optional:     !required,
		NestedObject: schema.NestedAttributeObject{Attributes: referenceAttributes(allowedType, true)},
	}
}

// referenceAttributes builds the { type, id, name } attribute map. When idRequired
// is true (typical for list/set members) the id must be supplied.
func referenceAttributes(allowedType string, idRequired bool) map[string]schema.Attribute {
	id := schema.StringAttribute{
		Computed: !idRequired,
		Optional: !idRequired,
		Required: idRequired,
	}
	if !idRequired {
		id.PlanModifiers = []planmodifier.String{stringplanmodifier.UseStateForUnknown()}
	}
	return map[string]schema.Attribute{
		"type": schema.StringAttribute{
			Computed: true,
			Optional: true,
			Default:  stringdefault.StaticString(allowedType),
			Validators: []validator.String{
				stringvalidator.OneOf(allowedType),
			},
		},
		"id": id,
		"name": schema.StringAttribute{
			Computed: true,
			Optional: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
	}
}
