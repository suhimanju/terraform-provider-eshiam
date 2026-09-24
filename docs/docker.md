# Dockerized Build — Beginner's Guide

**Who is this for?** Anyone who wants to build or test this Terraform provider
**without installing Go**. You only need **Docker** and **Git**.

This provider is the *tool* half of a two-part model:

```
  terraform-provider-eshiam   (THIS repo — the TOOL, a Go plugin)
            │  built here, then consumed as a versioned artifact
            ▼
  a linkage/consumer framework (Terraform/Terragrunt content that USES the provider)
            │  terraform / terragrunt apply
            ▼
        Your target API / tenant
```

Docker lets you compile the provider in a throwaway container, then hand the
binary (or a published version) to whatever consumes it.

---

## 1. Install Docker

| OS | How | Notes |
|----|-----|-------|
| **Windows** | Docker Desktop from docker.com | Needs WSL2 (installer sets it up). Start Docker Desktop first. |
| **macOS** | Docker Desktop (or `brew install --cask docker`) | Launch the app once so the engine starts. |
| **Linux** | Docker Engine + Compose plugin | May need `sudo` or the `docker` group. |

Verify it works:

```bash
docker version
docker compose version
```

---

## 2. Build the provider in a container

From the repo root:

```bash
docker build -t terraform-provider-eshiam-build .
```

This uses the multi-stage `Dockerfile`: it downloads modules, compiles
everything, and leaves the binary at `/out/terraform-provider-eshiam` inside
the image. No Go install on your machine.

### Or use Docker Compose (recommended)

```bash
docker compose build            # build the image
docker compose run --rm build   # compile all packages
docker compose run --rm test    # run the unit tests
```

---

## 3. Get the compiled binary onto your machine

The easiest way is the `extract` compose service, which writes the binary to
`./dist/` on the host via the mounted volume:

```bash
docker compose run --rm extract
# -> dist/terraform-provider-eshiam
```

Or copy it out of the built image directly:

```bash
docker create --name p terraform-provider-eshiam-build
# Linux/macOS:
docker cp p:/out/terraform-provider-eshiam ./terraform-provider-eshiam
# Windows (PowerShell):
docker cp p:/out/terraform-provider-eshiam .\terraform-provider-eshiam.exe
docker rm p
```

---

## 4. Run tests in the container

```bash
docker run --rm -v "${PWD}:/src" -w /src terraform-provider-eshiam-build go test ./...
# or
docker compose run --rm test
```

---

## 5. Linking to a consumer / linkage framework

A **linkage framework** (the *content* half — Terraform/Terragrunt modules that
call `resource "example_..."`) consumes this provider in one of two ways:

### Option A — `dev_overrides` (fastest for local dev)

1. Build and extract the binary (Part 3) into a folder, e.g.
   `~/go/bin` or `./dist`.
2. In the consumer repo, point Terraform at that folder via a `.terraformrc`:

   ```hcl
   provider_installation {
     dev_overrides {
       "eshiam-corp/eshiam" = "/absolute/path/to/dist"
     }
     direct {}
   }
   ```

3. Set `TF_CLI_CONFIG_FILE` to that file. No `terraform init` is needed with
   overrides.

> The linkage framework in this project ships a `docker/provider-builder.Dockerfile`
> and a `docs/docker.md` that automate exactly this. When using **that** repo,
> follow its guide — it builds this provider and wires up the override for you.

### Option B — registry / network mirror (best for teams & CI)

1. Publish a signed release (see `docs/publishing-to-the-registry.md`).
2. The consumer pins the published version in `required_providers` and downloads
   it normally — no binary copying. This is the recommended path for shared and
   Dockerized runners because the provider resolves over the network instead of a
   mounted binary.

### Which to use?

| Situation | Use |
|-----------|-----|
| Local experiments, fast iteration | Option A (`dev_overrides`) |
| Dockerized runner, team, or CI | Option B (registry/mirror) |

---

## 6. Files in this repo

| File | Purpose |
|------|---------|
| `Dockerfile` | Multi-stage build that compiles the provider binary |
| `docker-compose.yml` | `build` / `test` / `extract` convenience services |
| `.dockerignore` | Keeps the build context small and free of local artifacts |

---

## 7. Troubleshooting

| Symptom | Likely cause | Fix |
|---------|--------------|-----|
| `Cannot connect to the Docker daemon` | Docker not running | Start Docker Desktop (Win/mac); `sudo systemctl start docker` (Linux) |
| Build is slow every time | Module cache not reused | The `Dockerfile` copies `go.mod`/`go.sum` first — keep them in the context |
| `permission denied` on volume (Linux) | UID mismatch on the mount | Add `:z` to the volume, or run with a matching UID |
| Consumer can't find the provider | Wrong `dev_overrides` path or `source` | Ensure the path holds the binary and `source` = `eshiam-corp/eshiam` |

---

## References

- Publishing to a registry: `docs/publishing-to-the-registry.md`
- Local (non-Docker) development: `docs/local-development.md`
- Consumer-side Docker guide: the linkage framework's `docs/docker.md`
