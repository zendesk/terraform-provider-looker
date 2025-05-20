package provider

import (
	"context"

	"github.com/devoteamgcloud/terraform-provider-looker/pkg/lookergo"
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

	readBoolSetting := func(key string, value *bool) {
		if value != nil {
			d.Set(key, *value)
		}
	}

	readStringSetting := func(key string, value *string) {
		if value != nil {
			d.Set(key, *value)
		}
	}

	readBoolSetting("extension_framework_enabled", setting.ExtensionFrameworkEnabled)
	readBoolSetting("extension_load_url_enabled", setting.ExtensionLoadUrlEnabled)
	readBoolSetting("marketplace_auto_install_enabled", setting.MarketplaceAutoInstallEnabled)
	readBoolSetting("marketplace_enabled", setting.MarketplaceEnabled)
	readStringSetting("marketplace_site", setting.MarketplaceSite)
	readBoolSetting("marketplace_terms_accepted", setting.MarketplaceTermsAccepted)
	readBoolSetting("onboarding_enabled", setting.OnboardingEnabled)
	readStringSetting("timezone", setting.Timezone)
	readBoolSetting("allow_user_timezones", setting.AllowUserTimezones)
	readBoolSetting("data_connector_default_enabled", setting.DataConnectorDefaultEnabled)
	readStringSetting("host_url", setting.HostUrl)
	readBoolSetting("override_warnings", setting.OverrideWarnings)
	readBoolSetting("embed_cookieless_v2", setting.EmbedCookielessV2)
	readBoolSetting("embed_enabled", setting.EmbedEnabled)
	readBoolSetting("login_notification_enabled", setting.LoginNotificationEnabled)
	readStringSetting("login_notification_text", setting.LoginNotificationText)
	readBoolSetting("dashboard_autorefresh_restriction", setting.DashboardAutorefreshRestriction)
	readStringSetting("dashboard_auto_refresh_minimum_interval", setting.DashboardAutoRefreshMinimumInterval)

	// Lists
	d.Set("email_domain_allowlist", setting.EmailDomainAllowlist)
	d.Set("managed_certificate_uri", setting.ManagedCertificateUri)

	readInstanceConfig(d, setting)
	readMarketplaceAutomation(d, setting)
	readPrivatelabelConfiguration(d, setting)
	readCustomWelcomeEmail(d, setting)
	readEmbedConfig(d, setting)

	return diags
}

func readInstanceConfig(d *schema.ResourceData, setting *lookergo.Setting) {
	if setting.InstanceConfig != nil {
		instanceConfigMap := []map[string]any{{}}
		if setting.InstanceConfig.FeatureFlags != nil {
			instanceConfigMap[0]["feature_flags"] = *setting.InstanceConfig.FeatureFlags
		}
		if setting.InstanceConfig.LicenseFeatures != nil {
			instanceConfigMap[0]["license_features"] = *setting.InstanceConfig.LicenseFeatures
		}
		d.Set("instance_config", instanceConfigMap)
	}
}

func readMarketplaceAutomation(d *schema.ResourceData, setting *lookergo.Setting) {
	if setting.MarketplaceAutomation != nil {
		marketplaceAutomationMap := []map[string]any{{
			"install_enabled":            false,
			"update_looker_enabled":      false,
			"update_third_party_enabled": false,
		}}

		if setting.MarketplaceAutomation.InstallEnabled != nil {
			marketplaceAutomationMap[0]["install_enabled"] = *setting.MarketplaceAutomation.InstallEnabled
		}
		if setting.MarketplaceAutomation.UpdateLookerEnabled != nil {
			marketplaceAutomationMap[0]["update_looker_enabled"] = *setting.MarketplaceAutomation.UpdateLookerEnabled
		}
		if setting.MarketplaceAutomation.UpdateThirdPartyEnabled != nil {
			marketplaceAutomationMap[0]["update_third_party_enabled"] = *setting.MarketplaceAutomation.UpdateThirdPartyEnabled
		}

		d.Set("marketplace_automation", marketplaceAutomationMap)
	}
}

func readPrivatelabelConfiguration(d *schema.ResourceData, setting *lookergo.Setting) {
	if setting.PrivatelabelConfiguration != nil {
		plc := setting.PrivatelabelConfiguration
		plcMap := map[string]any{
			"logo_file":                     "",
			"favicon_file":                  "",
			"default_title":                 "",
			"show_help_menu":                false,
			"show_docs":                     false,
			"show_email_sub_options":        false,
			"allow_looker_mentions":         false,
			"allow_looker_links":            false,
			"custom_welcome_email_advanced": false,
			"setup_mentions":                false,
			"alerts_logo":                   false,
			"alerts_links":                  false,
			"folders_mentions":              false,
		}

		// Set computed values
		if plc.LogoUrl != nil {
			plcMap["logo_url"] = *plc.LogoUrl
		}
		if plc.FaviconUrl != nil {
			plcMap["favicon_url"] = *plc.FaviconUrl
		}

		// Set configurable values
		if plc.DefaultTitle != nil {
			plcMap["default_title"] = *plc.DefaultTitle
		}
		if plc.ShowHelpMenu != nil {
			plcMap["show_help_menu"] = *plc.ShowHelpMenu
		}
		if plc.ShowDocs != nil {
			plcMap["show_docs"] = *plc.ShowDocs
		}
		if plc.ShowEmailSubOptions != nil {
			plcMap["show_email_sub_options"] = *plc.ShowEmailSubOptions
		}
		if plc.AllowLookerMentions != nil {
			plcMap["allow_looker_mentions"] = *plc.AllowLookerMentions
		}
		if plc.AllowLookerLinks != nil {
			plcMap["allow_looker_links"] = *plc.AllowLookerLinks
		}
		if plc.CustomWelcomeEmailAdvanced != nil {
			plcMap["custom_welcome_email_advanced"] = *plc.CustomWelcomeEmailAdvanced
		}
		if plc.SetupMentions != nil {
			plcMap["setup_mentions"] = *plc.SetupMentions
		}
		if plc.AlertsLogo != nil {
			plcMap["alerts_logo"] = *plc.AlertsLogo
		}
		if plc.AlertsLinks != nil {
			plcMap["alerts_links"] = *plc.AlertsLinks
		}
		if plc.FoldersMentions != nil {
			plcMap["folders_mentions"] = *plc.FoldersMentions
		}

		d.Set("privatelabel_configuration", []map[string]any{plcMap})
	}
}

func readCustomWelcomeEmail(d *schema.ResourceData, setting *lookergo.Setting) {
	if setting.CustomWelcomeEmail != nil {
		cwe := setting.CustomWelcomeEmail
		cweMap := map[string]any{
			"enabled": false,
			"content": "",
			"subject": "",
			"header":  "",
		}

		if cwe.Enabled != nil {
			cweMap["enabled"] = *cwe.Enabled
		}
		if cwe.Content != nil {
			cweMap["content"] = *cwe.Content
		}
		if cwe.Subject != nil {
			cweMap["subject"] = *cwe.Subject
		}
		if cwe.Header != nil {
			cweMap["header"] = *cwe.Header
		}

		d.Set("custom_welcome_email", []map[string]any{cweMap})
	}
}

func readEmbedConfig(d *schema.ResourceData, setting *lookergo.Setting) {
	if setting.EmbedConfig != nil {
		ec := setting.EmbedConfig
		ecMap := map[string]any{
			"domain_allowlist":            ec.DomainAllowlist,
			"alert_url_allowlist":         ec.AlertUrlAllowlist,
			"alert_url_param_owner":       ec.AlertUrlParamOwner,
			"alert_url_label":             ec.AlertUrlLabel,
			"sso_auth_enabled":            false,
			"embed_cookieless_v2":         false,
			"embed_content_navigation":    false,
			"embed_content_management":    false,
			"strict_sameorigin_for_login": false,
			"look_filters":                false,
			"hide_look_navigation":        false,
			"embed_enabled":               false,
		}

		if ec.SsoAuthEnabled != nil {
			ecMap["sso_auth_enabled"] = *ec.SsoAuthEnabled
		}
		if ec.EmbedCookielessV2 != nil {
			ecMap["embed_cookieless_v2"] = *ec.EmbedCookielessV2
		}
		if ec.EmbedContentNavigation != nil {
			ecMap["embed_content_navigation"] = *ec.EmbedContentNavigation
		}
		if ec.EmbedContentManagement != nil {
			ecMap["embed_content_management"] = *ec.EmbedContentManagement
		}
		if ec.StrictSameoriginForLogin != nil {
			ecMap["strict_sameorigin_for_login"] = *ec.StrictSameoriginForLogin
		}
		if ec.LookFilters != nil {
			ecMap["look_filters"] = *ec.LookFilters
		}
		if ec.HideLookNavigation != nil {
			ecMap["hide_look_navigation"] = *ec.HideLookNavigation
		}
		if ec.EmbedEnabled != nil {
			ecMap["embed_enabled"] = *ec.EmbedEnabled
		}

		d.Set("embed_config", []map[string]any{ecMap})
	}
}

func resourceSettingUpdate(ctx context.Context, d *schema.ResourceData, m any) (diags diag.Diagnostics) {
	c := m.(*Config).Api
	setting, _, err := c.Setting.Get(ctx)
	if err != nil {
		return diag.FromErr(err)
	}

	setting.CleanFromReadOnly()

	updateBoolSetting := func(key string, target **bool) {
		if d.HasChange(key) {
			val := d.Get(key).(bool)
			*target = &val
		}
	}

	updateStringSetting := func(key string, target **string) {
		if d.HasChange(key) {
			val := d.Get(key).(string)
			*target = &val
		}
	}

	updateStringList := func(key string, target *[]string) {
		if d.HasChange(key) {
			values := d.Get(key).([]any)
			result := make([]string, len(values))
			for i, v := range values {
				result[i] = v.(string)
			}
			*target = result
		}
	}

	updateBoolSetting("extension_framework_enabled", &setting.ExtensionFrameworkEnabled)
	updateBoolSetting("extension_load_url_enabled", &setting.ExtensionLoadUrlEnabled)
	updateBoolSetting("marketplace_auto_install_enabled", &setting.MarketplaceAutoInstallEnabled)
	updateBoolSetting("marketplace_enabled", &setting.MarketplaceEnabled)
	updateBoolSetting("marketplace_terms_accepted", &setting.MarketplaceTermsAccepted)
	updateBoolSetting("onboarding_enabled", &setting.OnboardingEnabled)
	updateBoolSetting("allow_user_timezones", &setting.AllowUserTimezones)
	updateBoolSetting("data_connector_default_enabled", &setting.DataConnectorDefaultEnabled)
	updateBoolSetting("override_warnings", &setting.OverrideWarnings)
	updateBoolSetting("embed_cookieless_v2", &setting.EmbedCookielessV2)
	updateBoolSetting("dashboard_autorefresh_restriction", &setting.DashboardAutorefreshRestriction)

	updateStringSetting("timezone", &setting.Timezone)
	updateStringSetting("host_url", &setting.HostUrl)
	updateStringSetting("dashboard_auto_refresh_minimum_interval", &setting.DashboardAutoRefreshMinimumInterval)

	updateStringList("email_domain_allowlist", &setting.EmailDomainAllowlist)
	updateStringList("managed_certificate_uri", &setting.ManagedCertificateUri)

	updateMarketplaceAutomation(d, setting)
	updatePrivatelabelConfiguration(d, setting)
	updateCustomWelcomeEmail(d, setting)
	updateEmbedConfig(d, setting)

	handleCustomWelcomeEmail(ctx, setting)

	_, _, err = c.Setting.Update(ctx, setting)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceSettingRead(ctx, d, m)
}

func updateMarketplaceAutomation(d *schema.ResourceData, setting *lookergo.Setting) {
	if d.HasChange("marketplace_automation") {
		set := d.Get("marketplace_automation").(*schema.Set)
		if len(set.List()) > 0 {
			data := set.List()[0].(map[string]any)

			installEnabled := data["install_enabled"].(bool)
			updateLookerEnabled := data["update_looker_enabled"].(bool)
			updateThirdPartyEnabled := data["update_third_party_enabled"].(bool)

			setting.MarketplaceAutomation = &lookergo.MarketplaceAutomation{
				InstallEnabled:          &installEnabled,
				UpdateLookerEnabled:     &updateLookerEnabled,
				UpdateThirdPartyEnabled: &updateThirdPartyEnabled,
			}
		}
	}
}

func updatePrivatelabelConfiguration(d *schema.ResourceData, setting *lookergo.Setting) {
	if d.HasChange("privatelabel_configuration") {
		privatelabelConfigurationSet := d.Get("privatelabel_configuration").(*schema.Set)
		privatelabelConfiguration := privatelabelConfigurationSet.List()[0].(map[string]any)

		logoFile := privatelabelConfiguration["logo_file"].(string)
		faviconFile := privatelabelConfiguration["favicon_file"].(string)
		defaultTitle := privatelabelConfiguration["default_title"].(string)
		showHelpMenu := privatelabelConfiguration["show_help_menu"].(bool)
		showDocs := privatelabelConfiguration["show_docs"].(bool)
		showEmailSubOptions := privatelabelConfiguration["show_email_sub_options"].(bool)
		allowLookerMentions := privatelabelConfiguration["allow_looker_mentions"].(bool)
		allowLookerLinks := privatelabelConfiguration["allow_looker_links"].(bool)
		customWelcomeEmailAdvanced := privatelabelConfiguration["custom_welcome_email_advanced"].(bool)
		setupMentions := privatelabelConfiguration["setup_mentions"].(bool)
		alertsLogo := privatelabelConfiguration["alerts_logo"].(bool)
		alertsLinks := privatelabelConfiguration["alerts_links"].(bool)
		foldersMentions := privatelabelConfiguration["folders_mentions"].(bool)

		setting.PrivatelabelConfiguration = &lookergo.PrivatelabelConfiguration{
			LogoFile:                   &logoFile,
			FaviconFile:                &faviconFile,
			DefaultTitle:               &defaultTitle,
			ShowHelpMenu:               &showHelpMenu,
			ShowDocs:                   &showDocs,
			ShowEmailSubOptions:        &showEmailSubOptions,
			AllowLookerMentions:        &allowLookerMentions,
			AllowLookerLinks:           &allowLookerLinks,
			CustomWelcomeEmailAdvanced: &customWelcomeEmailAdvanced,
			SetupMentions:              &setupMentions,
			AlertsLogo:                 &alertsLogo,
			AlertsLinks:                &alertsLinks,
			FoldersMentions:            &foldersMentions,
		}
	}
}

func updateCustomWelcomeEmail(d *schema.ResourceData, setting *lookergo.Setting) {
	if d.HasChange("custom_welcome_email") {
		customWelcomeEmailSet := d.Get("custom_welcome_email").(*schema.Set)
		customWelcomeEmail := customWelcomeEmailSet.List()[0].(map[string]any)

		customWelcomeEmailEnabled := customWelcomeEmail["enabled"].(bool)
		customWelcomeEmailContent := customWelcomeEmail["content"].(string)
		customWelcomeEmailSubject := customWelcomeEmail["subject"].(string)
		customWelcomeEmailHeader := customWelcomeEmail["header"].(string)

		setting.CustomWelcomeEmail = &lookergo.CustomWelcomeEmail{
			Enabled: &customWelcomeEmailEnabled,
			Content: &customWelcomeEmailContent,
			Subject: &customWelcomeEmailSubject,
			Header:  &customWelcomeEmailHeader,
		}
	}
}

func updateEmbedConfig(d *schema.ResourceData, setting *lookergo.Setting) {
	if d.HasChange("embed_config") {
		embedConfigSet := d.Get("embed_config").(*schema.Set)
		embedConfig := embedConfigSet.List()[0].(map[string]any)

		domainAllowlist := embedConfig["domain_allowlist"].([]any)
		alertUrlAllowlist := embedConfig["alert_url_allowlist"].([]any)
		alertUrlParamOwner := embedConfig["alert_url_param_owner"].(string)
		alertUrlLabel := embedConfig["alert_url_label"].(string)
		ssoAuthEnabled := embedConfig["sso_auth_enabled"].(bool)
		embedCookielessV2 := embedConfig["embed_cookieless_v2"].(bool)
		embedContentNavigation := embedConfig["embed_content_navigation"].(bool)
		embedContentManagement := embedConfig["embed_content_management"].(bool)
		strictSameoriginForLogin := embedConfig["strict_sameorigin_for_login"].(bool)
		lookFilters := embedConfig["look_filters"].(bool)
		hideLookNavigation := embedConfig["hide_look_navigation"].(bool)
		embedEnabled := embedConfig["embed_enabled"].(bool)

		domainAllowlistString := make([]string, len(domainAllowlist))
		for i, v := range domainAllowlist {
			domainAllowlistString[i] = v.(string)
		}

		alertUrlAllowlistString := make([]string, len(alertUrlAllowlist))
		for i, v := range alertUrlAllowlist {
			alertUrlAllowlistString[i] = v.(string)
		}

		setting.EmbedConfig = &lookergo.EmbedConfig{
			DomainAllowlist:          domainAllowlistString,
			AlertUrlAllowlist:        alertUrlAllowlistString,
			AlertUrlParamOwner:       alertUrlParamOwner,
			AlertUrlLabel:            alertUrlLabel,
			SsoAuthEnabled:           &ssoAuthEnabled,
			EmbedCookielessV2:        &embedCookielessV2,
			EmbedContentNavigation:   &embedContentNavigation,
			EmbedContentManagement:   &embedContentManagement,
			StrictSameoriginForLogin: &strictSameoriginForLogin,
			LookFilters:              &lookFilters,
			HideLookNavigation:       &hideLookNavigation,
			EmbedEnabled:             &embedEnabled,
		}
	}
}

func handleCustomWelcomeEmail(ctx context.Context, setting *lookergo.Setting) {
	if setting.PrivatelabelConfiguration != nil &&
		setting.PrivatelabelConfiguration.CustomWelcomeEmailAdvanced != nil &&
		!*setting.PrivatelabelConfiguration.CustomWelcomeEmailAdvanced {

		tflog.Info(ctx, "Custom welcome email advanced is disabled, setting subject and header to nil")
		setting.CustomWelcomeEmail.Subject = nil
		setting.CustomWelcomeEmail.Header = nil
	}

	if setting.CustomWelcomeEmail != nil &&
		setting.CustomWelcomeEmail.Enabled != nil &&
		!*setting.CustomWelcomeEmail.Enabled {

		tflog.Info(ctx, "Custom welcome email is disabled, setting content, subject and header to nil")
		setting.CustomWelcomeEmail.Content = nil
		setting.CustomWelcomeEmail.Subject = nil
		setting.CustomWelcomeEmail.Header = nil
	}
}

func resourceSettingDelete(ctx context.Context, d *schema.ResourceData, m any) (diags diag.Diagnostics) {
	d.SetId("")
	return diags
}
