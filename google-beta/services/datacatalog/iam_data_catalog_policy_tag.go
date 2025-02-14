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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/datacatalog/PolicyTag.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/iam_policy.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package datacatalog

import (
	"fmt"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgiamresource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

var DataCatalogPolicyTagIamSchema = map[string]*schema.Schema{
	"policy_tag": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
	},
}

type DataCatalogPolicyTagIamUpdater struct {
	policyTag string
	d         tpgresource.TerraformResourceData
	Config    *transport_tpg.Config
}

func DataCatalogPolicyTagIamUpdaterProducer(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgiamresource.ResourceIamUpdater, error) {
	values := make(map[string]string)

	if v, ok := d.GetOk("policy_tag"); ok {
		values["policy_tag"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := tpgresource.GetImportIdQualifiers([]string{"(?P<policy_tag>.+)"}, d, config, d.Get("policy_tag").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &DataCatalogPolicyTagIamUpdater{
		policyTag: values["policy_tag"],
		d:         d,
		Config:    config,
	}

	if err := d.Set("policy_tag", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting policy_tag: %s", err)
	}

	return u, nil
}

func DataCatalogPolicyTagIdParseFunc(d *schema.ResourceData, config *transport_tpg.Config) error {
	values := make(map[string]string)

	m, err := tpgresource.GetImportIdQualifiers([]string{"(?P<policy_tag>.+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &DataCatalogPolicyTagIamUpdater{
		policyTag: values["policy_tag"],
		d:         d,
		Config:    config,
	}
	if err := d.Set("policy_tag", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting policy_tag: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *DataCatalogPolicyTagIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyPolicyTagUrl("getIamPolicy")
	if err != nil {
		return nil, err
	}

	var obj map[string]interface{}

	userAgent, err := tpgresource.GenerateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return nil, err
	}

	policy, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    u.Config,
		Method:    "POST",
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
	})
	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Error retrieving IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	out := &cloudresourcemanager.Policy{}
	err = tpgresource.Convert(policy, out)
	if err != nil {
		return nil, errwrap.Wrapf("Cannot convert a policy to a resource manager policy: {{err}}", err)
	}

	return out, nil
}

func (u *DataCatalogPolicyTagIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := tpgresource.ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyPolicyTagUrl("setIamPolicy")
	if err != nil {
		return err
	}

	userAgent, err := tpgresource.GenerateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return err
	}

	_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    u.Config,
		Method:    "POST",
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   u.d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *DataCatalogPolicyTagIamUpdater) qualifyPolicyTagUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{DataCatalogBasePath}}%s:%s", fmt.Sprintf("%s", u.policyTag), methodIdentifier)
	url, err := tpgresource.ReplaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *DataCatalogPolicyTagIamUpdater) GetResourceId() string {
	return fmt.Sprintf("%s", u.policyTag)
}

func (u *DataCatalogPolicyTagIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-datacatalog-policytag-%s", u.GetResourceId())
}

func (u *DataCatalogPolicyTagIamUpdater) DescribeResource() string {
	return fmt.Sprintf("datacatalog policytag %q", u.GetResourceId())
}
