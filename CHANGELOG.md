# Changelog

All notable changes to this project will be documented in this file.

Please choose versions by [Semantic Versioning](http://semver.org/).

* MAJOR version when you make incompatible API changes,
* MINOR version when you add functionality in a backwards-compatible manner, and
* PATCH version when you make backwards-compatible bug fixes.

## v1.8.2

- Bump bborbe/collection, errors, k8s dependencies
- Upgrade golangci-lint v2.11.4, counterfeiter v6.12.2, go-modtool v0.7.1
- Update Go toolchain to 1.26.2
- Add vuln ignore entries for bbolt and aws-sdk-go-v2 CVEs
- Improve vulncheck Makefile target with JSON filtering

## v1.8.1

- Update numerous indirect dependencies (docker, containerd, opentelemetry, go-openapi, moby/buildkit)
- Add replace directives for denis-tingaikin/go-header and opencontainers/runtime-spec
- Add go-openapi/swag sub-packages
- Remove josharian/intern and mailru/easyjson indirect deps

## v1.8.0

- upgrade k8s dependencies from v0.33 to v0.35
- migrate structured-merge-diff from v4 to v6
- add GetKind, GetAPIVersion, GetNamespace, IsApplyConfiguration methods to apply config types
- NewTypeConverter now returns managedfields.TypeConverter interface instead of concrete type
- update bborbe/* and other indirect dependencies

## v1.7.6

- chore: verified all tests pass, linting succeeds, and project meets Definition of Done criteria

## v1.7.5

- upgrade golangci-lint from v1 to v2
- standardize Makefile: add .PHONY declarations, multiline trivy, mocks mkdir
- update .golangci.yml to v2 format
- setup dark-factory config

## v1.7.4

- go mod update

## v1.7.3

- Update Go to 1.26.0

## v1.7.2

- Update Go to 1.25.7
- Update testing dependencies (ginkgo v2.28.1, gomega v1.39.1)
- Update bborbe dependencies (errors, k8s, math, parse, time, validation)
- Update various indirect dependencies for security and bug fixes

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
