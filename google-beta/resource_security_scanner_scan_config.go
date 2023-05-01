// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceSecurityScannerScanConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecurityScannerScanConfigCreate,
		Read:   resourceSecurityScannerScanConfigRead,
		Update: resourceSecurityScannerScanConfigUpdate,
		Delete: resourceSecurityScannerScanConfigDelete,

		Importer: &schema.ResourceImporter{
			State: resourceSecurityScannerScanConfigImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The user provider display name of the ScanConfig.`,
			},
			"starting_urls": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `The starting URLs from which the scanner finds site pages.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"authentication": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `The authentication configuration.
If specified, service will use the authentication configuration during scanning.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"custom_account": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Describes authentication configuration that uses a custom account.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"login_url": {
										Type:        schema.TypeString,
										Required:    true,
										Description: `The login form URL of the website.`,
									},
									"password": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
										Description: `The password of the custom account. The credential is stored encrypted
in GCP.`,
										Sensitive: true,
									},
									"username": {
										Type:        schema.TypeString,
										Required:    true,
										Description: `The user name of the custom account.`,
									},
								},
							},
							AtLeastOneOf: []string{"authentication.0.google_account", "authentication.0.custom_account"},
						},
						"google_account": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Describes authentication configuration that uses a Google account.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"password": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
										Description: `The password of the Google account. The credential is stored encrypted
in GCP.`,
										Sensitive: true,
									},
									"username": {
										Type:        schema.TypeString,
										Required:    true,
										Description: `The user name of the Google account.`,
									},
								},
							},
							AtLeastOneOf: []string{"authentication.0.google_account", "authentication.0.custom_account"},
						},
					},
				},
			},
			"blacklist_patterns": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `The blacklist URL patterns as described in
https://cloud.google.com/security-scanner/docs/excluded-urls`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"export_to_security_command_center": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateEnum([]string{"ENABLED", "DISABLED", ""}),
				Description:  `Controls export of scan configurations and results to Cloud Security Command Center. Default value: "ENABLED" Possible values: ["ENABLED", "DISABLED"]`,
				Default:      "ENABLED",
			},
			"max_qps": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntBetween(5, 20),
				Description: `The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively.
Defaults to 15.`,
				Default: 15,
			},
			"schedule": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `The schedule of the ScanConfig`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interval_duration_days": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: `The duration of time between executions in days`,
						},
						"schedule_time": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `A timestamp indicates when the next run will be scheduled. The value is refreshed
by the server after each run. If unspecified, it will default to current server time,
which means the scan will be scheduled to start immediately.`,
						},
					},
				},
			},
			"target_platforms": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Set of Cloud Platforms targeted by the scan. If empty, APP_ENGINE will be used as a default. Possible values: ["APP_ENGINE", "COMPUTE"]`,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validateEnum([]string{"APP_ENGINE", "COMPUTE"}),
				},
			},
			"user_agent": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateEnum([]string{"USER_AGENT_UNSPECIFIED", "CHROME_LINUX", "CHROME_ANDROID", "SAFARI_IPHONE", ""}),
				Description:  `Type of the user agents used for scanning Default value: "CHROME_LINUX" Possible values: ["USER_AGENT_UNSPECIFIED", "CHROME_LINUX", "CHROME_ANDROID", "SAFARI_IPHONE"]`,
				Default:      "CHROME_LINUX",
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `A server defined name for this index. Format:
'projects/{{project}}/scanConfigs/{{server_generated_id}}'`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceSecurityScannerScanConfigCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandSecurityScannerScanConfigDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	maxQpsProp, err := expandSecurityScannerScanConfigMaxQps(d.Get("max_qps"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("max_qps"); !isEmptyValue(reflect.ValueOf(maxQpsProp)) && (ok || !reflect.DeepEqual(v, maxQpsProp)) {
		obj["maxQps"] = maxQpsProp
	}
	startingUrlsProp, err := expandSecurityScannerScanConfigStartingUrls(d.Get("starting_urls"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("starting_urls"); !isEmptyValue(reflect.ValueOf(startingUrlsProp)) && (ok || !reflect.DeepEqual(v, startingUrlsProp)) {
		obj["startingUrls"] = startingUrlsProp
	}
	authenticationProp, err := expandSecurityScannerScanConfigAuthentication(d.Get("authentication"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("authentication"); !isEmptyValue(reflect.ValueOf(authenticationProp)) && (ok || !reflect.DeepEqual(v, authenticationProp)) {
		obj["authentication"] = authenticationProp
	}
	userAgentProp, err := expandSecurityScannerScanConfigUserAgent(d.Get("user_agent"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("user_agent"); !isEmptyValue(reflect.ValueOf(userAgentProp)) && (ok || !reflect.DeepEqual(v, userAgentProp)) {
		obj["userAgent"] = userAgentProp
	}
	blacklistPatternsProp, err := expandSecurityScannerScanConfigBlacklistPatterns(d.Get("blacklist_patterns"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("blacklist_patterns"); !isEmptyValue(reflect.ValueOf(blacklistPatternsProp)) && (ok || !reflect.DeepEqual(v, blacklistPatternsProp)) {
		obj["blacklistPatterns"] = blacklistPatternsProp
	}
	scheduleProp, err := expandSecurityScannerScanConfigSchedule(d.Get("schedule"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("schedule"); !isEmptyValue(reflect.ValueOf(scheduleProp)) && (ok || !reflect.DeepEqual(v, scheduleProp)) {
		obj["schedule"] = scheduleProp
	}
	targetPlatformsProp, err := expandSecurityScannerScanConfigTargetPlatforms(d.Get("target_platforms"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("target_platforms"); !isEmptyValue(reflect.ValueOf(targetPlatformsProp)) && (ok || !reflect.DeepEqual(v, targetPlatformsProp)) {
		obj["targetPlatforms"] = targetPlatformsProp
	}
	exportToSecurityCommandCenterProp, err := expandSecurityScannerScanConfigExportToSecurityCommandCenter(d.Get("export_to_security_command_center"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("export_to_security_command_center"); !isEmptyValue(reflect.ValueOf(exportToSecurityCommandCenterProp)) && (ok || !reflect.DeepEqual(v, exportToSecurityCommandCenterProp)) {
		obj["exportToSecurityCommandCenter"] = exportToSecurityCommandCenterProp
	}

	url, err := ReplaceVars(d, config, "{{SecurityScannerBasePath}}projects/{{project}}/scanConfigs")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ScanConfig: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ScanConfig: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating ScanConfig: %s", err)
	}
	if err := d.Set("name", flattenSecurityScannerScanConfigName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// `name` is autogenerated from the api so needs to be set post-create
	name, ok := res["name"]
	if !ok {
		respBody, ok := res["response"]
		if !ok {
			return fmt.Errorf("Create response didn't contain critical fields. Create may not have succeeded.")
		}

		name, ok = respBody.(map[string]interface{})["name"]
		if !ok {
			return fmt.Errorf("Create response didn't contain critical fields. Create may not have succeeded.")
		}
	}
	if err := d.Set("name", name.(string)); err != nil {
		return fmt.Errorf("Error setting name: %s", err)
	}
	d.SetId(name.(string))

	log.Printf("[DEBUG] Finished creating ScanConfig %q: %#v", d.Id(), res)

	return resourceSecurityScannerScanConfigRead(d, meta)
}

func resourceSecurityScannerScanConfigRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{SecurityScannerBasePath}}{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ScanConfig: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("SecurityScannerScanConfig %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ScanConfig: %s", err)
	}

	if err := d.Set("name", flattenSecurityScannerScanConfigName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading ScanConfig: %s", err)
	}
	if err := d.Set("display_name", flattenSecurityScannerScanConfigDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading ScanConfig: %s", err)
	}
	if err := d.Set("max_qps", flattenSecurityScannerScanConfigMaxQps(res["maxQps"], d, config)); err != nil {
		return fmt.Errorf("Error reading ScanConfig: %s", err)
	}
	if err := d.Set("starting_urls", flattenSecurityScannerScanConfigStartingUrls(res["startingUrls"], d, config)); err != nil {
		return fmt.Errorf("Error reading ScanConfig: %s", err)
	}
	if err := d.Set("authentication", flattenSecurityScannerScanConfigAuthentication(res["authentication"], d, config)); err != nil {
		return fmt.Errorf("Error reading ScanConfig: %s", err)
	}
	if err := d.Set("user_agent", flattenSecurityScannerScanConfigUserAgent(res["userAgent"], d, config)); err != nil {
		return fmt.Errorf("Error reading ScanConfig: %s", err)
	}
	if err := d.Set("blacklist_patterns", flattenSecurityScannerScanConfigBlacklistPatterns(res["blacklistPatterns"], d, config)); err != nil {
		return fmt.Errorf("Error reading ScanConfig: %s", err)
	}
	if err := d.Set("schedule", flattenSecurityScannerScanConfigSchedule(res["schedule"], d, config)); err != nil {
		return fmt.Errorf("Error reading ScanConfig: %s", err)
	}
	if err := d.Set("target_platforms", flattenSecurityScannerScanConfigTargetPlatforms(res["targetPlatforms"], d, config)); err != nil {
		return fmt.Errorf("Error reading ScanConfig: %s", err)
	}
	if err := d.Set("export_to_security_command_center", flattenSecurityScannerScanConfigExportToSecurityCommandCenter(res["exportToSecurityCommandCenter"], d, config)); err != nil {
		return fmt.Errorf("Error reading ScanConfig: %s", err)
	}

	return nil
}

func resourceSecurityScannerScanConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ScanConfig: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandSecurityScannerScanConfigDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	maxQpsProp, err := expandSecurityScannerScanConfigMaxQps(d.Get("max_qps"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("max_qps"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, maxQpsProp)) {
		obj["maxQps"] = maxQpsProp
	}
	startingUrlsProp, err := expandSecurityScannerScanConfigStartingUrls(d.Get("starting_urls"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("starting_urls"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, startingUrlsProp)) {
		obj["startingUrls"] = startingUrlsProp
	}
	authenticationProp, err := expandSecurityScannerScanConfigAuthentication(d.Get("authentication"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("authentication"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, authenticationProp)) {
		obj["authentication"] = authenticationProp
	}
	userAgentProp, err := expandSecurityScannerScanConfigUserAgent(d.Get("user_agent"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("user_agent"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, userAgentProp)) {
		obj["userAgent"] = userAgentProp
	}
	blacklistPatternsProp, err := expandSecurityScannerScanConfigBlacklistPatterns(d.Get("blacklist_patterns"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("blacklist_patterns"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, blacklistPatternsProp)) {
		obj["blacklistPatterns"] = blacklistPatternsProp
	}
	scheduleProp, err := expandSecurityScannerScanConfigSchedule(d.Get("schedule"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("schedule"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, scheduleProp)) {
		obj["schedule"] = scheduleProp
	}
	targetPlatformsProp, err := expandSecurityScannerScanConfigTargetPlatforms(d.Get("target_platforms"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("target_platforms"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, targetPlatformsProp)) {
		obj["targetPlatforms"] = targetPlatformsProp
	}
	exportToSecurityCommandCenterProp, err := expandSecurityScannerScanConfigExportToSecurityCommandCenter(d.Get("export_to_security_command_center"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("export_to_security_command_center"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, exportToSecurityCommandCenterProp)) {
		obj["exportToSecurityCommandCenter"] = exportToSecurityCommandCenterProp
	}

	url, err := ReplaceVars(d, config, "{{SecurityScannerBasePath}}{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating ScanConfig %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("max_qps") {
		updateMask = append(updateMask, "maxQps")
	}

	if d.HasChange("starting_urls") {
		updateMask = append(updateMask, "startingUrls")
	}

	if d.HasChange("authentication") {
		updateMask = append(updateMask, "authentication")
	}

	if d.HasChange("user_agent") {
		updateMask = append(updateMask, "userAgent")
	}

	if d.HasChange("blacklist_patterns") {
		updateMask = append(updateMask, "blacklistPatterns")
	}

	if d.HasChange("schedule") {
		updateMask = append(updateMask, "schedule")
	}

	if d.HasChange("target_platforms") {
		updateMask = append(updateMask, "targetPlatforms")
	}

	if d.HasChange("export_to_security_command_center") {
		updateMask = append(updateMask, "exportToSecurityCommandCenter")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating ScanConfig %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating ScanConfig %q: %#v", d.Id(), res)
	}

	return resourceSecurityScannerScanConfigRead(d, meta)
}

func resourceSecurityScannerScanConfigDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ScanConfig: %s", err)
	}
	billingProject = project

	url, err := ReplaceVars(d, config, "{{SecurityScannerBasePath}}{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting ScanConfig %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "ScanConfig")
	}

	log.Printf("[DEBUG] Finished deleting ScanConfig %q: %#v", d.Id(), res)
	return nil
}

func resourceSecurityScannerScanConfigImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	config := meta.(*transport_tpg.Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := ParseImportId([]string{"(?P<project>[^ ]+) (?P<name>[^ ]+)", "(?P<name>[^ ]+)"}, d, config); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

func flattenSecurityScannerScanConfigName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityScannerScanConfigDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityScannerScanConfigMaxQps(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenSecurityScannerScanConfigStartingUrls(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityScannerScanConfigAuthentication(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["google_account"] =
		flattenSecurityScannerScanConfigAuthenticationGoogleAccount(original["googleAccount"], d, config)
	transformed["custom_account"] =
		flattenSecurityScannerScanConfigAuthenticationCustomAccount(original["customAccount"], d, config)
	return []interface{}{transformed}
}
func flattenSecurityScannerScanConfigAuthenticationGoogleAccount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["username"] =
		flattenSecurityScannerScanConfigAuthenticationGoogleAccountUsername(original["username"], d, config)
	transformed["password"] =
		flattenSecurityScannerScanConfigAuthenticationGoogleAccountPassword(original["password"], d, config)
	return []interface{}{transformed}
}
func flattenSecurityScannerScanConfigAuthenticationGoogleAccountUsername(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityScannerScanConfigAuthenticationGoogleAccountPassword(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return d.Get("authentication.0.custom_account.0.password")
}

func flattenSecurityScannerScanConfigAuthenticationCustomAccount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["username"] =
		flattenSecurityScannerScanConfigAuthenticationCustomAccountUsername(original["username"], d, config)
	transformed["password"] =
		flattenSecurityScannerScanConfigAuthenticationCustomAccountPassword(original["password"], d, config)
	transformed["login_url"] =
		flattenSecurityScannerScanConfigAuthenticationCustomAccountLoginUrl(original["loginUrl"], d, config)
	return []interface{}{transformed}
}
func flattenSecurityScannerScanConfigAuthenticationCustomAccountUsername(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityScannerScanConfigAuthenticationCustomAccountPassword(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return d.Get("authentication.0.google_account.0.password")
}

func flattenSecurityScannerScanConfigAuthenticationCustomAccountLoginUrl(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityScannerScanConfigUserAgent(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityScannerScanConfigBlacklistPatterns(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityScannerScanConfigSchedule(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["schedule_time"] =
		flattenSecurityScannerScanConfigScheduleScheduleTime(original["scheduleTime"], d, config)
	transformed["interval_duration_days"] =
		flattenSecurityScannerScanConfigScheduleIntervalDurationDays(original["intervalDurationDays"], d, config)
	return []interface{}{transformed}
}
func flattenSecurityScannerScanConfigScheduleScheduleTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityScannerScanConfigScheduleIntervalDurationDays(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenSecurityScannerScanConfigTargetPlatforms(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityScannerScanConfigExportToSecurityCommandCenter(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandSecurityScannerScanConfigDisplayName(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityScannerScanConfigMaxQps(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityScannerScanConfigStartingUrls(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityScannerScanConfigAuthentication(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedGoogleAccount, err := expandSecurityScannerScanConfigAuthenticationGoogleAccount(original["google_account"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGoogleAccount); val.IsValid() && !isEmptyValue(val) {
		transformed["googleAccount"] = transformedGoogleAccount
	}

	transformedCustomAccount, err := expandSecurityScannerScanConfigAuthenticationCustomAccount(original["custom_account"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCustomAccount); val.IsValid() && !isEmptyValue(val) {
		transformed["customAccount"] = transformedCustomAccount
	}

	return transformed, nil
}

func expandSecurityScannerScanConfigAuthenticationGoogleAccount(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedUsername, err := expandSecurityScannerScanConfigAuthenticationGoogleAccountUsername(original["username"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUsername); val.IsValid() && !isEmptyValue(val) {
		transformed["username"] = transformedUsername
	}

	transformedPassword, err := expandSecurityScannerScanConfigAuthenticationGoogleAccountPassword(original["password"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPassword); val.IsValid() && !isEmptyValue(val) {
		transformed["password"] = transformedPassword
	}

	return transformed, nil
}

func expandSecurityScannerScanConfigAuthenticationGoogleAccountUsername(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityScannerScanConfigAuthenticationGoogleAccountPassword(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityScannerScanConfigAuthenticationCustomAccount(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedUsername, err := expandSecurityScannerScanConfigAuthenticationCustomAccountUsername(original["username"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUsername); val.IsValid() && !isEmptyValue(val) {
		transformed["username"] = transformedUsername
	}

	transformedPassword, err := expandSecurityScannerScanConfigAuthenticationCustomAccountPassword(original["password"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPassword); val.IsValid() && !isEmptyValue(val) {
		transformed["password"] = transformedPassword
	}

	transformedLoginUrl, err := expandSecurityScannerScanConfigAuthenticationCustomAccountLoginUrl(original["login_url"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLoginUrl); val.IsValid() && !isEmptyValue(val) {
		transformed["loginUrl"] = transformedLoginUrl
	}

	return transformed, nil
}

func expandSecurityScannerScanConfigAuthenticationCustomAccountUsername(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityScannerScanConfigAuthenticationCustomAccountPassword(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityScannerScanConfigAuthenticationCustomAccountLoginUrl(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityScannerScanConfigUserAgent(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityScannerScanConfigBlacklistPatterns(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityScannerScanConfigSchedule(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedScheduleTime, err := expandSecurityScannerScanConfigScheduleScheduleTime(original["schedule_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedScheduleTime); val.IsValid() && !isEmptyValue(val) {
		transformed["scheduleTime"] = transformedScheduleTime
	}

	transformedIntervalDurationDays, err := expandSecurityScannerScanConfigScheduleIntervalDurationDays(original["interval_duration_days"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIntervalDurationDays); val.IsValid() && !isEmptyValue(val) {
		transformed["intervalDurationDays"] = transformedIntervalDurationDays
	}

	return transformed, nil
}

func expandSecurityScannerScanConfigScheduleScheduleTime(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityScannerScanConfigScheduleIntervalDurationDays(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityScannerScanConfigTargetPlatforms(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityScannerScanConfigExportToSecurityCommandCenter(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
