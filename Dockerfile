# Dockerized build for the Terraform PROVIDER (this repo).
#
# Compiles the provider into a binary with no local Go toolchain required.
# A consumer/linkage framework then uses this binary via `dev_overrides` or by
# publishing it to a registry/mirror. See docs/docker.md.
#
# Build:   docker build -t terraform-provider-eshiam-build .
# Extract: docker create --name p terraform-provider-eshiam-build
#          docker cp p:/out/terraform-provider-eshiam ./terraform-provider-eshiam
#          docker rm p
# Test:    docker run --rm terraform-provider-eshiam-build go test ./...

FROM golang:1.23-alpine AS build

RUN apk add --no-cache git bash

WORKDIR /src

# Cache module downloads first for faster rebuilds.
COPY go.mod go.sum ./
RUN go mod download

# Then copy the rest of the source and build.
COPY . .
RUN go build -v ./... \
    && mkdir -p /out \
    && go build -o /out/terraform-provider-eshiam .

# Minimal final stage that just holds the compiled artifact.
FROM alpine:3.20
COPY --from=build /out/terraform-provider-eshiam /out/terraform-provider-eshiam
CMD ["sh", "-lc", "echo 'Provider binary is at /out/terraform-provider-eshiam. Use: docker cp <container>:/out/terraform-provider-eshiam .'"]
