package provider

import (
	"context"
	"fmt"
	"github.com/devoteamgcloud/terraform-provider-looker/pkg/lookergo"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceApiCredentials() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceApiCredentialsCreate,
		ReadContext:   resourceApiCredentialsRead,
		DeleteContext: resourceApiCredentialsDelete, // Assuming credentials can be deleted; if not, handle accordingly
		Schema: map[string]*schema.Schema{
			"user_id": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true, // User ID usually immutable for the credential
				Description: "ID of the user owning the API credential",
			},
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Type of API credential (e.g. api3)",
			},
			"client_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Client ID of the API credential",
			},
			"client_secret": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "Client secret of the API credential. Available only on creation.",
			},
			"is_disabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether the credential is disabled",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL of the API credential resource",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext, // Adjust if needed
		},
	}
}

func resourceApiCredentialsCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := m.(*Config).Api

	userID := d.Get("user_id").(int)

	req := &lookergo.ApiCredential{
		Type:       d.Get("type").(string),
		IsDisabled: d.Get("is_disabled").(bool),
	}

	cred, _, err := c.ApiCredentials.Create(ctx, userID, req)
	if err != nil {
		return diag.FromErr(fmt.Errorf("creating API credential: %w", err))
	}

	// Use the credential ID as the Terraform resource ID
	d.SetId(cred.ID)

	// Set returned attributes
	d.Set("client_id", cred.ClientID)
	d.Set("client_secret", cred.ClientSecret)
	d.Set("url", cred.URL)
	d.Set("is_disabled", cred.IsDisabled)
	d.Set("type", cred.Type)
	d.Set("user_id", userID)

	return resourceApiCredentialsRead(ctx, d, m)
}

func resourceApiCredentialsRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	c := m.(*Config).Api

	userID := d.Get("user_id").(int)
	credID := d.Id() // Terraform resource ID is credential ID

	cred, _, err := c.ApiCredentials.Get(ctx, userID, credID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("reading API credential: %w", err))
	}
	if cred == nil {
		d.SetId("")
		return nil
	}

	d.Set("client_id", cred.ClientID)
	d.Set("url", cred.URL)
	d.Set("is_disabled", cred.IsDisabled)
	d.Set("type", cred.Type)
	d.Set("user_id", userID)

	return nil
}

func resourceApiCredentialsDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	// If deleting is supported by your API, implement it here.
	// If not, you can just remove the resource from state:
	d.SetId("")
	return nil
}