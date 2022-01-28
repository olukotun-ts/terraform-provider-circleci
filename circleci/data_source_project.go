package circleci

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/olukotun-ts/circleci-client-go/circleci"
)

func dataSourceProject() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceProjectRead,
		Schema: map[string]*schema.Schema{
			"slug": {
				Type:     schema.TypeString,
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
		},
	}
}

func dataSourceProjectRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*circleci.Client)

	var diags diag.Diagnostics

	slug := d.Get("slug").(string)
	project, err := client.Projects.Get(ctx, slug)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(project.Name)
	d.Set("organization", project.Organization)
	d.Set("name", project.Name)
	d.Set("slug", project.Slug)

	return diags
}