package circleci

import (
	"context"

	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

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

// Todo: Fix auth -- get token in main.tf?
func resourceProjectCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// client := meta.(*circleci.Client)

	var diags diag.Diagnostics

	slug := d.Get("slug").(string)
	branch := d.Get("branch").(string)
	// name := d.Get("name").(string)

	// slug := "gh/olukotun-ts/name-button"
	// branch := "master"

	// ###############################

	url := fmt.Sprintf("https://circleci.com/api/v1.1/project/%s/follow", slug)

	reqBody, _ := json.Marshal(map[string]string{
		"branch": branch,
	})
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))

	req.Header.Set("circle-token", os.Getenv("CIRCLE_TOKEN"))
	req.Header.Set("content-type", "appliation/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		diag.FromErr(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return diag.Errorf("Expected 200 status code; got %v instead", res.StatusCode)
	}

	// get newly followed project
	// set ID
	// do stuff...
	d.Set("slug", slug)

	// Todo: To avoid errors from race condition, use d.Set() instead
	resourceProjectRead(ctx, d, meta)
	return diags

	// ###############################

	// _, err := client.Projects.Follow(ctx, slug, branch)
	// if err != nil {
	// 	return diag.FromErr(err)
	// }

	// resourceProjectRead(ctx, d, meta)

	// return diags
}

func resourceProjectRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func resourceProjectDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*circleci.Client)

	var diags diag.Diagnostics

	slug := d.Get("slug").(string)
	_, err := client.Projects.Unfollow(ctx, slug)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
