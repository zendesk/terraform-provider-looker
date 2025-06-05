package main

import (
	"github.com/zendesk/terraform-provider-looker/internal/provider"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary
	version string = "0.0.0"

	// goreleaser can also pass the specific commit if you want
	// commit  string = ""
)

func main() {
	var debugMode bool

	// flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	// flag.Parse()

	plugin.Serve(&plugin.ServeOpts{
		Debug: debugMode,

		// TODO: update this string with the full name of your provider as used in your configs
		ProviderAddr: "zdrepo.jfrog.io/artifactory/terraform/zendesk/looker",

		ProviderFunc: provider.New(version),
	})
}
