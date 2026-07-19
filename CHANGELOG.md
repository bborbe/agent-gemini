# Changelog

All notable changes to this project will be documented in this file.

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
