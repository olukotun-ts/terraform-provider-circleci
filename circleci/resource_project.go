package circleci

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/olukotun-ts/circleci-client-go/circleci"
)

func resourceProject() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceProjectCreate,
		ReadContext: resourceProjectRead,
		DeleteContext: resourceProjectDelete,
		Schema: map[string]*schema.Schema{
			"slug": {
				Type: schema.TypeString,
				Computed: true,
			},
			"organization": {
				Type: schema.TypeString,
				Computed: true,
			},
			"name": {
				Type: schema.TypeString,
				Computed: true,
			},
			"branch": {
				Type: schema.TypeString,
				Computed: true,
			},
		}
	}
}

func resourceProjectCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) error {
	client := meta.(*circleci.Client)

	_, err := client.Projects.Follow(ctx, d.Get("slug"), d.Get("branch"))
	if err != nil {
		return err
	}

	err := resourceProjectRead(ctx, d, meta)
	if err != nil {
		return err
	}

	return nil
}

func resourceProjectRead(ctx context.Context, d *schema.ResourceData, meta interface{}) error {
	client := meta.(*circleci.Client)

	project, err := client.Projects.Get(ctx, d.Get("slug"))
	if err != nil {
		return err
	}

	d.SetId(project.Name)
	d.Set("organization", project.Organization)
	d.Set("name", project.Name)
	d.Set("slug", project.Slug)

	return nil
}

func resourceProjectDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) error {
	client := meta.(*circleci.Client)

	_, err := client.Projects.Unfollow(ctx, d.Get("slug"))
	if err != nil {
		return err
	}

	return nil
}
