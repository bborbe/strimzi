// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strimzi

import (
	"context"

	"github.com/bborbe/errors"
	"github.com/golang/glog"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/bborbe/strimzi/k8s/apis/kafka.strimzi.io/v1beta2"
	"github.com/bborbe/strimzi/k8s/client/clientset/versioned"
)

//counterfeiter:generate -o mocks/topic-deployer.go --fake-name TopicDeployer . TopicDeployer

// TopicDeployer provides operations for deploying and managing Kafka topics in Kubernetes.
// It handles both creation and updates of KafkaTopic custom resources, as well as their removal.
type TopicDeployer interface {
	// Deploy creates or updates a KafkaTopic resource in Kubernetes.
	// If the topic doesn't exist, it will be created. If it exists, it will be updated
	// with the new configuration while preserving the resource version.
	Deploy(ctx context.Context, topic v1beta2.KafkaTopic) error

	// Undeploy removes a KafkaTopic resource from Kubernetes.
	// If the topic doesn't exist, the operation succeeds silently.
	Undeploy(ctx context.Context, namespace string, name string) error
}

// NewTopicDeployer creates a new TopicDeployer instance.
//
// Parameters:
//   - clientset: Strimzi clientset for interacting with KafkaTopic resources
//
// Returns:
//   - TopicDeployer: A new deployer instance for managing Kafka topics
func NewTopicDeployer(
	clientset versioned.Interface,
) TopicDeployer {
	return &topicDeployer{
		clientset: clientset,
	}
}

type topicDeployer struct {
	clientset versioned.Interface
}

func (t *topicDeployer) Deploy(ctx context.Context, topic v1beta2.KafkaTopic) error {
	currentTopic, err := t.clientset.KafkaV1beta2().
		KafkaTopics(topic.Namespace).
		Get(ctx, topic.Name, metav1.GetOptions{})
	if err != nil {
		glog.V(3).Infof("get topic %s failed: %s", topic.Name, err)
		_, err = t.clientset.KafkaV1beta2().
			KafkaTopics(topic.Namespace).
			Create(ctx, &topic, metav1.CreateOptions{})
		if err != nil {
			return errors.Wrap(ctx, err, "create topic failed")
		}
		glog.V(3).Infof("topic %s created successful", topic.Name)
		return nil
	}
	updateTopic := mergeTopic(*currentTopic, topic)
	_, err = t.clientset.KafkaV1beta2().
		KafkaTopics(topic.Namespace).
		Update(ctx, &updateTopic, metav1.UpdateOptions{})
	if err != nil {
		return errors.Wrap(ctx, err, "update topic failed")
	}
	glog.V(3).Infof("topic %s updated successful", topic.Name)
	return nil
}

func (t *topicDeployer) Undeploy(ctx context.Context, namespace string, name string) error {
	_, err := t.clientset.KafkaV1beta2().KafkaTopics(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		glog.V(3).Infof("topic '%s' not found => skip", name)
		return nil
	}
	if err := t.clientset.KafkaV1beta2().KafkaTopics(namespace).Delete(ctx, name, metav1.DeleteOptions{}); err != nil {
		return err
	}
	glog.V(3).Infof("delete %s completed", name)
	return nil
}

func mergeTopic(currentTopic, newTopic v1beta2.KafkaTopic) v1beta2.KafkaTopic {
	newTopic.ObjectMeta.ResourceVersion = currentTopic.ObjectMeta.ResourceVersion
	return newTopic
}
