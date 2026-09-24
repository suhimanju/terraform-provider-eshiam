package crud

import (
	"fmt"

	"terraform-provider-eshiam/internal/client"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// ClientFromProviderData extracts the *client.APIClient placed into provider
// data by the provider's Configure method. It returns nil when provider data is
// absent (during early configuration) or reports an error on a type mismatch.
func ClientFromProviderData(providerData any, diagnostics *diag.Diagnostics) *client.APIClient {
	if providerData == nil {
		return nil
	}
	c, ok := providerData.(*client.APIClient)
	if !ok {
		diagnostics.AddError(
			"Unexpected Configure Type",
			fmt.Sprintf("Expected *client.APIClient, got: %T. Please report this issue to the provider developers.", providerData),
		)
		return nil
	}
	return c
}
