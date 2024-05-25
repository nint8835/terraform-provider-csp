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
			{
				Config: testCspPolicyDataSourceGithub,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.csp_policy.github", "value", testCspPolicyDataSourceGithubValue),
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
const testCspPolicyDataSourceSimpleValue = "default-src 'self'; img-src 'self' cdn.example.com"

const testCspPolicyDataSourceGithub = `
data "csp_policy" "github" {
	directive {
		name     = "default-src"
		keywords = ["none"]
	}

	directive {
		name     = "base-uri"
		keywords = ["self"]
	}

	directive {
		name     = "child-src"
		hosts    = [
			"github.com/assets-cdn/worker/",
			"gist.github.com/assets-cdn/worker/",
		]
	}

	directive {
		name     = "connect-src"
		keywords = ["self"]
		hosts    = [
			"uploads.github.com",
			"www.githubstatus.com",
			"collector.github.com",
			"raw.githubusercontent.com",
			"api.github.com",
			"github-cloud.s3.amazonaws.com",
			"github-production-repository-file-5c1aeb.s3.amazonaws.com",
			"github-production-upload-manifest-file-7fdce7.s3.amazonaws.com",
			"github-production-user-asset-6210df.s3.amazonaws.com",
			"api.githubcopilot.com",
			"objects-origin.githubusercontent.com",
			"copilot-proxy.githubusercontent.com/v1/engines/copilot-codex/completions",
			"*.actions.githubusercontent.com",
			"wss://*.actions.githubusercontent.com",
			"productionresultssa0.blob.core.windows.net/",
			"productionresultssa1.blob.core.windows.net/",
			"productionresultssa2.blob.core.windows.net/",
			"productionresultssa3.blob.core.windows.net/",
			"productionresultssa4.blob.core.windows.net/",
			"productionresultssa5.blob.core.windows.net/",
			"productionresultssa6.blob.core.windows.net/",
			"productionresultssa7.blob.core.windows.net/",
			"productionresultssa8.blob.core.windows.net/",
			"productionresultssa9.blob.core.windows.net/",
			"productionresultssa10.blob.core.windows.net/",
			"productionresultssa11.blob.core.windows.net/",
			"productionresultssa12.blob.core.windows.net/",
			"productionresultssa13.blob.core.windows.net/",
			"productionresultssa14.blob.core.windows.net/",
			"productionresultssa15.blob.core.windows.net/",
			"productionresultssa16.blob.core.windows.net/",
			"productionresultssa17.blob.core.windows.net/",
			"productionresultssa18.blob.core.windows.net/",
			"productionresultssa19.blob.core.windows.net/",
			"github-production-repository-image-32fea6.s3.amazonaws.com",
			"github-production-release-asset-2e65be.s3.amazonaws.com",
			"insights.github.com",
			"wss://alive.github.com",
		]
	}

	directive {
		name     = "font-src"
		hosts    = ["github.githubassets.com"]
	}

	directive {
		name     = "form-action"
		keywords = ["self"]
		hosts    = [
			"github.com",
			"gist.github.com",
			"copilot-workspace.githubnext.com",
			"objects-origin.githubusercontent.com",
		]
	}

	directive {
		name     = "frame-ancestors"
		keywords = ["none"]
	}

	directive {
		name     = "frame-src"
		hosts    = [
			"viewscreen.githubusercontent.com",
			"notebooks.githubusercontent.com",
		]
	}

	directive {
		name     = "img-src"
		keywords = ["self"]
		schemes  = ["data"]
		hosts    = [
			"github.githubassets.com",
			"media.githubusercontent.com",
			"camo.githubusercontent.com",
			"identicons.github.com",
			"avatars.githubusercontent.com",
			"github-cloud.s3.amazonaws.com",
			"objects.githubusercontent.com",
			"secured-user-images.githubusercontent.com/",
			"user-images.githubusercontent.com/",
			"private-user-images.githubusercontent.com",
			"opengraph.githubassets.com",
			"github-production-user-asset-6210df.s3.amazonaws.com",
			"customer-stories-feed.github.com",
			"spotlights-feed.github.com",
			"objects-origin.githubusercontent.com",
			"*.githubusercontent.com",
		]
	}

	directive {
		name     = "manifest-src"
		keywords = ["self"]
	}

	directive {
		name     = "media-src"
		hosts    = [
			"github.com",
			"user-images.githubusercontent.com/",
			"secured-user-images.githubusercontent.com/",
			"private-user-images.githubusercontent.com",
			"github-production-user-asset-6210df.s3.amazonaws.com",
			"gist.github.com",
		]
	}

	directive {
		name     = "script-src"
		hosts    = ["github.githubassets.com"]
	}

	directive {
		name     = "style-src"
		keywords = ["unsafe-inline"]
		hosts    = ["github.githubassets.com"]
	}

	directive {
		name     = "upgrade-insecure-requests"
	}

	directive {
		name     = "worker-src"
		hosts    = [
			"github.com/assets-cdn/worker/",
			"gist.github.com/assets-cdn/worker/",
		]
	}
}
`

const testCspPolicyDataSourceGithubValue = "default-src 'none'; base-uri 'self'; child-src github.com/assets-cdn/worker/ gist.github.com/assets-cdn/worker/; connect-src 'self' uploads.github.com www.githubstatus.com collector.github.com raw.githubusercontent.com api.github.com github-cloud.s3.amazonaws.com github-production-repository-file-5c1aeb.s3.amazonaws.com github-production-upload-manifest-file-7fdce7.s3.amazonaws.com github-production-user-asset-6210df.s3.amazonaws.com api.githubcopilot.com objects-origin.githubusercontent.com copilot-proxy.githubusercontent.com/v1/engines/copilot-codex/completions *.actions.githubusercontent.com wss://*.actions.githubusercontent.com productionresultssa0.blob.core.windows.net/ productionresultssa1.blob.core.windows.net/ productionresultssa2.blob.core.windows.net/ productionresultssa3.blob.core.windows.net/ productionresultssa4.blob.core.windows.net/ productionresultssa5.blob.core.windows.net/ productionresultssa6.blob.core.windows.net/ productionresultssa7.blob.core.windows.net/ productionresultssa8.blob.core.windows.net/ productionresultssa9.blob.core.windows.net/ productionresultssa10.blob.core.windows.net/ productionresultssa11.blob.core.windows.net/ productionresultssa12.blob.core.windows.net/ productionresultssa13.blob.core.windows.net/ productionresultssa14.blob.core.windows.net/ productionresultssa15.blob.core.windows.net/ productionresultssa16.blob.core.windows.net/ productionresultssa17.blob.core.windows.net/ productionresultssa18.blob.core.windows.net/ productionresultssa19.blob.core.windows.net/ github-production-repository-image-32fea6.s3.amazonaws.com github-production-release-asset-2e65be.s3.amazonaws.com insights.github.com wss://alive.github.com; font-src github.githubassets.com; form-action 'self' github.com gist.github.com copilot-workspace.githubnext.com objects-origin.githubusercontent.com; frame-ancestors 'none'; frame-src viewscreen.githubusercontent.com notebooks.githubusercontent.com; img-src 'self' data: github.githubassets.com media.githubusercontent.com camo.githubusercontent.com identicons.github.com avatars.githubusercontent.com github-cloud.s3.amazonaws.com objects.githubusercontent.com secured-user-images.githubusercontent.com/ user-images.githubusercontent.com/ private-user-images.githubusercontent.com opengraph.githubassets.com github-production-user-asset-6210df.s3.amazonaws.com customer-stories-feed.github.com spotlights-feed.github.com objects-origin.githubusercontent.com *.githubusercontent.com; manifest-src 'self'; media-src github.com user-images.githubusercontent.com/ secured-user-images.githubusercontent.com/ private-user-images.githubusercontent.com github-production-user-asset-6210df.s3.amazonaws.com gist.github.com; script-src github.githubassets.com; style-src 'unsafe-inline' github.githubassets.com; upgrade-insecure-requests; worker-src github.com/assets-cdn/worker/ gist.github.com/assets-cdn/worker/"
