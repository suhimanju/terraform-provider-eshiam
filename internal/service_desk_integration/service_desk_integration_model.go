package service_desk_integration

import (
	"terraform-provider-eshiam/internal/util"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// serviceDeskIntegrationModel mirrors a service desk integration (SDIM) object.
type serviceDeskIntegrationModel struct {
	Id                     types.String         `tfsdk:"id"`
	Name                   types.String         `tfsdk:"name"`
	Description            types.String         `tfsdk:"description"`
	Type                   types.String         `tfsdk:"type"`
	OwnerRef               *util.ReferenceModel `tfsdk:"owner_ref"`
	ClusterRef             *util.ReferenceModel `tfsdk:"cluster_ref"`
	ProvisioningConfig     *provisioningConfig  `tfsdk:"provisioning_config"`
	Attributes             jsontypes.Normalized `tfsdk:"attributes"`
	AttributesCredentials  jsontypes.Exact      `tfsdk:"attributes_credentials"`
	BeforeProvisioningRule *util.ReferenceModel `tfsdk:"before_provisioning_rule"`
}

type provisioningConfig struct {
	ManagedResourceRefs           []util.ReferenceModel  `tfsdk:"managed_resource_refs"`
	PlanInitializerScript         *planInitializerScript `tfsdk:"plan_initializer_script"`
	NoProvisioningRequests        types.Bool             `tfsdk:"no_provisioning_requests"`
	ProvisioningRequestExpiration types.Int64            `tfsdk:"provisioning_request_expiration"`
}

type planInitializerScript struct {
	Source types.String `tfsdk:"source"`
}

func (m *serviceDeskIntegrationModel) GetID() string { return m.Id.ValueString() }
