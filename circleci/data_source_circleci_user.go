package circleci

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/olukotun-ts/circleci-client-go/circleci"
)

// Todo: Check schema
func dataSourceCircleCIUser() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceUserRead,
		Schema: map[string]*schema.Schema{
			"projects": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed: true,
			},
			"organization": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"branch": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vcs_provider": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*circleci.Client)

	user, err := client.Users.GetCurrentUser(ctx)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(user.Login)

	return nil
}
