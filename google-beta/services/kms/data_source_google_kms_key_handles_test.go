// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package kms_test

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccDataSourceGoogleKmsKeyHandles_basic(t *testing.T) {
	kmsAutokey := acctest.BootstrapKMSAutokeyKeyHandle(t)
	keyParts := strings.Split(kmsAutokey.KeyHandle.Name, "/")
	project := keyParts[1]
	location := keyParts[3]
	diskFilter := fmt.Sprintf("compute.googleapis.com/Disk")

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGoogleKmsKeyHandles_basic(project, location, diskFilter),
				Check: resource.ComposeTestCheckFunc(
					validateKeyHandleName(
						"data.google_kms_key_handles.mykeyhandles", kmsAutokey.KeyHandle.Name,
					),
				),
			},
		},
	})
}
func validateKeyHandleName(dataSourceName string, expectedKeyHandleName string) func(*terraform.State) error {
	return func(s *terraform.State) error {
		ds, ok := s.RootModule().Resources[dataSourceName]
		if !ok {
			return fmt.Errorf("can't find %s in state", dataSourceName)
		}

		var dsAttr map[string]string
		dsAttr = ds.Primary.Attributes

		totalKeyHandles, err := strconv.Atoi(dsAttr["key_handles.#"])
		if err != nil {
			return errors.New("Couldn't convert length of key_handles list to integer")
		}
		if totalKeyHandles != 1 {
			return errors.New(fmt.Sprintf("want 1 keyhandle, found %d", totalKeyHandles))
		}
		actualKeyHandleName := dsAttr["key_handles.0.name"]
		if actualKeyHandleName != expectedKeyHandleName {
			return errors.New(fmt.Sprintf("want keyhandle name %s, got: %s", expectedKeyHandleName, actualKeyHandleName))
		}
		return nil
	}
}
func testAccDataSourceGoogleKmsKeyHandles_basic(project string, location string, filter string) string {
	str := fmt.Sprintf(`
data "google_kms_key_handles" "mykeyhandles" {
  project = "%s"
  location = "%s"
  resource_type_selector = "%s"
}
`, project, location, filter)
	return str
}