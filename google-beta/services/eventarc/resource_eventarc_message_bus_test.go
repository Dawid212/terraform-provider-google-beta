// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package eventarc_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// We make sure not to run tests in parallel, since only one MessageBus per project is supported.
func TestAccEventarcMessageBus(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		"basic":     testAccEventarcMessageBus_basic,
		"cryptoKey": testAccEventarcMessageBus_cryptoKey,
		"update":    testAccEventarcMessageBus_update,
	}

	for name, tc := range testCases {
		// shadow the tc variable into scope so that when
		// the loop continues, if t.Run hasn't executed tc(t)
		// yet, we don't have a race condition
		// see https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		tc := tc
		t.Run(name, func(t *testing.T) {
			tc(t)
		})
	}
}

func testAccEventarcMessageBus_basic(t *testing.T) {
	context := map[string]interface{}{
		"region":        envvar.GetTestRegionFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckEventarcMessageBusDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccEventarcMessageBus_basicCfg(context),
			},
			{
				ResourceName:            "google_eventarc_message_bus.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "terraform_labels", "annotations"},
			},
		},
	})
}

func testAccEventarcMessageBus_basicCfg(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_eventarc_message_bus" "primary" {
  location       = "%{region}"
  message_bus_id = "tf-test-messagebus%{random_suffix}"
  display_name   = "basic bus"
  labels = {
    test_label = "test-eventarc-label"
  }
  annotations = {
    test_annotation = "test-eventarc-annotation"
  }
}
`, context)
}

func testAccEventarcMessageBus_cryptoKey(t *testing.T) {
	region := envvar.GetTestRegionFromEnv()

	context := map[string]interface{}{
		"project_number": envvar.GetTestProjectNumberFromEnv(),
		"region":         region,
		"key":            acctest.BootstrapKMSKeyWithPurposeInLocationAndName(t, "ENCRYPT_DECRYPT", region, "tf-bootstrap-eventarc-messagebus-key").CryptoKey.Name,
		"random_suffix":  acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckEventarcMessageBusDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccEventarcMessageBus_cryptoKeyCfg(context),
			},
			{
				ResourceName:            "google_eventarc_message_bus.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "terraform_labels", "annotations"},
			},
		},
	})
}

func testAccEventarcMessageBus_cryptoKeyCfg(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_kms_crypto_key_iam_member" "key_member" {
  crypto_key_id = "%{key}"
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member        = "serviceAccount:service-%{project_number}@gcp-sa-eventarc.iam.gserviceaccount.com"
}

resource "google_eventarc_message_bus" "primary" {
  location        = "%{region}"
  message_bus_id  = "tf-test-messagebus%{random_suffix}"
  crypto_key_name = "%{key}"
  logging_config {
    log_severity = "ALERT"
  }
  depends_on = [google_kms_crypto_key_iam_member.key_member]
}
`, context)
}

func testAccEventarcMessageBus_update(t *testing.T) {
	region := envvar.GetTestRegionFromEnv()

	context := map[string]interface{}{
		"project_number": envvar.GetTestProjectNumberFromEnv(),
		"region":         region,
		"key1":           acctest.BootstrapKMSKeyWithPurposeInLocationAndName(t, "ENCRYPT_DECRYPT", region, "tf-bootstrap-eventarc-messagebus-key1").CryptoKey.Name,
		"key2":           acctest.BootstrapKMSKeyWithPurposeInLocationAndName(t, "ENCRYPT_DECRYPT", region, "tf-bootstrap-eventarc-messagebus-key2").CryptoKey.Name,
		"random_suffix":  acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckEventarcMessageBusDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccEventarcMessageBus_setCfg(context),
			},
			{
				ResourceName:            "google_eventarc_message_bus.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "terraform_labels", "annotations"},
			},
			{
				Config: testAccEventarcMessageBus_updateCfg(context),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("google_eventarc_message_bus.primary", plancheck.ResourceActionUpdate),
					},
				},
			},
			{
				ResourceName:            "google_eventarc_message_bus.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "terraform_labels", "annotations"},
			},
			{
				Config: testAccEventarcMessageBus_deleteCfg(context),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("google_eventarc_message_bus.primary", plancheck.ResourceActionUpdate),
					},
				},
			},
			{
				ResourceName:            "google_eventarc_message_bus.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "terraform_labels", "annotations"},
			},
		},
	})
}

func testAccEventarcMessageBus_setCfg(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_kms_crypto_key_iam_member" "key1_member" {
  crypto_key_id = "%{key1}"
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member        = "serviceAccount:service-%{project_number}@gcp-sa-eventarc.iam.gserviceaccount.com"
}

resource "google_eventarc_message_bus" "primary" {
  location        = "%{region}"
  message_bus_id  = "tf-test-messagebus%{random_suffix}"
  crypto_key_name = "%{key1}"
  display_name    = "message bus"
  logging_config {
    log_severity = "ALERT"
  }
  depends_on = [google_kms_crypto_key_iam_member.key1_member]
}
`, context)
}

func testAccEventarcMessageBus_updateCfg(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_kms_crypto_key_iam_member" "key1_member" {
  crypto_key_id = "%{key1}"
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member        = "serviceAccount:service-%{project_number}@gcp-sa-eventarc.iam.gserviceaccount.com"
}

resource "google_kms_crypto_key_iam_member" "key2_member" {
  crypto_key_id = "%{key2}"
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member        = "serviceAccount:service-%{project_number}@gcp-sa-eventarc.iam.gserviceaccount.com"
}

resource "google_eventarc_message_bus" "primary" {
  location        = "%{region}"
  message_bus_id  = "tf-test-messagebus%{random_suffix}"
  crypto_key_name = "%{key2}"
  display_name    = "updated message bus"
  logging_config {
    log_severity = "DEBUG"
  }
  depends_on = [google_kms_crypto_key_iam_member.key1_member, google_kms_crypto_key_iam_member.key2_member]
}
`, context)
}

func testAccEventarcMessageBus_deleteCfg(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_kms_crypto_key_iam_member" "key1_member" {
  crypto_key_id = "%{key1}"
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member        = "serviceAccount:service-%{project_number}@gcp-sa-eventarc.iam.gserviceaccount.com"
}

resource "google_kms_crypto_key_iam_member" "key2_member" {
  crypto_key_id = "%{key2}"
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member        = "serviceAccount:service-%{project_number}@gcp-sa-eventarc.iam.gserviceaccount.com"
}

resource "google_eventarc_message_bus" "primary" {
  location        = "%{region}"
  message_bus_id  = "tf-test-messagebus%{random_suffix}"
  crypto_key_name = ""
  display_name    = "updated message bus"
  logging_config {
    log_severity = "DEBUG"
  }
}
`, context)
}

func testAccCheckEventarcMessageBusDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_eventarc_message_bus" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{EventarcBasePath}}projects/{{project}}/locations/{{location}}/messageBuses/{{name}}")
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
				return fmt.Errorf("EventarcMessageBus still exists at %s", url)
			}
		}

		return nil
	}
}
