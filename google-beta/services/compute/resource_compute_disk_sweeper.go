// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
// ----------------------------------------------------------------------------
//
//	***     AUTO GENERATED CODE    ***    Type: Handwritten     ***
//
// ----------------------------------------------------------------------------
//
//	This code is generated by Magic Modules using the following:
//
//	Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/services/compute/resource_compute_disk_sweeper.go
//
//	DO NOT EDIT this file directly. Any changes made to this file will be
//	overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------
package compute

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/sweeper"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// This will sweep GCE Disk resources
func init() {
	sweeper.AddTestSweepersLegacy("ComputeDisk", testSweepDisk)
}

// At the time of writing, the CI only passes us-central1 as the region
func testSweepDisk(region string) error {
	resourceName := "ComputeDisk"
	log.Printf("[INFO][SWEEPER_LOG] Starting sweeper for %s", resourceName)

	config, err := sweeper.SharedConfigForRegion(region)
	if err != nil {
		log.Printf("[INFO][SWEEPER_LOG] error getting shared config for region: %s", err)
		return err
	}

	err = config.LoadAndValidate(context.Background())
	if err != nil {
		log.Printf("[INFO][SWEEPER_LOG] error loading: %s", err)
		return err
	}

	zones := []string{"us-central1-a", "us-central1-b", "us-central1-c", "us-central1-f", "us-east1-b", "us-east1-c", "us-east1-d", "us-west1-a", "us-west1-b", "us-west1-c"}
	for _, zone := range zones {
		servicesUrl := "https://compute.googleapis.com/compute/v1/projects/" + config.Project + "/zones/" + zone + "/disks"
		// Page zero's URL is the raw list URL. Successive pages will return the token for the next page.
		pageUrl := servicesUrl
		for {

			res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   config.Project,
				RawURL:    pageUrl,
				UserAgent: config.UserAgent,
			})
			if err != nil {
				log.Printf("[INFO][SWEEPER_LOG] Error in response from request %s: %s", pageUrl, err)
				return nil
			}

			resourceList, ok := res["items"]
			if !ok {
				log.Printf("[INFO][SWEEPER_LOG] Nothing found in response.")
				return nil
			}

			rl := resourceList.([]interface{})

			log.Printf("[INFO][SWEEPER_LOG] Found %d items in %s list response.", len(rl), resourceName)
			// Count items that weren't sweeped.
			nonPrefixCount := 0
			for _, ri := range rl {
				obj := ri.(map[string]interface{})
				if obj["id"] == nil {
					log.Printf("[INFO][SWEEPER_LOG] %s resource id was nil", resourceName)
					return nil
				}

				id := obj["name"].(string)
				// Increment count and skip if resource is not sweepable.
				prefixes := []string{
					"pvc-", // b/291168201
				}
				if !sweeper.IsSweepableTestResource(id) && !sweeper.HasAnyPrefix(id, prefixes) {
					nonPrefixCount++
					continue
				}

				deleteUrl := servicesUrl + "/" + id
				// Don't wait on operations as we may have a lot to delete
				_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
					Config:    config,
					Method:    "DELETE",
					Project:   config.Project,
					RawURL:    deleteUrl,
					UserAgent: config.UserAgent,
				})
				if err != nil {
					log.Printf("[INFO][SWEEPER_LOG] Error deleting for url %s : %s", deleteUrl, err)
				} else {
					log.Printf("[INFO][SWEEPER_LOG] Sent delete request for %s resource: %s", resourceName, id)
				}
			}

			if nonPrefixCount > 0 {
				log.Printf("[INFO][SWEEPER_LOG] %d items without tf-test prefix remain for zone %s", nonPrefixCount, zone)
			}

			if res["nextPageToken"] == nil || res["nextPageToken"].(string) == "" {
				break
			}
			pageUrl, err = transport_tpg.AddQueryParams(servicesUrl, map[string]string{"pageToken": res["nextPageToken"].(string)})
		}

	}

	return nil
}
