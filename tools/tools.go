//go:build tools

// Package tools pins developer tooling as module dependencies so `go generate`
// and CI use a consistent version. Run `go generate ./...` to regenerate docs.
package tools

import (
	// Documentation generation for the Terraform Registry.
	_ "github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs"
)
