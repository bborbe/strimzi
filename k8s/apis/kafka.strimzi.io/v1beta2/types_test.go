// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v1beta2_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/bborbe/strimzi/k8s/apis/kafka.strimzi.io/v1beta2"
)

var _ = Describe("KafkaTopic", func() {
	Context("TopicName", func() {
		var kafkaTopic v1beta2.KafkaTopic
		var topicName string
		BeforeEach(func() {
			kafkaTopic = v1beta2.KafkaTopic{
				ObjectMeta: metav1.ObjectMeta{},
				Spec:       &v1beta2.KafkaTopicSpec{},
			}
		})
		JustBeforeEach(func() {
			topicName = kafkaTopic.TopicName()
		})
		Context("no topicName", func() {
			BeforeEach(func() {
				kafkaTopic.ObjectMeta.Name = "meta-name"
				kafkaTopic.Spec.TopicName = nil
			})
			It("returns no error", func() {
				Expect(topicName).To(Equal("meta-name"))
			})
		})
		Context("topicName", func() {
			BeforeEach(func() {
				kafkaTopic.ObjectMeta.Name = "meta-name"
				s := "spec-name"
				kafkaTopic.Spec.TopicName = &s
			})
			It("returns no error", func() {
				Expect(topicName).To(Equal("spec-name"))
			})
		})
	})
})
