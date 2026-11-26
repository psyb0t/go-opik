# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.7.2, 2026-11-26

- Added `RunAnalyticsQuery` and `RunAnalyticsQueryWithBody` to `Client` and `ClientWithResponses`, with `RunAnalyticsQueryWithResponse` and `RunAnalyticsQueryWithBodyWithResponse`. They call `POST /v1/private/analytics-queries` to run read-only analytics SQL scoped to the workspace. Added `RunAnalyticsQueryJSONRequestBody`, `RunAnalyticsQueryResponse`, `NewRunAnalyticsQueryRequest`, `NewRunAnalyticsQueryRequestWithBody` and `ParseRunAnalyticsQueryResponse`.
- Added `RunProjectTopics` and `ResetProjectTopics`, with their `WithResponse` methods, to run a project's topic discovery now or forget its topic list. Added `RunProjectTopicsParams`, `ResetProjectTopicsParams`, `RunProjectTopicsResponse`, `ResetProjectTopicsResponse`, the `New...Request` builders and the `Parse...Response` functions.
- Migration: the new methods are also on `ClientInterface` and `ClientWithResponsesInterface`. Code that implements either interface itself, for example a mock, must add them.

## v0.7.1, 2026-11-25

- Added guardrail model methods `FindGuardrailModels`, `CreateGuardrailModelWithBody`, `GetGuardrailModelById`, `UpdateGuardrailModel`, `UpdateGuardrailModelWithBody`, `DeleteGuardrailModel` and `RetrainGuardrailModelWithBody` to `Client` and `ClientWithResponses`, with the matching `...WithResponse` methods, `New...Request` builders and `Parse...Response` functions.
- Added guardrail training run methods `FindGuardrailTrainingRuns`, `GetGuardrailTrainingRunById` and `CancelGuardrailTrainingRun`, with their `WithResponse` methods, request builders and parsers.
- Added `GetProjectTopics`, `GetProjectTopicsWithResponse`, `GetProjectTopicsParams`, `ProjectTopics`, `TopicStats`, `OtherStats` and `DayCount`, plus the status enum `ProjectTopicsStatus` (`collecting`, `disabled`, `ready`, `too_large`).
- Added the types `GuardrailModel`, `GuardrailModelPage`, `GuardrailModelUpdate`, `GuardrailTrainingRun`, `GuardrailTrainingRunPage`, `Attempt`, and the enums `GuardrailTrainingRunStatus`, `AttemptStatus` and `FindGuardrailModelsParamsLatestRunStatus`.
- Added the optional fields `GuardrailsTrainingEnabled` and `TopicsEnabled` to `ServiceTogglesConfig`, and `DatasetItemData` to `ExperimentItem` and `ExperimentItemCompare`.
- Migration: the new methods are also on `ClientInterface` and `ClientWithResponsesInterface`. Code that implements either interface itself, for example a mock, must add them.

## v0.7.0, 2026-11-24

- Added view endpoints to `ClientInterface`, `Client`, `ClientWithResponsesInterface` and `ClientWithResponses`: `CreateView`, `GetViewById`, `UpdateView`, `DeleteView`, `GetViewVersions`, `GetViewVersion`, `CreateViewVersion`, `DeleteViewVersion` and `SetViewCurrentVersion`, each with a `WithBody` variant where it takes a body and a `WithResponse` form. Matching `New...Request` builders and `Parse...Response` functions are included.
- Added the view types `View`, `ViewCreate`, `ViewUpdate`, `ViewInput`, `ViewCurrentVersion`, `ViewVersion`, `ViewVersionCreate`, `ViewVersionSummary` and `ViewVersionPage`, with their enum types and constants. Added `GetViewVersionsParams` for the version list query.
- Added `StreamDatasetExperimentItems`, `StreamDatasetExperimentItemsWithBody` and their `WithResponse` forms. They call `POST /v1/private/datasets/items/experiments/items/stream`. Added `DatasetExperimentItemStreamRequest`, `StreamDatasetExperimentItemsJSONRequestBody`, `StreamDatasetExperimentItemsResponse`, `NewStreamDatasetExperimentItemsRequest` and `ParseStreamDatasetExperimentItemsResponse`.
- Added the field `ModelPricesEnabled` (`bool`) to `ServiceTogglesConfig`.
- Removed the field `DatasetVersion` (`*string`) from `FindDatasetItemsWithExperimentItemsParams`. The `dataset_version` query parameter is gone from the upstream spec. `DatasetVersions` replaces it in the new stream request. Code that sets the field must drop it.
- Added methods to the `ClientInterface` and `ClientWithResponsesInterface` interfaces. Custom implementations or mocks of these interfaces must add them.
- Migration: delete uses of `FindDatasetItemsWithExperimentItemsParams.DatasetVersion` and implement the new interface methods in any custom client or mock.

## v0.6.1, 2026-11-23

- Added the field `JSON409` (`*ErrorMessage`) and the method `GetJSON409` to `UpdateAutomationRuleEvaluatorResponse`. `ParseUpdateAutomationRuleEvaluatorResponse` now decodes an HTTP 409 `application/json` reply into `JSON409`.
- Updated the doc comments of `DatasetWrite.ProjectName` and `PromptWrite.ProjectName`. Without a project, the item is now created in the Default Project instead of at workspace level.

## v0.6.0, 2026-11-18

- Added guardrail policy endpoints to `ClientInterface`, `Client`, `ClientWithResponsesInterface` and `ClientWithResponses`: `CreateGuardrailPolicy`, `FindGuardrailPolicies`, `GetGuardrailPolicyById`, `UpdateGuardrailPolicy`, `DeleteGuardrailPolicy` and `ValidateGuardrails`, each with a `WithBody` variant and a `WithResponse` form. Matching `New...Request` builders and `Parse...Response` functions are included.
- Added the guardrail types `GuardrailPolicyPublic`, `GuardrailPolicyPagePublic`, `GuardrailPolicyWrite`, `GuardrailPolicyRulePublic`, `GuardrailPolicyRuleWrite`, `GuardrailsValidationRequest`, `GuardrailsValidationResponse`, `ValidationDescriptor` and `ValidationResult`, with their enum types and constants. Added `FindGuardrailPoliciesParams` for the list query.
- Added the field `GuardrailsValidationsEnabled` to `ServiceTogglesConfig`.
- Removed the field `Slug` (`*string`) from `DashboardPublic`. The upstream spec no longer has it. Code that reads the field must drop it.
- Added methods to the `ClientInterface` and `ClientWithResponsesInterface` interfaces. Custom implementations or mocks of these interfaces must add the new methods.
- Raised the minimum Go version to 1.26.11 in `go.mod` and the development image, to pick up the `net/http` fix for GO-2026-6859.
- Migration: delete uses of `DashboardPublic.Slug` and implement the new interface methods in any custom client or mock.

## v0.5.0, 2026-11-07

- Removed the field `OnlineScoringTracingEnabled` (`bool`) from `ServiceTogglesConfig`. The upstream spec no longer has `onlineScoringTracingEnabled`. Code that reads or sets the field must drop it.
- Updated the doc comments of `ProviderAuthConfig.SendAs`, `ProviderAuthConfigPublic.SendAs`, `ProviderAuthConfigWrite.SendAs` and their `SendAs` types. Basic auth now documents `client_id`/`client_secret` or `username`/`password` in the HTTP Basic header.
- Updated the doc comments of `TokenField` on `ProviderAuthConfig`, `ProviderAuthConfigPublic` and `ProviderAuthConfigWrite`. A reply that is the bare token is used as-is and needs `fallback_ttl_seconds`.
- Migration: delete uses of `ServiceTogglesConfig.OnlineScoringTracingEnabled`.

## v0.4.0, 2026-11-05

- Changed `FindDatasetItemsWithExperimentItemsParams.ExperimentIds` from `string` to `*string`. The `experiment_ids` query parameter is now optional, and `NewFindDatasetItemsWithExperimentItemsRequest` omits it when the pointer is nil. Callers must pass a pointer.
- Added the optional field `DatasetVersion` (`*string`) to `FindDatasetItemsWithExperimentItemsParams`. It sets the `dataset_version` query parameter.
- Added the optional field `OnlyRoute` (`*string`) to `ModelParams`.
- Updated the doc comments of `BreakdownConfig.Field`, `BreakdownConfigField`, `BreakdownConfigPublic.Field` and `BreakdownConfigPublicField`. They now say that tags are overlapping groups and that adding tag groups can exceed actual cost.
- Migration: change `ExperimentIds: "a,b"` to `ExperimentIds: &ids` in `FindDatasetItemsWithExperimentItemsParams`.

## v0.3.0, 2026-10-29

- Added `GetLlmModelProfiles` to `Client` and `GetLlmModelProfilesWithResponse` to `ClientWithResponses`. They call `GET /v1/private/llm/models/profiles`.
- Added `GetLlmModelProfilesResponse`, `NewGetLlmModelProfilesRequest` and `ParseGetLlmModelProfilesResponse`. The response type exposes a `map[string]ModelParams` in `JSON200`.
- Added the types `ModelParams`, `EffortParam`, `FlagParam`, `NumberParam`, `ToolsParam` and `ToolsParamChatCompletions`, with the values `ToolsParamChatCompletionsEffortNoneOnly` and `ToolsParamChatCompletionsRefused`.
- Migration: the two new methods are also on `ClientInterface` and `ClientWithResponsesInterface`. Code that implements either interface itself, for example a mock, must add them.

## v0.2.0, 2026-10-28

- Added `CancelExperiments`, `CancelExperimentsWithBody`, `CancelExperimentsWithResponse` and `CancelExperimentsWithBodyWithResponse` to `Client` and `ClientWithResponses`. They call `PUT /v1/private/experiments/cancel`.
- Added `CancelExperimentsResponse` (with `JSON403` and `JSON429` fields), `CancelExperimentsJSONRequestBody`, `IdsHolder`, `NewCancelExperimentsRequest`, `NewCancelExperimentsRequestWithBody` and `ParseCancelExperimentsResponse`.
- Migration: the four new methods are also on `ClientInterface` and `ClientWithResponsesInterface`. Code that implements either interface itself, for example a mock, must add them.
- Added `JSON403`, `JSON429`, `GetJSON403` and `GetJSON429` to `ExecuteExperimentResponse`.
- Added the optional fields `Filters`, `ProjectId` and `SelectedRuleIds` to `ExperimentExecutionRequest`, `TotalItems` to `ExperimentInfo` and `FinishedAt` to `ExperimentPublic`.
- Added the status value `failed`: `ExperimentPublicStatusFailed`, `ExperimentUpdateStatusFailed` and `ExperimentWriteStatusFailed`.
- Added the types `AnnotationQueueAnnotatorView`, `AnnotationQueueAnnotatorViewPublic` and `AnnotationQueueAnnotatorViewWrite`, and the optional field `AnnotatorView` on `AnnotationQueuePublic`, `AnnotationQueueUpdate` and `AnnotationQueueWrite`.

## v0.1.7, 2026-10-28

- Updated Semgrep in the development image to 1.182.1 and refreshed its locked dependencies. certifi 2026.10.23 drops the emSign root CAs whose intermediate keys were compromised (GHSA-6q48-6c7w-9gw2).
- Dependabot no longer opens pull requests for `mcp` or the OpenTelemetry packages in `requirements-semgrep.txt`. Semgrep pins their versions itself, so bumping one of them on its own cannot install.
- Updated kin-openapi, which the code generator uses, from 0.142.0 to 0.149.0. The newer version fixes several panics on malformed specs reported in GitHub security advisories.
- The security scan merges its two gitleaks runs into one SARIF run. GitHub code scanning rejects a report with two runs from the same tool, so no scan results had been uploaded.

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
