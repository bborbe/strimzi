// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v1beta2_test

import (
	"github.com/bborbe/collection"
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
		Context("empty topicName", func() {
			BeforeEach(func() {
				kafkaTopic.ObjectMeta.Name = "meta-name"
				empty := ""
				kafkaTopic.Spec.TopicName = &empty
			})
			It("returns metadata name", func() {
				Expect(topicName).To(Equal("meta-name"))
			})
		})
	})

	Context("Equal", func() {
		var kafkaTopic1, kafkaTopic2 v1beta2.KafkaTopic
		var isEqual bool

		BeforeEach(func() {
			partitions := int32(3)
			replicas := int32(2)
			kafkaTopic1 = v1beta2.KafkaTopic{
				ObjectMeta: metav1.ObjectMeta{Name: "test-topic"},
				Spec: &v1beta2.KafkaTopicSpec{
					Partitions: &partitions,
					Replicas:   &replicas,
					Config:     map[string]string{"cleanup.policy": "delete"},
				},
			}
			kafkaTopic2 = v1beta2.KafkaTopic{
				ObjectMeta: metav1.ObjectMeta{Name: "test-topic"},
				Spec: &v1beta2.KafkaTopicSpec{
					Partitions: &partitions,
					Replicas:   &replicas,
					Config:     map[string]string{"cleanup.policy": "delete"},
				},
			}
		})

		JustBeforeEach(func() {
			isEqual = kafkaTopic1.Equal(kafkaTopic2)
		})

		Context("identical topics", func() {
			It("returns true", func() {
				Expect(isEqual).To(BeTrue())
			})
		})

		Context("different topic names", func() {
			BeforeEach(func() {
				kafkaTopic2.ObjectMeta.Name = "different-topic"
			})
			It("returns false", func() {
				Expect(isEqual).To(BeFalse())
			})
		})

		Context("different partitions", func() {
			BeforeEach(func() {
				differentPartitions := int32(5)
				kafkaTopic2.Spec.Partitions = &differentPartitions
			})
			It("returns false", func() {
				Expect(isEqual).To(BeFalse())
			})
		})

		Context("different replicas", func() {
			BeforeEach(func() {
				differentReplicas := int32(3)
				kafkaTopic2.Spec.Replicas = &differentReplicas
			})
			It("returns false", func() {
				Expect(isEqual).To(BeFalse())
			})
		})

		Context("different config", func() {
			BeforeEach(func() {
				kafkaTopic2.Spec.Config = map[string]string{"cleanup.policy": "compact"}
			})
			It("returns false", func() {
				Expect(isEqual).To(BeFalse())
			})
		})

		Context("nil partitions", func() {
			BeforeEach(func() {
				kafkaTopic1.Spec.Partitions = nil
				kafkaTopic2.Spec.Partitions = nil
			})
			It("returns true", func() {
				Expect(isEqual).To(BeTrue())
			})
		})

		Context("nil replicas", func() {
			BeforeEach(func() {
				kafkaTopic1.Spec.Replicas = nil
				kafkaTopic2.Spec.Replicas = nil
			})
			It("returns true", func() {
				Expect(isEqual).To(BeTrue())
			})
		})

		Context("nil config", func() {
			BeforeEach(func() {
				kafkaTopic1.Spec.Config = nil
				kafkaTopic2.Spec.Config = nil
			})
			It("returns true", func() {
				Expect(isEqual).To(BeTrue())
			})
		})

		Context("one nil partitions, one set", func() {
			BeforeEach(func() {
				kafkaTopic1.Spec.Partitions = nil
				partitions := int32(3)
				kafkaTopic2.Spec.Partitions = &partitions
			})
			It("returns false", func() {
				Expect(isEqual).To(BeFalse())
			})
		})

		Context("custom topic names", func() {
			BeforeEach(func() {
				customName := "custom-topic"
				kafkaTopic1.Spec.TopicName = &customName
				kafkaTopic2.Spec.TopicName = &customName
			})
			It("returns true", func() {
				Expect(isEqual).To(BeTrue())
			})
		})

		Context("different custom topic names", func() {
			BeforeEach(func() {
				customName1 := "custom-topic-1"
				customName2 := "custom-topic-2"
				kafkaTopic1.Spec.TopicName = &customName1
				kafkaTopic2.Spec.TopicName = &customName2
			})
			It("returns false", func() {
				Expect(isEqual).To(BeFalse())
			})
		})

		Context("nil spec comparison", func() {
			BeforeEach(func() {
				kafkaTopic1.Spec = nil
				kafkaTopic2.Spec = nil
			})
			It("returns true", func() {
				Expect(isEqual).To(BeTrue())
			})
		})

		Context("one nil spec, one set", func() {
			BeforeEach(func() {
				kafkaTopic1.Spec = nil
			})
			It("returns false", func() {
				Expect(isEqual).To(BeFalse())
			})
		})

		Context("one nil replicas, one set", func() {
			BeforeEach(func() {
				kafkaTopic1.Spec.Replicas = nil
				replicas := int32(2)
				kafkaTopic2.Spec.Replicas = &replicas
			})
			It("returns false", func() {
				Expect(isEqual).To(BeFalse())
			})
		})

		Context("empty vs nil config", func() {
			BeforeEach(func() {
				kafkaTopic1.Spec.Config = map[string]string{}
				kafkaTopic2.Spec.Config = nil
			})
			It("returns false", func() {
				Expect(isEqual).To(BeFalse())
			})
		})

		Context("zero value partitions", func() {
			BeforeEach(func() {
				zero := int32(0)
				kafkaTopic1.Spec.Partitions = &zero
				kafkaTopic2.Spec.Partitions = &zero
			})
			It("returns true", func() {
				Expect(isEqual).To(BeTrue())
			})
		})

		Context("one custom topic name, one metadata name", func() {
			BeforeEach(func() {
				kafkaTopic1.ObjectMeta.Name = "metadata-name"
				kafkaTopic2.ObjectMeta.Name = "different-name"
				customName := "metadata-name"
				kafkaTopic2.Spec.TopicName = &customName
			})
			It("returns true", func() {
				Expect(isEqual).To(BeTrue())
			})
		})
	})
})

var _ = Describe("KafkaTopicSpec", func() {
	Context("creation", func() {
		It("can be created with all fields", func() {
			partitions := int32(3)
			replicas := int32(2)
			topicName := "test-topic"
			config := map[string]string{"cleanup.policy": "delete"}

			spec := v1beta2.KafkaTopicSpec{
				Partitions: &partitions,
				Replicas:   &replicas,
				TopicName:  &topicName,
				Config:     config,
			}

			Expect(spec.Partitions).To(Equal(&partitions))
			Expect(spec.Replicas).To(Equal(&replicas))
			Expect(spec.TopicName).To(Equal(&topicName))
			Expect(spec.Config).To(Equal(config))
		})

		It("can be created with nil fields", func() {
			spec := v1beta2.KafkaTopicSpec{
				Partitions: nil,
				Replicas:   nil,
				TopicName:  nil,
				Config:     nil,
			}

			Expect(spec.Partitions).To(BeNil())
			Expect(spec.Replicas).To(BeNil())
			Expect(spec.TopicName).To(BeNil())
			Expect(spec.Config).To(BeNil())
		})
	})
})

var _ = Describe("KafkaTopicStatus", func() {
	Context("creation", func() {
		It("can be created with all fields", func() {
			generation := int32(1)
			topicName := "test-topic"
			condition := v1beta2.KafkaTopicStatusConditionsElem{
				Type:   collection.Ptr("Ready"),
				Status: collection.Ptr("True"),
			}

			status := v1beta2.KafkaTopicStatus{
				ObservedGeneration: &generation,
				TopicName:          &topicName,
				Conditions:         []v1beta2.KafkaTopicStatusConditionsElem{condition},
			}

			Expect(status.ObservedGeneration).To(Equal(&generation))
			Expect(status.TopicName).To(Equal(&topicName))
			Expect(status.Conditions).To(HaveLen(1))
			Expect(status.Conditions[0].Type).To(Equal(collection.Ptr("Ready")))
		})
	})
})

var _ = Describe("KafkaTopicStatusConditionsElem", func() {
	Context("creation", func() {
		It("can be created with all fields", func() {
			condition := v1beta2.KafkaTopicStatusConditionsElem{
				Type:               collection.Ptr("Ready"),
				Status:             collection.Ptr("True"),
				Reason:             collection.Ptr("TopicCreated"),
				Message:            collection.Ptr("Topic created successfully"),
				LastTransitionTime: collection.Ptr("2023-01-01T00:00:00Z"),
			}

			Expect(condition.Type).To(Equal(collection.Ptr("Ready")))
			Expect(condition.Status).To(Equal(collection.Ptr("True")))
			Expect(condition.Reason).To(Equal(collection.Ptr("TopicCreated")))
			Expect(condition.Message).To(Equal(collection.Ptr("Topic created successfully")))
			Expect(condition.LastTransitionTime).To(Equal(collection.Ptr("2023-01-01T00:00:00Z")))
		})
	})
})

var _ = Describe("KafkaTopicList", func() {
	Context("creation", func() {
		It("can be created with topics", func() {
			topic1 := v1beta2.KafkaTopic{
				ObjectMeta: metav1.ObjectMeta{Name: "topic1"},
				Spec:       &v1beta2.KafkaTopicSpec{},
			}
			topic2 := v1beta2.KafkaTopic{
				ObjectMeta: metav1.ObjectMeta{Name: "topic2"},
				Spec:       &v1beta2.KafkaTopicSpec{},
			}

			list := v1beta2.KafkaTopicList{
				Items: []v1beta2.KafkaTopic{topic1, topic2},
			}

			Expect(list.Items).To(HaveLen(2))
			Expect(list.Items[0].Name).To(Equal("topic1"))
			Expect(list.Items[1].Name).To(Equal("topic2"))
		})
	})
})

var _ = Describe("KafkaTopics", func() {
	Context("slice operations", func() {
		It("can be created and accessed", func() {
			topic1 := v1beta2.KafkaTopic{
				ObjectMeta: metav1.ObjectMeta{Name: "topic1"},
				Spec:       &v1beta2.KafkaTopicSpec{},
			}
			topic2 := v1beta2.KafkaTopic{
				ObjectMeta: metav1.ObjectMeta{Name: "topic2"},
				Spec:       &v1beta2.KafkaTopicSpec{},
			}

			topics := v1beta2.KafkaTopics{topic1, topic2}

			Expect(topics).To(HaveLen(2))
			Expect(topics[0].Name).To(Equal("topic1"))
			Expect(topics[1].Name).To(Equal("topic2"))
		})
	})
})

var _ = Describe("KafkaTopicSpecs", func() {
	Context("slice operations", func() {
		It("can be created and accessed", func() {
			partitions := int32(3)
			spec1 := v1beta2.KafkaTopicSpec{
				Partitions: &partitions,
			}
			spec2 := v1beta2.KafkaTopicSpec{
				Partitions: &partitions,
			}

			specs := v1beta2.KafkaTopicSpecs{spec1, spec2}

			Expect(specs).To(HaveLen(2))
			Expect(specs[0].Partitions).To(Equal(&partitions))
			Expect(specs[1].Partitions).To(Equal(&partitions))
		})
	})
})
