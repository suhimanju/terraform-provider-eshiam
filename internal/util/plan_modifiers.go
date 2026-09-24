package util

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// RetainAPIValueIfStateIsEmpty is a string plan modifier that keeps the value
// currently in state when the configuration is empty. This is useful for
// attributes the API populates and that Terraform should not reset to null when
// the user leaves them blank.
type RetainAPIValueIfStateIsEmpty struct{}

func (m RetainAPIValueIfStateIsEmpty) Description(_ context.Context) string {
	return "Retains the API-returned value when the configuration is empty."
}

func (m RetainAPIValueIfStateIsEmpty) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m RetainAPIValueIfStateIsEmpty) PlanModifyString(
	ctx context.Context,
	req planmodifier.StringRequest,
	resp *planmodifier.StringResponse,
) {
	// If config is empty but state holds a value, retain the state value.
	if req.ConfigValue.ValueString() == "" && !req.StateValue.IsNull() && req.StateValue.ValueString() != "" {
		tflog.Info(ctx, "[RetainAPIValueIfStateIsEmpty] retaining state value over empty config")
		resp.PlanValue = req.StateValue
		return
	}

	// If config is empty and there is no state, treat as null.
	if req.ConfigValue.ValueString() == "" {
		resp.PlanValue = types.StringNull()
		return
	}

	// Otherwise keep the planned value.
	resp.PlanValue = req.PlanValue
}
