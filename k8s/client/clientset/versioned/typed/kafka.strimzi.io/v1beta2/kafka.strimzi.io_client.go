// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Code generated by client-gen. DO NOT EDIT.

package v1beta2

import (
	http "net/http"

	rest "k8s.io/client-go/rest"

	kafkastrimziiov1beta2 "github.com/bborbe/strimzi/k8s/apis/kafka.strimzi.io/v1beta2"
	scheme "github.com/bborbe/strimzi/k8s/client/clientset/versioned/scheme"
)

type KafkaV1beta2Interface interface {
	RESTClient() rest.Interface
	KafkaTopicsGetter
}

// KafkaV1beta2Client is used to interact with features provided by the kafka group.
type KafkaV1beta2Client struct {
	restClient rest.Interface
}

func (c *KafkaV1beta2Client) KafkaTopics(namespace string) KafkaTopicInterface {
	return newKafkaTopics(c, namespace)
}

// NewForConfig creates a new KafkaV1beta2Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*KafkaV1beta2Client, error) {
	config := *c
	setConfigDefaults(&config)
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new KafkaV1beta2Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*KafkaV1beta2Client, error) {
	config := *c
	setConfigDefaults(&config)
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &KafkaV1beta2Client{client}, nil
}

// NewForConfigOrDie creates a new KafkaV1beta2Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *KafkaV1beta2Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new KafkaV1beta2Client for the given RESTClient.
func New(c rest.Interface) *KafkaV1beta2Client {
	return &KafkaV1beta2Client{c}
}

func setConfigDefaults(config *rest.Config) {
	gv := kafkastrimziiov1beta2.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = rest.CodecFactoryForGeneratedClient(scheme.Scheme, scheme.Codecs).WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *KafkaV1beta2Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
