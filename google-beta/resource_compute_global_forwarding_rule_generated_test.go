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

func TestAccComputeGlobalForwardingRule_externalTcpProxyLbMigBackendExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeGlobalForwardingRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeGlobalForwardingRule_externalTcpProxyLbMigBackendExample(context),
			},
			{
				ResourceName:            "google_compute_global_forwarding_rule.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"network", "port_range", "target", "ip_address"},
			},
		},
	})
}

func testAccComputeGlobalForwardingRule_externalTcpProxyLbMigBackendExample(context map[string]interface{}) string {
	return Nprintf(`
# External TCP proxy load balancer with managed instance group backend

# VPC
resource "google_compute_network" "default" {
  name                    = "tf-test-tcp-proxy-xlb-network%{random_suffix}"
  provider                = google-beta
  auto_create_subnetworks = false
}

# backend subnet
resource "google_compute_subnetwork" "default" {
  name          = "tf-test-tcp-proxy-xlb-subnet%{random_suffix}"
  provider      = google-beta
  ip_cidr_range = "10.0.1.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.id
}

# reserved IP address
resource "google_compute_global_address" "default" {
  provider = google-beta
  name = "tf-test-tcp-proxy-xlb-ip%{random_suffix}"
}

# forwarding rule
resource "google_compute_global_forwarding_rule" "default" {
  name                  = "tf-test-tcp-proxy-xlb-forwarding-rule%{random_suffix}"
  provider              = google-beta
  ip_protocol           = "TCP"
  load_balancing_scheme = "EXTERNAL"
  port_range            = "110"
  target                = google_compute_target_tcp_proxy.default.id
  ip_address            = google_compute_global_address.default.id
}

resource "google_compute_target_tcp_proxy" "default" {
  provider = google-beta
  name            = "tf-test-test-proxy-health-check%{random_suffix}"
  backend_service = google_compute_backend_service.default.id
}

# backend service
resource "google_compute_backend_service" "default" {
  provider = google-beta
  name                  = "tf-test-tcp-proxy-xlb-backend-service%{random_suffix}"
  protocol              = "TCP"
  port_name             = "tcp"
  load_balancing_scheme = "EXTERNAL"
  timeout_sec           = 10
  health_checks         = [google_compute_health_check.default.id]
  backend {
    group           = google_compute_instance_group_manager.default.instance_group
    balancing_mode  = "UTILIZATION"
    max_utilization = 1.0
    capacity_scaler = 1.0
  }
}

resource "google_compute_health_check" "default" {
  provider = google-beta
  name               = "tf-test-tcp-proxy-health-check%{random_suffix}"
  timeout_sec        = 1
  check_interval_sec = 1

  tcp_health_check {
    port = "80"
  }
}

# instance template
resource "google_compute_instance_template" "default" {
  name         = "tf-test-tcp-proxy-xlb-mig-template%{random_suffix}"
  provider     = google-beta
  machine_type = "e2-small"
  tags         = ["allow-health-check"]

  network_interface {
    network    = google_compute_network.default.id
    subnetwork = google_compute_subnetwork.default.id
    access_config {
      # add external ip to fetch packages
    }
  }
  disk {
    source_image = "debian-cloud/debian-10"
    auto_delete  = true
    boot         = true
  }

  # install nginx and serve a simple web page
  metadata = {
    startup-script = <<-EOF1
      #! /bin/bash
      set -euo pipefail
      export DEBIAN_FRONTEND=noninteractive
      apt-get update
      apt-get install -y nginx-light jq
      NAME=$(curl -H "Metadata-Flavor: Google" "http://metadata.google.internal/computeMetadata/v1/instance/hostname")
      IP=$(curl -H "Metadata-Flavor: Google" "http://metadata.google.internal/computeMetadata/v1/instance/network-interfaces/0/ip")
      METADATA=$(curl -f -H "Metadata-Flavor: Google" "http://metadata.google.internal/computeMetadata/v1/instance/attributes/?recursive=True" | jq 'del(.["startup-script"])')
      cat <<EOF > /var/www/html/index.html
      <pre>
      Name: $NAME
      IP: $IP
      Metadata: $METADATA
      </pre>
      EOF
    EOF1
  }
  lifecycle {
    create_before_destroy = true
  }
}

# MIG
resource "google_compute_instance_group_manager" "default" {
  name     = "tf-test-tcp-proxy-xlb-mig1%{random_suffix}"
  provider = google-beta
  zone     = "us-central1-c"
  named_port {
    name = "tcp"
    port = 80
  }
  version {
    instance_template = google_compute_instance_template.default.id
    name              = "primary"
  }
  base_instance_name = "vm"
  target_size        = 2
}

# allow access from health check ranges
resource "google_compute_firewall" "default" {
  name          = "tf-test-tcp-proxy-xlb-fw-allow-hc%{random_suffix}"
  provider      = google-beta
  direction     = "INGRESS"
  network       = google_compute_network.default.id
  source_ranges = ["130.211.0.0/22", "35.191.0.0/16"]
  allow {
    protocol = "tcp"
  }
  target_tags = ["allow-health-check"]
}
`, context)
}

func TestAccComputeGlobalForwardingRule_externalHttpLbMigBackendCustomHeaderExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeGlobalForwardingRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeGlobalForwardingRule_externalHttpLbMigBackendCustomHeaderExample(context),
			},
			{
				ResourceName:            "google_compute_global_forwarding_rule.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"network", "port_range", "target", "ip_address"},
			},
		},
	})
}

func testAccComputeGlobalForwardingRule_externalHttpLbMigBackendCustomHeaderExample(context map[string]interface{}) string {
	return Nprintf(`
# External HTTP load balancer with a CDN-enabled managed instance group backend
# and custom request and response headers

# VPC
resource "google_compute_network" "default" {
  name                    = "tf-test-l7-xlb-network%{random_suffix}"
  provider                = google-beta
  auto_create_subnetworks = false
}

# backend subnet
resource "google_compute_subnetwork" "default" {
  name          = "tf-test-l7-xlb-subnet%{random_suffix}"
  provider      = google-beta
  ip_cidr_range = "10.0.1.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.id
}

# reserved IP address
resource "google_compute_global_address" "default" {
  provider = google-beta
  name = "tf-test-l7-xlb-static-ip%{random_suffix}"
}

# forwarding rule
resource "google_compute_global_forwarding_rule" "default" {
  name                  = "tf-test-l7-xlb-forwarding-rule%{random_suffix}"
  provider              = google-beta
  ip_protocol           = "TCP"
  load_balancing_scheme = "EXTERNAL"
  port_range            = "80"
  target                = google_compute_target_http_proxy.default.id
  ip_address            = google_compute_global_address.default.id
}

# http proxy
resource "google_compute_target_http_proxy" "default" {
  name     = "tf-test-l7-xlb-target-http-proxy%{random_suffix}"
  provider = google-beta
  url_map  = google_compute_url_map.default.id
}

# url map
resource "google_compute_url_map" "default" {
  name            = "tf-test-l7-xlb-url-map%{random_suffix}"
  provider        = google-beta
  default_service = google_compute_backend_service.default.id
}

# backend service with custom request and response headers
resource "google_compute_backend_service" "default" {
  name                     = "tf-test-l7-xlb-backend-service%{random_suffix}"
  provider                 = google-beta
  protocol                 = "HTTP"
  port_name                = "my-port"
  load_balancing_scheme    = "EXTERNAL"
  timeout_sec              = 10
  enable_cdn               = true
  custom_request_headers   = ["X-Client-Geo-Location: {client_region_subdivision}, {client_city}"]
  custom_response_headers  = ["X-Cache-Hit: {cdn_cache_status}"]
  health_checks            = [google_compute_health_check.default.id]
  backend {
    group           = google_compute_instance_group_manager.default.instance_group
    balancing_mode  = "UTILIZATION"
    capacity_scaler = 1.0
  }
}

# instance template
resource "google_compute_instance_template" "default" {
  name         = "tf-test-l7-xlb-mig-template%{random_suffix}"
  provider     = google-beta
  machine_type = "e2-small"
  tags         = ["allow-health-check"]

  network_interface {
    network    = google_compute_network.default.id
    subnetwork = google_compute_subnetwork.default.id
    access_config {
      # add external ip to fetch packages
    }
  }
  disk {
    source_image = "debian-cloud/debian-10"
    auto_delete  = true
    boot         = true
  }

  # install nginx and serve a simple web page
  metadata = {
    startup-script = <<-EOF1
      #! /bin/bash
      set -euo pipefail

      export DEBIAN_FRONTEND=noninteractive
      apt-get update
      apt-get install -y nginx-light jq

      NAME=$(curl -H "Metadata-Flavor: Google" "http://metadata.google.internal/computeMetadata/v1/instance/hostname")
      IP=$(curl -H "Metadata-Flavor: Google" "http://metadata.google.internal/computeMetadata/v1/instance/network-interfaces/0/ip")
      METADATA=$(curl -f -H "Metadata-Flavor: Google" "http://metadata.google.internal/computeMetadata/v1/instance/attributes/?recursive=True" | jq 'del(.["startup-script"])')

      cat <<EOF > /var/www/html/index.html
      <pre>
      Name: $NAME
      IP: $IP
      Metadata: $METADATA
      </pre>
      EOF
    EOF1
  }
  lifecycle {
    create_before_destroy = true
  }
}

# health check
resource "google_compute_health_check" "default" {
  name     = "tf-test-l7-xlb-hc%{random_suffix}"
  provider = google-beta
  http_health_check {
    port_specification = "USE_SERVING_PORT"
  }
}

# MIG
resource "google_compute_instance_group_manager" "default" {
  name     = "tf-test-l7-xlb-mig1%{random_suffix}"
  provider = google-beta
  zone     = "us-central1-c"
  named_port {
    name = "http"
    port = 8080
  }
  version {
    instance_template = google_compute_instance_template.default.id
    name              = "primary"
  }
  base_instance_name = "vm"
  target_size        = 2
}

# allow access from health check ranges
resource "google_compute_firewall" "default" {
  name          = "tf-test-l7-xlb-fw-allow-hc%{random_suffix}"
  provider      = google-beta
  direction     = "INGRESS"
  network       = google_compute_network.default.id
  source_ranges = ["130.211.0.0/22", "35.191.0.0/16"]
  allow {
    protocol = "tcp"
  }
  target_tags = ["allow-health-check"]
}
`, context)
}

func TestAccComputeGlobalForwardingRule_globalForwardingRuleHttpExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeGlobalForwardingRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeGlobalForwardingRule_globalForwardingRuleHttpExample(context),
			},
			{
				ResourceName:            "google_compute_global_forwarding_rule.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"network", "port_range", "target"},
			},
		},
	})
}

func testAccComputeGlobalForwardingRule_globalForwardingRuleHttpExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_global_forwarding_rule" "default" {
  name       = "tf-test-global-rule%{random_suffix}"
  target     = google_compute_target_http_proxy.default.id
  port_range = "80"
}

resource "google_compute_target_http_proxy" "default" {
  name        = "tf-test-target-proxy%{random_suffix}"
  description = "a description"
  url_map     = google_compute_url_map.default.id
}

resource "google_compute_url_map" "default" {
  name            = "url-map-tf-test-target-proxy%{random_suffix}"
  description     = "a description"
  default_service = google_compute_backend_service.default.id

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name            = "allpaths"
    default_service = google_compute_backend_service.default.id

    path_rule {
      paths   = ["/*"]
      service = google_compute_backend_service.default.id
    }
  }
}

resource "google_compute_backend_service" "default" {
  name        = "backend%{random_suffix}"
  port_name   = "http"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_http_health_check" "default" {
  name               = "check-backend%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}
`, context)
}

func TestAccComputeGlobalForwardingRule_globalForwardingRuleInternalExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeGlobalForwardingRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeGlobalForwardingRule_globalForwardingRuleInternalExample(context),
			},
			{
				ResourceName:            "google_compute_global_forwarding_rule.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"network", "port_range", "target"},
			},
		},
	})
}

func testAccComputeGlobalForwardingRule_globalForwardingRuleInternalExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_global_forwarding_rule" "default" {
  provider              = google-beta
  name                  = "tf-test-global-rule%{random_suffix}"
  target                = google_compute_target_http_proxy.default.id
  port_range            = "80"
  load_balancing_scheme = "INTERNAL_SELF_MANAGED"
  ip_address            = "0.0.0.0"
  metadata_filters {
    filter_match_criteria = "MATCH_ANY"
    filter_labels {
      name  = "PLANET"
      value = "MARS"
    }
  }
}

resource "google_compute_target_http_proxy" "default" {
  provider    = google-beta
  name        = "tf-test-target-proxy%{random_suffix}"
  description = "a description"
  url_map     = google_compute_url_map.default.id
}

resource "google_compute_url_map" "default" {
  provider        = google-beta
  name            = "url-map-tf-test-target-proxy%{random_suffix}"
  description     = "a description"
  default_service = google_compute_backend_service.default.id

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name            = "allpaths"
    default_service = google_compute_backend_service.default.id

    path_rule {
      paths   = ["/*"]
      service = google_compute_backend_service.default.id
    }
  }
}

resource "google_compute_backend_service" "default" {
  provider              = google-beta
  name                  = "backend%{random_suffix}"
  port_name             = "http"
  protocol              = "HTTP"
  timeout_sec           = 10
  load_balancing_scheme = "INTERNAL_SELF_MANAGED"

  backend {
    group                 = google_compute_instance_group_manager.igm.instance_group
    balancing_mode        = "RATE"
    capacity_scaler       = 0.4
    max_rate_per_instance = 50
  }

  health_checks = [google_compute_health_check.default.id]
}

data "google_compute_image" "debian_image" {
  provider = google-beta
  family   = "debian-11"
  project  = "debian-cloud"
}

resource "google_compute_instance_group_manager" "igm" {
  provider = google-beta
  name     = "tf-test-igm-internal%{random_suffix}"
  version {
    instance_template = google_compute_instance_template.instance_template.id
    name              = "primary"
  }
  base_instance_name = "internal-glb"
  zone               = "us-central1-f"
  target_size        = 1
}

resource "google_compute_instance_template" "instance_template" {
  provider     = google-beta
  name         = "template-backend%{random_suffix}"
  machine_type = "e2-medium"

  network_interface {
    network = "default"
  }

  disk {
    source_image = data.google_compute_image.debian_image.self_link
    auto_delete  = true
    boot         = true
  }
}

resource "google_compute_health_check" "default" {
  provider           = google-beta
  name               = "check-backend%{random_suffix}"
  check_interval_sec = 1
  timeout_sec        = 1

  tcp_health_check {
    port = "80"
  }
}
`, context)
}

func TestAccComputeGlobalForwardingRule_globalForwardingRuleExternalManagedExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeGlobalForwardingRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeGlobalForwardingRule_globalForwardingRuleExternalManagedExample(context),
			},
			{
				ResourceName:            "google_compute_global_forwarding_rule.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"network", "port_range", "target"},
			},
		},
	})
}

func testAccComputeGlobalForwardingRule_globalForwardingRuleExternalManagedExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_global_forwarding_rule" "default" {
  name                  = "tf-test-global-rule%{random_suffix}"
  target                = google_compute_target_http_proxy.default.id
  port_range            = "80"
  load_balancing_scheme = "EXTERNAL_MANAGED"
}

resource "google_compute_target_http_proxy" "default" {
  name        = "tf-test-target-proxy%{random_suffix}"
  description = "a description"
  url_map     = google_compute_url_map.default.id
}

resource "google_compute_url_map" "default" {
  name            = "url-map-tf-test-target-proxy%{random_suffix}"
  description     = "a description"
  default_service = google_compute_backend_service.default.id

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name            = "allpaths"
    default_service = google_compute_backend_service.default.id

    path_rule {
      paths   = ["/*"]
      service = google_compute_backend_service.default.id
    }
  }
}

resource "google_compute_backend_service" "default" {
  name                  = "backend%{random_suffix}"
  port_name             = "http"
  protocol              = "HTTP"
  timeout_sec           = 10
  load_balancing_scheme = "EXTERNAL_MANAGED"
}
`, context)
}

func TestAccComputeGlobalForwardingRule_globalForwardingRuleHybridExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeGlobalForwardingRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeGlobalForwardingRule_globalForwardingRuleHybridExample(context),
			},
			{
				ResourceName:            "google_compute_global_forwarding_rule.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"network", "port_range", "target"},
			},
		},
	})
}

func testAccComputeGlobalForwardingRule_globalForwardingRuleHybridExample(context map[string]interface{}) string {
	return Nprintf(`
// Roughly mirrors https://cloud.google.com/load-balancing/docs/https/setting-up-ext-https-hybrid
variable "subnetwork_cidr" {
  default = "10.0.0.0/24"
}

resource "google_compute_network" "default" {
  name                    = "tf-test-my-network%{random_suffix}"
}

resource "google_compute_network" "internal" {
  name                    = "tf-test-my-internal-network%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "internal"{
  name                    = "tf-test-my-subnetwork%{random_suffix}"
  network                 = google_compute_network.internal.id
  ip_cidr_range           = var.subnetwork_cidr
  region                  = "us-central1"
  private_ip_google_access= true
}

// Zonal NEG with GCE_VM_IP_PORT
resource "google_compute_network_endpoint_group" "default" {
  name                  = "tf-test-default-neg%{random_suffix}"
  network               = google_compute_network.default.id
  default_port          = "90"
  zone                  = "us-central1-a"
  network_endpoint_type = "GCE_VM_IP_PORT"
}

// Zonal NEG with GCE_VM_IP
resource "google_compute_network_endpoint_group" "internal" {
  name                  = "tf-test-internal-neg%{random_suffix}"
  network               = google_compute_network.internal.id
  subnetwork            = google_compute_subnetwork.internal.id
  zone                  = "us-central1-a"
  network_endpoint_type = "GCE_VM_IP"
}

// Hybrid connectivity NEG
resource "google_compute_network_endpoint_group" "hybrid" {
  name                  = "tf-test-hybrid-neg%{random_suffix}"
  network               = google_compute_network.default.id
  default_port          = "90"
  zone                  = "us-central1-a"
  network_endpoint_type = "NON_GCP_PRIVATE_IP_PORT"
}

resource "google_compute_network_endpoint" "hybrid-endpoint" {
  network_endpoint_group = google_compute_network_endpoint_group.hybrid.name
  port       = google_compute_network_endpoint_group.hybrid.default_port
  ip_address = "127.0.0.1"
}

// Backend service for Zonal NEG
resource "google_compute_backend_service" "default" {
  name                  = "tf-test-backend-default%{random_suffix}"
  port_name             = "http"
  protocol              = "HTTP"
  timeout_sec           = 10
  backend {
    group = google_compute_network_endpoint_group.default.id
    balancing_mode               = "RATE"
    max_rate_per_endpoint        = 10
  }
  health_checks = [google_compute_health_check.default.id]
}

// Backgend service for Hybrid NEG
resource "google_compute_backend_service" "hybrid" {
  name                  = "tf-test-backend-hybrid%{random_suffix}"
  port_name             = "http"
  protocol              = "HTTP"
  timeout_sec           = 10
  backend {
    group                        = google_compute_network_endpoint_group.hybrid.id
    balancing_mode               = "RATE"
    max_rate_per_endpoint = 10
  }
  health_checks = [google_compute_health_check.default.id]
}

resource "google_compute_health_check" "default" {
  name               = "tf-test-health-check%{random_suffix}"
  timeout_sec        = 1
  check_interval_sec = 1

  tcp_health_check {
    port = "80"
  }
}

resource "google_compute_url_map" "default" {
  name            = "url-map-tf-test-target-proxy%{random_suffix}"
  description     = "a description"
  default_service = google_compute_backend_service.default.id

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name            = "allpaths"
    default_service = google_compute_backend_service.default.id

    path_rule {
      paths   = ["/*"]
      service = google_compute_backend_service.default.id
    }

    path_rule {
      paths   = ["/hybrid"]
      service = google_compute_backend_service.hybrid.id
    }
  }
}

resource "google_compute_target_http_proxy" "default" {
  name        = "tf-test-target-proxy%{random_suffix}"
  description = "a description"
  url_map     = google_compute_url_map.default.id
}

resource "google_compute_global_forwarding_rule" "default" {
  name       = "tf-test-global-rule%{random_suffix}"
  target     = google_compute_target_http_proxy.default.id
  port_range = "80"
}
`, context)
}

func TestAccComputeGlobalForwardingRule_privateServiceConnectGoogleApisExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       GetTestProjectFromEnv(),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeGlobalForwardingRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeGlobalForwardingRule_privateServiceConnectGoogleApisExample(context),
			},
			{
				ResourceName:            "google_compute_global_forwarding_rule.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"network", "ip_address"},
			},
		},
	})
}

func testAccComputeGlobalForwardingRule_privateServiceConnectGoogleApisExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_network" "network" {
  provider      = google-beta
  project       = "%{project}"
  name          = "tf-test-my-network%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "vpc_subnetwork" {
  provider                 = google-beta
  project                  = google_compute_network.network.project
  name                     = "tf-test-my-subnetwork%{random_suffix}"
  ip_cidr_range            = "10.2.0.0/16"
  region                   = "us-central1"
  network                  = google_compute_network.network.id
  private_ip_google_access = true
}

resource "google_compute_global_address" "default" {
  provider      = google-beta
  project       = google_compute_network.network.project
  name          = "tf-test-global-psconnect-ip%{random_suffix}"
  address_type  = "INTERNAL"
  purpose       = "PRIVATE_SERVICE_CONNECT"
  network       = google_compute_network.network.id
  address       = "100.100.100.106"
}

resource "google_compute_global_forwarding_rule" "default" {
  provider      = google-beta
  project       = google_compute_network.network.project
  name          = "globalrule%{random_suffix}"
  target        = "all-apis"
  network       = google_compute_network.network.id
  ip_address    = google_compute_global_address.default.id
  load_balancing_scheme = ""
}
`, context)
}

func testAccCheckComputeGlobalForwardingRuleDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_global_forwarding_rule" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/global/forwardingRules/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("ComputeGlobalForwardingRule still exists at %s", url)
			}
		}

		return nil
	}
}
