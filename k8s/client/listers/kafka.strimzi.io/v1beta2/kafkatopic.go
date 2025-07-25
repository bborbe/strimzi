// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Code generated by lister-gen. DO NOT EDIT.

package v1beta2

import (
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"

	kafkastrimziiov1beta2 "github.com/bborbe/strimzi/k8s/apis/kafka.strimzi.io/v1beta2"
)

// KafkaTopicLister helps list KafkaTopics.
// All objects returned here must be treated as read-only.
type KafkaTopicLister interface {
	// List lists all KafkaTopics in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*kafkastrimziiov1beta2.KafkaTopic, err error)
	// KafkaTopics returns an object that can list and get KafkaTopics.
	KafkaTopics(namespace string) KafkaTopicNamespaceLister
	KafkaTopicListerExpansion
}

// kafkaTopicLister implements the KafkaTopicLister interface.
type kafkaTopicLister struct {
	listers.ResourceIndexer[*kafkastrimziiov1beta2.KafkaTopic]
}

// NewKafkaTopicLister returns a new KafkaTopicLister.
func NewKafkaTopicLister(indexer cache.Indexer) KafkaTopicLister {
	return &kafkaTopicLister{listers.New[*kafkastrimziiov1beta2.KafkaTopic](indexer, kafkastrimziiov1beta2.Resource("kafkatopic"))}
}

// KafkaTopics returns an object that can list and get KafkaTopics.
func (s *kafkaTopicLister) KafkaTopics(namespace string) KafkaTopicNamespaceLister {
	return kafkaTopicNamespaceLister{listers.NewNamespaced[*kafkastrimziiov1beta2.KafkaTopic](s.ResourceIndexer, namespace)}
}

// KafkaTopicNamespaceLister helps list and get KafkaTopics.
// All objects returned here must be treated as read-only.
type KafkaTopicNamespaceLister interface {
	// List lists all KafkaTopics in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*kafkastrimziiov1beta2.KafkaTopic, err error)
	// Get retrieves the KafkaTopic from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*kafkastrimziiov1beta2.KafkaTopic, error)
	KafkaTopicNamespaceListerExpansion
}

// kafkaTopicNamespaceLister implements the KafkaTopicNamespaceLister
// interface.
type kafkaTopicNamespaceLister struct {
	listers.ResourceIndexer[*kafkastrimziiov1beta2.KafkaTopic]
}
