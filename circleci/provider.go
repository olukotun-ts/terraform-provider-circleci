package circleci

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/olukotun-ts/circleci-client-go/circleci"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap:   map[string]*schema.Resource{
			"project": resourceProject(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"project": dataSourceProject()
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, error) {
	client, err := circleci.NewClient(nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
