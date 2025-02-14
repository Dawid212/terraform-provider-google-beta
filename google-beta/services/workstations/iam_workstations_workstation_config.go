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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/workstations/WorkstationConfig.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/iam_policy.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package workstations

import (
	"fmt"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgiamresource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

var WorkstationsWorkstationConfigIamSchema = map[string]*schema.Schema{
	"project": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"location": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"workstation_cluster_id": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	"workstation_config_id": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
	},
}

type WorkstationsWorkstationConfigIamUpdater struct {
	project              string
	location             string
	workstationClusterId string
	workstationConfigId  string
	d                    tpgresource.TerraformResourceData
	Config               *transport_tpg.Config
}

func WorkstationsWorkstationConfigIamUpdaterProducer(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgiamresource.ResourceIamUpdater, error) {
	values := make(map[string]string)

	project, _ := tpgresource.GetProject(d, config)
	if project != "" {
		if err := d.Set("project", project); err != nil {
			return nil, fmt.Errorf("Error setting project: %s", err)
		}
	}
	values["project"] = project
	location, _ := tpgresource.GetLocation(d, config)
	if location != "" {
		if err := d.Set("location", location); err != nil {
			return nil, fmt.Errorf("Error setting location: %s", err)
		}
	}
	values["location"] = location
	if v, ok := d.GetOk("workstation_cluster_id"); ok {
		values["workstation_cluster_id"] = v.(string)
	}

	if v, ok := d.GetOk("workstation_config_id"); ok {
		values["workstation_config_id"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := tpgresource.GetImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/workstationClusters/(?P<workstation_cluster_id>[^/]+)/workstationConfigs/(?P<workstation_config_id>[^/]+)", "(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<workstation_cluster_id>[^/]+)/(?P<workstation_config_id>[^/]+)", "(?P<location>[^/]+)/(?P<workstation_cluster_id>[^/]+)/(?P<workstation_config_id>[^/]+)", "(?P<workstation_config_id>[^/]+)"}, d, config, d.Get("workstation_config_id").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &WorkstationsWorkstationConfigIamUpdater{
		project:              values["project"],
		location:             values["location"],
		workstationClusterId: values["workstation_cluster_id"],
		workstationConfigId:  values["workstation_config_id"],
		d:                    d,
		Config:               config,
	}

	if err := d.Set("project", u.project); err != nil {
		return nil, fmt.Errorf("Error setting project: %s", err)
	}
	if err := d.Set("location", u.location); err != nil {
		return nil, fmt.Errorf("Error setting location: %s", err)
	}
	if err := d.Set("workstation_cluster_id", u.workstationClusterId); err != nil {
		return nil, fmt.Errorf("Error setting workstation_cluster_id: %s", err)
	}
	if err := d.Set("workstation_config_id", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting workstation_config_id: %s", err)
	}

	return u, nil
}

func WorkstationsWorkstationConfigIdParseFunc(d *schema.ResourceData, config *transport_tpg.Config) error {
	values := make(map[string]string)

	project, _ := tpgresource.GetProject(d, config)
	if project != "" {
		values["project"] = project
	}

	location, _ := tpgresource.GetLocation(d, config)
	if location != "" {
		values["location"] = location
	}

	m, err := tpgresource.GetImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/workstationClusters/(?P<workstation_cluster_id>[^/]+)/workstationConfigs/(?P<workstation_config_id>[^/]+)", "(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<workstation_cluster_id>[^/]+)/(?P<workstation_config_id>[^/]+)", "(?P<location>[^/]+)/(?P<workstation_cluster_id>[^/]+)/(?P<workstation_config_id>[^/]+)", "(?P<workstation_config_id>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &WorkstationsWorkstationConfigIamUpdater{
		project:              values["project"],
		location:             values["location"],
		workstationClusterId: values["workstation_cluster_id"],
		workstationConfigId:  values["workstation_config_id"],
		d:                    d,
		Config:               config,
	}
	if err := d.Set("workstation_config_id", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting workstation_config_id: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *WorkstationsWorkstationConfigIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyWorkstationConfigUrl("getIamPolicy")
	if err != nil {
		return nil, err
	}

	project, err := tpgresource.GetProject(u.d, u.Config)
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
		Method:    "GET",
		Project:   project,
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

func (u *WorkstationsWorkstationConfigIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := tpgresource.ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyWorkstationConfigUrl("setIamPolicy")
	if err != nil {
		return err
	}
	project, err := tpgresource.GetProject(u.d, u.Config)
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
		Project:   project,
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

func (u *WorkstationsWorkstationConfigIamUpdater) qualifyWorkstationConfigUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{WorkstationsBasePath}}%s:%s", fmt.Sprintf("projects/%s/locations/%s/workstationClusters/%s/workstationConfigs/%s", u.project, u.location, u.workstationClusterId, u.workstationConfigId), methodIdentifier)
	url, err := tpgresource.ReplaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *WorkstationsWorkstationConfigIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/locations/%s/workstationClusters/%s/workstationConfigs/%s", u.project, u.location, u.workstationClusterId, u.workstationConfigId)
}

func (u *WorkstationsWorkstationConfigIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-workstations-workstationconfig-%s", u.GetResourceId())
}

func (u *WorkstationsWorkstationConfigIamUpdater) DescribeResource() string {
	return fmt.Sprintf("workstations workstationconfig %q", u.GetResourceId())
}
