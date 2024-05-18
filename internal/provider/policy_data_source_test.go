package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccPolicyDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: testCspPolicyDataSourceSimple,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.csp_policy.test", "value", testCspPolicyDataSourceSimpleValue),
				),
			},
		},
	})
}

const testCspPolicyDataSourceSimple = `
data "csp_policy" "test" {
	directive {
		name     = "default-src"
		keywords = ["self"]
	}
	directive {
		name     = "img-src"
		keywords = ["self"]
		hosts    = ["cdn.example.com"]
	}
}
`
const testCspPolicyDataSourceSimpleValue = "default-src 'self'; img-src 'self' cdn.example.com;"
