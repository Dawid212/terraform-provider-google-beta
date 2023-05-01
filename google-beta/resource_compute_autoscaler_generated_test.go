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
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccComputeAutoscaler_autoscalerSingleInstanceExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"provider_name":  "google-beta.us-central1",
		"provider_alias": "alias  = \"us-central1\"",
		"random_suffix":  RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeAutoscalerDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeAutoscaler_autoscalerSingleInstanceExample(context),
			},
			{
				ResourceName:            "google_compute_autoscaler.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"target", "zone"},
			},
		},
	})
}

func testAccComputeAutoscaler_autoscalerSingleInstanceExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_autoscaler" "default" {
  provider = %{provider_name}

  name   = "tf-test-my-autoscaler%{random_suffix}"
  zone   = "us-central1-f"
  target = google_compute_instance_group_manager.default.id

  autoscaling_policy {
    max_replicas    = 5
    min_replicas    = 1
    cooldown_period = 60

    metric {
      name                       = "pubsub.googleapis.com/subscription/num_undelivered_messages"
      filter                     = "resource.type = pubsub_subscription AND resource.label.subscription_id = our-subscription"
      single_instance_assignment = 65535
    }
  }
}

resource "google_compute_instance_template" "default" {
  provider = %{provider_name}

  name           = "tf-test-my-instance-template%{random_suffix}"
  machine_type   = "e2-medium"
  can_ip_forward = false

  tags = ["foo", "bar"]

  disk {
    source_image = data.google_compute_image.debian_9.id
  }

  network_interface {
    network = "default"
  }

  metadata = {
    foo = "bar"
  }

  service_account {
    scopes = ["userinfo-email", "compute-ro", "storage-ro"]
  }
}

resource "google_compute_target_pool" "default" {
  provider = %{provider_name}

  name = "tf-test-my-target-pool%{random_suffix}"
}

resource "google_compute_instance_group_manager" "default" {
  provider = %{provider_name}

  name = "tf-test-my-igm%{random_suffix}"
  zone = "us-central1-f"

  version {
    instance_template = google_compute_instance_template.default.id
    name              = "primary"
  }

  target_pools       = [google_compute_target_pool.default.id]
  base_instance_name = "autoscaler-sample"
}

data "google_compute_image" "debian_9" {
  provider = %{provider_name}

  family  = "debian-11"
  project = "debian-cloud"
}

provider "google-beta" {
  region = "us-central1"
  zone   = "us-central1-a"
  %{provider_alias}
}
`, context)
}

func TestAccComputeAutoscaler_autoscalerBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeAutoscalerDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeAutoscaler_autoscalerBasicExample(context),
			},
			{
				ResourceName:            "google_compute_autoscaler.foobar",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"target", "zone"},
			},
		},
	})
}

func testAccComputeAutoscaler_autoscalerBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_autoscaler" "foobar" {
  name   = "tf-test-my-autoscaler%{random_suffix}"
  zone   = "us-central1-f"
  target = google_compute_instance_group_manager.foobar.id

  autoscaling_policy {
    max_replicas    = 5
    min_replicas    = 1
    cooldown_period = 60

    cpu_utilization {
      target = 0.5
    }
  }
}

resource "google_compute_instance_template" "foobar" {
  name           = "tf-test-my-instance-template%{random_suffix}"
  machine_type   = "e2-medium"
  can_ip_forward = false

  tags = ["foo", "bar"]

  disk {
    source_image = data.google_compute_image.debian_9.id
  }

  network_interface {
    network = "default"
  }

  metadata = {
    foo = "bar"
  }

  service_account {
    scopes = ["userinfo-email", "compute-ro", "storage-ro"]
  }
}

resource "google_compute_target_pool" "foobar" {
  name = "tf-test-my-target-pool%{random_suffix}"
}

resource "google_compute_instance_group_manager" "foobar" {
  name = "tf-test-my-igm%{random_suffix}"
  zone = "us-central1-f"

  version {
    instance_template  = google_compute_instance_template.foobar.id
    name               = "primary"
  }

  target_pools       = [google_compute_target_pool.foobar.id]
  base_instance_name = "foobar"
}

data "google_compute_image" "debian_9" {
  family  = "debian-11"
  project = "debian-cloud"
}
`, context)
}

func testAccCheckComputeAutoscalerDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_autoscaler" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/autoscalers/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("ComputeAutoscaler still exists at %s", url)
			}
		}

		return nil
	}
}
