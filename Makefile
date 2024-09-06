GO ?= go

GOFLAGS :=

VERSION_MAJOR ?= x
VERSION_MINOR ?= x
VERSION_REV ?= x

BUILD_TIME ?=$(shell date -I'seconds')
BUILD_BRANCH ?= ---
BUILD_COMMIT ?= ---
BUILD_CREATOR ?= a person via makefile
BUILD_MACHINE ?= $(shell hostname)

LDFLAGS += -X 'sigstore/pkg/config.VersionMajor=$(VERSION_MAJOR)'
LDFLAGS += -X 'sigstore/pkg/config.VersionMinor=$(VERSION_MINOR)'
LDFLAGS += -X 'sigstore/pkg/config.BuildNumber=$(VERSION_REV)'
LDFLAGS += -X 'sigstore/pkg/config.BuildTime=$(BUILD_TIME)'
LDFLAGS += -X 'sigstore/pkg/config.BuildBranch=$(BUILD_BRANCH)'
LDFLAGS += -X 'sigstore/pkg/config.BuildCommit=$(BUILD_COMMIT)'
LDFLAGS += -X 'sigstore/pkg/config.Builder=$(BUILD_CREATOR)'
LDFLAGS += -X 'sigstore/pkg/config.BuildMachine=$(BUILD_MACHINE)'

.PHONY: all
all: app test

.PHONY: build
build: app

.PHONY: clean
clean:
	CGO_ENABLED=0 $(GO) clean
	CGO_ENABLED=0 $(GO) clean -cache

.PHONY: deps
deps:
	CGO_ENABLED=0 $(GO) get -t -v ./...

.PHONY: test
test:
	CGO_ENABLED=0 $(GO) install github.com/jstemmer/go-junit-report/v2@latest
	CGO_ENABLED=0 $(GO) install github.com/boumenot/gocover-cobertura@v1.2.0
	CGO_ENABLED=0 $(GO) test -v --cover -coverprofile=coverage.out -covermode=count ./... 2>&1 | go-junit-report -set-exit-code > test_report.xml
	CGO_ENABLED=0 gocover-cobertura < coverage.out > coverage.xml

.PHONY: lint
lint:
	CGO_ENABLED=0 golangci-lint run --out-format "checkstyle:lint_report.xml,code-climate:gl-code-quality-report.json,tab:stdout" --timeout=10m

.PHONY: app
app:
	CGO_ENABLED=0 $(GO) build $(GOFLAGS) -ldflags "$(LDFLAGS)" -v -o ./sigstore-app ./cmd/sigstore