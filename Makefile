
default: precommit

precommit: deps ensure format generate test check addlicense
	@echo "ready to commit"

deps:
	go install github.com/onsi/ginkgo/v2/ginkgo

ensure:
	go mod tidy
	go mod verify
	rm -rf vendor

format:
	find . -type f -name '*.go' -not -path './vendor/*' -exec gofmt -w "{}" +
	find . -type f -name '*.go' -not -path './vendor/*' -exec go run -mod=mod github.com/incu6us/goimports-reviser -project-name github.com/bborbe/strimzi -file-path "{}" \;

generate:
	rm -rf mocks avro
	go generate -mod=mod ./...

test:
	go test -mod=mod -p=$${GO_TEST_PARALLEL:-1} -cover -race $(shell go list -mod=mod ./... | grep -v /vendor/)

check: vet errcheck vulncheck

vet:
	go vet -mod=mod $(shell go list -mod=mod ./... | grep -v /vendor/)

errcheck:
	go run -mod=mod github.com/kisielk/errcheck -ignore '(Close|Write|Fprint)' $(shell go list -mod=mod ./... | grep -v /vendor/ | grep -v k8s/client/informers/externalversions)

addlicense:
	go run -mod=mod github.com/google/addlicense -c "Benjamin Borbe" -y $$(date +'%Y') -l bsd $$(find . -name "*.go" -not -path './vendor/*')

vulncheck:
	go run -mod=mod golang.org/x/vuln/cmd/govulncheck $(shell go list -mod=mod ./... | grep -v /vendor/)

generatek8s:
	go mod vendor
	bash hack/update-codegen.sh
	rm -rf vendor
