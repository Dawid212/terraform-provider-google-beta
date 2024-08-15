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

package compute_test

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

func TestAccComputeSubnetwork_subnetworkBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeSubnetworkDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeSubnetwork_subnetworkBasicExample(context),
			},
			{
				ResourceName:            "google_compute_subnetwork.network-with-private-secondary-ip-ranges",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"network", "region", "reserved_internal_range"},
			},
		},
	})
}

func testAccComputeSubnetwork_subnetworkBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_subnetwork" "network-with-private-secondary-ip-ranges" {
  name          = "tf-test-test-subnetwork%{random_suffix}"
  ip_cidr_range = "10.2.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.custom-test.id
  secondary_ip_range {
    range_name    = "tf-test-secondary-range-update1"
    ip_cidr_range = "192.168.10.0/24"
  }
}

resource "google_compute_network" "custom-test" {
  name                    = "tf-test-test-network%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func TestAccComputeSubnetwork_subnetworkLoggingConfigExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeSubnetworkDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeSubnetwork_subnetworkLoggingConfigExample(context),
			},
			{
				ResourceName:            "google_compute_subnetwork.subnet-with-logging",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"network", "region", "reserved_internal_range"},
			},
		},
	})
}

func testAccComputeSubnetwork_subnetworkLoggingConfigExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_subnetwork" "subnet-with-logging" {
  name          = "tf-test-log-test-subnetwork%{random_suffix}"
  ip_cidr_range = "10.2.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.custom-test.id

  log_config {
    aggregation_interval = "INTERVAL_10_MIN"
    flow_sampling        = 0.5
    metadata             = "INCLUDE_ALL_METADATA"
  }
}

resource "google_compute_network" "custom-test" {
  name                    = "tf-test-log-test-network%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func TestAccComputeSubnetwork_subnetworkInternalL7lbExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeSubnetworkDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeSubnetwork_subnetworkInternalL7lbExample(context),
			},
			{
				ResourceName:            "google_compute_subnetwork.network-for-l7lb",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"network", "region", "reserved_internal_range"},
			},
		},
	})
}

func testAccComputeSubnetwork_subnetworkInternalL7lbExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_subnetwork" "network-for-l7lb" {
  provider = google-beta

  name          = "tf-test-l7lb-test-subnetwork%{random_suffix}"
  ip_cidr_range = "10.0.0.0/22"
  region        = "us-central1"
  purpose       = "REGIONAL_MANAGED_PROXY"
  role          = "ACTIVE"
  network       = google_compute_network.custom-test.id
}

resource "google_compute_network" "custom-test" {
  provider = google-beta

  name                    = "tf-test-l7lb-test-network%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func TestAccComputeSubnetwork_subnetworkIpv6Example(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeSubnetworkDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeSubnetwork_subnetworkIpv6Example(context),
			},
			{
				ResourceName:            "google_compute_subnetwork.subnetwork-ipv6",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"network", "region", "reserved_internal_range"},
			},
		},
	})
}

func testAccComputeSubnetwork_subnetworkIpv6Example(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_subnetwork" "subnetwork-ipv6" {
  name          = "tf-test-ipv6-test-subnetwork%{random_suffix}"
  
  ip_cidr_range = "10.0.0.0/22"
  region        = "us-west2"
  
  stack_type       = "IPV4_IPV6"
  ipv6_access_type = "EXTERNAL"

  network       = google_compute_network.custom-test.id
}

resource "google_compute_network" "custom-test" {
  name                    = "tf-test-ipv6-test-network%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func TestAccComputeSubnetwork_subnetworkInternalIpv6Example(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeSubnetworkDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeSubnetwork_subnetworkInternalIpv6Example(context),
			},
			{
				ResourceName:            "google_compute_subnetwork.subnetwork-internal-ipv6",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"network", "region", "reserved_internal_range"},
			},
		},
	})
}

func testAccComputeSubnetwork_subnetworkInternalIpv6Example(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_subnetwork" "subnetwork-internal-ipv6" {
  name          = "tf-test-internal-ipv6-test-subnetwork%{random_suffix}"
  
  ip_cidr_range = "10.0.0.0/22"
  region        = "us-west2"
  
  stack_type       = "IPV4_IPV6"
  ipv6_access_type = "INTERNAL"

  network       = google_compute_network.custom-test.id
}

resource "google_compute_network" "custom-test" {
  name                    = "tf-test-internal-ipv6-test-network%{random_suffix}"
  auto_create_subnetworks = false
  enable_ula_internal_ipv6 = true
}
`, context)
}

func TestAccComputeSubnetwork_subnetworkPurposePrivateNatExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeSubnetworkDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeSubnetwork_subnetworkPurposePrivateNatExample(context),
			},
			{
				ResourceName:            "google_compute_subnetwork.subnetwork-purpose-private-nat",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"network", "region", "reserved_internal_range"},
			},
		},
	})
}

func testAccComputeSubnetwork_subnetworkPurposePrivateNatExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_subnetwork" "subnetwork-purpose-private-nat" {
  provider         = google-beta

  name             = "tf-test-subnet-purpose-test-subnetwork%{random_suffix}"
  region           = "us-west2"
  ip_cidr_range    = "192.168.1.0/24"
  purpose          = "PRIVATE_NAT"
  network          = google_compute_network.custom-test.id
}

resource "google_compute_network" "custom-test" {
  provider                = google-beta

  name                    = "tf-test-subnet-purpose-test-network%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func TestAccComputeSubnetwork_subnetworkCidrOverlapExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeSubnetworkDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeSubnetwork_subnetworkCidrOverlapExample(context),
			},
			{
				ResourceName:            "google_compute_subnetwork.subnetwork-cidr-overlap",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"network", "region", "reserved_internal_range"},
			},
		},
	})
}

func testAccComputeSubnetwork_subnetworkCidrOverlapExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_subnetwork" "subnetwork-cidr-overlap" {
  provider = google-beta

  name                             = "tf-test-subnet-cidr-overlap%{random_suffix}"
  region                           = "us-west2"
  ip_cidr_range                    = "192.168.1.0/24"
  allow_subnet_cidr_routes_overlap = true
  network                          = google_compute_network.net-cidr-overlap.id
}

resource "google_compute_network" "net-cidr-overlap" {
  provider                = google-beta

  name                    = "tf-test-net-cidr-overlap%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func TestAccComputeSubnetwork_subnetworkReservedInternalRangeExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeSubnetworkDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeSubnetwork_subnetworkReservedInternalRangeExample(context),
			},
			{
				ResourceName:            "google_compute_subnetwork.subnetwork-reserved-internal-range",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"network", "region", "reserved_internal_range"},
			},
		},
	})
}

func testAccComputeSubnetwork_subnetworkReservedInternalRangeExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_subnetwork" "subnetwork-reserved-internal-range" {
  provider                = google-beta
  name                    = "tf-test-subnetwork-reserved-internal-range%{random_suffix}"
  region                  = "us-central1"
  network                 = google_compute_network.default.id
  reserved_internal_range = "networkconnectivity.googleapis.com/${google_network_connectivity_internal_range.reserved.id}"
}

resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "tf-test-network-reserved-internal-range%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_network_connectivity_internal_range" "reserved" {
  provider          = google-beta
  name              = "reserved"
  network           = google_compute_network.default.id
  usage             = "FOR_VPC"
  peering           = "FOR_SELF"
  prefix_length     = 24
  target_cidr_range = [
    "10.0.0.0/8"
  ]
}
`, context)
}

func TestAccComputeSubnetwork_subnetworkReservedSecondaryRangeExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeSubnetworkDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeSubnetwork_subnetworkReservedSecondaryRangeExample(context),
			},
			{
				ResourceName:            "google_compute_subnetwork.subnetwork-reserved-secondary-range",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"network", "region", "reserved_internal_range"},
			},
		},
	})
}

func testAccComputeSubnetwork_subnetworkReservedSecondaryRangeExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_subnetwork" "subnetwork-reserved-secondary-range" {
  provider                = google-beta
  name                    = "tf-test-subnetwork-reserved-secondary-range%{random_suffix}"
  region                  = "us-central1"
  network                 = google_compute_network.default.id
  reserved_internal_range = "networkconnectivity.googleapis.com/${google_network_connectivity_internal_range.reserved.id}"

  secondary_ip_range {
    range_name              = "secondary"
    reserved_internal_range = "networkconnectivity.googleapis.com/${google_network_connectivity_internal_range.reserved_secondary.id}"
  }
}

resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "tf-test-network-reserved-secondary-range%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_network_connectivity_internal_range" "reserved" {
  provider          = google-beta
  name              = "reserved"
  network           = google_compute_network.default.id
  usage             = "FOR_VPC"
  peering           = "FOR_SELF"
  prefix_length     = 24
  target_cidr_range = [
    "10.0.0.0/8"
  ]
}

resource "google_network_connectivity_internal_range" "reserved_secondary" {
  provider          = google-beta
  name              = "reserved-secondary"
  network           = google_compute_network.default.id
  usage             = "FOR_VPC"
  peering           = "FOR_SELF"
  prefix_length     = 16
  target_cidr_range = [
    "10.0.0.0/8"
  ]
}
`, context)
}

func testAccCheckComputeSubnetworkDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_subnetwork" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/subnetworks/{{name}}")
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
				return fmt.Errorf("ComputeSubnetwork still exists at %s", url)
			}
		}

		return nil
	}
}
