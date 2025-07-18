---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/monitoring/UptimeCheckConfig.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Cloud (Stackdriver) Monitoring"
description: |-
  This message configures which resources and services to monitor for availability.
---

# google_monitoring_uptime_check_config

This message configures which resources and services to monitor for availability.


To get more information about UptimeCheckConfig, see:

* [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.uptimeCheckConfigs)
* How-to Guides
    * [Official Documentation](https://cloud.google.com/monitoring/uptime-checks/)

~> **Warning:** All arguments including the following potentially sensitive
values will be stored in the raw state as plain text: `http_check.auth_info.password`.
[Read more about sensitive data in state](https://www.terraform.io/language/state/sensitive-data).

~> **Note:**  All arguments marked as write-only values will not be stored in the state: `http_check.auth_info.password_wo`.
[Read more about Write-only Attributes](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/write-only-arguments).

## Example Usage - Uptime Check Config Http


```hcl
resource "google_monitoring_uptime_check_config" "http" {
  display_name       = "http-uptime-check"
  timeout            = "60s"
  log_check_failures = true
  user_labels  = {
    example-key = "example-value"
  }

  http_check {
    path = "some-path"
    port = "8010"
    request_method = "POST"
    content_type = "USER_PROVIDED"
    custom_content_type = "application/json"
    body = "Zm9vJTI1M0RiYXI="
    ping_config {
      pings_count = 1
    }
  }

  monitored_resource {
    type = "uptime_url"
    labels = {
      project_id = "my-project-name"
      host       = "192.168.1.1"
    }
  }

  content_matchers {
    content = "\"example\""
    matcher = "MATCHES_JSON_PATH"
    json_path_matcher {
      json_path = "$.path"
      json_matcher = "EXACT_MATCH"
    }
  }

  checker_type = "STATIC_IP_CHECKERS"
}
```
## Example Usage - Uptime Check Config Http Password Wo


```hcl
resource "google_monitoring_uptime_check_config" "http" {
  display_name = "http-uptime-check"
  timeout      = "60s"
  user_labels  = {
    example-key = "example-value"
  }

  http_check {
    path = "some-path"
    port = "8010"
    request_method = "POST"
    content_type = "USER_PROVIDED"
    custom_content_type = "application/json"
    body = "Zm9vJTI1M0RiYXI="
    ping_config {
      pings_count = 1
    }
    auth_info {
      username = "name"
      password_wo = "password1"
      password_wo_version = "1"
    }
  }

  monitored_resource {
    type = "uptime_url"
    labels = {
      project_id = "my-project-name"
      host       = "192.168.1.1"
    }
  }

  content_matchers {
    content = "\"example\""
    matcher = "MATCHES_JSON_PATH"
    json_path_matcher {
      json_path = "$.path"
      json_matcher = "EXACT_MATCH"
    }
  }

  checker_type = "STATIC_IP_CHECKERS"
}
```
## Example Usage - Uptime Check Config Status Code


```hcl
resource "google_monitoring_uptime_check_config" "status_code" {
  display_name = "http-uptime-check"
  timeout      = "60s"

  http_check {
    path = "some-path"
    port = "8010"
    request_method = "POST"
    content_type = "URL_ENCODED"
    body = "Zm9vJTI1M0RiYXI="

    accepted_response_status_codes {
      status_class = "STATUS_CLASS_2XX"
    }
    accepted_response_status_codes {
            status_value = 301
    }
    accepted_response_status_codes {
            status_value = 302
    }
  }

  monitored_resource {
    type = "uptime_url"
    labels = {
      project_id = "my-project-name"
      host       = "192.168.1.1"
    }
  }

  content_matchers {
    content = "\"example\""
    matcher = "MATCHES_JSON_PATH"
    json_path_matcher {
      json_path = "$.path"
      json_matcher = "EXACT_MATCH"
    }
  }

  checker_type = "STATIC_IP_CHECKERS"
}
```
## Example Usage - Uptime Check Config Https


```hcl
resource "google_monitoring_uptime_check_config" "https" {
  display_name = "https-uptime-check"
  timeout = "60s"

  http_check {
    path = "/some-path"
    port = "443"
    use_ssl = true
    validate_ssl = true
    service_agent_authentication {
      type = "OIDC_TOKEN"
    }
  }

  monitored_resource {
    type = "uptime_url"
    labels = {
      project_id = "my-project-name"
      host = "192.168.1.1"
    }
  }

  content_matchers {
    content = "example"
    matcher = "MATCHES_JSON_PATH"
    json_path_matcher {
      json_path = "$.path"
      json_matcher = "REGEX_MATCH"
    }
  }
}
```
<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=uptime_check_tcp&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Uptime Check Tcp


```hcl
resource "google_monitoring_uptime_check_config" "tcp_group" {
  display_name = "tcp-uptime-check"
  timeout      = "60s"

  tcp_check {
    port = 888
    ping_config {
      pings_count = 2
    }
  }

  resource_group {
    resource_type = "INSTANCE"
    group_id      = google_monitoring_group.check.name
  }
}

resource "google_monitoring_group" "check" {
  display_name = "uptime-check-group"
  filter       = "resource.metadata.name=has_substring(\"foo\")"
}
```
## Example Usage - Uptime Check Config Synthetic Monitor


```hcl
resource "google_storage_bucket" "bucket" {
  name     = "my-project-name-gcf-source"  # Every bucket name must be globally unique
  location = "US"
  uniform_bucket_level_access = true
}
 
resource "google_storage_bucket_object" "object" {
  name   = "function-source.zip"
  bucket = google_storage_bucket.bucket.name
  source = "synthetic-fn-source.zip"  # Add path to the zipped function source code
}
 
resource "google_cloudfunctions2_function" "function" {
  name = "synthetic_function"
  location = "us-central1"
 
  build_config {
    runtime = "nodejs20"
    entry_point = "SyntheticFunction"  # Set the entry point 
    source {
      storage_source {
        bucket = google_storage_bucket.bucket.name
        object = google_storage_bucket_object.object.name
      }
    }
  }
 
  service_config {
    max_instance_count  = 1
    available_memory    = "256M"
    timeout_seconds     = 60
  }
}

resource "google_monitoring_uptime_check_config" "synthetic_monitor" {
  display_name = "synthetic_monitor"
  timeout = "60s"

  synthetic_monitor {
    cloud_function_v2 {
      name = google_cloudfunctions2_function.function.id
    }
  }
}
```

## Argument Reference

The following arguments are supported:


* `display_name` -
  (Required)
  A human-friendly name for the uptime check configuration. The display name should be unique within a Stackdriver Workspace in order to make it easier to identify; however, uniqueness is not enforced.

* `timeout` -
  (Required)
  The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). [See the accepted formats]( https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#google.protobuf.Duration)


* `period` -
  (Optional)
  How often, in seconds, the uptime check is performed. Currently, the only supported values are 60s (1 minute), 300s (5 minutes), 600s (10 minutes), and 900s (15 minutes). Optional, defaults to 300s.

* `content_matchers` -
  (Optional)
  The expected content on the page the check is run against. Currently, only the first entry in the list is supported, and other entries will be ignored. The server will look for an exact match of the string in the page response's content. This field is optional and should only be specified if a content match is required.
  Structure is [documented below](#nested_content_matchers).

* `selected_regions` -
  (Optional)
  The list of regions from which the check will be run. Some regions contain one location, and others contain more than one. If this field is specified, enough regions to include a minimum of 3 locations must be provided, or an error message is returned. Not specifying this field will result in uptime checks running from all regions.

* `log_check_failures` -
  (Optional)
  Specifies whether to log the results of failed probes to Cloud Logging.

* `checker_type` -
  (Optional)
  The checker type to use for the check. If the monitored resource type is `servicedirectory_service`, `checker_type` must be set to `VPC_CHECKERS`.
  Possible values are: `STATIC_IP_CHECKERS`, `VPC_CHECKERS`.

* `user_labels` -
  (Optional)
  User-supplied key/value data to be used for organizing and identifying the `UptimeCheckConfig` objects. The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.

* `http_check` -
  (Optional)
  Contains information needed to make an HTTP or HTTPS check.
  Structure is [documented below](#nested_http_check).

* `tcp_check` -
  (Optional)
  Contains information needed to make a TCP check.
  Structure is [documented below](#nested_tcp_check).

* `resource_group` -
  (Optional)
  The group resource associated with the configuration.
  Structure is [documented below](#nested_resource_group).

* `monitored_resource` -
  (Optional)
  The [monitored resource]
  (https://cloud.google.com/monitoring/api/resources) associated with the
  configuration. The following monitored resource types are supported for
  uptime checks:
  * `aws_ec2_instance`
  * `aws_elb_load_balancer`
  * `gae_app`
  * `gce_instance`
  * `k8s_service`
  * `servicedirectory_service`
  * `uptime_url`
  Structure is [documented below](#nested_monitored_resource).

* `synthetic_monitor` -
  (Optional)
  A Synthetic Monitor deployed to a Cloud Functions V2 instance.
  Structure is [documented below](#nested_synthetic_monitor).

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.



<a name="nested_content_matchers"></a>The `content_matchers` block supports:

* `content` -
  (Required)
  String or regex content to match (max 1024 bytes)

* `matcher` -
  (Optional)
  The type of content matcher that will be applied to the server output, compared to the content string when the check is run.
  Default value is `CONTAINS_STRING`.
  Possible values are: `CONTAINS_STRING`, `NOT_CONTAINS_STRING`, `MATCHES_REGEX`, `NOT_MATCHES_REGEX`, `MATCHES_JSON_PATH`, `NOT_MATCHES_JSON_PATH`.

* `json_path_matcher` -
  (Optional)
  Information needed to perform a JSONPath content match. Used for `ContentMatcherOption::MATCHES_JSON_PATH` and `ContentMatcherOption::NOT_MATCHES_JSON_PATH`.
  Structure is [documented below](#nested_content_matchers_content_matchers_json_path_matcher).


<a name="nested_content_matchers_content_matchers_json_path_matcher"></a>The `json_path_matcher` block supports:

* `json_path` -
  (Required)
  JSONPath within the response output pointing to the expected `ContentMatcher::content` to match against.

* `json_matcher` -
  (Optional)
  Options to perform JSONPath content matching.
  Default value is `EXACT_MATCH`.
  Possible values are: `EXACT_MATCH`, `REGEX_MATCH`.

<a name="nested_http_check"></a>The `http_check` block supports:

* `request_method` -
  (Optional)
  The HTTP request method to use for the check. If set to `METHOD_UNSPECIFIED` then `request_method` defaults to `GET`.
  Default value is `GET`.
  Possible values are: `METHOD_UNSPECIFIED`, `GET`, `POST`.

* `content_type` -
  (Optional)
  The content type to use for the check.
  Possible values are: `TYPE_UNSPECIFIED`, `URL_ENCODED`, `USER_PROVIDED`.

* `custom_content_type` -
  (Optional)
  A user provided content type header to use for the check. The invalid configurations outlined in the `content_type` field apply to custom_content_type`, as well as the following 1. `content_type` is `URL_ENCODED` and `custom_content_type` is set. 2. `content_type` is `USER_PROVIDED` and `custom_content_type` is not set.

* `auth_info` -
  (Optional)
  The authentication information using username and password. Optional when creating an HTTP check; defaults to empty. Do not use with other authentication fields.
  Structure is [documented below](#nested_http_check_auth_info).

* `service_agent_authentication` -
  (Optional)
  The authentication information using the Monitoring Service Agent. Optional when creating an HTTPS check; defaults to empty. Do not use with other authentication fields.
  Structure is [documented below](#nested_http_check_service_agent_authentication).

* `port` -
  (Optional)
  The port to the page to run the check against. Will be combined with `host` (specified within the [`monitored_resource`](#nested_monitored_resource)) and path to construct the full URL. Optional (defaults to 80 without SSL, or 443 with SSL).

* `headers` -
  (Optional)
  The list of headers to send as part of the uptime check request. If two headers have the same key and different values, they should be entered as a single header, with the value being a comma-separated list of all the desired values as described in [RFC 2616 (page 31)](https://www.w3.org/Protocols/rfc2616/rfc2616.txt). Entering two separate headers with the same key in a Create call will cause the first to be overwritten by the second. The maximum number of headers allowed is 100.

* `path` -
  (Optional)
  The path to the page to run the check against. Will be combined with the host (specified within the MonitoredResource) and port to construct the full URL. If the provided path does not begin with `/`, a `/` will be prepended automatically. Optional (defaults to `/`).

* `use_ssl` -
  (Optional)
  If true, use HTTPS instead of HTTP to run the check.

* `validate_ssl` -
  (Optional)
  Boolean specifying whether to include SSL certificate validation as a part of the Uptime check. Only applies to checks where `monitored_resource` is set to `uptime_url`. If `use_ssl` is `false`, setting `validate_ssl` to `true` has no effect.

* `mask_headers` -
  (Optional)
  Boolean specifying whether to encrypt the header information. Encryption should be specified for any headers related to authentication that you do not wish to be seen when retrieving the configuration. The server will be responsible for encrypting the headers. On Get/List calls, if `mask_headers` is set to `true` then the headers will be obscured with `******`.

* `body` -
  (Optional)
  The request body associated with the HTTP POST request. If `content_type` is `URL_ENCODED`, the body passed in must be URL-encoded. Users can provide a `Content-Length` header via the `headers` field or the API will do so. If the `request_method` is `GET` and `body` is not empty, the API will return an error. The maximum byte size is 1 megabyte. Note - As with all bytes fields JSON representations are base64 encoded. e.g. `foo=bar` in URL-encoded form is `foo%3Dbar` and in base64 encoding is `Zm9vJTI1M0RiYXI=`.

* `accepted_response_status_codes` -
  (Optional)
  If present, the check will only pass if the HTTP response status code is in this set of status codes. If empty, the HTTP status code will only pass if the HTTP status code is 200-299.
  Structure is [documented below](#nested_http_check_accepted_response_status_codes).

* `ping_config` -
  (Optional)
  Contains information needed to add pings to an HTTP check.
  Structure is [documented below](#nested_http_check_ping_config).


<a name="nested_http_check_service_agent_authentication"></a>The `service_agent_authentication` block supports:

* `type` -
  (Optional)
  The type of authentication to use.
  Possible values are: `SERVICE_AGENT_AUTHENTICATION_TYPE_UNSPECIFIED`, `OIDC_TOKEN`.

<a name="nested_http_check_accepted_response_status_codes"></a>The `accepted_response_status_codes` block supports:

* `status_value` -
  (Optional)
  A status code to accept.

* `status_class` -
  (Optional)
  A class of status codes to accept.
  Possible values are: `STATUS_CLASS_1XX`, `STATUS_CLASS_2XX`, `STATUS_CLASS_3XX`, `STATUS_CLASS_4XX`, `STATUS_CLASS_5XX`, `STATUS_CLASS_ANY`.

<a name="nested_http_check_ping_config"></a>The `ping_config` block supports:

* `pings_count` -
  (Required)
  Number of ICMP pings. A maximum of 3 ICMP pings is currently supported.

<a name="nested_tcp_check"></a>The `tcp_check` block supports:

* `port` -
  (Required)
  The port to the page to run the check against. Will be combined with host (specified within the `monitored_resource`) to construct the full URL.

* `ping_config` -
  (Optional)
  Contains information needed to add pings to a TCP check.
  Structure is [documented below](#nested_tcp_check_ping_config).


<a name="nested_tcp_check_ping_config"></a>The `ping_config` block supports:

* `pings_count` -
  (Required)
  Number of ICMP pings. A maximum of 3 ICMP pings is currently supported.

<a name="nested_resource_group"></a>The `resource_group` block supports:

* `resource_type` -
  (Optional)
  The resource type of the group members.
  Possible values are: `RESOURCE_TYPE_UNSPECIFIED`, `INSTANCE`, `AWS_ELB_LOAD_BALANCER`.

* `group_id` -
  (Optional)
  The group of resources being monitored. Should be the `name` of a group

<a name="nested_monitored_resource"></a>The `monitored_resource` block supports:

* `type` -
  (Required)
  The monitored resource type. This field must match the type field of a [`MonitoredResourceDescriptor`](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.monitoredResourceDescriptors#MonitoredResourceDescriptor) object. For example, the type of a Compute Engine VM instance is `gce_instance`. For a list of types, see [Monitoring resource types](https://cloud.google.com/monitoring/api/resources) and [Logging resource types](https://cloud.google.com/logging/docs/api/v2/resource-list).

* `labels` -
  (Required)
  Values for all of the labels listed in the associated monitored resource descriptor. For example, Compute Engine VM instances use the labels `project_id`, `instance_id`, and `zone`.

<a name="nested_synthetic_monitor"></a>The `synthetic_monitor` block supports:

* `cloud_function_v2` -
  (Required)
  Target a Synthetic Monitor GCFv2 Instance
  Structure is [documented below](#nested_synthetic_monitor_cloud_function_v2).


<a name="nested_synthetic_monitor_cloud_function_v2"></a>The `cloud_function_v2` block supports:

* `name` -
  (Required)
  The fully qualified name of the cloud function resource.

## Ephemeral Attributes Reference

The following write-only attributes are supported:


## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `{{name}}`

* `name` -
  A unique resource name for this UptimeCheckConfig. The format is `projects/[PROJECT_ID]/uptimeCheckConfigs/[UPTIME_CHECK_ID]`.

* `uptime_check_id` -
  The id of the uptime check


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


UptimeCheckConfig can be imported using any of these accepted formats:

* `{{project}}/{{name}}`
* `{{project}} {{name}}`
* `{{name}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import UptimeCheckConfig using one of the formats above. For example:

```tf
import {
  id = "{{project}}/{{name}}"
  to = google_monitoring_uptime_check_config.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), UptimeCheckConfig can be imported using one of the formats above. For example:

```
$ terraform import google_monitoring_uptime_check_config.default {{project}}/{{name}}
$ terraform import google_monitoring_uptime_check_config.default "{{project}} {{name}}"
$ terraform import google_monitoring_uptime_check_config.default {{name}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
