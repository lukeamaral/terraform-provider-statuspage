package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/lukeamaral/terraform-provider-statuspage/statuspage"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: statuspage.Provider,
	})
}
