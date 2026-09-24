// Package main is the entry point for the example Terraform provider.
//
// This is a generic, vendor-neutral boilerplate extracted from a
// terraform-plugin-framework provider. Replace "example" throughout with your
// own provider name and wire in your API client and resources.
package main

import (
	"context"
	"flag"
	"log"

	"terraform-provider-eshiam/internal/provider"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

// version is set by the release tooling (e.g. goreleaser) at build time.
var version = "dev"

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		// Update this to your registry namespace, e.g.
		// registry.terraform.io/<namespace>/<name>.
		Address: "registry.terraform.io/eshiam-corp/eshiam",
		Debug:   debug,
	}

	if err := providerserver.Serve(context.Background(), provider.New(version), opts); err != nil {
		log.Fatal(err.Error())
	}
}
