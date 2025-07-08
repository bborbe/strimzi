// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strimzi_test

import (
	"context"
	"os"
	"path/filepath"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/strimzi"
)

var _ = Describe("CreateClientset", func() {
	var ctx context.Context

	BeforeEach(func() {
		ctx = context.Background()
	})

	Context("with valid kubeconfig", func() {
		It("returns clientset successfully", func() {
			Skip("requires valid kubeconfig - integration test")
			kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
			clientset, err := strimzi.CreateClientset(ctx, kubeconfig)
			Expect(err).To(BeNil())
			Expect(clientset).ToNot(BeNil())
		})
	})

	Context("with empty kubeconfig path", func() {
		It("uses in-cluster config", func() {
			Skip("requires running in cluster - integration test")
			clientset, err := strimzi.CreateClientset(ctx, "")
			Expect(err).To(BeNil())
			Expect(clientset).ToNot(BeNil())
		})
	})

	Context("with invalid kubeconfig path", func() {
		It("returns error for non-existent file", func() {
			clientset, err := strimzi.CreateClientset(ctx, "/non/existent/path")
			Expect(err).To(HaveOccurred())
			Expect(clientset).To(BeNil())
			Expect(err.Error()).To(ContainSubstring("create k8s config failed"))
		})
	})

	Context("with invalid kubeconfig content", func() {
		var tempFile string

		BeforeEach(func() {
			tmpFile, err := os.CreateTemp("", "invalid-kubeconfig-*.yaml")
			Expect(err).To(BeNil())
			tempFile = tmpFile.Name()

			_, err = tmpFile.WriteString("invalid yaml content: [")
			Expect(err).To(BeNil())
			tmpFile.Close()
		})

		AfterEach(func() {
			if tempFile != "" {
				_ = os.Remove(tempFile)
			}
		})

		It("returns error for invalid kubeconfig", func() {
			clientset, err := strimzi.CreateClientset(ctx, tempFile)
			Expect(err).To(HaveOccurred())
			Expect(clientset).To(BeNil())
			Expect(err.Error()).To(ContainSubstring("create k8s config failed"))
		})
	})

	Context("with malformed but valid YAML", func() {
		var tempFile string

		BeforeEach(func() {
			tmpFile, err := os.CreateTemp("", "malformed-kubeconfig-*.yaml")
			Expect(err).To(BeNil())
			tempFile = tmpFile.Name()

			// Valid YAML but not a k8s config
			_, err = tmpFile.WriteString(`
apiVersion: v1
kind: Config
metadata:
  name: not-a-k8s-config
spec:
  someField: value
`)
			Expect(err).To(BeNil())
			tmpFile.Close()
		})

		AfterEach(func() {
			if tempFile != "" {
				_ = os.Remove(tempFile)
			}
		})

		It("returns error for malformed k8s config", func() {
			clientset, err := strimzi.CreateClientset(ctx, tempFile)
			Expect(err).To(HaveOccurred())
			Expect(clientset).To(BeNil())
		})
	})

	Context("with context cancellation", func() {
		It("handles cancelled context gracefully", func() {
			cancelledCtx, cancel := context.WithCancel(context.Background())
			cancel() // Cancel immediately

			clientset, err := strimzi.CreateClientset(cancelledCtx, "/some/path")
			Expect(err).To(HaveOccurred())
			Expect(clientset).To(BeNil())
		})
	})

	Context("with context timeout", func() {
		It("handles context timeout gracefully", func() {
			timeoutCtx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
			defer cancel()

			// Give time for timeout to occur
			time.Sleep(1 * time.Millisecond)

			clientset, err := strimzi.CreateClientset(timeoutCtx, "/some/path")
			Expect(err).To(HaveOccurred())
			Expect(clientset).To(BeNil())
		})
	})
})
