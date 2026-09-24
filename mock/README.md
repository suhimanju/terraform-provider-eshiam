# Mock API server

Unit tests can run against a local [Mockoon](https://mockoon.com/) mock server
instead of a live tenant. This keeps tests fast, deterministic, and free of
customer-specific data.

## Usage

Start the mock (Docker):

```bash
docker run -d --name provider-mock \
  --mount type=bind,source="$(pwd)/mock/mock_api.json",target=/data,readonly \
  -p 3000:3000 mockoon/cli:latest -d data -p 3000
```

Point the provider at it and run the tests:

```bash
export EXAMPLE_HOST="http://localhost:3000"
export EXAMPLE_CLIENT_ID="test"
export EXAMPLE_CLIENT_SECRET="test"
TF_ACC=1 go test ./...
```

Tear it down:

```bash
docker stop provider-mock && docker rm provider-mock
```

## Extending

`mock_api.json` is a Mockoon environment file. Add a new route for each
endpoint your resource needs (create/read/update/delete). Open the file in the
Mockoon desktop app to edit it visually, or hand-edit the JSON. Keep responses
generic — do not embed real tenant identifiers or credentials.
