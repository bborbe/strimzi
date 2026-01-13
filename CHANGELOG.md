# Changelog

All notable changes to this project will be documented in this file.

Please choose versions by [Semantic Versioning](http://semver.org/).

* MAJOR version when you make incompatible API changes,
* MINOR version when you add functionality in a backwards-compatible manner, and
* PATCH version when you make backwards-compatible bug fixes.

## v1.7.1

- Update Go to 1.25.5
- Update golang.org/x/crypto to v0.47.0
- Update dependencies

## v1.7.0

- update go and deps

## v1.6.1

- Add mock generation for KafkaV1beta2Interface and KafkaTopicInterface
- Consolidate mock type aliases into strimzi_mocks.go
- Improve code organization with proper GoDoc comments

## v1.6.0

- Add golangci-lint configuration and linting to build pipeline
- Enhance CI workflow with Trivy security scanning
- Update Go version to 1.25.2
- Add comprehensive security checks (gosec, osv-scanner, trivy)
- Integrate golines for improved code formatting
- Improve code readability with better line wrapping
- Update dependencies to latest versions

## v1.5.5

- Enhance README.md with library documentation
- Add detailed GoDoc comments to main functions and interfaces

## v1.5.4

- add mock for StrimziClientset 

## v1.5.3

- update dependencies to latest versions
- update Go to 1.24.5
- add license headers 
- regenerate mocks with latest counterfeiter

## v1.5.2

- add tests

## v1.5.1

- ignore claude files
- add make deps
- add tests

## v1.5.0

- update k8s generate
- go mod update

## v1.4.0

- remove vendor
- go mod update

## v1.3.1

- return versioned.Interface

## v1.3.2

- return *versioned.Clientset

## v1.3.0

- add strimzi k8s clientset

## v1.2.2

- go mod update

## v1.2.1

- Update to k8s 1.29

## v1.1.0

- Add topic deployer

## v1.0.0

- Initial Version
