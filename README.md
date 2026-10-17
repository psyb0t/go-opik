# go-opik

[![Go Reference](https://pkg.go.dev/badge/github.com/psyb0t/go-opik.svg)](https://pkg.go.dev/github.com/psyb0t/go-opik)
[![CI](https://github.com/psyb0t/go-opik/actions/workflows/pipeline.yml/badge.svg?branch=main)](https://github.com/psyb0t/go-opik/actions/workflows/pipeline.yml)
[![version](https://raw.githubusercontent.com/psyb0t/go-opik/badges/version.svg)](https://github.com/psyb0t/go-opik/tags)
[![license](https://raw.githubusercontent.com/psyb0t/go-opik/badges/license.svg)](LICENSE)
[![imported by](https://raw.githubusercontent.com/psyb0t/go-opik/badges/importers.svg)](https://github.com/psyb0t/go-opik/blob/badges/importers.md)

Unofficial Go client for the [Opik REST API](https://github.com/comet-ml/opik), generated with [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) from the OpenAPI document Comet publishes. I am not affiliated with Comet.

The package is `opik`, in the repository root. Everything in `client_*.gen.go` is generated. Don't edit it by hand, regenerate it.

## Install

```sh
go get github.com/psyb0t/go-opik
```

## Usage

Every endpoint in the spec has a typed method on `ClientWithResponses`. Responses decode into `JSON200` and the other status fields.

```go
package main

import (
	"context"
	"net/http"
	"os"

	"github.com/psyb0t/go-opik"
)

func main() {
	auth := opik.WithRequestEditorFn(func(_ context.Context, req *http.Request) error {
		req.Header.Set("Authorization", os.Getenv("OPIK_API_KEY"))
		req.Header.Set("Comet-Workspace", os.Getenv("OPIK_WORKSPACE"))

		return nil
	})

	client, err := opik.NewClientWithResponses("https://www.comet.com/opik/api", auth)
	if err != nil {
		panic(err)
	}

	res, err := client.FindProjectsWithResponse(
		context.Background(),
		&opik.FindProjectsParams{},
	)
	if err != nil {
		panic(err)
	}

	_ = res.JSON200
}
```

The same code is in `example_test.go`, so `make test` checks that it compiles.

Check `StatusCode()` and the relevant decoded response field before using a response. HTTP error status codes do not produce a Go error by themselves. Method names come from the spec's operation IDs.

## Auth

Opik Cloud takes the bare API key in `Authorization`, with no `Bearer ` prefix, and the workspace name in `Comet-Workspace`. Set both in a `WithRequestEditorFn`, as above. The spec describes these headers in its introduction rather than declaring a security scheme, so the generated client does not add them automatically.

| Deployment | Base URL |
|---|---|
| Opik Cloud | `https://www.comet.com/opik/api` |
| Local self-hosted server | `http://localhost:5173/api` |

For a self-hosted deployment, use its API URL and authentication policy. Keep the `/api` suffix in the base URL; the main resource paths start with `/v1/private/` and OAuth paths use `/oauth/`.

## Regenerating

The spec is not committed. `make generate` downloads the current OpenAPI document into the gitignored `api/` directory and regenerates the client:

```sh
make generate
```

`gen.go` declares the steps, so `make generate` is just `go generate ./...`:

- `scripts/fetch-spec.sh` downloads `https://raw.githubusercontent.com/comet-ml/opik/main/sdks/code_generation/fern/openapi/openapi.yaml` to `api/openapi.yaml`.
- `scripts/patch-spec.sh` runs the Go patch tool and writes `api/openapi.fixed.yaml`. It retags bare YAML `=` enum values as strings so strict loaders can read the document. It also replaces dash punctuation in descriptions and summaries with commas for generated comments.
- oapi-codegen runs with `config.yaml`. The config maps unformatted `number` fields to `float64`.
- `scripts/split-gen.sh` splits the generated output at declaration boundaries into `client_*.gen.go` files and fixes their imports.

`.github/workflows/regen.yml` does the same every day in CI. When the new spec changes the client, it runs `make lint`, `make test` and `make sec`. If a check fails, Claude edits the code and the checks run again, up to three times, and a review step reads the result. A passing change is committed with a CHANGELOG entry and tagged as the next release: a patch version for fixes and additions, a minor version for breaking changes.

## Agent integrations

The [skill](.agents/skills/go-opik) teaches a coding agent how to use go-opik when it writes Go. It works in any agent that reads `.agents/skills/`, and installs natively in the clients below.

### Claude Code

```bash
claude plugin marketplace add psyb0t/agents
claude plugin install go-opik@psyb0t
```

### Codex

```bash
codex plugin marketplace add psyb0t/agents
codex plugin add go-opik@psyb0t
```

Codex also picks the skill up with no install in any checkout of this repo.

### OpenClaw

The skill is published to ClawHub on every release:

```bash
openclaw skills install @psyb0t/go-opik
```

## Development

Every command runs in the `Dockerfile.dev` image. The host needs only Docker and Make.

```sh
make help      # list targets
make dep       # go mod tidy + vendor
make lint      # shellcheck, shfmt, go fix, golangci-lint
make test      # go test -race
make sec       # gitleaks, govulncheck, semgrep
make generate  # download the spec and regenerate
```

## License

Repository tooling and tests use [MIT](LICENSE). Generated definitions and descriptions derive from Comet's [Apache 2.0 licensed specification](LICENSE.opik); see [NOTICE](NOTICE). The specification itself is not redistributed here.
