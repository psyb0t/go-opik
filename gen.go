// Package opik is a Go client for the Opik REST API,
// generated with oapi-codegen from the OpenAPI document Opik publishes.
package opik

//go:generate bash scripts/fetch-spec.sh
//go:generate bash scripts/patch-spec.sh
//go:generate go tool oapi-codegen --config=config.yaml api/openapi.fixed.yaml
//go:generate bash scripts/split-gen.sh
