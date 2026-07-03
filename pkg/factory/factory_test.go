// Copyright (c) 2026 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package factory_test

import (
	"context"

	agentlib "github.com/bborbe/agent"
	"github.com/bborbe/cqrs/base"
	kafkamocks "github.com/bborbe/kafka/mocks"
	libtime "github.com/bborbe/time"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/agent-gemini/pkg/factory"
)

var _ = Describe("CreateAgentProvider", func() {
	var (
		ctx      context.Context
		provider agentlib.AgentProvider
	)

	BeforeEach(func() {
		ctx = context.Background()
		// nil parser is safe — NewGeminiStep stores the parser without invoking it.
		provider = factory.CreateAgentProvider(nil)
	})

	It("returns a non-nil provider", func() {
		Expect(provider).NotTo(BeNil())
	})

	It("Get returns the liveness agent for TaskTypeHealthcheck", func() {
		agent, err := provider.Get(ctx, agentlib.TaskTypeHealthcheck)
		Expect(err).To(BeNil())
		Expect(agent).NotTo(BeNil())
	})

	Describe("Get with unknown task_type", func() {
		DescribeTable(
			"error shape",
			func(taskType agentlib.TaskType, expectedSubstr string) {
				agent, err := provider.Get(ctx, taskType)
				Expect(agent).To(BeNil())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("unknown task_type"))
				Expect(err.Error()).To(ContainSubstring(expectedSubstr))
				Expect(err.Error()).To(ContainSubstring("agent-gemini"))
				Expect(err.Error()).To(ContainSubstring("[healthcheck]"))
			},
			Entry(
				"literal gemini rejected (no implicit domain type)",
				agentlib.TaskType("gemini"),
				`"gemini"`,
			),
			Entry("bogus value", agentlib.TaskType("bogus"), `"bogus"`),
			Entry("empty value", agentlib.TaskType(""), `""`),
		)
	})
})

var _ = Describe("CreateKafkaResultDeliverer", func() {
	DescribeTable(
		"publishes the task update to the topic derived from the topic prefix",
		func(topicPrefix base.TopicPrefix, expectedTopic string) {
			ctx := context.Background()
			fakeProducer := &kafkamocks.KafkaSyncProducer{}

			deliverer := factory.CreateKafkaResultDeliverer(
				fakeProducer,
				topicPrefix,
				agentlib.TaskIdentifier("task-1"),
				"## Plan\n\nsome content\n",
				libtime.NewCurrentDateTime(),
			)

			err := deliverer.DeliverResult(ctx, agentlib.AgentResultInfo{
				Status: agentlib.AgentStatusDone,
				Output: "## Plan\n\nsome content\n",
			})
			Expect(err).To(BeNil())

			Expect(fakeProducer.SendMessageCallCount()).To(Equal(1))
			_, msg := fakeProducer.SendMessageArgsForCall(0)
			Expect(msg.Topic).To(Equal(expectedTopic))
		},
		Entry("develop prefix", base.TopicPrefix("develop"), "develop-agent-task-v1-request"),
		Entry("master prefix", base.TopicPrefix("master"), "master-agent-task-v1-request"),
		Entry("empty prefix", base.TopicPrefix(""), "agent-task-v1-request"),
	)
})
