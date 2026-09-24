# Publishing to the Terraform Registry

This guide walks through publishing the provider to the
[Terraform Registry](https://registry.terraform.io) and wiring up a CI/CD
pipeline (Jenkins or any other platform) to automate releases.

If you are new to the project, read `../README.md` first for the scope and
architecture, then come back here when you are ready to ship a release.

---

## 1. Scope & purpose

This repository is a **generic, vendor-neutral boilerplate** for a Terraform
provider built on the
[terraform-plugin-framework](https://developer.hashicorp.com/terraform/plugin/framework).

- **What it is:** a ready-to-fork starting point with a working provider, an
  OAuth2 API client, generic CRUD bases, 26 example resources, and 4 data
  sources — all using the placeholder name `example`.
- **What you do with it:** rename `example`/`example-org` to your own names,
  point the client at your API, keep the resources you need, and publish.
- **What the registry gives you:** a public (or private) home so users can run
  `terraform init` and automatically download your provider.

> The Terraform Registry only hosts **public** providers from **public GitHub
> repositories**. For private distribution, use a
> [private registry](https://developer.hashicorp.com/terraform/cloud-docs/registry)
> (Terraform Cloud/Enterprise) or a network mirror — the release artifacts
> produced here work with all of them.

---

## 2. One-time prerequisites

### 2.1 Rename the placeholders

Replace these everywhere before publishing:

| Placeholder | Replace with | Example |
|-------------|--------------|---------|
| `terraform-provider-eshiam` | `terraform-provider-<name>` | `terraform-provider-acme` |
| `example` (provider type) | your provider type | `acme` |
| `eshiam-corp/eshiam` | `<namespace>/<name>` | `acme-corp/acme` |

The GitHub repository **must** be named `terraform-provider-<name>` — the
registry requires this exact pattern.

### 2.2 Generate a GPG signing key

The registry requires every release to be **GPG-signed** so Terraform can verify
authenticity.

```bash
# Generate a key (choose RSA 4096, no expiry or a long one).
gpg --full-generate-key

# Find the key's fingerprint (the long hex string).
gpg --list-secret-keys --keyid-format=long

# Export the PUBLIC key — you paste this into the registry.
gpg --armor --export <FINGERPRINT> > public.asc

# Export the PRIVATE key — this becomes a CI secret. Keep it safe.
gpg --armor --export-secret-keys <FINGERPRINT> > private.asc
```

### 2.3 Add the public key to your registry account

1. Sign in to <https://registry.terraform.io> with your GitHub account.
2. Go to **User Settings → Signing Keys → New GPG Key**.
3. Paste the contents of `public.asc`.

### 2.4 Register the provider on the registry

1. On the registry, click **Publish → Provider**.
2. Authorize the registry's GitHub app and select your
   `terraform-provider-<name>` repository.
3. The registry now watches the repo for new **GitHub Releases**.

---

## 3. How a release is built

Releases are produced by [GoReleaser](https://goreleaser.com) using the
`.goreleaser.yaml` in this repo. On a version tag it:

1. Cross-compiles binaries for windows/linux/darwin × amd64/arm64.
2. Zips each binary with the docs.
3. Writes a `SHA256SUMS` checksum file.
4. **GPG-signs** the checksum file (needs `GPG_FINGERPRINT`).
5. Publishes everything as a GitHub Release.

The registry then detects the release, reads `terraform-registry-manifest.json`
(which declares Protocol 6), and makes the new version available.

To cut a release manually:

```bash
git tag v0.1.0
git push origin v0.1.0
# then run GoReleaser locally, or let CI do it (recommended — see below).
```

Version tags **must** follow semantic versioning with a `v` prefix: `v1.2.3`.

---

## 4. Automating releases in CI

Every pipeline template in this repo already has a **tag-gated release stage**.
You only need to store three secrets in your platform's secret manager:

| Secret | What it is |
|--------|------------|
| `GPG_PRIVATE_KEY` | Contents of `private.asc` from step 2.2 |
| `GPG_PASSPHRASE` | The passphrase you set on the key (omit if none) |
| `GITHUB_TOKEN` | A token with `repo` scope to publish the GitHub Release |

See `pipelines/README.md` for the exact secret mechanism per platform. **Never**
commit these values.

---

## 5. Setting up Jenkins (step by step)

The repo ships two Jenkins pipelines:

- `Jenkinsfile` — for a **Linux/Unix** agent (uses `sh`).
- `Jenkinsfile.windows` — for a **Windows** agent (uses `powershell`).

### 5.1 Add the credentials

In Jenkins: **Manage Jenkins → Credentials → (global) → Add Credentials**.

| Kind | ID | Value |
|------|----|-------|
| Secret file | `gpg-private-key` | upload `private.asc` |
| Secret text | `gpg-passphrase` | your key passphrase |
| Secret text | `github-token` | a GitHub token with `repo` scope |

(Optional, for integration tests: `example-host`, `example-client-id`,
`example-client-secret` as Secret text.)

The pipelines bind these with `withCredentials`, so they are masked in logs and
scoped to the release stage only.

### 5.2 Create the pipeline job

1. **New Item → Pipeline** (or **Multibranch Pipeline** to build every branch/tag).
2. Under **Pipeline → Definition**, choose **Pipeline script from SCM**.
3. Set **SCM = Git** and point it at your repository.
4. Set **Script Path**:
   - `Jenkinsfile` for a Linux agent, or
   - `Jenkinsfile.windows` for a Windows agent.
5. For a Multibranch job, enable **Discover tags** so `v*` tags trigger the
   release stage.

### 5.3 Run it

- Push to a branch → build, vet, lint, and unit tests run.
- Push a `v*` tag → the release stage imports the GPG key, runs GoReleaser, and
  publishes the GitHub Release, which the registry picks up automatically.

---

## 6. Setting up any other pipeline

The flow is identical everywhere — only the secret store and trigger syntax
differ. Pick the file for your platform from `pipelines/README.md`:

| Platform | CI file(s) | Trigger for release |
|----------|-----------|---------------------|
| GitHub Actions | `.github/workflows/*.yml` | push tag `v*` |
| Azure DevOps | `pipelines/azure-pipelines.yml` | tag `v*` |
| GitLab CI/CD | `.gitlab-ci.yml` | tag matching `^v` |
| AWS CodeBuild | `buildspec*.yml` | tag `v*` |
| CircleCI | `.circleci/config.yml` | tag `^v.*` |
| Travis CI | `.travis.yml` | tag `^v` |
| Bitbucket | `bitbucket-pipelines.yml` | tag `v*` |
| Drone / Woodpecker | `.drone.yml` / `.woodpecker.yml` | tag `v*` |

Steps for any platform:

1. Store `GPG_PRIVATE_KEY`, `GPG_PASSPHRASE`, `GITHUB_TOKEN` in its secret store.
2. Commit the CI file (already provided).
3. Push a `v*` tag to trigger a signed release.

---

## 7. Verifying the published provider

Once the registry shows your version, verify from a clean directory:

```terraform
terraform {
  required_providers {
    acme = {
      source  = "acme-corp/acme"
      version = ">= 0.1.0"
    }
  }
}
```

```bash
terraform init   # should download your provider from the registry
```

---

## 8. Troubleshooting

| Symptom | Likely cause |
|---------|--------------|
| Registry doesn't see the release | Repo not named `terraform-provider-<name>`, or the GitHub Release is a draft. |
| "signature verification failed" | The public key in your registry account doesn't match the private key used in CI. |
| GoReleaser: `GPG_FINGERPRINT` empty | Key wasn't imported before signing, or the fingerprint lookup failed. |
| Version rejected | Tag isn't semver with a `v` prefix (`v1.2.3`). |
| `terraform init` can't find it | Wrong `source` (`<namespace>/<name>`) or version not yet indexed (wait a few minutes). |

## References

- [Publishing Providers](https://developer.hashicorp.com/terraform/registry/providers/publishing)
- [Preparing a provider release](https://developer.hashicorp.com/terraform/registry/providers/publishing#preparing-and-adding-a-signing-key)
- [GoReleaser for Terraform providers](https://goreleaser.com/customization/builds/)
