
all: build


# Tool versions
GOLANGCI_LINT_VERSION?=v1.50.1

.PHONY: build
build: dt
	go build -o opgame .

.PHONY: clean
clean:
	- rm -f opgame

.PHONY: test
test:

.PHONY: test-go
test: test-go
test-go:
	go test ./...

.PHONY: test-sample
test: test-sample
test-sample:
	go run . < testdata/basic.txt
	
.PHONY: lint
test: lint
lint: tool/golangci-lint
	tool/golangci-lint run

.PHONY: fix
fix: tool/golangci-lint
	tool/golangci-lint run --fix

tool/golangci-lint: tool/.golangci-lint.$(GOLANGCI_LINT_VERSION)
	@mkdir -p tool
	GOBIN=$(PWD)/tool go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)

tool/.golangci-lint.$(GOLANGCI_LINT_VERSION):
	@rm -f tool/.golangci-lint.*
	@mkdir -p tool
	touch $@

.PHONY: tool
tool: tool/controller-gen tool/crd-ref-docs tool/ko tool/golangci-lint

.PHONY: apidocs
apidocs: tool/crd-ref-docs $(wildcard pkg/apis/*/*/*_types.go)
	@mkdir -p docs/apis
	tool/crd-ref-docs --source-path=pkg/apis --config=apidocs.yaml --renderer=markdown --output-path=docs/apis
