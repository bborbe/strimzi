// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strimzi

import "github.com/bborbe/strimzi/k8s/client/clientset/versioned/typed/kafka.strimzi.io/v1beta2"

//counterfeiter:generate -o mocks/kafka-v1-beta2-interface.go --fake-name KafkaV1beta2Interface . KafkaV1beta2Interface

// KafkaV1beta2Interface is a type alias for v1beta2.KafkaV1beta2Interface
// that enables mock generation using counterfeiter.
type KafkaV1beta2Interface v1beta2.KafkaV1beta2Interface

//counterfeiter:generate -o mocks/kafka-topic-interface.go --fake-name KafkaTopicInterface . KafkaTopicInterface

// KafkaTopicInterface is a type alias for v1beta2.KafkaTopicInterface
// that enables mock generation using counterfeiter.
type KafkaTopicInterface = v1beta2.KafkaTopicInterface
