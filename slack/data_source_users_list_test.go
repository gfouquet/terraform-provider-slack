package slack

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccSlackUsersListDataSource_basic(t *testing.T) {
	t.Parallel()
	dataSourceName := "data.slack_users_list.test"

	var providers []*schema.Provider

	t.Run("list all", func(t *testing.T) {
		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { testAccPreCheck(t) },
			ProviderFactories: testAccProviderFactories(&providers),
			Steps: []resource.TestStep{
				{
					Config: testAccCheckSlackUsersListDataSourceConfigListAll,
					Check: resource.ComposeTestCheckFunc(
						testAccCheckSlackUsersListDataSourceID(dataSourceName),
						resource.TestCheckResourceAttrSet(dataSourceName, "users.#"),
					),
				},
			},
		})
	})

	t.Run("list by email", func(t *testing.T) {
		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { testAccPreCheck(t) },
			ProviderFactories: testAccProviderFactories(&providers),
			Steps: []resource.TestStep{
				{
					Config: testAccCheckSlackUsersListDataSourceConfigListByEmail,
					Check: resource.ComposeTestCheckFunc(
						testAccCheckSlackUsersListDataSourceID(dataSourceName),
						resource.TestCheckResourceAttr(dataSourceName, "users.#", "1"),
						resource.TestCheckResourceAttr(dataSourceName, "users.0.name", testUser00.name),
						resource.TestCheckResourceAttr(dataSourceName, "users.0.email", testUser00.email),
						resource.TestCheckResourceAttr(dataSourceName, "users.0.id", testUser00.id),
					),
				},
			},
		})
	})
}

func testAccCheckSlackUsersListDataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("can't find slack users list datasource: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("slack  users list datasource id not set")
		}
		return nil
	}
}

const (
	testAccCheckSlackUsersListDataSourceConfigListAll = `
data slack_users_list test {
}
`
)

var (
	testAccCheckSlackUsersListDataSourceConfigListByEmail = fmt.Sprintf(`
data slack_users_list test {
  email = "%s"
}
`, testUser00.email)
)
