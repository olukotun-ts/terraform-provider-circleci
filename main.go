package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"github.com/olukotun-ts/terraform-provider-circleci/circleci"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: circleci.Provider,
	})
}
