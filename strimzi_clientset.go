// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strimzi

import (
	"context"
	"github.com/bborbe/strimzi/k8s/client/clientset/versioned"

	"github.com/bborbe/k8s"
	"github.com/pkg/errors"
)

func CreateClientset(ctx context.Context, kubeconfig string) (versioned.Interface, error) {
	config, err := k8s.CreateConfig(kubeconfig)
	if err != nil {
		return nil, errors.Wrap(err, "create k8s config failed")
	}
	return versioned.NewForConfig(config)
}
