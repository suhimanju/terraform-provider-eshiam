# Common development tasks for the provider.

BINARY := terraform-provider-eshiam

.PHONY: build install test testacc vet fmt lint docs clean

build:
	go build -v ./...

install:
	go install .

test:
	go test -v ./...

testacc:
	TF_ACC=1 go test -v -tags=integration -timeout 120m ./internal/...

vet:
	go vet ./...

fmt:
	gofmt -w .

lint:
	golangci-lint run ./...

docs:
	go generate ./...

clean:
	go clean
	rm -rf dist bin
