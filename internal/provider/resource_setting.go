package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSetting() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSettingUpdate,
		ReadContext:   resourceSettingRead,
		UpdateContext: resourceSettingUpdate,
		DeleteContext: resourceSettingDelete, // Setting is not deletable, but DeleteContext is required
		Schema: map[string]*schema.Schema{
			"instance_config": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Externally available instance configuration information",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"feature_flags": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Feature flags for the instance",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"license_features": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "License features enabled on the instance",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"extension_framework_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Toggle extension framework on or off",
			},
			"extension_load_url_enabled": {
				Type:       schema.TypeBool,
				Optional:   true,
				Computed:   true,
				Deprecated: "(DEPRECATED) Toggle extension load url on or off. Do not use. This is temporary setting that will eventually become a noop and subsequently deleted",
			},
			"marketplace_auto_install_enabled": {
				Type:       schema.TypeBool,
				Optional:   true,
				Computed:   true,
				Deprecated: "(DEPRECATED) Toggle marketplace auto install on or off. Deprecated - do not use. Auto install can now be enabled via marketplace automation settings",
			},
			"marketplace_automation": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Description: "Marketplace automation settings",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"install_enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "Whether marketplace auto installation is enabled",
						},
						"update_looker_enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "Whether marketplace auto update is enabled for looker extensions",
						},
						"update_third_party_enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "Whether marketplace auto update is enabled for third party extensions",
						},
					},
				},
			},
			"marketplace_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Toggle marketplace on or off",
			},
			"marketplace_site": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Location of Looker marketplace CDN",
			},
			"marketplace_terms_accepted": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Accept marketplace terms by setting this value to true, or get the current status. Marketplace terms CANNOT be declined once accepted. Accepting marketplace terms automatically enables the marketplace. The marketplace can still be disabled after it has been enabled",
			},
			"privatelabel_configuration": {
				Type:        schema.TypeSet,
				Computed:    true,
				Optional:    true,
				Description: "Private label configuration",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"logo_file": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Customer logo image. Expected base64 encoded data (write-only)",
						},
						"logo_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Logo image url (read-only)",
						},
						"favicon_file": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Customer favicon image. Expected base64 encoded data (write-only)",
						},
						"favicon_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Favicon image url (read-only)",
						},
						"default_title": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Default page title",
						},
						"show_help_menu": {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "Boolean to toggle showing help menus",
						},
						"show_docs": {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "Boolean to toggle showing docs",
						},
						"show_email_sub_options": {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "Boolean to toggle showing email subscription options",
						},
						"allow_looker_mentions": {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "Boolean to toggle mentions of Looker in emails",
						},
						"allow_looker_links": {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "Boolean to toggle links to Looker in emails",
						},
						"custom_welcome_email_advanced": {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "Allow subject line and email heading customization in customized emails",
						},
						"setup_mentions": {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "Remove the word Looker from appearing in the account setup page",
						},
						"alerts_logo": {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "Remove Looker logo from Alerts",
						},
						"alerts_links": {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "Remove Looker links from Alerts",
						},
						"folders_mentions": {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "Remove Looker mentions in home folder page when you don't have any items saved",
						},
					},
				},
			},
			"custom_welcome_email": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Description: "Custom welcome email configuration",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "If true, custom email content will replace the default body of welcome emails",
						},
						"content": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Requires custom_welcome_email to be enabled. The HTML to use as custom content for welcome emails. Script elements and other potentially dangerous markup will be removed",
						},
						"subject": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Requires custom_welcome_email and privatelabel_configuration.custom_welcome_email_advanced to be enabled. The text to appear in the email subject line. Only available with a whitelabel license and whitelabel_configuration.advanced_custom_welcome_email enabled",
						},
						"header": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Requires custom_welcome_email and privatelabel_configuration.custom_welcome_email_advanced to be enabled. The text to appear in the header line of the email body. Only available with a whitelabel license and whitelabel_configuration.advanced_custom_welcome_email enabled",
						},
					},
				},
			},
			"onboarding_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Toggle onboarding on or off",
			},
			"timezone": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Change instance-wide default timezone",
			},
			"allow_user_timezones": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Toggle user-specific timezones on or off",
			},
			"data_connector_default_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Toggle default future connectors on or off",
			},
			"host_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Change the base portion of your Looker instance URL setting",
			},
			"override_warnings": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "(Write-Only) If warnings are preventing a host URL change, this parameter allows for overriding warnings to force update the setting. Does not directly change any Looker settings",
			},
			"email_domain_allowlist": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "List of email domains that are allowed to be used for user creation",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"embed_cookieless_v2": {
				Type:       schema.TypeBool,
				Optional:   true,
				Computed:   true,
				Deprecated: "(DEPRECATED) Use embed_config.embed_cookieless_v2 instead. If embed_config.embed_cookieless_v2 is specified, it overrides this value",
			},
			"embed_enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "True if embedding is enabled https://cloud.google.com/looker/docs/r/looker-core-feature-embed, false otherwise",
			},
			"embed_config": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Description: "Embed configuration. Requires embedding to be enabled https://cloud.google.com/looker/docs/r/looker-core-feature-embed",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain_allowlist": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"alert_url_allowlist": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"alert_url_param_owner": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Owner of who defines the alert/schedule params on the base url",
						},
						"alert_url_label": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Label for the alert/schedule url",
						},
						"sso_auth_enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "Is SSO embedding enabled for this Looker",
						},
						"embed_cookieless_v2": {
							Type:         schema.TypeBool,
							Optional:     true,
							Computed:     true,
							RequiredWith: []string{"embed_enabled"},
							Description:  "Is Cookieless embedding enabled for this Looker",
						},
						"embed_content_navigation": {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "Is embed content navigation enabled for this looker",
						},
						"embed_content_management": {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "Is embed content management enabled for this Looker",
						},
						"strict_sameorigin_for_login": {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "When true, prohibits the use of Looker login pages in non-Looker iframes. When false, Looker login pages may be used in non-Looker hosted iframes",
						},
						"look_filters": {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "When true, filters are enabled on embedded Looks",
						},
						"hide_look_navigation": {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "When true, removes navigation to Looks from embedded dashboards and explores",
						},
						"embed_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "True if embedding is licensed for this Looker instance",
						},
					},
				},
			},
			"login_notification_enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Toggle login notification on or off",
			},
			"login_notification_text": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Text to display in the login notification banner",
			},
			"dashboard_autorefresh_restriction": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Toggle dashboard auto refresh restriction",
			},
			"dashboard_auto_refresh_minimum_interval": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Minimum time interval for dashboard element automatic refresh. Examples: (30 seconds, 1 minute)",
			},
			"managed_certificate_uri": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceSettingRead(ctx context.Context, d *schema.ResourceData, m any) (diags diag.Diagnostics) {
	c := m.(*Config).Api

	d.SetId("looker_settings")

	setting, _, err := c.Setting.Get(ctx)
	if err != nil {
		return diag.FromErr(err)
	}

	settingItems, err := setting.ToMap()
	if err != nil {
		return diag.FromErr(err)
	}

	for key, val := range settingItems {
		if err := d.Set(key, val); err != nil {
			return diag.FromErr(err)
		}
	}

	return diags
}

func resourceSettingUpdate(ctx context.Context, d *schema.ResourceData, m any) (diags diag.Diagnostics) {
	c := m.(*Config).Api
	setting, _, err := c.Setting.Get(ctx)
	if err != nil {
		return diag.FromErr(err)
	}

	setting.CleanFromReadOnly()

	settingItems, err := setting.ToMap()
	if err != nil {
		return diag.FromErr(err)
	}

	// Checks all fields from `setting` and compare them with values in `d`
	for key := range settingItems {
		if d.HasChange(key) {
			tflog.Info(ctx, "Updating Looker Setting", map[string]any{"key": key, "value": d.Get(key)})
			settingItems[key] = d.Get(key)
		}
	}

	if privatelabelConfiguration, ok := settingItems["privatelabel_configuration"]; ok && privatelabelConfiguration != nil {
		privatelabelConfiguration := privatelabelConfiguration.(map[string]any)

		if customWelcomeEmailAdvanced, ok := privatelabelConfiguration["custom_welcome_email_advanced"]; ok && customWelcomeEmailAdvanced != nil {
			customWelcomeEmailAdvanced := customWelcomeEmailAdvanced.(bool)

			if !customWelcomeEmailAdvanced {
				// If custom_welcome_email_advanced is set to false, remove subject and header from the custom_welcome_email configuration

				if customWelcomeEmail, ok := settingItems["custom_welcome_email"]; ok && customWelcomeEmail != nil {
					customWelcomeEmail := customWelcomeEmail.(map[string]any)

					tflog.Info(ctx, "Removing custom_welcome_email subject and header - privatelabel_configuration.custom_welcome_email_advanced is false")

					customWelcomeEmail["subject"] = ""
					customWelcomeEmail["header"] = ""
				}
			}
		}
	}

	if customWelcomeEmail, ok := settingItems["custom_welcome_email"]; ok && customWelcomeEmail != nil {
		customWelcomeEmail := customWelcomeEmail.(map[string]any)

		if enabled, ok := customWelcomeEmail["enabled"]; ok && enabled != nil {
			enabled := enabled.(bool)

			if !enabled {
				tflog.Info(ctx, "Removing custom_welcome_email content, subject and header - custom_welcome_email.enabled is false")

				customWelcomeEmail["content"] = ""
				customWelcomeEmail["subject"] = ""
				customWelcomeEmail["header"] = ""
			}
		}
	}

	setting.FromMap(settingItems)

	// Checks specifically for write-only fields in `d`
	if d.HasChange("override_warnings") {
		overrideWarnings := d.Get("override_warnings").(bool)
		setting.OverrideWarnings = &overrideWarnings
	}

	_, _, err = c.Setting.Update(ctx, setting)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceSettingRead(ctx, d, m)
}

func resourceSettingDelete(ctx context.Context, d *schema.ResourceData, m any) (diags diag.Diagnostics) {
	d.SetId("")
	return diags
}
