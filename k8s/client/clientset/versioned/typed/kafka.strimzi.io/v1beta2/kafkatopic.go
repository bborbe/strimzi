// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Code generated by client-gen. DO NOT EDIT.

package v1beta2

import (
	context "context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"

	kafkastrimziiov1beta2 "github.com/bborbe/strimzi/k8s/apis/kafka.strimzi.io/v1beta2"
	applyconfigurationkafkastrimziiov1beta2 "github.com/bborbe/strimzi/k8s/client/applyconfiguration/kafka.strimzi.io/v1beta2"
	scheme "github.com/bborbe/strimzi/k8s/client/clientset/versioned/scheme"
)

// KafkaTopicsGetter has a method to return a KafkaTopicInterface.
// A group's client should implement this interface.
type KafkaTopicsGetter interface {
	KafkaTopics(namespace string) KafkaTopicInterface
}

// KafkaTopicInterface has methods to work with KafkaTopic resources.
type KafkaTopicInterface interface {
	Create(ctx context.Context, kafkaTopic *kafkastrimziiov1beta2.KafkaTopic, opts v1.CreateOptions) (*kafkastrimziiov1beta2.KafkaTopic, error)
	Update(ctx context.Context, kafkaTopic *kafkastrimziiov1beta2.KafkaTopic, opts v1.UpdateOptions) (*kafkastrimziiov1beta2.KafkaTopic, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*kafkastrimziiov1beta2.KafkaTopic, error)
	List(ctx context.Context, opts v1.ListOptions) (*kafkastrimziiov1beta2.KafkaTopicList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *kafkastrimziiov1beta2.KafkaTopic, err error)
	Apply(ctx context.Context, kafkaTopic *applyconfigurationkafkastrimziiov1beta2.KafkaTopicApplyConfiguration, opts v1.ApplyOptions) (result *kafkastrimziiov1beta2.KafkaTopic, err error)
	KafkaTopicExpansion
}

// kafkaTopics implements KafkaTopicInterface
type kafkaTopics struct {
	*gentype.ClientWithListAndApply[*kafkastrimziiov1beta2.KafkaTopic, *kafkastrimziiov1beta2.KafkaTopicList, *applyconfigurationkafkastrimziiov1beta2.KafkaTopicApplyConfiguration]
}

// newKafkaTopics returns a KafkaTopics
func newKafkaTopics(c *KafkaV1beta2Client, namespace string) *kafkaTopics {
	return &kafkaTopics{
		gentype.NewClientWithListAndApply[*kafkastrimziiov1beta2.KafkaTopic, *kafkastrimziiov1beta2.KafkaTopicList, *applyconfigurationkafkastrimziiov1beta2.KafkaTopicApplyConfiguration](
			"kafkatopics",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *kafkastrimziiov1beta2.KafkaTopic { return &kafkastrimziiov1beta2.KafkaTopic{} },
			func() *kafkastrimziiov1beta2.KafkaTopicList { return &kafkastrimziiov1beta2.KafkaTopicList{} },
		),
	}
}
