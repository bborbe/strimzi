
default: precommit

precommit: ensure format generate test check addlicense
	@echo "ready to commit"

ensure:
	go mod tidy
	go mod verify
	go mod vendor

format:
	find . -type f -name '*.go' -not -path './vendor/*' -exec gofmt -w "{}" +
	find . -type f -name '*.go' -not -path './vendor/*' -exec go run -mod=vendor github.com/incu6us/goimports-reviser -project-name github.com/bborbe/strimzi -file-path "{}" \;

generate:
	rm -rf mocks avro
	go generate -mod=vendor ./...

test:
	go test -mod=vendor -p=$${GO_TEST_PARALLEL:-1} -cover -race $(shell go list -mod=vendor ./... | grep -v /vendor/)

check: lint vet errcheck vulncheck

lint:
	go run -mod=vendor golang.org/x/lint/golint -min_confidence 1 $(shell go list -mod=vendor ./... | grep -v /vendor/)

vet:
	go vet -mod=vendor $(shell go list -mod=vendor ./... | grep -v /vendor/)

errcheck:
	go run -mod=vendor github.com/kisielk/errcheck -ignore '(Close|Write|Fprint)' $(shell go list -mod=vendor ./... | grep -v /vendor/)

addlicense:
	go run -mod=vendor github.com/google/addlicense -c "Benjamin Borbe" -y $$(date +'%Y') -l bsd $$(find . -name "*.go" -not -path './vendor/*')

vulncheck:
	go run -mod=vendor golang.org/x/vuln/cmd/govulncheck $(shell go list -mod=vendor ./... | grep -v /vendor/)

generatek8s:
	rm -rf k8s/client ${GOPATH}/src/github.com/bborbe/strimzi
	chmod a+x vendor/k8s.io/code-generator/*.sh
	bash vendor/k8s.io/code-generator/generate-groups.sh applyconfiguration,client,deepcopy,informer,lister \
	github.com/bborbe/strimzi/k8s/client github.com/bborbe/strimzi/k8s/apis \
	kafka.strimzi.io:v1beta2
	cp -R ${GOPATH}/src/github.com/bborbe/strimzi/k8s .
