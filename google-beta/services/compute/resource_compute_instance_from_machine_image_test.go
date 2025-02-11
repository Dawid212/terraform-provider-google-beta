// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package compute_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"

	compute "google.golang.org/api/compute/v0.beta"
	"google.golang.org/api/googleapi"
)

func TestAccComputeInstanceFromMachineImage_basic(t *testing.T) {
	t.Parallel()

	var instance compute.Instance
	instanceName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	generatedInstanceName := fmt.Sprintf("tf-test-generated-%s", acctest.RandString(t, 10))
	resourceName := "google_compute_instance_from_machine_image.foobar"

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeInstanceFromMachineImageDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeInstanceFromMachineImage_basic(instanceName, generatedInstanceName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeInstanceExists(t, resourceName, &instance),

					// Check that fields were set based on the template
					resource.TestCheckResourceAttr(resourceName, "machine_type", "n1-standard-1"),
					resource.TestCheckResourceAttr(resourceName, "attached_disk.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "scheduling.0.automatic_restart", "false"),
				),
			},
		},
	})
}

func TestAccComputeInstanceFromMachineImage_maxRunDuration(t *testing.T) {
	t.Parallel()

	var instance compute.Instance
	instanceName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	generatedInstanceName := fmt.Sprintf("tf-test-generated-%s", acctest.RandString(t, 10))
	resourceName := "google_compute_instance_from_machine_image.foobar"
	var expectedMaxRunDuration = compute.Duration{}
	// Define in testAccComputeInstanceFromMachineImage_maxRunDuration
	expectedMaxRunDuration.Nanos = 123
	expectedMaxRunDuration.Seconds = 60

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeInstanceFromMachineImageDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeInstanceFromMachineImage_maxRunDuration(instanceName, generatedInstanceName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeInstanceExists(t, resourceName, &instance),

					// Check that fields were set based on the template
					resource.TestCheckResourceAttr(resourceName, "machine_type", "n1-standard-1"),
					resource.TestCheckResourceAttr(resourceName, "attached_disk.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "scheduling.0.automatic_restart", "false"),
					testAccCheckComputeInstanceMaxRunDuration(&instance, expectedMaxRunDuration),
				),
			},
		},
	})
}

func TestAccComputeInstanceFromMachineImage_localSsdRecoveryTimeout(t *testing.T) {
	t.Parallel()

	var instance compute.Instance
	instanceName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	generatedInstanceName := fmt.Sprintf("tf-test-generated-%s", acctest.RandString(t, 10))
	resourceName := "google_compute_instance_from_machine_image.foobar"
	var expectedLocalSsdRecoveryTimeout = compute.Duration{}
	expectedLocalSsdRecoveryTimeout.Nanos = 0
	expectedLocalSsdRecoveryTimeout.Seconds = 3600

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeInstanceFromMachineImageDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeInstanceFromMachineImage_localSsdRecoveryTimeout(instanceName, generatedInstanceName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeInstanceExists(t, resourceName, &instance),
					testAccCheckComputeInstanceLocalSsdRecoveryTimeout(&instance, expectedLocalSsdRecoveryTimeout),
				),
			},
		},
	})
}

func TestAccComputeInstanceFromMachineImageWithOverride_localSsdRecoveryTimeout(t *testing.T) {
	t.Parallel()

	var instance compute.Instance
	instanceName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	generatedInstanceName := fmt.Sprintf("tf-test-generated-%s", acctest.RandString(t, 10))
	resourceName := "google_compute_instance_from_machine_image.foobar"
	var expectedLocalSsdRecoveryTimeout = compute.Duration{}
	expectedLocalSsdRecoveryTimeout.Nanos = 0
	expectedLocalSsdRecoveryTimeout.Seconds = 7200

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeInstanceFromMachineImageDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeInstanceFromMachineImageWithOverride_localSsdRecoveryTimeout(instanceName, generatedInstanceName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeInstanceExists(t, resourceName, &instance),
					testAccCheckComputeInstanceLocalSsdRecoveryTimeout(&instance, expectedLocalSsdRecoveryTimeout),
				),
			},
		},
	})
}

func TestAccComputeInstanceFromMachineImageWithOverride_partnerMetadata(t *testing.T) {
	t.Parallel()

	var instance compute.Instance
	instanceName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	generatedInstanceName := fmt.Sprintf("tf-test-generated-%s", acctest.RandString(t, 10))
	resourceName := "google_compute_instance_from_machine_image.foobar"
	var namespace = "test.compute.googleapis.com"
	expectedPartnerMetadata := make(map[string]compute.StructuredEntries)
	expectedPartnerMetadata[namespace] = compute.StructuredEntries{
		Entries: googleapi.RawMessage(`{"key1": "value1", "key2": 2,"key3": {"key31":"value31"}}`),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeInstanceFromMachineImageDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeInstanceFromMachineImageWithOverride_partnerMetadata(instanceName, generatedInstanceName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeInstanceExists(t, resourceName, &instance),
					testAccCheckComputeInstancePartnerMetadata(&instance, expectedPartnerMetadata),
				),
			},
		},
	})
}

func TestAccComputeInstanceFromMachineImage_overrideMetadataDotStartupScript(t *testing.T) {
	t.Parallel()

	var instance compute.Instance
	instanceName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	generatedInstanceName := fmt.Sprintf("tf-test-generated-%s", acctest.RandString(t, 10))
	resourceName := "google_compute_instance_from_machine_image.foobar"

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeInstanceFromMachineImageDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeInstanceFromMachineImage_overrideMetadataDotStartupScript(instanceName, generatedInstanceName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeInstanceExists(t, resourceName, &instance),
					resource.TestCheckResourceAttr(resourceName, "metadata.startup-script", ""),
				),
			},
		},
	})

}

func TestAccComputeInstanceFromMachineImage_diffProject(t *testing.T) {
	t.Parallel()

	var instance compute.Instance
	instanceName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	generatedInstanceName := fmt.Sprintf("tf-test-generated-%s", acctest.RandString(t, 10))
	resourceName := "google_compute_instance_from_machine_image.foobar"
	org := envvar.GetTestOrgFromEnv(t)
	billingId := envvar.GetTestBillingAccountFromEnv(t)
	projectID := fmt.Sprintf("tf-test-%d", acctest.RandInt(t))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeInstanceFromMachineImageDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeInstanceFromMachineImage_diffProject(projectID, org, billingId, instanceName, generatedInstanceName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeInstanceExists(t, resourceName, &instance),

					// Check that fields were set based on the template
					resource.TestCheckResourceAttr(resourceName, "machine_type", "n1-standard-1"),
					resource.TestCheckResourceAttr(resourceName, "attached_disk.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "scheduling.0.automatic_restart", "false"),
				),
			},
		},
	})
}

func TestAccComputeInstanceFromMachineImage_confidentialInstanceConfigMain(t *testing.T) {
	t.Parallel()

	var instance compute.Instance

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeInstanceFromMachineImageDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeInstanceFromMachineImage_ConfidentialInstanceConfigEnable(fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10)), fmt.Sprintf("tf-test-generated-%s", acctest.RandString(t, 10))),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeInstanceExists(t, "google_compute_instance_from_machine_image.foobar1", &instance),
					resource.TestCheckResourceAttr("google_compute_instance_from_machine_image.foobar1", "machine_type", "n2d-standard-2"),
					resource.TestCheckResourceAttr("google_compute_instance_from_machine_image.foobar1", "scheduling.0.on_host_maintenance", "TERMINATE"),
					testAccCheckComputeInstanceHasConfidentialInstanceConfig(&instance, true, ""),
				),
			},
			{
				Config: testAccComputeInstanceFromMachineImage_ConfidentialInstanceConfigEnableSev(fmt.Sprintf("tf-test-sev0-%s", acctest.RandString(t, 10)), fmt.Sprintf("tf-test-sev0-generated-%s", acctest.RandString(t, 10)), "SEV"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeInstanceExists(t, "google_compute_instance_from_machine_image.foobar2", &instance),
					resource.TestCheckResourceAttr("google_compute_instance_from_machine_image.foobar2", "machine_type", "n2d-standard-2"),
					resource.TestCheckResourceAttr("google_compute_instance_from_machine_image.foobar2", "scheduling.0.on_host_maintenance", "TERMINATE"),
					testAccCheckComputeInstanceHasConfidentialInstanceConfig(&instance, true, "SEV"),
				),
			},
			{
				Config: testAccComputeInstanceFromMachineImage_ConfidentialInstanceConfigNoEnableSev(fmt.Sprintf("tf-test-sev1-%s", acctest.RandString(t, 10)), fmt.Sprintf("tf-test-sev1-generated-%s", acctest.RandString(t, 10)), "SEV"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeInstanceExists(t, "google_compute_instance_from_machine_image.foobar3", &instance),
					resource.TestCheckResourceAttr("google_compute_instance_from_machine_image.foobar3", "min_cpu_platform", "AMD Milan"),
					resource.TestCheckResourceAttr("google_compute_instance_from_machine_image.foobar3", "scheduling.0.on_host_maintenance", "TERMINATE"),
					testAccCheckComputeInstanceHasConfidentialInstanceConfig(&instance, false, "SEV"),
				),
			},
			{
				Config: testAccComputeInstanceFromMachineImage_ConfidentialInstanceConfigSevSnp(fmt.Sprintf("tf-test-sev-snp-%s", acctest.RandString(t, 10)), fmt.Sprintf("tf-test-sev-snp-generated-%s", acctest.RandString(t, 10)), "SEV_SNP"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeInstanceExists(t, "google_compute_instance_from_machine_image.foobar4", &instance),
					resource.TestCheckResourceAttr("google_compute_instance_from_machine_image.foobar4", "min_cpu_platform", "AMD Milan"),
					resource.TestCheckResourceAttr("google_compute_instance_from_machine_image.foobar4", "scheduling.0.on_host_maintenance", "TERMINATE"),
					testAccCheckComputeInstanceHasConfidentialInstanceConfig(&instance, false, "SEV_SNP"),
				),
			},
			{
				Config: testAccComputeInstanceFromMachineImage_ConfidentialInstanceConfigTdx(fmt.Sprintf("tf-test-tdx-%s", acctest.RandString(t, 10)), fmt.Sprintf("tf-test-tdx-generated-%s", acctest.RandString(t, 10)), "TDX"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeInstanceExists(t, "google_compute_instance_from_machine_image.foobar5", &instance),
					// Check that fields were set based on the template
					resource.TestCheckResourceAttr("google_compute_instance_from_machine_image.foobar5", "machine_type", "c3-standard-4"),
					resource.TestCheckResourceAttr("google_compute_instance_from_machine_image.foobar5", "scheduling.0.on_host_maintenance", "TERMINATE"),
					testAccCheckComputeInstanceHasConfidentialInstanceConfig(&instance, false, "TDX"),
				),
			},
		},
	})
}

func TestAccComputeInstanceFromMachineImage_encryptionKeyKMS(t *testing.T) {
	t.Parallel()

	var instance compute.Instance
	instanceName := fmt.Sprintf("tf-encinst-%s", acctest.RandString(t, 10))
	machineImageName := fmt.Sprintf("tf-encimg-%s", acctest.RandString(t, 10))
	kmsKey := acctest.BootstrapKMSKeyInLocation(t, "us-central1")

	context := map[string]interface{}{
		"instance_name":           instanceName,
		"image_name":              machineImageName,
		"kms_key_name":            kmsKey.CryptoKey.Name,
		"kms_key_service_account": "service-compute@compute.iam.gserviceaccount.com",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeInstanceFromMachineImageDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeInstanceFromMachineImage_encryptionKeyKMS(context),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeInstanceExists(t, "google_compute_instance_from_machine_image.foobar", &instance),
					resource.TestCheckResourceAttr("google_compute_instance_from_machine_image.foobar",
						"source_machine_image_encryption_key.0.kms_key_name", context["kms_key_name"].(string)),
					resource.TestCheckResourceAttr("google_compute_instance_from_machine_image.foobar",
						"source_machine_image_encryption_key.0.kms_key_service_account",
						context["kms_key_service_account"].(string)),
				),
			},
		},
	})
}

func TestAccComputeInstanceFromMachineImage_encryptionKeyRaw(t *testing.T) {
	t.Parallel()

	var instance compute.Instance
	instanceName := fmt.Sprintf("tf-encinst-%s", acctest.RandString(t, 10))
	machineImageName := fmt.Sprintf("tf-encimg-%s", acctest.RandString(t, 10))

	context := map[string]interface{}{
		"instance_name": instanceName,
		"image_name":    machineImageName,
		"raw_key":       "SGVsbG8gZnJvbSBHb29nbGUgQ2xvdWQgUGxhdGZvcm0=",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeInstanceFromMachineImageDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeInstanceFromMachineImage_encryptionKeyRaw(context),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeInstanceExists(t, "google_compute_instance_from_machine_image.foobar", &instance),
					resource.TestCheckResourceAttr("google_compute_instance_from_machine_image.foobar",
						"source_machine_image_encryption_key.0.raw_key", context["raw_key"].(string)),
				),
			},
		},
	})
}

func TestAccComputeInstanceFromMachineImage_encryptionKeySha(t *testing.T) {
	t.Parallel()

	var instance compute.Instance
	instanceName := fmt.Sprintf("tf-encinst-%s", acctest.RandString(t, 10))
	machineImageName := fmt.Sprintf("tf-encimg-%s", acctest.RandString(t, 10))

	context := map[string]interface{}{
		"instance_name": instanceName,
		"image_name":    machineImageName,
		"sha256":        "LKV0jGTF7BXTrQqiwcN4kKBZZiYHcQBX1Pg4gNPvVs8=",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeInstanceFromMachineImageDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeInstanceFromMachineImage_encryptionKeySha(context),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeInstanceExists(t, "google_compute_instance_from_machine_image.foobar", &instance),
					resource.TestCheckResourceAttr("google_compute_instance_from_machine_image.foobar",
						"source_machine_image_encryption_key.0.sha256", context["sha256"].(string)),
				),
			},
		},
	})
}

func TestAccComputeInstanceFromMachineImage_encryptionKeyRSA(t *testing.T) {
	t.Parallel()

	var instance compute.Instance
	instanceName := fmt.Sprintf("tf-encinst-%s", acctest.RandString(t, 10))
	machineImageName := fmt.Sprintf("tf-encimg-%s", acctest.RandString(t, 10))

	context := map[string]interface{}{
		"instance_name":     instanceName,
		"image_name":        machineImageName,
		"rsa_encrypted_key": "ieCx/NcW06PcT7Ep1X6LUTc/hLvUDYyzSZPPVCVPTVEohpeHASqC8uw5TzyO9U+Fka9JFHz0mBibXUInrC/jEk014kCK/NPjYgEMOyssZ4ZINPKxlUh2zn1bV+MCaTICrdmuSBTWlUUiFoDD6PYznLwh8ZNdaheCeZ8ewEXgFQ8V+sDroLaN3Xs3MDTXQEMMoNUXMCZEIpg9Vtp9x2oeQ5lAbtt7bYAAHf5l+gJWw3sUfs0/Glw5fpdjT8Uggrr+RMZezGrltJEF293rvTIjWOEB3z5OHyHwQkvdrPDFcTqsLfh+8Hr8g+mf+7zVPEC8nEbqpdl3GPv3A7AwpFp7MA==",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeInstanceFromMachineImageDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeInstanceFromMachineImage_encryptionKeyRSA(context),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeInstanceExists(t, "google_compute_instance_from_machine_image.foobar", &instance),
					resource.TestCheckResourceAttr("google_compute_instance_from_machine_image.foobar",
						"source_machine_image_encryption_key.0.rsa_encrypted_key", context["rsa_encrypted_key"].(string)),
				),
			},
		},
	})
}

func testAccCheckComputeInstanceFromMachineImageDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		config := acctest.GoogleProviderConfig(t)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_instance_from_machine_image" {
				continue
			}

			_, err := config.NewComputeClient(config.UserAgent).Instances.Get(
				config.Project, rs.Primary.Attributes["zone"], rs.Primary.ID).Do()
			if err == nil {
				return fmt.Errorf("Instance still exists")
			}
		}

		return nil
	}
}

func testAccComputeInstanceFromMachineImage_basic(instance, newInstance string) string {
	return fmt.Sprintf(`
resource "google_compute_instance" "vm" {
  provider     = google-beta

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-12"
    }
  }

  name         = "%s"
  machine_type = "n1-standard-1"

  network_interface {
    network = "default"
  }

  metadata = {
    foo = "bar"
  }

  scheduling {
    automatic_restart = true
  }

  can_ip_forward = true
}

resource "google_compute_machine_image" "foobar" {
  provider        = google-beta
  name            = "%s"
  source_instance = google_compute_instance.vm.self_link
}

resource "google_compute_instance_from_machine_image" "foobar" {
  provider = google-beta
  name = "%s"
  zone = "us-central1-a"

  source_machine_image = google_compute_machine_image.foobar.self_link

  // Overrides
  can_ip_forward = false
  labels = {
    my_key = "my_value"
  }
  scheduling {
    automatic_restart = false
  }
}
`, instance, instance, newInstance)
}

func testAccComputeInstanceFromMachineImage_ConfidentialInstanceConfigEnable(instance string, newInstance string) string {
	return fmt.Sprintf(`
resource "google_compute_instance" "vm1" {
  provider     = google-beta

  boot_disk {
    initialize_params {
      image = "ubuntu-os-cloud/ubuntu-2004-lts"
    }
  }

  name         = "%s"
  machine_type = "n2d-standard-2"

  network_interface {
    network = "default"
  }

  metadata = {
    foo = "bar"
  }

  confidential_instance_config {
    enable_confidential_compute = true
  }

  scheduling {
    on_host_maintenance = "TERMINATE"
  }
}

resource "google_compute_machine_image" "foobar1" {
  provider        = google-beta
  name            = "%s"
  source_instance = google_compute_instance.vm1.self_link
}

resource "google_compute_instance_from_machine_image" "foobar1" {
  provider = google-beta
  name = "%s"
  zone = "us-central1-a"

  source_machine_image = google_compute_machine_image.foobar1.self_link

  labels = {
    my_key = "my_value"
  }
  confidential_instance_config {
    enable_confidential_compute = true
  }
  scheduling {
    on_host_maintenance = "TERMINATE"
  }
}
`, instance, instance, newInstance)
}

func testAccComputeInstanceFromMachineImage_ConfidentialInstanceConfigEnableSev(instance string, newInstance string, confidentialInstanceType string) string {
	return fmt.Sprintf(`
resource "google_compute_instance" "vm2" {
  provider     = google-beta

  boot_disk {
    initialize_params {
      image = "ubuntu-os-cloud/ubuntu-2004-lts"
    }
  }

  name         = "%s"
  machine_type = "n2d-standard-2"

  network_interface {
    network = "default"
  }

  metadata = {
    foo = "bar"
  }

  confidential_instance_config {
    enable_confidential_compute     = true
    confidential_instance_type      = %q
  }

  scheduling {
    on_host_maintenance = "TERMINATE"
  }
}

resource "google_compute_machine_image" "foobar2" {
  provider        = google-beta
  name            = "%s"
  source_instance = google_compute_instance.vm2.self_link
}

resource "google_compute_instance_from_machine_image" "foobar2" {
  provider = google-beta
  name = "%s"
  zone = "us-central1-a"

  source_machine_image = google_compute_machine_image.foobar2.self_link

  labels = {
    my_key = "my_value"
  }
  confidential_instance_config {
    enable_confidential_compute     = true
    confidential_instance_type      = %q
  }
}
`, instance, confidentialInstanceType, instance, newInstance, confidentialInstanceType)
}

func testAccComputeInstanceFromMachineImage_ConfidentialInstanceConfigNoEnableSev(instance string, newInstance string, confidentialInstanceType string) string {
	return fmt.Sprintf(`
resource "google_compute_instance" "vm3" {
  provider     = google-beta

  boot_disk {
    initialize_params {
      image = "ubuntu-os-cloud/ubuntu-2004-lts"
    }
  }

  name         = "%s"
  machine_type = "n2d-standard-2"

  network_interface {
    network = "default"
  }

  metadata = {
    foo = "bar"
  }

  min_cpu_platform = "AMD Milan"

  confidential_instance_config {
    enable_confidential_compute     = false
    confidential_instance_type      = %q
  }
}

resource "google_compute_machine_image" "foobar3" {
  provider        = google-beta
  name            = "%s"
  source_instance = google_compute_instance.vm3.self_link
}

resource "google_compute_instance_from_machine_image" "foobar3" {
  provider = google-beta
  name = "%s"
  zone = "us-central1-a"

  source_machine_image = google_compute_machine_image.foobar3.self_link

  labels = {
    my_key = "my_value"
  }

  confidential_instance_config {
    enable_confidential_compute     = false
    confidential_instance_type      = %q
  }

  scheduling {
    on_host_maintenance = "TERMINATE"
  }
}
`, instance, confidentialInstanceType, instance, newInstance, confidentialInstanceType)
}

func testAccComputeInstanceFromMachineImage_ConfidentialInstanceConfigSevSnp(instance string, newInstance string, confidentialInstanceType string) string {
	return fmt.Sprintf(`
resource "google_compute_instance" "vm4" {
  provider     = google-beta

  boot_disk {
    initialize_params {
      image = "ubuntu-os-cloud/ubuntu-2004-lts"
    }
  }

  name         = "%s"
  machine_type = "n2d-standard-2"

  network_interface {
    network = "default"
  }

  metadata = {
    foo = "bar"
  }

  min_cpu_platform = "AMD Milan"

  confidential_instance_config {
    enable_confidential_compute     = false
    confidential_instance_type      = %q
  }

  scheduling {
    on_host_maintenance = "TERMINATE"
  }
}

resource "google_compute_machine_image" "foobar4" {
  provider        = google-beta
  name            = "%s"
  source_instance = google_compute_instance.vm4.self_link
}

resource "google_compute_instance_from_machine_image" "foobar4" {
  provider = google-beta
  name = "%s"
  zone = "us-central1-a"

  source_machine_image = google_compute_machine_image.foobar4.self_link

  labels = {
    my_key = "my_value"
  }
  confidential_instance_config {
    enable_confidential_compute     = false
    confidential_instance_type      = %q
  }
  scheduling {
    on_host_maintenance = "TERMINATE"
  }
}
`, instance, confidentialInstanceType, instance, newInstance, confidentialInstanceType)
}

func testAccComputeInstanceFromMachineImage_ConfidentialInstanceConfigTdx(instance string, newInstance string, confidentialInstanceType string) string {
	return fmt.Sprintf(`
resource "google_compute_instance" "vm5" {
  provider     = google-beta

  boot_disk {
    initialize_params {
      image = "ubuntu-os-cloud/ubuntu-2204-jammy-v20240927"
    }
  }

  name         = "%s"
  machine_type = "c3-standard-4"

  network_interface {
    network = "default"
  }

  metadata = {
    foo = "bar"
  }

  confidential_instance_config {
    confidential_instance_type = %q
  }

  scheduling {
    on_host_maintenance = "TERMINATE"
  }
}

resource "google_compute_machine_image" "foobar5" {
  provider        = google-beta
  name            = "%s"
  source_instance = google_compute_instance.vm5.self_link
}

resource "google_compute_instance_from_machine_image" "foobar5" {
  provider = google-beta
  name = "%s"
  zone = "us-central1-a"

  source_machine_image = google_compute_machine_image.foobar5.self_link

  labels = {
    my_key = "my_value"
  }
  confidential_instance_config {
    confidential_instance_type  = %q
  }
  scheduling {
    on_host_maintenance = "TERMINATE"
  }
}
`, instance, confidentialInstanceType, instance, newInstance, confidentialInstanceType)
}

func testAccComputeInstanceFromMachineImage_maxRunDuration(instance, newInstance string) string {
	return fmt.Sprintf(`
resource "google_compute_instance" "vm" {
  provider     = google-beta

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-12"
    }
  }

  name         = "%s"
  machine_type = "n1-standard-1"

  network_interface {
    network = "default"
  }

  metadata = {
    foo = "bar"
  }

  scheduling {
    automatic_restart = false
  }

}

resource "google_compute_machine_image" "foobar" {
  provider        = google-beta
  name            = "%s"
  source_instance = google_compute_instance.vm.self_link
}

resource "google_compute_instance_from_machine_image" "foobar" {
  provider = google-beta
  name = "%s"
  zone = "us-central1-a"

  source_machine_image = google_compute_machine_image.foobar.self_link

  labels = {
    my_key = "my_value"
  }
  scheduling {
    automatic_restart = false
    provisioning_model = "STANDARD"
    instance_termination_action = "STOP"
    max_run_duration {
	nanos = 123
	seconds = 60
    }
	on_instance_stop_action {
		discard_local_ssd = true
	}
  }
}
`, instance, instance, newInstance)
}

func testAccComputeInstanceFromMachineImage_localSsdRecoveryTimeout(instance, newInstance string) string {
	return fmt.Sprintf(`
resource "google_compute_instance" "vm" {
  provider     = google-beta

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-12"
    }
  }

  name         = "%s"
  machine_type = "n1-standard-1"

  network_interface {
    network = "default"
  }

  metadata = {
    foo = "bar"
  }

  scheduling {
    automatic_restart = true
    local_ssd_recovery_timeout {
			nanos = 0
			seconds = 3600
		}
  }
}

resource "google_compute_machine_image" "foobar" {
  provider        = google-beta
  name            = "%s"
  source_instance = google_compute_instance.vm.self_link
}

resource "google_compute_instance_from_machine_image" "foobar" {
  provider = google-beta
  name = "%s"
  zone = "us-central1-a"

  source_machine_image = google_compute_machine_image.foobar.self_link

  labels = {
    my_key = "my_value"
  }
  scheduling {
    automatic_restart = false
  }
}
`, instance, instance, newInstance)
}

func testAccComputeInstanceFromMachineImageWithOverride_localSsdRecoveryTimeout(instance, newInstance string) string {
	return fmt.Sprintf(`
resource "google_compute_instance" "vm" {
  provider     = google-beta

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-12"
    }
  }

  name         = "%s"
  machine_type = "n1-standard-1"

  network_interface {
    network = "default"
  }

  metadata = {
    foo = "bar"
  }

  scheduling {
    automatic_restart = true
    local_ssd_recovery_timeout {
			nanos = 0
			seconds = 3600
		}
  }
}

resource "google_compute_machine_image" "foobar" {
  provider        = google-beta
  name            = "%s"
  source_instance = google_compute_instance.vm.self_link
}

resource "google_compute_instance_from_machine_image" "foobar" {
  provider = google-beta
  name = "%s"
  zone = "us-central1-a"

  source_machine_image = google_compute_machine_image.foobar.self_link

  labels = {
    my_key = "my_value"
  }
  scheduling {
    automatic_restart = false
    local_ssd_recovery_timeout {
			nanos = 0
			seconds = 7200
		}
  }
}
`, instance, instance, newInstance)
}

func testAccComputeInstanceFromMachineImageWithOverride_partnerMetadata(instance, newInstance string) string {
	return fmt.Sprintf(`
resource "google_compute_instance" "vm" {
  provider     = google-beta

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-12"
    }
  }

  name         = "%s"
  machine_type = "n1-standard-1"

  network_interface {
    network = "default"
  }

  metadata = {
    foo = "bar"
  }

  scheduling {
    automatic_restart = true
  }

  partner_metadata = {
  	"test.compute.googleapis.com" = jsonencode({
  		entries = {
  			key = "value"
  		}
  	})
  }
}

resource "google_compute_machine_image" "foobar" {
  provider        = google-beta
  name            = "%s"
  source_instance = google_compute_instance.vm.self_link
}

resource "google_compute_instance_from_machine_image" "foobar" {
  provider = google-beta
  name = "%s"
  zone = "us-central1-a"

  source_machine_image = google_compute_machine_image.foobar.self_link

  labels = {
    my_key = "my_value"
  }
  scheduling {
    automatic_restart = false
  }

  partner_metadata = {
  	"test.compute.googleapis.com" = jsonencode({
  		entries = {
  			key1 = "value1"
  			key2 = 2
  			key3 = {
  				key31 = "value31"
  			}
  		}
  	})
  }
}
`, instance, instance, newInstance)
}

func testAccComputeInstanceFromMachineImage_overrideMetadataDotStartupScript(instanceName, generatedInstanceName string) string {
	return fmt.Sprintf(`
resource "google_compute_instance" "vm" {
  provider     = google-beta

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-12"
    }
  }

  name         = "%s"
  machine_type = "n1-standard-1"

  network_interface {
    network = "default"
  }

  metadata = {
    startup-script = "#!/bin/bash\necho Hello"
  }

}

resource "google_compute_machine_image" "foobar" {
  provider        = google-beta
  name            = "%s"
  source_instance = google_compute_instance.vm.self_link
}

resource "google_compute_instance_from_machine_image" "foobar" {
  provider = google-beta
  name = "%s"
  zone = "us-central1-a"

  source_machine_image = google_compute_machine_image.foobar.self_link

  // Overrides
  metadata = {
    startup-script = ""
  }
}
`, instanceName, instanceName, generatedInstanceName)
}

func testAccComputeInstanceFromMachineImage_diffProject(projectID, org, billingId, instance, newInstance string) string {
	return fmt.Sprintf(`
resource "google_project" "project" {
	provider     = google-beta
	project_id      = "%s"
	name            = "%s"
	org_id          = "%s"
	billing_account = "%s"
	deletion_policy = "DELETE"
}

resource "google_project_service" "service" {
	provider     = google-beta
	project = google_project.project.project_id
	service = "compute.googleapis.com"
	timeouts {
	  create = "30m"
	  update = "40m"
	}
	disable_dependent_services = true
}

resource "google_project_service" "monitoring" {
	provider     = google-beta
	project = google_project.project.project_id
	service = "monitoring.googleapis.com"
	timeouts {
	  create = "30m"
	  update = "40m"
	}
	disable_dependent_services = true

	depends_on = [google_project_service.service]
}

resource "google_compute_instance" "vm" {
  provider     = google-beta
  project = google_project.project.project_id
  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-12"
    }
  }

  name         = "%s"
  machine_type = "n1-standard-1"

  network_interface {
    network = "default"
  }

  metadata = {
    foo = "bar"
  }

  scheduling {
    automatic_restart = true
  }

  can_ip_forward = true

  depends_on = [google_project_service.monitoring]
}

resource "google_compute_machine_image" "foobar" {
  provider        = google-beta
  project = google_project.project.project_id
  name            = "%s"
  source_instance = google_compute_instance.vm.self_link
}

resource "google_compute_instance_from_machine_image" "foobar" {
  provider = google-beta
  name = "%s"
  zone = "us-central1-a"

  source_machine_image = google_compute_machine_image.foobar.self_link

  // Overrides
  can_ip_forward = false
  labels = {
    my_key = "my_value"
  }
  scheduling {
    automatic_restart = false
  }
}
`, projectID, projectID, org, billingId, instance, instance, newInstance)
}

func testAccComputeInstanceFromMachineImage_encryptionKeyKMS(context map[string]interface{}) string {
	return fmt.Sprintf(`
resource "google_compute_instance_from_machine_image" "foobar" {
		provider = google-beta
		name = "%s"
		zone = "us-central1-a"
		source_machine_image = google_compute_machine_image.foobar.self_link
		source_machine_image_encryption_key {
			kms_key_name = "%s"
			kms_key_service_account = "%s"
		}
	}
`, context["instance_name"].(string), context["kms_key_name"].(string), context["kms_key_service_account"].(string))
}

func testAccComputeInstanceFromMachineImage_encryptionKeyRaw(context map[string]interface{}) string {
	return fmt.Sprintf(`
resource "google_compute_instance_from_machine_image" "foobar" {
		provider = google-beta
		name = "%s"
		zone = "us-central1-a"
		source_machine_image = google_compute_machine_image.foobar.self_link
		source_machine_image_encryption_key {
			raw_key = "%s"
		}
	}
`, context["instance_name"].(string), context["raw_key"].(string))
}

func testAccComputeInstanceFromMachineImage_encryptionKeySha(context map[string]interface{}) string {
	return fmt.Sprintf(`
resource "google_compute_instance_from_machine_image" "foobar" {
		provider = google-beta
		name = "%s"
		zone = "us-central1-a"
		source_machine_image = google_compute_machine_image.foobar.self_link
		source_machine_image_encryption_key {
		}
	}
`, context["instance_name"].(string))
}

func testAccComputeInstanceFromMachineImage_encryptionKeyRSA(context map[string]interface{}) string {
	return fmt.Sprintf(`
resource "google_compute_instance_from_machine_image" "foobar" {
		provider = google-beta
		name = "%s"
		zone = "us-central1-a"
		source_machine_image = google_compute_machine_image.foobar.self_link
		source_machine_image_encryption_key {
			rsa_encrypted_key = "%s"
		}
	}
`, context["instance_name"].(string), context["rsa_encrypted_key"].(string))
}
