---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/gkehub2/Feature.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "GKEHub"
description: |-
  Feature represents the settings and status of any Hub Feature.
---

# google_gke_hub_feature

Feature represents the settings and status of any Hub Feature.


To get more information about Feature, see:

* [API documentation](https://cloud.google.com/anthos/fleet-management/docs/reference/rest/v1/projects.locations.features)
* How-to Guides
    * [Registering a Cluster](https://cloud.google.com/anthos/multicluster-management/connect/registering-a-cluster#register_cluster)

## Example Usage - Gkehub Feature Multi Cluster Ingress


```hcl
resource "google_container_cluster" "cluster" {
  name               = "my-cluster"
  location           = "us-central1-a"
  initial_node_count = 1
}

resource "google_gke_hub_membership" "membership" {
  membership_id = "my-membership"
  endpoint {
    gke_cluster {
      resource_link = "//container.googleapis.com/${google_container_cluster.cluster.id}"
    }
  }
}

resource "google_gke_hub_feature" "feature" {
  name = "multiclusteringress"
  location = "global"
  spec {
    multiclusteringress {
      config_membership = google_gke_hub_membership.membership.id
    }
  }
}
```
## Example Usage - Gkehub Feature Multi Cluster Service Discovery


```hcl
resource "google_gke_hub_feature" "feature" {
  name = "multiclusterservicediscovery"
  location = "global"
  labels = {
    foo = "bar"
  }
}
```
## Example Usage - Gkehub Feature Anthos Service Mesh


```hcl
resource "google_gke_hub_feature" "feature" {
  name = "servicemesh"
  location = "global"
}
```
## Example Usage - Enable Fleet Observability For Default Logs With Copy


```hcl
resource "google_gke_hub_feature" "feature" {
  name = "fleetobservability"
  location = "global"
  spec {
    fleetobservability {
      logging_config {
        default_config {
          mode = "COPY"
        }
      }
    }
  }
}
```
## Example Usage - Enable Fleet Observability For Scope Logs With Move


```hcl
resource "google_gke_hub_feature" "feature" {
  name = "fleetobservability"
  location = "global"
  spec {
    fleetobservability {
      logging_config {
        fleet_scope_logs_config {
          mode = "MOVE"
        }
      }
    }
  }
}
```
## Example Usage - Enable Fleet Observability For Both Default And Scope Logs


```hcl
resource "google_gke_hub_feature" "feature" {
  name = "fleetobservability"
  location = "global"
  spec {
    fleetobservability {
      logging_config {
        default_config {
          mode = "COPY"
        }
        fleet_scope_logs_config {
          mode = "MOVE"
        }
      }
    }
  }
}
```
## Example Usage - Enable Fleet Default Member Config Service Mesh


```hcl
resource "google_gke_hub_feature" "feature" {
  name = "servicemesh"
  location = "global"
  fleet_default_member_config {
    mesh {
      management = "MANAGEMENT_AUTOMATIC"
    }
  }
}
```
## Example Usage - Enable Fleet Default Member Config Configmanagement


```hcl
resource "google_gke_hub_feature" "feature" {
  name = "configmanagement"
  location = "global"
  fleet_default_member_config {
    configmanagement {
      config_sync {
        git {
          sync_repo = "https://github.com/hashicorp/terraform"
        }
      }
    }
  }
}
```
## Example Usage - Enable Fleet Default Member Config Policycontroller


```hcl
resource "google_gke_hub_feature" "feature" {
  name = "policycontroller"
  location = "global"
  fleet_default_member_config {
    policycontroller {
      policy_controller_hub_config {
        install_spec = "INSTALL_SPEC_ENABLED"
        exemptable_namespaces = ["foo"]
        policy_content {
          bundles {
            bundle = "policy-essentials-v2022"
            exempted_namespaces = ["foo", "bar"]
          }
          template_library {
            installation = "ALL"
          }
        }
        audit_interval_seconds = 30
        referential_rules_enabled = true
      }
    }
  }
}
```
## Example Usage - Enable Fleet Default Member Config Policycontroller Full


```hcl
resource "google_gke_hub_feature" "feature" {
  name = "policycontroller"
  location = "global"
  fleet_default_member_config {
    policycontroller {
      policy_controller_hub_config {
        install_spec = "INSTALL_SPEC_SUSPENDED"
        policy_content {
          bundles {
            bundle = "pci-dss-v3.2.1"
            exempted_namespaces = ["baz", "bar"]
          }
          bundles {
            bundle = "nist-sp-800-190"
            exempted_namespaces = []
          }
          template_library {
            installation = "ALL"
          }
        }
        constraint_violation_limit = 50
        referential_rules_enabled = true
        log_denies_enabled = true
        mutation_enabled = true
        deployment_configs {
          component = "admission"
          replica_count = 2
          pod_affinity = "ANTI_AFFINITY"
        }
        deployment_configs {
          component = "audit"
          container_resources {
            limits {
              memory = "1Gi"
              cpu = "1.5"
            }
            requests {
              memory = "500Mi"
              cpu = "150m"
            }
          }
          pod_toleration {
            key = "key1"
            operator = "Equal"
            value = "value1"
            effect = "NoSchedule"
          }
        }
        monitoring {
          backends = [
            "PROMETHEUS"
          ]
        }
      }
    }
  }
}
```
## Example Usage - Enable Fleet Default Member Config Policycontroller Minimal


```hcl
resource "google_gke_hub_feature" "feature" {
  name = "policycontroller"
  location = "global"
  fleet_default_member_config {
    policycontroller {
      policy_controller_hub_config {
        install_spec = "INSTALL_SPEC_ENABLED"
        policy_content {}
        constraint_violation_limit = 50
        referential_rules_enabled = true
        log_denies_enabled = true
        mutation_enabled = true
        deployment_configs {
          component = "admission"
        }
        monitoring {}
      }
    }
  }
}
```
## Example Usage - Gkehub Feature Clusterupgrade


```hcl
resource "google_gke_hub_feature" "feature" {
  name = "clusterupgrade"
  location = "global"
  spec {
    clusterupgrade {
      upstream_fleets = []
      post_conditions {
        soaking = "60s"
      }
    }
  }
}
```
## Example Usage - Gkehub Feature Rbacrolebinding Actuation


```hcl
resource "google_gke_hub_feature" "feature" {
  name = "rbacrolebindingactuation"
  location = "global"
  spec {
    rbacrolebindingactuation {
      allowed_custom_roles = ["custom-role1","custom-role2","custom-role3"]
    }
  }
}
```

## Argument Reference

The following arguments are supported:


* `location` -
  (Required)
  The location for the resource


* `name` -
  (Optional)
  The full, unique name of this Feature resource

* `labels` -
  (Optional)
  GCP labels for this Feature.
  **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
  Please refer to the field `effective_labels` for all of the labels present on the resource.

* `spec` -
  (Optional)
  Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
  Structure is [documented below](#nested_spec).

* `fleet_default_member_config` -
  (Optional)
  Optional. Fleet Default Membership Configuration.
  Structure is [documented below](#nested_fleet_default_member_config).

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.



<a name="nested_spec"></a>The `spec` block supports:

* `multiclusteringress` -
  (Optional)
  Multicluster Ingress-specific spec.
  Structure is [documented below](#nested_spec_multiclusteringress).

* `fleetobservability` -
  (Optional)
  Fleet Observability feature spec.
  Structure is [documented below](#nested_spec_fleetobservability).

* `clusterupgrade` -
  (Optional)
  Clusterupgrade feature spec.
  Structure is [documented below](#nested_spec_clusterupgrade).

* `rbacrolebindingactuation` -
  (Optional)
  RBACRolebinding Actuation feature spec.
  Structure is [documented below](#nested_spec_rbacrolebindingactuation).


<a name="nested_spec_multiclusteringress"></a>The `multiclusteringress` block supports:

* `config_membership` -
  (Required)
  Fully-qualified Membership name which hosts the MultiClusterIngress CRD. Example: `projects/foo-proj/locations/global/memberships/bar`

<a name="nested_spec_fleetobservability"></a>The `fleetobservability` block supports:

* `logging_config` -
  (Optional)
  Specified if fleet logging feature is enabled for the entire fleet. If UNSPECIFIED, fleet logging feature is disabled for the entire fleet.
  Structure is [documented below](#nested_spec_fleetobservability_logging_config).


<a name="nested_spec_fleetobservability_logging_config"></a>The `logging_config` block supports:

* `default_config` -
  (Optional)
  Specified if applying the default routing config to logs not specified in other configs.
  Structure is [documented below](#nested_spec_fleetobservability_logging_config_default_config).

* `fleet_scope_logs_config` -
  (Optional)
  Specified if applying the routing config to all logs for all fleet scopes.
  Structure is [documented below](#nested_spec_fleetobservability_logging_config_fleet_scope_logs_config).


<a name="nested_spec_fleetobservability_logging_config_default_config"></a>The `default_config` block supports:

* `mode` -
  (Optional)
  Specified if fleet logging feature is enabled.
  Possible values are: `MODE_UNSPECIFIED`, `COPY`, `MOVE`.

<a name="nested_spec_fleetobservability_logging_config_fleet_scope_logs_config"></a>The `fleet_scope_logs_config` block supports:

* `mode` -
  (Optional)
  Specified if fleet logging feature is enabled.
  Possible values are: `MODE_UNSPECIFIED`, `COPY`, `MOVE`.

<a name="nested_spec_clusterupgrade"></a>The `clusterupgrade` block supports:

* `upstream_fleets` -
  (Required)
  Specified if other fleet should be considered as a source of upgrades. Currently, at most one upstream fleet is allowed. The fleet name should be either fleet project number or id.

* `post_conditions` -
  (Required)
  Post conditions to override for the specified upgrade.
  Structure is [documented below](#nested_spec_clusterupgrade_post_conditions).

* `gke_upgrade_overrides` -
  (Optional)
  Configuration overrides for individual upgrades.
  Structure is [documented below](#nested_spec_clusterupgrade_gke_upgrade_overrides).


<a name="nested_spec_clusterupgrade_post_conditions"></a>The `post_conditions` block supports:

* `soaking` -
  (Required)
  Amount of time to "soak" after a rollout has been finished before marking it COMPLETE. Cannot exceed 30 days.

<a name="nested_spec_clusterupgrade_gke_upgrade_overrides"></a>The `gke_upgrade_overrides` block supports:

* `upgrade` -
  (Required)
  Which upgrade to override.
  Structure is [documented below](#nested_spec_clusterupgrade_gke_upgrade_overrides_gke_upgrade_overrides_upgrade).

* `post_conditions` -
  (Required)
  Post conditions to override for the specified upgrade.
  Structure is [documented below](#nested_spec_clusterupgrade_gke_upgrade_overrides_gke_upgrade_overrides_post_conditions).


<a name="nested_spec_clusterupgrade_gke_upgrade_overrides_gke_upgrade_overrides_upgrade"></a>The `upgrade` block supports:

* `name` -
  (Required)
  Name of the upgrade, e.g., "k8s_control_plane". It should be a valid upgrade name. It must not exceet 99 characters.

* `version` -
  (Required)
  Version of the upgrade, e.g., "1.22.1-gke.100". It should be a valid version. It must not exceet 99 characters.

<a name="nested_spec_clusterupgrade_gke_upgrade_overrides_gke_upgrade_overrides_post_conditions"></a>The `post_conditions` block supports:

* `soaking` -
  (Required)
  Amount of time to "soak" after a rollout has been finished before marking it COMPLETE. Cannot exceed 30 days.

<a name="nested_spec_rbacrolebindingactuation"></a>The `rbacrolebindingactuation` block supports:

* `allowed_custom_roles` -
  (Optional)
  The list of allowed custom roles (ClusterRoles). If a custom role is not part of this list, it cannot be used in a fleet scope RBACRoleBinding. If a custom role in this list is in use, it cannot be removed from the list until the scope RBACRolebindings using it are deleted.

<a name="nested_fleet_default_member_config"></a>The `fleet_default_member_config` block supports:

* `mesh` -
  (Optional)
  Service Mesh spec
  Structure is [documented below](#nested_fleet_default_member_config_mesh).

* `configmanagement` -
  (Optional)
  Config Management spec
  Structure is [documented below](#nested_fleet_default_member_config_configmanagement).

* `policycontroller` -
  (Optional)
  Policy Controller spec
  Structure is [documented below](#nested_fleet_default_member_config_policycontroller).


<a name="nested_fleet_default_member_config_mesh"></a>The `mesh` block supports:

* `management` -
  (Required)
  Whether to automatically manage Service Mesh
  Possible values are: `MANAGEMENT_UNSPECIFIED`, `MANAGEMENT_AUTOMATIC`, `MANAGEMENT_MANUAL`.

<a name="nested_fleet_default_member_config_configmanagement"></a>The `configmanagement` block supports:

* `version` -
  (Optional)
  Version of Config Sync installed

* `management` -
  (Optional)
  Set this field to MANAGEMENT_AUTOMATIC to enable Config Sync auto-upgrades, and set this field to MANAGEMENT_MANUAL or MANAGEMENT_UNSPECIFIED to disable Config Sync auto-upgrades.
  Possible values are: `MANAGEMENT_UNSPECIFIED`, `MANAGEMENT_AUTOMATIC`, `MANAGEMENT_MANUAL`.

* `config_sync` -
  (Optional)
  ConfigSync configuration for the cluster
  Structure is [documented below](#nested_fleet_default_member_config_configmanagement_config_sync).


<a name="nested_fleet_default_member_config_configmanagement_config_sync"></a>The `config_sync` block supports:

* `source_format` -
  (Optional)
  Specifies whether the Config Sync Repo is in hierarchical or unstructured mode

* `enabled` -
  (Optional)
  Enables the installation of ConfigSync. If set to true, ConfigSync resources will be created and the other ConfigSync fields will be applied if exist. If set to false, all other ConfigSync fields will be ignored, ConfigSync resources will be deleted. If omitted, ConfigSync resources will be managed depends on the presence of the git or oci field.

* `prevent_drift` -
  (Optional)
  Set to true to enable the Config Sync admission webhook to prevent drifts. If set to `false`, disables the Config Sync admission webhook and does not prevent drifts.

* `metrics_gcp_service_account_email` -
  (Optional)
  The Email of the Google Cloud Service Account (GSA) used for exporting Config Sync metrics to Cloud Monitoring. The GSA should have the Monitoring Metric Writer(roles/monitoring.metricWriter) IAM role. The Kubernetes ServiceAccount `default` in the namespace `config-management-monitoring` should be bound to the GSA.

* `git` -
  (Optional)
  Git repo configuration for the cluster
  Structure is [documented below](#nested_fleet_default_member_config_configmanagement_config_sync_git).

* `oci` -
  (Optional)
  OCI repo configuration for the cluster
  Structure is [documented below](#nested_fleet_default_member_config_configmanagement_config_sync_oci).


<a name="nested_fleet_default_member_config_configmanagement_config_sync_git"></a>The `git` block supports:

* `sync_repo` -
  (Optional)
  The URL of the Git repository to use as the source of truth

* `sync_branch` -
  (Optional)
  The branch of the repository to sync from. Default: master

* `policy_dir` -
  (Optional)
  The path within the Git repository that represents the top level of the repo to sync

* `sync_rev` -
  (Optional)
  Git revision (tag or hash) to check out. Default HEAD

* `secret_type` -
  (Required)
  Type of secret configured for access to the Git repo

* `https_proxy` -
  (Optional)
  URL for the HTTPS Proxy to be used when communicating with the Git repo

* `gcp_service_account_email` -
  (Optional)
  The Google Cloud Service Account Email used for auth when secretType is gcpServiceAccount

* `sync_wait_secs` -
  (Optional)
  Period in seconds between consecutive syncs. Default: 15

<a name="nested_fleet_default_member_config_configmanagement_config_sync_oci"></a>The `oci` block supports:

* `sync_repo` -
  (Optional)
  The OCI image repository URL for the package to sync from

* `policy_dir` -
  (Optional)
  The absolute path of the directory that contains the local resources. Default: the root directory of the image

* `secret_type` -
  (Required)
  Type of secret configured for access to the Git repo

* `gcp_service_account_email` -
  (Optional)
  The Google Cloud Service Account Email used for auth when secretType is gcpServiceAccount

* `sync_wait_secs` -
  (Optional)
  Period in seconds between consecutive syncs. Default: 15

* `version` -
  (Optional, Deprecated)
  Version of Config Sync installed

  ~> **Warning:** The `configmanagement.config_sync.oci.version` field is deprecated and will be removed in a future major release. Please use `configmanagement.version` field to specify the version of Config Sync installed instead.

<a name="nested_fleet_default_member_config_policycontroller"></a>The `policycontroller` block supports:

* `version` -
  (Optional)
  Configures the version of Policy Controller

* `policy_controller_hub_config` -
  (Required)
  Configuration of Policy Controller
  Structure is [documented below](#nested_fleet_default_member_config_policycontroller_policy_controller_hub_config).


<a name="nested_fleet_default_member_config_policycontroller_policy_controller_hub_config"></a>The `policy_controller_hub_config` block supports:

* `install_spec` -
  (Required)
  Configures the mode of the Policy Controller installation
  Possible values are: `INSTALL_SPEC_UNSPECIFIED`, `INSTALL_SPEC_NOT_INSTALLED`, `INSTALL_SPEC_ENABLED`, `INSTALL_SPEC_SUSPENDED`, `INSTALL_SPEC_DETACHED`.

* `audit_interval_seconds` -
  (Optional)
  Interval for Policy Controller Audit scans (in seconds). When set to 0, this disables audit functionality altogether.

* `exemptable_namespaces` -
  (Optional)
  The set of namespaces that are excluded from Policy Controller checks. Namespaces do not need to currently exist on the cluster.

* `log_denies_enabled` -
  (Optional)
  Logs all denies and dry run failures.

* `mutation_enabled` -
  (Optional)
  Enables the ability to mutate resources using Policy Controller.

* `referential_rules_enabled` -
  (Optional)
  Enables the ability to use Constraint Templates that reference to objects other than the object currently being evaluated.

* `monitoring` -
  (Optional)
  Monitoring specifies the configuration of monitoring Policy Controller.
  Structure is [documented below](#nested_fleet_default_member_config_policycontroller_policy_controller_hub_config_monitoring).

* `constraint_violation_limit` -
  (Optional)
  The maximum number of audit violations to be stored in a constraint. If not set, the internal default of 20 will be used.

* `deployment_configs` -
  (Optional)
  Map of deployment configs to deployments ("admission", "audit", "mutation").
  Structure is [documented below](#nested_fleet_default_member_config_policycontroller_policy_controller_hub_config_deployment_configs).

* `policy_content` -
  (Optional)
  Specifies the desired policy content on the cluster.
  Structure is [documented below](#nested_fleet_default_member_config_policycontroller_policy_controller_hub_config_policy_content).


<a name="nested_fleet_default_member_config_policycontroller_policy_controller_hub_config_monitoring"></a>The `monitoring` block supports:

* `backends` -
  (Optional)
  Specifies the list of backends Policy Controller will export to. An empty list would effectively disable metrics export.
  Each value may be one of: `MONITORING_BACKEND_UNSPECIFIED`, `PROMETHEUS`, `CLOUD_MONITORING`.

<a name="nested_fleet_default_member_config_policycontroller_policy_controller_hub_config_deployment_configs"></a>The `deployment_configs` block supports:

* `component` - (Required) The identifier for this object. Format specified above.

* `replica_count` -
  (Optional)
  Pod replica count.

* `container_resources` -
  (Optional)
  Container resource requirements.
  Structure is [documented below](#nested_fleet_default_member_config_policycontroller_policy_controller_hub_config_deployment_configs_deployment_config_container_resources).

* `pod_affinity` -
  (Optional)
  Pod affinity configuration.
  Possible values are: `AFFINITY_UNSPECIFIED`, `NO_AFFINITY`, `ANTI_AFFINITY`.

* `pod_toleration` -
  (Optional)
  Pod tolerations of node taints.
  Structure is [documented below](#nested_fleet_default_member_config_policycontroller_policy_controller_hub_config_deployment_configs_deployment_config_pod_toleration).


<a name="nested_fleet_default_member_config_policycontroller_policy_controller_hub_config_deployment_configs_deployment_config_container_resources"></a>The `container_resources` block supports:

* `limits` -
  (Optional)
  Limits describes the maximum amount of compute resources allowed for use by the running container.
  Structure is [documented below](#nested_fleet_default_member_config_policycontroller_policy_controller_hub_config_deployment_configs_deployment_config_container_resources_limits).

* `requests` -
  (Optional)
  Requests describes the amount of compute resources reserved for the container by the kube-scheduler.
  Structure is [documented below](#nested_fleet_default_member_config_policycontroller_policy_controller_hub_config_deployment_configs_deployment_config_container_resources_requests).


<a name="nested_fleet_default_member_config_policycontroller_policy_controller_hub_config_deployment_configs_deployment_config_container_resources_limits"></a>The `limits` block supports:

* `memory` -
  (Optional)
  Memory requirement expressed in Kubernetes resource units.

* `cpu` -
  (Optional)
  CPU requirement expressed in Kubernetes resource units.

<a name="nested_fleet_default_member_config_policycontroller_policy_controller_hub_config_deployment_configs_deployment_config_container_resources_requests"></a>The `requests` block supports:

* `memory` -
  (Optional)
  Memory requirement expressed in Kubernetes resource units.

* `cpu` -
  (Optional)
  CPU requirement expressed in Kubernetes resource units.

<a name="nested_fleet_default_member_config_policycontroller_policy_controller_hub_config_deployment_configs_deployment_config_pod_toleration"></a>The `pod_toleration` block supports:

* `key` -
  (Optional)
  Matches a taint key (not necessarily unique).

* `operator` -
  (Optional)
  Matches a taint operator.

* `value` -
  (Optional)
  Matches a taint value.

* `effect` -
  (Optional)
  Matches a taint effect.

<a name="nested_fleet_default_member_config_policycontroller_policy_controller_hub_config_policy_content"></a>The `policy_content` block supports:

* `template_library` -
  (Optional)
  Configures the installation of the Template Library.
  Structure is [documented below](#nested_fleet_default_member_config_policycontroller_policy_controller_hub_config_policy_content_template_library).

* `bundles` -
  (Optional)
  Configures which bundles to install and their corresponding install specs.
  Structure is [documented below](#nested_fleet_default_member_config_policycontroller_policy_controller_hub_config_policy_content_bundles).


<a name="nested_fleet_default_member_config_policycontroller_policy_controller_hub_config_policy_content_template_library"></a>The `template_library` block supports:

* `installation` -
  (Optional)
  Configures the manner in which the template library is installed on the cluster.
  Possible values are: `INSTALLATION_UNSPECIFIED`, `NOT_INSTALLED`, `ALL`.

<a name="nested_fleet_default_member_config_policycontroller_policy_controller_hub_config_policy_content_bundles"></a>The `bundles` block supports:

* `bundle` - (Required) The identifier for this object. Format specified above.

* `exempted_namespaces` -
  (Optional)
  The set of namespaces to be exempted from the bundle.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/locations/{{location}}/features/{{name}}`

* `resource_state` -
  State of the Feature resource itself.
  Structure is [documented below](#nested_resource_state).

* `state` -
  Output only. The Hub-wide Feature state
  Structure is [documented below](#nested_state).

* `create_time` -
  Output only. When the Feature resource was created.

* `update_time` -
  Output only. When the Feature resource was last updated.

* `delete_time` -
  Output only. When the Feature resource was deleted.

* `terraform_labels` -
  The combination of labels configured directly on the resource
   and default labels configured on the provider.

* `effective_labels` -
  All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.


<a name="nested_resource_state"></a>The `resource_state` block contains:

* `state` -
  (Output)
  The current state of the Feature resource in the Hub API.

* `has_resources` -
  (Output)
  Whether this Feature has outstanding resources that need to be cleaned up before it can be disabled.

<a name="nested_state"></a>The `state` block contains:

* `state` -
  (Output)
  Output only. The "running state" of the Feature in this Hub.
  Structure is [documented below](#nested_state_state).


<a name="nested_state_state"></a>The `state` block contains:

* `code` -
  (Output)
  The high-level, machine-readable status of this Feature.

* `description` -
  (Output)
  A human-readable description of the current status.

* `update_time` -
  (Output)
  The time this status and any related Feature-specific details were updated. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z"

## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


Feature can be imported using any of these accepted formats:

* `projects/{{project}}/locations/{{location}}/features/{{name}}`
* `{{project}}/{{location}}/{{name}}`
* `{{location}}/{{name}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Feature using one of the formats above. For example:

```tf
import {
  id = "projects/{{project}}/locations/{{location}}/features/{{name}}"
  to = google_gke_hub_feature.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), Feature can be imported using one of the formats above. For example:

```
$ terraform import google_gke_hub_feature.default projects/{{project}}/locations/{{location}}/features/{{name}}
$ terraform import google_gke_hub_feature.default {{project}}/{{location}}/{{name}}
$ terraform import google_gke_hub_feature.default {{location}}/{{name}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
