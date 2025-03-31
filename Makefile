LOCAL_BIN:=$(CURDIR)/bin

export PATH:=$(LOCAL_BIN):$(PATH)

run-tests:
	$(info Running...)
	$(BUILD_ENVPARMS) go run ./test/run.go

GOLANGCI_BIN:=$(LOCAL_BIN)/golangci-lint
GOLANGCI_TAG:=1.48.0

.PHONY: install-lint
install-lint:
ifeq ($(wildcard $(GOLANGCI_BIN)),)
	$(info Downloading golangci-lint v$(GOLANGCI_TAG))
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v$(GOLANGCI_TAG)
GOLANGCI_BIN:=$(LOCAL_BIN)/golangci-lint
endif

lint: install-lint
	$(info Running lint...)
	$(GOLANGCI_BIN) run --config=.cfg/lint.yaml ./...

.PHONY: test
test:
	$(info Running tests...)
	go test ./...
