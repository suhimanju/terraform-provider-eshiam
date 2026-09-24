# Setup & installation guide — Go Terraform provider framework

A complete, first-time setup guide for the **Go Terraform provider framework**
(`terraform-provider-eshiam`). Written from the framework's actual code —
`main.go`, `go.mod`, and `internal/provider/provider.go` — so the addresses,
environment variables, and commands here match what the provider really does.

> **New here?** The end goal is: build the provider binary, tell Terraform to
> use your local build (a "dev override"), set three env vars, and run
> `terraform plan`. Sections 2–8 get you there.

**Contents**
1. [What you are building](#1-what-you-are-building)
2. [Prerequisites](#2-prerequisites)
3. [Get the code](#3-get-the-code)
4. [Build & install the provider](#4-build--install-the-provider)
5. [Point Terraform at your local build (dev overrides)](#5-point-terraform-at-your-local-build-dev-overrides)
6. [Provide credentials (EXAMPLE_* env vars)](#6-provide-credentials-example_-env-vars)
7. [Write a tiny test configuration](#7-write-a-tiny-test-configuration)
8. [Run it](#8-run-it)
9. [Run against the offline mock API](#9-run-against-the-offline-mock-api)
10. [Install mode B — Docker build/extract](#10-install-mode-b--docker-buildextract)
11. [Install mode C — consume from a registry / mirror](#11-install-mode-c--consume-from-a-registry--mirror)
12. [Rename the placeholders for your own provider](#12-rename-the-placeholders-for-your-own-provider)
13. [Editor / IDE setup](#13-editor--ide-setup)
14. [Troubleshooting](#14-troubleshooting)

---

## 1. What you are building

| Piece | What it is | Where |
|-------|------------|-------|
| `terraform-provider-eshiam` | A Terraform **provider** binary (plugin) | built from `main.go` |
| Provider address | `registry.terraform.io/eshiam-corp/eshiam` | set in `main.go` → `ServeOpts.Address` |
| Provider type name | `example` | set in `provider.go` → `Metadata` |
| Resources / data sources | 26 resources + 4 data sources | registered in `provider.go` |

The provider speaks **terraform-plugin-framework** (Protocol 6) and
authenticates with an **OAuth2 client-credentials** flow
(`host` + `client_id` + `client_secret`, with env-var fallback).

You have **three install modes**:

| Mode | Best for | Section |
|------|----------|---------|
| **A. Local build + dev override** | Development, trying it out | [4](#4-build--install-the-provider)–[8](#8-run-it) |
| **B. Docker build/extract** | No local Go toolchain, CI | [10](#10-install-mode-b--docker-buildextract) |
| **C. Registry / network mirror** | Teams consuming a published build | [11](#11-install-mode-c--consume-from-a-registry--mirror) |

---

## 2. Prerequisites

| Tool | Version | Why | Check |
|------|---------|-----|-------|
| **Go** | **1.25+** (see `go.mod` → `go 1.25.0`) | Compiles the provider | `go version` |
| **Terraform** | **1.5+** | Runs the provider | `terraform version` |
| **git** | any | Clone the repo | `git --version` |
| golangci-lint | latest | Optional: linting | `golangci-lint --version` |
| Docker | 20.10+ | Only for install mode B | `docker --version` |

### Installing Go

- **Windows:** download the MSI from [go.dev/dl](https://go.dev/dl/); it adds Go
  to PATH. Open a **new** terminal and run `go version`.
- **macOS:** `brew install go` (or the go.dev pkg installer).
- **Linux:** follow [go.dev/doc/install](https://go.dev/doc/install) (untar into
  `/usr/local/go`, add `/usr/local/go/bin` to PATH).

### Installing Terraform

- **Windows:** `choco install terraform` or unzip the binary and add it to PATH.
- **macOS:** `brew tap hashicorp/tap && brew install hashicorp/tap/terraform`.
- **Linux:** use HashiCorp's apt/yum repo (see
  [developer.hashicorp.com/terraform/install](https://developer.hashicorp.com/terraform/install)).

Verify both are on PATH before continuing:

```bash
go version
terraform version
```

> **Network note:** the first build downloads Go modules. Behind a proxy set
> `GOPROXY` (default `https://proxy.golang.org,direct`) and `GOFLAGS=-mod=mod`
> if module resolution is restricted.

---

## 3. Get the code

```bash
git clone <your-fork-or-repo-url> go-terraform-framework
cd go-terraform-framework
```

---

## 4. Build & install the provider

**Compile everything** (no output = success):

```bash
go build ./...
```

**Install the binary** into your Go bin directory so the dev override can find
it:

```bash
go install .
```

Find where it landed — you need this path next:

```bash
go env GOBIN     # if empty, the path is: <go env GOPATH>\bin
```

On Windows that's typically `C:\Users\<you>\go\bin\terraform-provider-eshiam.exe`.

---

## 5. Point Terraform at your local build (dev overrides)

A **dev override** tells Terraform to use your freshly built binary instead of
downloading the provider from a registry.

Create a Terraform CLI config file. Two equivalent approaches:

**A) Use the default location**
- **Windows:** `%APPDATA%\terraform.rc`
- **Linux/macOS:** `~/.terraformrc`

**B) Point at a repo-local file (no global state)**

```powershell
# Windows PowerShell
$env:TF_CLI_CONFIG_FILE = Join-Path (Get-Location) ".terraformrc"
```
```bash
# Linux/macOS
export TF_CLI_CONFIG_FILE="$PWD/.terraformrc"
```

Put this inside the file, replacing the path with your Go bin folder from
Section 4:

```hcl
provider_installation {
  dev_overrides {
    # Linux/macOS:
    "eshiam-corp/eshiam" = "/home/you/go/bin"
    # Windows (double backslashes):
    # "eshiam-corp/eshiam" = "C:\\Users\\you\\go\\bin"
  }
  # Everything else installs normally.
  direct {}
}
```

> A ready-to-edit template is at
> `examples/provider-install-verification/.terraformrc`.
> With a dev override active, **do not run `terraform init`** — overrides skip
> it. Terraform prints a warning that an override is in effect; that is expected.

---

## 6. Provide credentials (EXAMPLE_* env vars)

`provider.go` → `Configure` reads these environment variables as defaults
(explicit values in the provider block override them):

| Variable | Meaning |
|----------|---------|
| `EXAMPLE_HOST` | Base URL of the API (token URL is `host + /oauth/token`) |
| `EXAMPLE_CLIENT_ID` | OAuth2 client ID |
| `EXAMPLE_CLIENT_SECRET` | OAuth2 client secret (sensitive) |

**Windows (PowerShell):**

```powershell
$env:EXAMPLE_HOST          = "https://api.example.com"
$env:EXAMPLE_CLIENT_ID     = "your-client-id"
$env:EXAMPLE_CLIENT_SECRET = "your-client-secret"
```

**macOS / Linux:**

```bash
export EXAMPLE_HOST="https://api.example.com"
export EXAMPLE_CLIENT_ID="your-client-id"
export EXAMPLE_CLIENT_SECRET="your-client-secret"
```

> Prefer env vars over hard-coding secrets in `.tf` files. If any of the three
> is missing, `Configure` returns a clear "Missing API ..." diagnostic naming
> the field.

---

## 7. Write a tiny test configuration

Create a throwaway folder (or reuse `examples/provider-install-verification/`):

```hcl
# main.tf
terraform {
  required_providers {
    example = {
      source = "eshiam-corp/eshiam"
    }
  }
}

provider "example" {
  # host/client_id/client_secret fall back to EXAMPLE_* env vars (Section 6).
}

# Add a resource once your API is reachable, e.g.:
# resource "example_transform" "demo" {
#   name = "demo"
# }
```

---

## 8. Run it

```bash
# Do NOT run `terraform init` when a dev override is active.
terraform plan
terraform apply
```

Terraform loads your local binary, calls `Configure` (which builds the OAuth2
client), and executes the plan. The warning about dev overrides is normal.

**Sanity-check the provider without an API** using the verification example:

```bash
cd examples/provider-install-verification
terraform plan   # should load the provider and report no resources
```

---

## 9. Run against the offline mock API

You don't need a live tenant to exercise the provider. The repo ships a mock API
definition under `mock/`. See `mock/README.md` for how to start it (e.g. with
Mockoon) and point `EXAMPLE_HOST` at `http://localhost:<port>`.

Unit tests need no network at all:

```bash
go test ./...
```

---

## 10. Install mode B — Docker build/extract

When you don't have a Go toolchain locally, build the binary in a container and
extract it. Full details in [docker.md](docker.md); the essentials:

```bash
# Build the provider inside an image:
docker build -t terraform-provider-eshiam:latest .

# Extract the compiled binary to ./out on the host:
docker compose run --rm extract     # writes ./out/terraform-provider-eshiam

# Then use that path in your dev_overrides (Section 5) and run terraform plan.
```

This is also how CI produces artifacts without installing Go on the runner.

---

## 11. Install mode C — consume from a registry / mirror

For teams, publish the provider and let Terraform download it normally — no dev
override needed:

- **Public/private Terraform Registry:** declare it in `required_providers` and
  run `terraform init`. Publishing is covered step-by-step in
  [publishing-to-the-registry.md](publishing-to-the-registry.md).
- **Network mirror (e.g. internal Artifactory):** configure a
  `network_mirror` block in `.terraformrc` and exclude the provider from
  `direct` so it resolves from the mirror.

```hcl
provider_installation {
  network_mirror { url = "https://artifacts.example.com/api/terraform/providers/" }
  direct { exclude = ["registry.terraform.io/eshiam-corp/eshiam"] }
}
```

---

## 12. Rename the placeholders for your own provider

This repo is a **generic boilerplate**. Before publishing, replace the
placeholders everywhere:

| Placeholder | Replace with | Set in |
|-------------|--------------|--------|
| `terraform-provider-eshiam` | `terraform-provider-<name>` | module name, repo name |
| `example` (type) | your type | `provider.go` → `Metadata` |
| `eshiam-corp/eshiam` | `<namespace>/<name>` | `main.go` → `ServeOpts.Address` |
| `EXAMPLE_HOST/CLIENT_ID/CLIENT_SECRET` | your `UPPER_*` vars | `provider.go` → `Configure` |

The GitHub repo **must** be named `terraform-provider-<name>` for the registry.
See [publishing-to-the-registry.md](publishing-to-the-registry.md) for the full
rename + release checklist.

---

## 13. Editor / IDE setup

- **VS Code:** install the *Go* extension; accept its prompt to install
  `gopls`, `dlv` (debugger), etc. The `.editorconfig` enforces tabs for Go and
  2-space indent for HCL/JSON/YAML.
- **GoLand:** opens the module automatically from `go.mod`.

Pre-commit hygiene:

```bash
gofmt -w .
golangci-lint run ./...
go test ./...
```

Debugging: `go run . -debug` starts the provider in debug mode (see the `-debug`
flag in `main.go`); attach Delve and use `TF_REATTACH_PROVIDERS` as printed.

---

## 14. Troubleshooting

| Symptom | Cause | Fix |
|---------|-------|-----|
| `terraform` downloads the provider anyway | Dev override not picked up | Ensure `TF_CLI_CONFIG_FILE` points at your file / correct default path; check `dev_overrides` address is `eshiam-corp/eshiam` |
| Warning: "Provider development overrides are in effect" | Expected with dev overrides | Ignore; it confirms the override works |
| `Error: Missing API Host/client_id/client_secret` | `EXAMPLE_*` not set | Export all three (Section 6) |
| `could not connect` / `oauth token` errors | Wrong host or credentials | Verify `EXAMPLE_HOST` and PAT; token URL is `host + /oauth/token` |
| `go: updates to go.mod needed` | Restricted module mode | `set GOFLAGS=-mod=mod` (PowerShell: `$env:GOFLAGS="-mod=mod"`) |
| `go install` binary not found by Terraform | Wrong dev-override path | Use `go env GOBIN` (or `GOPATH\bin`) exactly |
| Build downloads fail behind proxy | Blocked module proxy | Set `GOPROXY` / `HTTPS_PROXY` |
| `terraform init` errors under dev override | `init` isn't allowed with overrides | Skip `init`; run `plan` directly |

---

## Next steps

- [Local development](local-development.md) — the fast inner loop.
- [Adding a resource](adding-a-resource.md) — extend the provider.
- [Docker guide](docker.md) — containerized build & the linkage workflow.
- [Publishing to the registry](publishing-to-the-registry.md).
- [Choosing Go or Python](../../python-terraform-framework/docs/choosing-go-or-python.md)
  (also mirrored in the Python framework).
