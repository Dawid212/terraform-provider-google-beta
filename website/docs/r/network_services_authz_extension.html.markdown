---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/networkservices/AuthzExtension.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Network Services"
description: |-
  AuthzExtension is a resource that allows traffic forwarding to a callout backend service to make an authorization decision.
---

# google_network_services_authz_extension

AuthzExtension is a resource that allows traffic forwarding to a callout backend service to make an authorization decision.


To get more information about AuthzExtension, see:

* [API documentation](https://cloud.google.com/service-extensions/docs/reference/rest/v1beta1/projects.locations.authzExtensions)

## Example Usage - Network Services Authz Extension Basic


```hcl
resource "google_compute_region_backend_service" "default" {
  name                  = "authz-service"
  project               = "my-project-name"
  region                = "us-west1"

  protocol              = "HTTP2"
  load_balancing_scheme = "INTERNAL_MANAGED"
  port_name             = "grpc"
}

resource "google_network_services_authz_extension" "default" {
  name     = "my-authz-ext"
  project  = "my-project-name"
  location = "us-west1"

  description           = "my description"
  load_balancing_scheme = "INTERNAL_MANAGED"
  authority             = "ext11.com"
  service               = google_compute_region_backend_service.default.self_link
  timeout               = "0.1s"
  fail_open             = false
  forward_headers       = ["Authorization"]
}
```

## Argument Reference

The following arguments are supported:


* `load_balancing_scheme` -
  (Required)
  All backend services and forwarding rules referenced by this extension must share the same load balancing scheme.
  For more information, refer to [Backend services overview](https://cloud.google.com/load-balancing/docs/backend-service).
  Possible values are: `INTERNAL_MANAGED`, `EXTERNAL_MANAGED`.

* `authority` -
  (Required)
  The :authority header in the gRPC request sent from Envoy to the extension service.

* `service` -
  (Required)
  The reference to the service that runs the extension.
  To configure a callout extension, service must be a fully-qualified reference to a [backend service](https://cloud.google.com/compute/docs/reference/rest/v1/backendServices) in the format:
  https://www.googleapis.com/compute/v1/projects/{project}/regions/{region}/backendServices/{backendService} or https://www.googleapis.com/compute/v1/projects/{project}/global/backendServices/{backendService}.

* `timeout` -
  (Required)
  Specifies the timeout for each individual message on the stream. The timeout must be between 10-10000 milliseconds.

* `name` -
  (Required)
  Identifier. Name of the AuthzExtension resource.

* `location` -
  (Required)
  The location of the resource.


* `description` -
  (Optional)
  A human-readable description of the resource.

* `labels` -
  (Optional)
  Set of labels associated with the AuthzExtension resource.

  **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
  Please refer to the field `effective_labels` for all of the labels present on the resource.

* `fail_open` -
  (Optional)
  Determines how the proxy behaves if the call to the extension fails or times out.
  When set to TRUE, request or response processing continues without error. Any subsequent extensions in the extension chain are also executed. When set to FALSE or the default setting of FALSE is used, one of the following happens:
  * If response headers have not been delivered to the downstream client, a generic 500 error is returned to the client. The error response can be tailored by configuring a custom error response in the load balancer.
  * If response headers have been delivered, then the HTTP stream to the downstream client is reset.

* `metadata` -
  (Optional)
  The metadata provided here is included as part of the metadata_context (of type google.protobuf.Struct) in the ProcessingRequest message sent to the extension server. The metadata is available under the namespace com.google.authz_extension.<resourceName>. The following variables are supported in the metadata Struct:
  {forwarding_rule_id} - substituted with the forwarding rule's fully qualified resource name.

* `forward_headers` -
  (Optional)
  List of the HTTP headers to forward to the extension (from the client). If omitted, all headers are sent. Each element is a string indicating the header name.

* `wire_format` -
  (Optional)
  The format of communication supported by the callout extension. Will be set to EXT_PROC_GRPC by the backend if no value is set.
  Possible values are: `WIRE_FORMAT_UNSPECIFIED`, `EXT_PROC_GRPC`.

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.



## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/locations/{{location}}/authzExtensions/{{name}}`

* `create_time` -
  The timestamp when the resource was created.

* `update_time` -
  The timestamp when the resource was updated.

* `terraform_labels` -
  The combination of labels configured directly on the resource
   and default labels configured on the provider.

* `effective_labels` -
  All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 30 minutes.
- `update` - Default is 30 minutes.
- `delete` - Default is 30 minutes.

## Import


AuthzExtension can be imported using any of these accepted formats:

* `projects/{{project}}/locations/{{location}}/authzExtensions/{{name}}`
* `{{project}}/{{location}}/{{name}}`
* `{{location}}/{{name}}`
* `{{name}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import AuthzExtension using one of the formats above. For example:

```tf
import {
  id = "projects/{{project}}/locations/{{location}}/authzExtensions/{{name}}"
  to = google_network_services_authz_extension.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), AuthzExtension can be imported using one of the formats above. For example:

```
$ terraform import google_network_services_authz_extension.default projects/{{project}}/locations/{{location}}/authzExtensions/{{name}}
$ terraform import google_network_services_authz_extension.default {{project}}/{{location}}/{{name}}
$ terraform import google_network_services_authz_extension.default {{location}}/{{name}}
$ terraform import google_network_services_authz_extension.default {{name}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
