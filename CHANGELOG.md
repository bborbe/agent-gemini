# Changelog

All notable changes to this project will be documented in this file.

## Unreleased

- chore: update github.com/bborbe/agent to v0.86.0, github.com/bborbe/cqrs to v0.6.10, github.com/bborbe/errors to v1.6.0, github.com/bborbe/kafka to v1.25.11, github.com/bborbe/sentry to v1.10.1, github.com/bborbe/service to v1.10.11, github.com/bborbe/time to v1.27.12, github.com/bborbe/vault-cli to v0.121.2, github.com/onsi/gomega to v1.43.0, google.golang.org/genai to v1.71.0

## v0.1.8

- chore: update github.com/bborbe/agent to v0.83.1, github.com/bborbe/cqrs to v0.6.8, github.com/bborbe/errors to v1.5.21, github.com/bborbe/kafka to v1.25.9, github.com/bborbe/sentry to v1.9.27, github.com/bborbe/service to v1.10.9, github.com/bborbe/time to v1.27.10, github.com/bborbe/vault-cli to v0.116.2, google.golang.org/genai to v1.70.0

## v0.1.7

- chore: update Go to 1.27.0

## v0.1.6

- chore: Bump errcheck to v1.20.0 and golangci-lint to v2.13.1 for Go 1.27 support

## v0.1.5

- update Go to 1.26.6 and update dependencies (GO-2026-5026, GO-2026-5972, GO-2026-6090, GO-2026-6218)

## v0.1.4

- fix(deps): bump google.golang.org/grpc to v1.82.1 (GHSA-hrxh-6v49-42gf)
- fix(deps): bump go.opentelemetry.io/otel to v1.44.0 (GO-2026-5158)

## v0.1.3

- Bump `golang.org/x/text` to v0.39.0 (CVE-2026-56852)

## v0.1.2

- Bump Go toolchain 1.26.4 -> 1.26.5 and Alpine 3.23 -> 3.24 in Dockerfile
- Update bborbe/* and google.golang.org/genai dependencies
- Ignore no-fix vuln advisory GO-2026-5932 in osv-scanner and trivy configs

## v0.1.1

- refactor: converge build to bborbe/kafka-topic-reader publish-only model — make buca publishes docker.io/bborbe/agent-gemini:$(VERSION); deploy machinery removed.

## v0.1.0

- Bump `github.com/bborbe/agent` to v0.72.0 and `github.com/bborbe/cqrs` to v0.6.0.
- Add optional `TopicPrefix` config (`TOPIC_PREFIX` env / `--topic-prefix`) for explicit Kafka
  topic prefixing, independent of the existing `Branch` field. `NewKafkaResultDeliverer` now
  takes `base.TopicPrefix` instead of `base.Branch`; `Branch` is kept for its other uses.
