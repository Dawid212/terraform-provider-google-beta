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

package iap_test

import (
	"log"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccIapSettings_iapSettingsBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckIapSettingsDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIapSettings_iapSettingsBasicExample(context),
			},
			{
				ResourceName:            "google_iap_settings.iap_settings",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"access_settings.0.workforce_identity_settings.0.oauth2.0.client_secret", "name"},
			},
		},
	})
}

func testAccIapSettings_iapSettingsBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_project" "project" {
}

resource "google_compute_region_backend_service" "default" {
  name                            = "tf-test-iap-settings-tf%{random_suffix}"
  region                          = "us-central1"
  health_checks                   = [google_compute_health_check.default.id]
  connection_draining_timeout_sec = 10
  session_affinity                = "CLIENT_IP"
}

resource "google_compute_health_check" "default" {
  name               = "tf-test-iap-bs-health-check%{random_suffix}"
  check_interval_sec = 1
  timeout_sec        = 1

  tcp_health_check {
    port = "80"
  }
}

resource "google_iap_settings" "iap_settings" {
  name = "projects/${data.google_project.project.number}/iap_web/compute-us-central1/services/${google_compute_region_backend_service.default.name}"
  access_settings {
    identity_sources = ["WORKFORCE_IDENTITY_FEDERATION"]
    allowed_domains_settings {
      domains = ["test.abc.com"]
      enable  = true
    }
    cors_settings {
      allow_http_options = true
    }
    reauth_settings {
      method = "SECURE_KEY"
      max_age = "305s"
      policy_type = "MINIMUM"
    }
    gcip_settings {
      login_page_uri = "https://test.com/?apiKey=abc"
    }
    oauth_settings {
      login_hint = "test"
    }
    workforce_identity_settings {
      workforce_pools = ["wif-pool"]
      oauth2 {
        client_id = "test-client-id"
        client_secret = "test-client-secret"
      }
    }    
  }
  application_settings {
    cookie_domain = "test.abc.com"
    csm_settings {
      rctoken_aud = "test-aud-set"
    }
    access_denied_page_settings {
      access_denied_page_uri = "test-uri"
      generate_troubleshooting_uri = true
      remediation_token_generation_enabled = false
    }
    attribute_propagation_settings {
      output_credentials = ["HEADER"]
      expression = "attributes.saml_attributes.filter(attribute, attribute.name in [\"test1\", \"test2\"])"
      enable = false
    }
  }
}
`, context)
}

func testAccCheckIapSettingsDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_iap_settings" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			log.Printf("[DEBUG] Ignoring destroy during test")
		}

		return nil
	}
}
