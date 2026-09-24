# Pipelines

Vendor-neutral CI/CD templates for the provider. Pick the file that matches
your CI platform — every one runs the same `build → vet → lint → test` flow and
releases via GoReleaser on `v*` tags.

| Platform | File(s) |
|----------|---------|
| GitHub Actions | `.github/workflows/ci.yml`, `.github/workflows/release.yml` |
| Azure DevOps | `pipelines/azure-pipelines.yml` |
| Jenkins | `Jenkinsfile` (Linux/Unix agent), `Jenkinsfile.windows` (Windows agent) |
| GitLab CI/CD | `.gitlab-ci.yml` |
| AWS CodeBuild | `buildspec.yml` (CI), `buildspec-release.yml` (release) |
| CircleCI | `.circleci/config.yml` |
| Travis CI | `.travis.yml` |
| Bitbucket Pipelines | `bitbucket-pipelines.yml` |
| Drone CI | `.drone.yml` |
| Woodpecker CI | `.woodpecker.yml` |

## GitHub Actions

- `.github/workflows/ci.yml` — build, vet, test, and lint on push/PR.
- `.github/workflows/release.yml` — GoReleaser build + signed release on
  `v*` tags. Requires `GPG_PRIVATE_KEY` and `GPG_PASSPHRASE` secrets for
  Terraform Registry publishing.

## Azure DevOps

- `pipelines/azure-pipelines.yml` — build + unit tests with JUnit results.

## Jenkins

- `Jenkinsfile` — declarative pipeline with build, vet, lint, test (JUnit +
  coverage), and a tag-gated release stage. Runs `sh` steps on a **Linux/Unix**
  agent.
- `Jenkinsfile.windows` — the same stages for a **Windows** agent using
  `powershell` steps. Point your pipeline job at this file (or rename it to
  `Jenkinsfile`) and target an agent labelled `windows`.

Both use `withCredentials` to bind secrets to the narrowest possible scope: the
GPG signing key is a **Secret file** credential (its contents never touch an
environment variable) and the passphrase/token are masked Secret Text.

## GitLab CI/CD

- `.gitlab-ci.yml` — staged build/test/release with JUnit reports and coverage
  parsing.

## AWS CodeBuild

- `buildspec.yml` — CI build + tests with JUnit reports.
- `buildspec-release.yml` — signed GoReleaser release; pulls `GPG_PRIVATE_KEY`,
  `GPG_PASSPHRASE`, and `GITHUB_TOKEN` from AWS Secrets Manager.

## CircleCI

- `.circleci/config.yml` — build/test workflow plus a tag-only release job.

## Travis CI

- `.travis.yml` — build/test with a tag-gated GoReleaser deploy.

## Bitbucket Pipelines

- `bitbucket-pipelines.yml` — build/test on branches, release on `v*` tags.

## Drone / Woodpecker CI

- `.drone.yml`, `.woodpecker.yml` — build/test with tag-gated release steps.

## Releasing to the Terraform Registry

1. Configure a GPG signing key and add its private key + passphrase as CI
   secrets.
2. Tag a release: `git tag v0.1.0 && git push origin v0.1.0`.
3. GoReleaser builds cross-platform binaries, a `SHA256SUMS` file, and a
   detached signature — the artifacts the Terraform Registry expects.

> For the full walkthrough — generating a signing key, registering the provider,
> and per-platform release setup — see
> [../docs/publishing-to-the-registry.md](../docs/publishing-to-the-registry.md).

## Integration tests

`integration_tests.sh` runs Go tests tagged `integration` against a live tenant.
Export `EXAMPLE_HOST`, `EXAMPLE_CLIENT_ID`, and `EXAMPLE_CLIENT_SECRET` first.

## Handling secrets

**Unit tests use the offline mock in `mock/` and need no credentials**, so the
default build/test jobs carry no secrets at all. Secrets are only needed for two
optional things: integration tests (tenant credentials) and releases (a GPG
signing key). Every template pulls these from the platform's **native secret
store** — nothing is ever written into the pipeline files.

| Platform | Secret mechanism |
|----------|------------------|
| GitHub Actions | Actions **Secrets** (`${{ secrets.* }}`); not exposed to fork PRs |
| Azure DevOps | **Variable Groups** / Key Vault, mapped explicitly via `env:` |
| Jenkins | **Credentials Manager** (`credentials('...')`) |
| GitLab CI/CD | **masked + protected** CI/CD variables |
| AWS CodeBuild | **AWS Secrets Manager** (`secrets-manager:` block) |
| CircleCI | **Contexts** (`context: provider-release`) |
| Travis CI | encrypted **environment variables** (value hidden in logs) |
| Bitbucket | **Secured** repository/workspace variables |
| Drone / Woodpecker | secret store via `from_secret` |

### Secrets used

| Secret | Purpose | Required for |
|--------|---------|--------------|
| `EXAMPLE_HOST` | API base URL | Integration tests |
| `EXAMPLE_CLIENT_ID` | OAuth2 client ID | Integration tests |
| `EXAMPLE_CLIENT_SECRET` | OAuth2 client secret | Integration tests |
| `GPG_PRIVATE_KEY` | Terraform Registry signing key | Release |
| `GPG_PASSPHRASE` | Passphrase for the signing key | Release |
| `GITHUB_TOKEN` | Publishing the release | Release |

### Rules of thumb

- **Never** hardcode a secret in a pipeline file or commit it to the repo.
- Prefer **masked/secured** variable types so values are redacted in build logs.
- Restrict tenant credentials to protected branches / manual jobs, and avoid
  running integration tests on pull requests from forks.
- Rotate the GPG key and `GITHUB_TOKEN` periodically.
