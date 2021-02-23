package huaweicloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIdentityRoleDataSource_basic(t *testing.T) {
	resourceName := "data.huaweicloud_identity_role.role_1"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAdminOnly(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIdentityRoleDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIdentityDataSourceID(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "secu_admin"),
				),
			},
		},
	})
}

func testAccCheckIdentityDataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find role data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Role data source ID not set")
		}

		return nil
	}
}

const testAccIdentityRoleDataSource_basic = `
data "huaweicloud_identity_role" "role_1" {
  name = "secu_admin"
}
`
