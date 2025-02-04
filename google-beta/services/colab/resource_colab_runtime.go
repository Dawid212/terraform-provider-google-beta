// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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

package colab

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceColabRuntime() *schema.Resource {
	return &schema.Resource{
		Create: resourceColabRuntimeCreate,
		Read:   resourceColabRuntimeRead,
		Update: resourceColabRuntimeUpdate,
		Delete: resourceColabRuntimeDelete,

		Importer: &schema.ResourceImporter{
			State: resourceColabRuntimeImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Required. The display name of the Runtime.`,
			},
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The location for the resource: https://cloud.google.com/colab/docs/locations`,
			},
			"runtime_user": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The user email of the NotebookRuntime.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The description of the Runtime.`,
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The resource name of the Runtime`,
			},
			"notebook_runtime_template_ref": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `'Runtime specific information used for NotebookRuntime creation.'`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"notebook_runtime_template": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tpgresource.ProjectNumberDiffSuppress,
							Description:      `The resource name of the NotebookRuntimeTemplate based on which a NotebookRuntime will be created.`,
						},
					},
				},
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

func resourceColabRuntimeCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	notebookRuntimeTemplateRefProp, err := expandColabRuntimeNotebookRuntimeTemplateRef(d.Get("notebook_runtime_template_ref"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("notebook_runtime_template_ref"); !tpgresource.IsEmptyValue(reflect.ValueOf(notebookRuntimeTemplateRefProp)) && (ok || !reflect.DeepEqual(v, notebookRuntimeTemplateRefProp)) {
		obj["notebookRuntimeTemplateRef"] = notebookRuntimeTemplateRefProp
	}
	runtimeUserProp, err := expandColabRuntimeRuntimeUser(d.Get("runtime_user"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("runtime_user"); !tpgresource.IsEmptyValue(reflect.ValueOf(runtimeUserProp)) && (ok || !reflect.DeepEqual(v, runtimeUserProp)) {
		obj["runtimeUser"] = runtimeUserProp
	}
	displayNameProp, err := expandColabRuntimeDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandColabRuntimeDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}

	obj, err = resourceColabRuntimeEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ColabBasePath}}projects/{{project}}/locations/{{location}}/notebookRuntimes:assign?notebook_runtime_id={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Runtime: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Runtime: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating Runtime: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/notebookRuntimes/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = ColabOperationWaitTime(
		config, res, project, "Creating Runtime", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Runtime: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Runtime %q: %#v", d.Id(), res)

	return resourceColabRuntimeRead(d, meta)
}

func resourceColabRuntimeRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ColabBasePath}}projects/{{project}}/locations/{{location}}/notebookRuntimes/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Runtime: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ColabRuntime %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Runtime: %s", err)
	}

	if err := d.Set("notebook_runtime_template_ref", flattenColabRuntimeNotebookRuntimeTemplateRef(res["notebookRuntimeTemplateRef"], d, config)); err != nil {
		return fmt.Errorf("Error reading Runtime: %s", err)
	}
	if err := d.Set("runtime_user", flattenColabRuntimeRuntimeUser(res["runtimeUser"], d, config)); err != nil {
		return fmt.Errorf("Error reading Runtime: %s", err)
	}
	if err := d.Set("display_name", flattenColabRuntimeDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Runtime: %s", err)
	}
	if err := d.Set("description", flattenColabRuntimeDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Runtime: %s", err)
	}

	return nil
}

func resourceColabRuntimeUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Runtime: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	notebookRuntimeTemplateRefProp, err := expandColabRuntimeNotebookRuntimeTemplateRef(d.Get("notebook_runtime_template_ref"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("notebook_runtime_template_ref"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, notebookRuntimeTemplateRefProp)) {
		obj["notebookRuntimeTemplateRef"] = notebookRuntimeTemplateRefProp
	}
	runtimeUserProp, err := expandColabRuntimeRuntimeUser(d.Get("runtime_user"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("runtime_user"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, runtimeUserProp)) {
		obj["runtimeUser"] = runtimeUserProp
	}
	displayNameProp, err := expandColabRuntimeDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandColabRuntimeDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}

	obj, err = resourceColabRuntimeEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ColabBasePath}}projects/{{project}}/locations/{{location}}/notebookRuntimes/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Runtime %q: %#v", d.Id(), obj)
	headers := make(http.Header)

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PUT",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
		Headers:   headers,
	})

	if err != nil {
		return fmt.Errorf("Error updating Runtime %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Runtime %q: %#v", d.Id(), res)
	}

	err = ColabOperationWaitTime(
		config, res, project, "Updating Runtime", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceColabRuntimeRead(d, meta)
}

func resourceColabRuntimeDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Runtime: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{ColabBasePath}}projects/{{project}}/locations/{{location}}/notebookRuntimes/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting Runtime %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "Runtime")
	}

	err = ColabOperationWaitTime(
		config, res, project, "Deleting Runtime", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Runtime %q: %#v", d.Id(), res)
	return nil
}

func resourceColabRuntimeImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/notebookRuntimes/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/notebookRuntimes/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenColabRuntimeNotebookRuntimeTemplateRef(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["notebook_runtime_template"] =
		flattenColabRuntimeNotebookRuntimeTemplateRefNotebookRuntimeTemplate(original["notebookRuntimeTemplate"], d, config)
	return []interface{}{transformed}
}
func flattenColabRuntimeNotebookRuntimeTemplateRefNotebookRuntimeTemplate(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenColabRuntimeRuntimeUser(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenColabRuntimeDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenColabRuntimeDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandColabRuntimeNotebookRuntimeTemplateRef(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNotebookRuntimeTemplate, err := expandColabRuntimeNotebookRuntimeTemplateRefNotebookRuntimeTemplate(original["notebook_runtime_template"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNotebookRuntimeTemplate); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["notebookRuntimeTemplate"] = transformedNotebookRuntimeTemplate
	}

	return transformed, nil
}

func expandColabRuntimeNotebookRuntimeTemplateRefNotebookRuntimeTemplate(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandColabRuntimeRuntimeUser(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandColabRuntimeDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandColabRuntimeDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceColabRuntimeEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	newObj := make(map[string]interface{})
	newObj["notebookRuntimeTemplate"], _ = d.GetOk("notebook_runtime_template_ref.0.notebook_runtime_template")

	delete(obj, "notebookRuntimeTemplateRef")

	newObj["notebookRuntime"] = obj
	return newObj, nil
}
