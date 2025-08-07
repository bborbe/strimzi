# Strimzi Go Client Library

[![Go Reference](https://pkg.go.dev/badge/github.com/bborbe/strimzi.svg)](https://pkg.go.dev/github.com/bborbe/strimzi)
[![Go Report Card](https://goreportcard.com/badge/github.com/bborbe/strimzi)](https://goreportcard.com/report/github.com/bborbe/strimzi)
[![License: BSD](https://img.shields.io/badge/License-BSD-blue.svg)](LICENSE)

A Go library that provides Kubernetes client bindings for Strimzi Kafka custom resources. This library generates typed Go clients for Kafka CRDs (Custom Resource Definitions) using Kubernetes code generation tools, enabling programmatic management of Strimzi Kafka resources in your Go applications.

## Features

- 🎯 **Typed Kubernetes Clients**: Auto-generated clientsets, informers, and listers for Strimzi Kafka CRDs
- 🔧 **Code Generation**: Built using Kubernetes client-go code generation tools
- 🧪 **Well Tested**: Comprehensive test coverage using Ginkgo v2 and Gomega
- 📦 **Easy Integration**: Simple API for creating and managing Kafka resources
- 🏗️ **Production Ready**: Follows Kubernetes client-go patterns and conventions

## Installation

```bash
go get github.com/bborbe/strimzi
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/bborbe/strimzi"
)

func main() {
    ctx := context.Background()
    
    // Create a clientset (uses default kubeconfig)
    clientset, err := strimzi.CreateClientset(ctx, "")
    if err != nil {
        log.Fatal(err)
    }
    
    // List Kafka topics
    topics, err := clientset.KafkaV1beta2().KafkaTopics("kafka").List(ctx, metav1.ListOptions{})
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Found %d Kafka topics\n", len(topics.Items))
}
```

## API Documentation

For comprehensive API documentation, visit [pkg.go.dev/github.com/bborbe/strimzi](https://pkg.go.dev/github.com/bborbe/strimzi).

## Development

This project uses standard Go development tools and follows Benjamin Borbe's coding guidelines.

### Prerequisites

- Go 1.24+
- Kubernetes cluster with Strimzi operator installed (for integration testing)

### Building

```bash
make precommit  # Run all quality checks
make test       # Run tests
make generatek8s # Regenerate Kubernetes client code
```

### Testing

```bash
# Run all tests
make test

# Run specific tests
ginkgo run pkg/

# Generate mocks
go generate ./...
```

## Project Structure

- `k8s/apis/kafka.strimzi.io/v1beta2/` - API definitions and types for Kafka custom resources
- `k8s/client/` - Auto-generated Kubernetes clients (clientset, informers, listers, apply configurations)
- `strimzi_clientset.go` - Main entry point for creating clientsets
- `hack/update-codegen.sh` - Kubernetes code generation script

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Run quality checks (`make precommit`)
4. Commit your changes (`git commit -m 'Add amazing feature'`)
5. Push to the branch (`git push origin feature/amazing-feature`)
6. Open a Pull Request

## License

This project is licensed under the BSD License - see the [LICENSE](LICENSE) file for details.
