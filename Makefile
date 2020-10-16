GO_PACKAGE := github.com/6543/go-version

GO ?= go
GOFMT ?= gofmt -s
SHASUM ?= shasum -a 256

GO_SOURCES := $(shell find . -type f -name "*.go")

.PHONY: all
all: clean test

.PHONY: clean
clean:
	$(GO) clean -i ./...

.PHONY: fmt
fmt:
	$(GOFMT) -w $(GO_SOURCES)

.PHONY: vet
vet:
	$(GO) vet $(GO_PACKAGE)

.PHONY: revive
revive:
	$(GO) get -u github.com/mgechev/revive; \
	revive -config .revive.toml || exit 1

.PHONY: misspell
misspell-check:
	$(GO) get -u github.com/client9/misspell/cmd/misspell; \
	misspell -error -i unknwon,destory $(GO_SOURCES)

.PHONY: fmt-check
fmt-check:
	# get all go files and run go fmt on them
	@diff=$$($(GOFMT) -d $(GO_SOURCES)); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;

.PHONY: lint
lint: misspell golangci-lint revive vet

.PHONY: test
test:
	$(GO) test -cover -coverprofile coverage.out $(GO_PACKAGE) || exit 1

.PHONY: golangci-lint
golangci-lint:
	@hash golangci-lint > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		export BINARY="golangci-lint"; \
		curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(GOPATH)/bin v1.31.0; \
	fi
	golangci-lint run --timeout 5m
