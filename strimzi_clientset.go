// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package strimzi provides Kubernetes client bindings for Strimzi Kafka custom resources.
//
// This package generates typed Go clients for Kafka CRDs (Custom Resource Definitions)
// using Kubernetes code generation tools, enabling programmatic management of Strimzi
// Kafka resources in your Go applications.
//
// Example usage:
//
//	ctx := context.Background()
//	clientset, err := strimzi.CreateClientset(ctx, "")
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	topics, err := clientset.KafkaV1beta2().KafkaTopics("kafka").List(ctx, metav1.ListOptions{})
//	if err != nil {
//	    log.Fatal(err)
//	}
package strimzi

import (
	"context"

	"github.com/bborbe/errors"
	"github.com/bborbe/k8s"

	"github.com/bborbe/strimzi/k8s/client/clientset/versioned"
)

// StrimziClientset provides access to all Strimzi Kafka custom resource clients.
// It includes typed clients for KafkaTopics and other Strimzi resources.
//
//counterfeiter:generate -o mocks/strimzi-clientset.go --fake-name StrimziClientset . StrimziClientset
//nolint:revive // StrimziClientset name is intentional for clarity when used as a standalone package
type StrimziClientset = versioned.Interface

// CreateClientset creates a new Strimzi clientset for interacting with Kafka custom resources.
//
// Parameters:
//   - ctx: Context for the operation
//   - kubeconfig: Path to kubeconfig file. If empty, uses default kubeconfig resolution
//     (in-cluster config if running in pod, or ~/.kube/config if running locally)
//
// Returns:
//   - StrimziClientset: A clientset for accessing Strimzi Kafka resources
//   - error: Any error that occurred during clientset creation
//
// Example:
//
//	// Use default kubeconfig
//	clientset, err := strimzi.CreateClientset(ctx, "")
//
//	// Use specific kubeconfig file
//	clientset, err := strimzi.CreateClientset(ctx, "/path/to/kubeconfig")
func CreateClientset(ctx context.Context, kubeconfig string) (StrimziClientset, error) {
	config, err := k8s.CreateConfig(kubeconfig)
	if err != nil {
		return nil, errors.Wrap(ctx, err, "create k8s config failed")
	}
	return versioned.NewForConfig(config)
}
