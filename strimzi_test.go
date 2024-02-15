// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package topic_test

import (
	"encoding/json"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/strimzi/k8s/apis/kafka.strimzi.io/v1beta2"
)

const topicJSON = `
{
  "apiVersion": "kafka.strimzi.io/v1beta2beta2",
  "kind": "KafkaTopic",
  "metadata": {
    "name": "my-topic"
  },
  "spec": {
    "partitions": 1,
    "replicas": 3,
    "config": {
      "retention.ms": "-1",
      "retention.bytes": "-1",
      "cleanup.policy": "compact"
    }
  }
}
`

var _ = Describe("KafkaTopic", func() {
	var topic v1beta2.KafkaTopic
	var err error

	BeforeEach(func() {
		topic = v1beta2.KafkaTopic{}
		err = json.Unmarshal([]byte(topicJSON), &topic)
	})
	It("returns no error", func() {
		Expect(err).To(BeNil())
	})
	It("contains config values", func() {
		Expect(topic.Spec.Config).To(HaveKeyWithValue("retention.ms", "-1"))
		Expect(topic.Spec.Config).To(HaveKeyWithValue("retention.bytes", "-1"))
		Expect(topic.Spec.Config).To(HaveKeyWithValue("cleanup.policy", "compact"))
	})
	It("contains partitions", func() {
		Expect(*topic.Spec.Replicas).To(Equal(int32(3)))
	})
	It("contains replicas", func() {
		Expect(*topic.Spec.Partitions).To(Equal(int32(1)))
	})
})
