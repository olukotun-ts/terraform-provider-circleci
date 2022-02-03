package circleci

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/olukotun-ts/circleci-client-go/circleci"
)

func resourceCircleCIUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceUserCreate,
		ReadContext:   resourceUserRead,
		DeleteContext: resourceUserDelete,
		Schema: map[string]*schema.Schema{
			"projects": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
				ForceNew: true,
			},
			"organization": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"branch": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vcs_provider": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceUserCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := meta.(*circleci.Client)

	org_name := d.Get("organization").(string)
	vcs := d.Get("vcs_provider").(string)
	vcs_slug := ""
	switch vcs {
	case "github":
		vcs_slug = "gh"
	default:
		// Todo: Implement as validator in schema
		diag_summary := fmt.Sprintf("Please enter a supported vcs_provider")
		diag_detail := fmt.Sprintf("%s is currently not supported. Defaulting to github.", vcs)
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  diag_summary,
			Detail:   diag_detail,
		})

		vcs_slug = "gh"
	}

	projects := d.Get("projects").([]interface{})
	projects_slugs := []string{}
	for _, project := range projects {
		project_slug := fmt.Sprintf("%s/%s/%s", vcs_slug, org_name, project)
		projects_slugs = append(projects_slugs, project_slug)
	}

	branch := d.Get("branch").(string)
	_, err := client.Projects.FollowMany(ctx, projects_slugs, branch)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceUserRead(ctx, d, meta)
}

func resourceUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*circleci.Client)

	user, err := client.Users.GetCurrentUser(ctx)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(user.Login)

	return nil
}

func resourceUserDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := meta.(*circleci.Client)

	org_name := d.Get("organization").(string)
	vcs := d.Get("vcs_provider").(string)
	vcs_slug := ""
	switch vcs {
	case "github":
		vcs_slug = "gh"
	default:
		// Todo: Implement as validator in schema
		diag_summary := fmt.Sprintf("Please enter a supported vcs_provider")
		diag_detail := fmt.Sprintf("%s is currently not supported. Defaulting to github.", vcs)
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  diag_summary,
			Detail:   diag_detail,
		})

		vcs_slug = "gh"
	}

	projects := d.Get("projects").([]interface{})
	projects_slugs := []string{}
	for _, project := range projects {
		project_slug := fmt.Sprintf("%s/%s/%s", vcs_slug, org_name, project)
		projects_slugs = append(projects_slugs, project_slug)
	}

	_, err := client.Projects.UnfollowMany(ctx, projects_slugs)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceUserRead(ctx, d, meta)
}
