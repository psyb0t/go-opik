# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.1.6, 2026-10-27

- The development image is based on Debian (`golang:1.26.10-trixie`) instead of Alpine. Alpine keeps only the latest build of each package, so the exact pins broke the image build whenever a package was updated. System packages now come from the snapshot.debian.org date the base image was built from, which keeps them fixed for a given base image digest. jq 1.8.2, shellcheck 0.11.0 and shfmt 3.13.1 keep their versions, and Semgrep runs in its own virtual environment.
- Raised the minimum Go version in `go.mod` from 1.26.9 to 1.26.10, which fixes the `crypto/x509` vulnerability GO-2026-6809 found by govulncheck.

## v0.1.5, 2026-10-23

- Added the field `JSON400` (`*ErrorMessage`) and the method `GetJSON400` to `ExecuteExperimentResponse`. `ParseExecuteExperimentResponse` now decodes an HTTP 400 JSON body into `JSON400`.

## v0.1.4, 2026-10-22

- Added the optional field `CustomFields` (`*map[string]string`) to `AlertPublic` and `AlertWrite`. It holds user-defined key/value pairs added as top-level keys of the alert's EventBridge event detail.
- Updated the doc comments of `FindEnvironments` and `FindEnvironmentsWithResponse` on the client and interface types: the workspace limit default is now documented as 100, not 20.
- The development image `Dockerfile.dev` uses Go 1.26.10, which fixes the `crypto/x509` vulnerability GO-2026-6809 reported by `make sec`, and `ca-certificates` 20261020-r0.

## v0.1.3, 2026-10-20

- Added `GetDatasetVersionByHash` and `GetDatasetVersionByHashWithResponse` to `Client` and `ClientWithResponses`, and to `ClientInterface` and `ClientWithResponsesInterface`. They call `GET /v1/private/datasets/{id}/versions/hash/{versionHash}`.
- Added `GetDatasetVersionByHashResponse`, `NewGetDatasetVersionByHashRequest` and `ParseGetDatasetVersionByHashResponse`. The response type exposes `DatasetVersionPublic` on 200 and `ErrorMessagePublic` on 404.
- Added the optional field `ResetProjectIds` to `Response`, a list of project IDs whose automatic run never finished.
- Code that implements `ClientInterface` or `ClientWithResponsesInterface` itself, for example a mock, must add the new methods.

## v0.1.2, 2026-10-17

- Documentation only. The agent skill explains how to look up every endpoint, parameter and model with `go doc` or on pkg.go.dev, and declares its license.
- The skill no longer lists the API credentials under OpenClaw's `requires.env`, which kept OpenClaw from loading it until they were set. It requires only `go`; the credentials are still documented as what your program reads.
- The README points to the full API reference, and the examples in the README and the skill were checked against the current code.

## v0.1.1, 2026-10-17

- Generated doc comments keep the API's own description text. The spec patch no longer rewrites dashes or wording in descriptions.
- `make sec` uses Semgrep 1.181.0. Its pinned requirements now hash every published file, so the development image also builds on arm64.

## v0.1.0, 2026-10-17

- Added `opik`, a REST client generated with oapi-codegen from Comet's Opik REST API specification, version 1.0.0, OpenAPI 3.1.0. It covers all 256 paths and 318 operations, including projects, traces, spans, datasets, experiments and prompts.
- Number fields without a format use `float64` instead of oapi-codegen's default `float32`.
- `make generate` downloads the specification from `https://raw.githubusercontent.com/comet-ml/opik/main/sdks/code_generation/fern/openapi/openapi.yaml`, normalizes YAML `=` enum values for strict loaders and description punctuation for generated comments, then regenerates the client. The specification is not committed.
- Added the `go-opik` agent skill and plugin manifests.
