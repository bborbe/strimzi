// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strimzi

import "github.com/bborbe/strimzi/k8s/client/clientset/versioned/typed/kafka.strimzi.io/v1beta2"

//counterfeiter:generate -o mocks/kafka-topic-interface.go --fake-name KafkaTopicInterface . KafkaTopicInterface
type KafkaTopicInterface = v1beta2.KafkaTopicInterface
