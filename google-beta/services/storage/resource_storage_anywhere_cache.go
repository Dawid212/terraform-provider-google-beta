// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/storage/AnywhereCache.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package storage

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceStorageAnywhereCache() *schema.Resource {
	return &schema.Resource{
		Create: resourceStorageAnywhereCacheCreate,
		Read:   resourceStorageAnywhereCacheRead,
		Update: resourceStorageAnywhereCacheUpdate,
		Delete: resourceStorageAnywhereCacheDelete,

		Importer: &schema.ResourceImporter{
			State: resourceStorageAnywhereCacheImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(240 * time.Minute),
			Update: schema.DefaultTimeout(240 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"bucket": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `A reference to Bucket resource`,
			},
			"zone": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The zone in which the cache instance needs to be created. For example, 'us-central1-a.'`,
			},
			"admission_policy": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: verify.ValidateEnum([]string{"admit-on-first-miss", "admit-on-second-miss", ""}),
				Description:  `The cache admission policy dictates whether a block should be inserted upon a cache miss. Default value: "admit-on-first-miss" Possible values: ["admit-on-first-miss", "admit-on-second-miss"]`,
				Default:      "admit-on-first-miss",
			},
			"ttl": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The TTL of all cache entries in whole seconds. e.g., "7200s". It defaults to '86400s'`,
				Default:     "86400s",
			},
			"anywhere_cache_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The ID of the Anywhere cache instance.`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The creation time of the cache instance in RFC 3339 format.`,
			},
			"pending_update": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: `True if the cache instance has an active Update long-running operation.`,
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The current state of the cache instance.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The modification time of the cache instance metadata in RFC 3339 format.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceStorageAnywhereCacheCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	zoneProp, err := expandStorageAnywhereCacheZone(d.Get("zone"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("zone"); !tpgresource.IsEmptyValue(reflect.ValueOf(zoneProp)) && (ok || !reflect.DeepEqual(v, zoneProp)) {
		obj["zone"] = zoneProp
	}
	admissionPolicyProp, err := expandStorageAnywhereCacheAdmissionPolicy(d.Get("admission_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("admission_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(admissionPolicyProp)) && (ok || !reflect.DeepEqual(v, admissionPolicyProp)) {
		obj["admissionPolicy"] = admissionPolicyProp
	}
	ttlProp, err := expandStorageAnywhereCacheTtl(d.Get("ttl"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ttl"); !tpgresource.IsEmptyValue(reflect.ValueOf(ttlProp)) && (ok || !reflect.DeepEqual(v, ttlProp)) {
		obj["ttl"] = ttlProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{StorageBasePath}}b/{{bucket}}/anywhereCaches/")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new AnywhereCache: %#v", obj)
	billingProject := ""

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
		return fmt.Errorf("Error creating AnywhereCache: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{bucket}}/{{anywhere_cache_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = StorageOperationWaitTimeWithResponse(
		config, res, &opRes, "Creating AnywhereCache", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create AnywhereCache: %s", err)
	}

	if err := d.Set("anywhere_cache_id", flattenStorageAnywhereCacheAnywhereCacheId(opRes["anywhereCacheId"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "{{bucket}}/{{anywhere_cache_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	nameVal, ok := opRes["name"].(string)
	if !ok {
		return fmt.Errorf("opRes['name'] is not a string: %v", opRes["name"])
	}

	nameParts := strings.Split(nameVal, "/")
	if len(nameParts) != 6 || nameParts[0] != "projects" || nameParts[2] != "buckets" || nameParts[4] != "anywhereCaches" {
		return fmt.Errorf("error parsing the anywhereCacheId from %s", nameVal)
	}

	anywhereCacheID := nameParts[5]
	if err := d.Set("anywhere_cache_id", anywhereCacheID); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "{{bucket}}/{{anywhere_cache_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating AnywhereCache %q: %#v", d.Id(), res)

	return resourceStorageAnywhereCacheRead(d, meta)
}

func resourceStorageAnywhereCacheRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{StorageBasePath}}b/{{bucket}}/anywhereCaches/{{anywhere_cache_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("StorageAnywhereCache %q", d.Id()))
	}

	if err := d.Set("zone", flattenStorageAnywhereCacheZone(res["zone"], d, config)); err != nil {
		return fmt.Errorf("Error reading AnywhereCache: %s", err)
	}
	if err := d.Set("admission_policy", flattenStorageAnywhereCacheAdmissionPolicy(res["admissionPolicy"], d, config)); err != nil {
		return fmt.Errorf("Error reading AnywhereCache: %s", err)
	}
	if err := d.Set("ttl", flattenStorageAnywhereCacheTtl(res["ttl"], d, config)); err != nil {
		return fmt.Errorf("Error reading AnywhereCache: %s", err)
	}
	if err := d.Set("anywhere_cache_id", flattenStorageAnywhereCacheAnywhereCacheId(res["anywhereCacheId"], d, config)); err != nil {
		return fmt.Errorf("Error reading AnywhereCache: %s", err)
	}
	if err := d.Set("create_time", flattenStorageAnywhereCacheCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading AnywhereCache: %s", err)
	}
	if err := d.Set("update_time", flattenStorageAnywhereCacheUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading AnywhereCache: %s", err)
	}
	if err := d.Set("pending_update", flattenStorageAnywhereCachePendingUpdate(res["pendingUpdate"], d, config)); err != nil {
		return fmt.Errorf("Error reading AnywhereCache: %s", err)
	}
	if err := d.Set("state", flattenStorageAnywhereCacheState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading AnywhereCache: %s", err)
	}

	return nil
}

func resourceStorageAnywhereCacheUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	admissionPolicyProp, err := expandStorageAnywhereCacheAdmissionPolicy(d.Get("admission_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("admission_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, admissionPolicyProp)) {
		obj["admissionPolicy"] = admissionPolicyProp
	}
	ttlProp, err := expandStorageAnywhereCacheTtl(d.Get("ttl"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ttl"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, ttlProp)) {
		obj["ttl"] = ttlProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{StorageBasePath}}b/{{bucket}}/anywhereCaches/{{anywhere_cache_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating AnywhereCache %q: %#v", d.Id(), obj)
	headers := make(http.Header)

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
		Headers:   headers,
	})

	if err != nil {
		return fmt.Errorf("Error updating AnywhereCache %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating AnywhereCache %q: %#v", d.Id(), res)
	}

	err = StorageOperationWaitTime(
		config, res, "Updating AnywhereCache", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceStorageAnywhereCacheRead(d, meta)
}

func resourceStorageAnywhereCacheDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{StorageBasePath}}b/{{bucket}}/anywhereCaches/{{anywhere_cache_id}}/disable")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting AnywhereCache %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "AnywhereCache")
	}

	log.Printf("[DEBUG] Finished deleting AnywhereCache %q: %#v", d.Id(), res)
	return nil
}

func resourceStorageAnywhereCacheImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^b/(?P<bucket>[^/]+)/anywhereCaches/(?P<anywhere_cache_id>[^/]+)$",
		"^(?P<bucket>[^/]+)/(?P<anywhere_cache_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "{{bucket}}/{{anywhere_cache_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenStorageAnywhereCacheZone(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenStorageAnywhereCacheAdmissionPolicy(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenStorageAnywhereCacheTtl(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenStorageAnywhereCacheAnywhereCacheId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenStorageAnywhereCacheCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenStorageAnywhereCacheUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenStorageAnywhereCachePendingUpdate(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenStorageAnywhereCacheState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandStorageAnywhereCacheZone(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandStorageAnywhereCacheAdmissionPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandStorageAnywhereCacheTtl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
