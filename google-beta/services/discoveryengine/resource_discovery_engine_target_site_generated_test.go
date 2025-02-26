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

package discoveryengine_test

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

func TestAccDiscoveryEngineTargetSite_discoveryengineTargetsiteBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDiscoveryEngineTargetSiteDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDiscoveryEngineTargetSite_discoveryengineTargetsiteBasicExample(context),
			},
			{
				ResourceName:            "google_discovery_engine_target_site.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"data_store_id", "location", "project", "provided_uri_pattern", "target_site_id"},
			},
		},
	})
}

func testAccDiscoveryEngineTargetSite_discoveryengineTargetsiteBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_discovery_engine_target_site" "basic" {
  location                    = google_discovery_engine_data_store.basic.location
  data_store_id               = google_discovery_engine_data_store.basic.data_store_id
  provided_uri_pattern        = "cloud.google.com/docs/*"
  type                        = "INCLUDE"
  exact_match                 = false
}

resource "google_discovery_engine_data_store" "basic" {
  location                     = "global"
  data_store_id                = "tf-test-data-store-id%{random_suffix}"
  display_name                 = "tf-test-basic-site-search-datastore"
  industry_vertical            = "GENERIC"
  content_config               = "PUBLIC_WEBSITE"
  solution_types               = ["SOLUTION_TYPE_SEARCH"]
  create_advanced_site_search  = false
  skip_default_schema_creation = false
}
`, context)
}

func TestAccDiscoveryEngineTargetSite_discoveryengineTargetsiteAdvancedExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDiscoveryEngineTargetSiteDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDiscoveryEngineTargetSite_discoveryengineTargetsiteAdvancedExample(context),
			},
			{
				ResourceName:            "google_discovery_engine_target_site.advanced",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"data_store_id", "location", "project", "provided_uri_pattern", "target_site_id"},
			},
		},
	})
}

func testAccDiscoveryEngineTargetSite_discoveryengineTargetsiteAdvancedExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_discovery_engine_target_site" "advanced" {
  location                    = google_discovery_engine_data_store.advanced.location
  data_store_id               = google_discovery_engine_data_store.advanced.data_store_id
  provided_uri_pattern        = "cloud.google.com/docs/*"
  type                        = "INCLUDE"
  exact_match                 = false
}

resource "google_discovery_engine_data_store" "advanced" {
  location                     = "global"
  data_store_id                = "tf-test-data-store-id%{random_suffix}"
  display_name                 = "tf-test-advanced-site-search-datastore"
  industry_vertical            = "GENERIC"
  content_config               = "PUBLIC_WEBSITE"
  solution_types               = ["SOLUTION_TYPE_SEARCH"]
  create_advanced_site_search  = true
  skip_default_schema_creation = false
}
`, context)
}

func testAccCheckDiscoveryEngineTargetSiteDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_discovery_engine_target_site" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{DiscoveryEngineBasePath}}{{name}}")
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
				return fmt.Errorf("DiscoveryEngineTargetSite still exists at %s", url)
			}
		}

		return nil
	}
}
