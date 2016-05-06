package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/mdb/terraform-provider-example"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: example.Provider,
	})
}
