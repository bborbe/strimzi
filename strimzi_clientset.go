// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strimzi

import (
	"context"

	"github.com/bborbe/errors"
	"github.com/bborbe/k8s"

	"github.com/bborbe/strimzi/k8s/client/clientset/versioned"
)

//counterfeiter:generate -o mocks/strimzi-clientset.go --fake-name StrimziClientset . StrimziClientset
type StrimziClientset = versioned.Interface

func CreateClientset(ctx context.Context, kubeconfig string) (StrimziClientset, error) {
	config, err := k8s.CreateConfig(kubeconfig)
	if err != nil {
		return nil, errors.Wrap(ctx, err, "create k8s config failed")
	}
	return versioned.NewForConfig(config)
}
