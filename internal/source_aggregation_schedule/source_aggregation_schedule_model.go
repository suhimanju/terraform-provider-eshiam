package source_aggregation_schedule

import "github.com/hashicorp/terraform-plugin-framework/types"

// sourceAggregationScheduleModel mirrors the account-aggregation schedule of a
// source. It is a singleton sub-resource keyed by the schedule type.
type sourceAggregationScheduleModel struct {
	SourceId        types.String `tfsdk:"source_id"`
	CronExpression  types.String `tfsdk:"cron_expression"`
	AggregationType types.String `tfsdk:"type"`
}

func (m *sourceAggregationScheduleModel) GetID() string { return m.SourceId.ValueString() }

func (m *sourceAggregationScheduleModel) path() string {
	return "/beta/sources/" + m.SourceId.ValueString() + "/schedules/" + m.AggregationType.ValueString()
}
