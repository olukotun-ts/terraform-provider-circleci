package circleci

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/olukotun-ts/circleci-client-go/circleci"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"circleci_project": resourceCircleCIProject(),
			"circleci_user":    resourceCircleCIUser(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"circleci_project": dataSourceCircleCIProject(),
			"circleci_user":    dataSourceCircleCIUser(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	client := circleci.NewClient(nil)

	return client, diags
}
