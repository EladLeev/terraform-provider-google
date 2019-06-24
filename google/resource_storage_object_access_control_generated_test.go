// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccStorageObjectAccessControl_storageObjectAccessControlPublicObjectExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckStorageObjectAccessControlDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccStorageObjectAccessControl_storageObjectAccessControlPublicObjectExample(context),
			},
			{
				ResourceName:      "google_storage_object_access_control.public_rule",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccStorageObjectAccessControl_storageObjectAccessControlPublicObjectExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_object_access_control" "public_rule" {
  object = "${google_storage_bucket_object.object.output_name}"
  bucket = "${google_storage_bucket.bucket.name}"
  role   = "READER"
  entity = "allUsers"
}

resource "google_storage_bucket" "bucket" {
	name = "static-content-bucket-%{random_suffix}"
}

 resource "google_storage_bucket_object" "object" {
	name   = "public-object-%{random_suffix}"
	bucket = "${google_storage_bucket.bucket.name}"
	source = "test-fixtures/header-logo.png"
}
`, context)
}

func testAccCheckStorageObjectAccessControlDestroy(s *terraform.State) error {
	for name, rs := range s.RootModule().Resources {
		if rs.Type != "google_storage_object_access_control" {
			continue
		}
		if strings.HasPrefix(name, "data.") {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(config, rs, "{{StorageBasePath}}b/{{bucket}}/o/{{object}}/acl/{{entity}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", url, nil)
		if err == nil {
			return fmt.Errorf("StorageObjectAccessControl still exists at %s", url)
		}
	}

	return nil
}
