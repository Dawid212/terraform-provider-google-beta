// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package fwprovider_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

// TestAccFwProvider_user_project_override is a series of acc tests asserting how the plugin-framework provider handles credentials arguments
// It is PF specific because the HCL used uses a PF-implemented data source
// It is a counterpart to TestAccSdkProvider_user_project_override
func TestAccFwProvider_user_project_override(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		// Configuring the provider using inputs
		"config takes precedence over environment variables":                                testAccFwProvider_user_project_override_configPrecedenceOverEnvironmentVariables,
		"when user_project_override is unset in the config, environment variables are used": testAccFwProvider_user_project_override_precedenceOrderEnvironmentVariables,

		// Schema-level validation
		"when user_project_override is set in the config the value can be a boolean (true/false) or a string (true/false/1/0)": testAccFwProvider_user_project_override_booleansInConfigOnly,
		"when user_project_override is set via environment variables any of these values can be used: true/false/1/0":          testAccFwProvider_user_project_override_envStringsAccepted,

		// Usage
		"user_project_override uses a resource's project argument to control which project is used for quota and billing purposes": testAccFwProvider_UserProjectOverride,
		// We cannot currently implement this test case in a PF-specific way: "user_project_override works for resources that don't take a project argument (provider-level default project value is used)"
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

func testAccFwProvider_user_project_override_configPrecedenceOverEnvironmentVariables(t *testing.T) {
	acctest.SkipIfVcr(t) // Test doesn't interact with API

	override := "true"
	providerOverride := false

	// ensure all possible region env vars set; show they aren't used
	t.Setenv("USER_PROJECT_OVERRIDE", override)

	context := map[string]interface{}{
		"user_project_override": providerOverride,
	}

	acctest.VcrTest(t, resource.TestCase{
		// No PreCheck for checking ENVs
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Apply-time error; bad value in config is used over of good values in ENVs
				Config: testAccFwProvider_user_project_overrideInProviderBlock_boolean(context),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.google_provider_config_plugin_framework.default", "user_project_override", fmt.Sprintf("%v", providerOverride)),
				),
			},
		},
	})
}

func testAccFwProvider_user_project_override_precedenceOrderEnvironmentVariables(t *testing.T) {
	acctest.SkipIfVcr(t) // Test doesn't interact with API
	/*
		These are all the ENVs for region, and they are in order of precedence.
		USER_PROJECT_OVERRIDE
	*/

	context := map[string]interface{}{}

	acctest.VcrTest(t, resource.TestCase{
		// No PreCheck for checking ENVs
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					t.Setenv("USER_PROJECT_OVERRIDE", "") // unset
				},
				Config: testAccFwProvider_user_project_overrideInEnvsOnly(context),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("data.google_provider_config_plugin_framework.default", "user_project_override"),
				),
			},
			{
				PreConfig: func() {
					t.Setenv("USER_PROJECT_OVERRIDE", "true")
				},
				Config: testAccFwProvider_user_project_overrideInEnvsOnly(context),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.google_provider_config_plugin_framework.default", "user_project_override", "true"),
				),
			},
		},
	})
}

func testAccFwProvider_user_project_override_booleansInConfigOnly(t *testing.T) {
	acctest.SkipIfVcr(t) // Test doesn't interact with API

	context_true := map[string]interface{}{
		"user_project_override": true,
	}
	context_false := map[string]interface{}{
		"user_project_override": false,
	}

	context_1 := map[string]interface{}{
		"user_project_override": "1",
	}
	context_0 := map[string]interface{}{
		"user_project_override": "0",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFwProvider_user_project_overrideInProviderBlock_boolean(context_true),
				// No error expected
			},
			{
				Config: testAccFwProvider_user_project_overrideInProviderBlock_boolean(context_false),
				// No error expected
			},
			{
				Config: testAccFwProvider_user_project_overrideInProviderBlock_string(context_true),
				// No error expected
			},
			{
				Config: testAccFwProvider_user_project_overrideInProviderBlock_string(context_false),
				// No error expected
			},
			{
				Config: testAccFwProvider_user_project_overrideInProviderBlock_string(context_1),
				// No error expected
			},
			{
				Config: testAccFwProvider_user_project_overrideInProviderBlock_string(context_0),
				// No error expected
			},
		},
	})
}

func testAccFwProvider_user_project_override_envStringsAccepted(t *testing.T) {
	acctest.SkipIfVcr(t) // Test doesn't interact with API

	context := map[string]interface{}{}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					t.Setenv("USER_PROJECT_OVERRIDE", "true")
				},
				Config: testAccFwProvider_user_project_overrideInEnvsOnly(context),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.google_provider_config_plugin_framework.default", "user_project_override", "true"),
				),
			},
			{
				PreConfig: func() {
					t.Setenv("USER_PROJECT_OVERRIDE", "1")
				},
				Config: testAccFwProvider_user_project_overrideInEnvsOnly(context),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.google_provider_config_plugin_framework.default", "user_project_override", "true"),
				),
			},
			{
				PreConfig: func() {
					t.Setenv("USER_PROJECT_OVERRIDE", "false")
				},
				Config: testAccFwProvider_user_project_overrideInEnvsOnly(context),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.google_provider_config_plugin_framework.default", "user_project_override", "false"),
				),
			},
			{
				PreConfig: func() {
					t.Setenv("USER_PROJECT_OVERRIDE", "0")
				},
				Config: testAccFwProvider_user_project_overrideInEnvsOnly(context),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.google_provider_config_plugin_framework.default", "user_project_override", "false"),
				),
			},
		},
	})
}

// testAccFwProvider_user_project_overrideInProviderBlock allows setting the user_project_override argument in a provider block.
// This function uses data.google_provider_config_plugin_framework because it is implemented with the plugin-framework
func testAccFwProvider_user_project_overrideInProviderBlock_boolean(context map[string]interface{}) string {
	v := acctest.Nprintf(`
provider "google" {
	user_project_override = %{user_project_override}
}

data "google_provider_config_plugin_framework" "default" {}
`, context)
	return v
}

func testAccFwProvider_user_project_overrideInProviderBlock_string(context map[string]interface{}) string {
	return acctest.Nprintf(`
provider "google" {
	user_project_override = "%{user_project_override}"
}

data "google_provider_config_plugin_framework" "default" {}
`, context)
}

// testAccFwProvider_user_project_overrideInEnvsOnly allows testing when the user_project_override argument
// is only supplied via ENVs
func testAccFwProvider_user_project_overrideInEnvsOnly(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_provider_config_plugin_framework" "default" {}
`, context)
}

func testAccFwProvider_UserProjectOverride(t *testing.T) {
	// Parallel fine-grained resource creation
	acctest.SkipIfVcr(t)
	t.Parallel()

	org := envvar.GetTestOrgFromEnv(t)
	billing := envvar.GetTestBillingAccountFromEnv(t)
	pid := "tf-test-" + acctest.RandString(t, 10)

	config := acctest.BootstrapConfig(t)
	accessToken, err := acctest.SetupProjectsAndGetAccessToken(org, billing, pid, "firebase", config)
	if err != nil || accessToken == "" {
		if err == nil {
			t.Fatal("error when setting up projects and retrieving access token: access token is an empty string")
		}
		t.Fatalf("error when setting up projects and retrieving access token: %s", err)
	}

	context := map[string]interface{}{
		"user_project_override": false, // changed in config functions
		"access_token":          accessToken,
		"project_2":             pid + "-2", // acctest.SetupProjectsAndGetAccessToken creates project 1 with id = `pid`, and 2 = `pid` + "-2"
		"random_suffix":         acctest.RandString(t, 5),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		// No TestDestroy since that's not really the point of this test
		Steps: []resource.TestStep{
			{
				Config:      testAccFwProvider_UserProjectOverride_step2(context, false),
				ExpectError: regexp.MustCompile("Firebase Management API has not been used"),
			},
			{
				Config: testAccFwProvider_UserProjectOverride_step2(context, true),
				Check: resource.ComposeTestCheckFunc(
					// Firebase resources and data sources are linked to project-2, where the API is enabled
					resource.TestCheckResourceAttr("google_firebase_apple_app.project_2_app", "project", context["project_2"].(string)),
					resource.TestCheckResourceAttr("data.google_firebase_apple_app_config.project_2_app", "project", context["project_2"].(string)),
				),
			},
			{
				Config: testAccFwProvider_UserProjectOverride_step3(context, true),
			},
		},
	})
}

func testAccFwProvider_UserProjectOverride_step2(context map[string]interface{}, override bool) string {
	return testAccFwProvider_UserProjectOverride_step3(context, override) +
		acctest.Nprintf(`
// See step 3 appended above, which is really step 2 minus the pubsub topic.
// Step 3 exists because provider configurations can't be removed while objects
// created by that provider still exist in state. Step 3 will remove the
// pubsub topic so the whole config can be deleted.

resource "google_firebase_project" "default" {
  provider = google.project-1-token
  project  = "%{project_2}"
}

resource "google_firebase_apple_app" "project_2_app" {
  provider = google.project-1-token
  project  = google_firebase_project.default.project // add dependency, also uses project 2
  bundle_id = "apple.app.%{random_suffix}"
  display_name = "tf-test Display Name AppleAppConfig DataSource"
  app_store_id = "12345"
  team_id = "1234567890"
}

// This is implemented with plugin-framework so tests our use of user_project_override in a PF specific way
data "google_firebase_apple_app_config" "project_2_app" {
  project = google_firebase_apple_app.project_2_app.project
  app_id = google_firebase_apple_app.project_2_app.app_id
}
`, context)
}

func testAccFwProvider_UserProjectOverride_step3(context map[string]interface{}, override bool) string {
	context["user_project_override"] = override

	return acctest.Nprintf(`
provider "google" {
	alias  = "project-1-token"
	access_token = "%{access_token}"
	user_project_override = "%{user_project_override}"
}
`, context)
}