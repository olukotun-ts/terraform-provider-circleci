package circleci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/olukotun-ts/circleci-client-go/circleci"
)

func resourceCircleCIProject() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceProjectCreate,
		ReadContext:   resourceProjectRead,
		DeleteContext: resourceProjectDelete,
		Schema: map[string]*schema.Schema{
			"slug": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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
				Optional: true,
				Default:  "master",
				ForceNew: true,
			},
		},
	}
}

// Todo: Explore error vs Diagnostics
func resourceProjectCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*circleci.Client)

	slug := d.Get("slug").(string)
	branch := d.Get("branch").(string)

	_, err := client.Projects.Follow(ctx, slug, branch)
	if err != nil {
		return diag.FromErr(err)
	}

	// Todo: To avoid errors from race condition, use d.Set() instead. How likely of an issue is this?
	// Ideally, client.Projects.Follow should include metadata for followed project; would eliminate need for another call.
	return resourceProjectRead(ctx, d, meta)
}

func resourceProjectRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*circleci.Client)

	slug := d.Get("slug").(string)
	project, err := client.Projects.Get(ctx, slug)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(project.Name)
	d.Set("organization", project.Organization)
	d.Set("name", project.Name)
	d.Set("slug", project.Slug)

	return nil
}

func resourceProjectDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*circleci.Client)

	slug := d.Get("slug").(string)
	_, err := client.Projects.Unfollow(ctx, slug)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
