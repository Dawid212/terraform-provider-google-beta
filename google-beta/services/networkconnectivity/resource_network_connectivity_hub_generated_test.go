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

package networkconnectivity_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccNetworkConnectivityHub_networkConnectivityHubBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkConnectivityHubDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkConnectivityHub_networkConnectivityHubBasicExample(context),
			},
			{
				ResourceName:            "google_network_connectivity_hub.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "terraform_labels"},
			},
		},
	})
}

func testAccNetworkConnectivityHub_networkConnectivityHubBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_network_connectivity_hub" "primary"  {
 name        = "basic%{random_suffix}"
 description = "A sample hub"
 labels = {
    label-one = "value-one"
  }
}
`, context)
}

func TestAccNetworkConnectivityHub_networkConnectivityHubWithExportPscExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkConnectivityHubDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkConnectivityHub_networkConnectivityHubWithExportPscExample(context),
			},
			{
				ResourceName:            "google_network_connectivity_hub.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "terraform_labels"},
			},
		},
	})
}

func testAccNetworkConnectivityHub_networkConnectivityHubWithExportPscExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_network_connectivity_hub" "primary"  {
 name        = "basic%{random_suffix}"
 description = "A sample hub with Private Service Connect transitivity is enabled"
 export_psc = true
}
`, context)
}

func TestAccNetworkConnectivityHub_networkConnectivityHubMeshTopologyExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkConnectivityHubDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkConnectivityHub_networkConnectivityHubMeshTopologyExample(context),
			},
			{
				ResourceName:            "google_network_connectivity_hub.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "terraform_labels"},
			},
		},
	})
}

func testAccNetworkConnectivityHub_networkConnectivityHubMeshTopologyExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_network_connectivity_hub" "primary"  {
 name        = "mesh%{random_suffix}"
 description = "A sample mesh hub"
 labels = {
    label-one = "value-one"
  }
}
`, context)
}

func TestAccNetworkConnectivityHub_networkConnectivityHubStarTopologyExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkConnectivityHubDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkConnectivityHub_networkConnectivityHubStarTopologyExample(context),
			},
			{
				ResourceName:            "google_network_connectivity_hub.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "terraform_labels"},
			},
		},
	})
}

func testAccNetworkConnectivityHub_networkConnectivityHubStarTopologyExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_network_connectivity_hub" "primary"  {
 name        = "star%{random_suffix}"
 description = "A sample star hub"
 labels = {
    label-one = "value-one"
  }
 preset_topology = "STAR"
  
}
`, context)
}

func TestAccNetworkConnectivityHub_networkConnectivityHubPolicyModeExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkConnectivityHubDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkConnectivityHub_networkConnectivityHubPolicyModeExample(context),
			},
			{
				ResourceName:            "google_network_connectivity_hub.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "terraform_labels"},
			},
		},
	})
}

func testAccNetworkConnectivityHub_networkConnectivityHubPolicyModeExample(context map[string]interface{}) string {
	return acctest.Nprintf(`

resource "google_network_connectivity_hub" "primary" {
 name            = "policy%{random_suffix}"
 description     = "A sample hub with PRESET policy_mode and STAR topology"
 policy_mode     = "PRESET"
 preset_topology = "STAR"
 labels = {
    label-one = "value-one"
  }
}
`, context)
}

func testAccCheckNetworkConnectivityHubDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_network_connectivity_hub" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{NetworkConnectivityBasePath}}projects/{{project}}/locations/global/hubs/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("NetworkConnectivityHub still exists at %s", url)
			}
		}

		return nil
	}
}
