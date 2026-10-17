# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.1.1, 2026-10-17

- Generated doc comments keep the API's own description text. The spec patch no longer rewrites dashes or wording in descriptions.
- `make sec` uses Semgrep 1.181.0. Its pinned requirements now hash every published file, so the development image also builds on arm64.

## v0.1.0, 2026-10-17

- Added `opik`, a REST client generated with oapi-codegen from Comet's Opik REST API specification, version 1.0.0, OpenAPI 3.1.0. It covers all 256 paths and 318 operations, including projects, traces, spans, datasets, experiments and prompts.
- Number fields without a format use `float64` instead of oapi-codegen's default `float32`.
- `make generate` downloads the specification from `https://raw.githubusercontent.com/comet-ml/opik/main/sdks/code_generation/fern/openapi/openapi.yaml`, normalizes YAML `=` enum values for strict loaders and description punctuation for generated comments, then regenerates the client. The specification is not committed.
- Added the `go-opik` agent skill and plugin manifests.
