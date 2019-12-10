package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-vra7/utils"
	"github.com/terraform-providers/terraform-provider-vra7/vra7"
)

func main() {
	utils.InitLog()
	opts := plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return vra7.Provider()
		},
	}

	plugin.Serve(&opts)
}
