// Package source_aggregation_schedule implements the
// example_source_aggregation_schedule sub-resource.
package source_aggregation_schedule

import (
	"context"

	"terraform-provider-eshiam/internal/crud"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var (
	_ resource.Resource              = &sourceAggregationScheduleResource{}
	_ resource.ResourceWithConfigure = &sourceAggregationScheduleResource{}
)

// NewSourceAggregationScheduleResource is the constructor registered with the provider.
func NewSourceAggregationScheduleResource() resource.Resource {
	return &sourceAggregationScheduleResource{crud.SubResource[sourceAggregationScheduleModel, *sourceAggregationScheduleModel]{
		TypeNameSuffix: "_source_aggregation_schedule",
		CollectionPath: func(m *sourceAggregationScheduleModel) string { return m.path() },
		ObjectPath:     func(m *sourceAggregationScheduleModel) string { return m.path() },
		Singleton:      true,
		SchemaFn:       sourceAggregationScheduleSchema,
	}}
}

type sourceAggregationScheduleResource struct {
	crud.SubResource[sourceAggregationScheduleModel, *sourceAggregationScheduleModel]
}

func sourceAggregationScheduleSchema(_ context.Context) schema.Schema {
	return schema.Schema{
		Description: "Manages the aggregation schedule (cron) of a source.",
		Attributes: map[string]schema.Attribute{
			"source_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"cron_expression": schema.StringAttribute{Required: true},
			"type": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
	}
}
