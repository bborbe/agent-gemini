# Changelog

All notable changes to this project will be documented in this file.

## Unreleased

- refactor: converge build to bborbe/kafka-topic-reader publish-only model — make buca publishes docker.io/bborbe/agent-gemini:$(VERSION); deploy machinery removed.

## v0.1.0

- Bump `github.com/bborbe/agent` to v0.72.0 and `github.com/bborbe/cqrs` to v0.6.0.
- Add optional `TopicPrefix` config (`TOPIC_PREFIX` env / `--topic-prefix`) for explicit Kafka
  topic prefixing, independent of the existing `Branch` field. `NewKafkaResultDeliverer` now
  takes `base.TopicPrefix` instead of `base.Branch`; `Branch` is kept for its other uses.
