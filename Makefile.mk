HOMEDIR := $(shell pwd)
OUTDIR := $(HOMEDIR)/build
BINDIR := $(OUTDIR)/bin
GO := go
GOPATH := $(shell $(GO) env GOPATH)
GOBIN := $(GOPATH)/bin

.PHONY: all
all: prepare compile check test package

.PHONY: build
build: prepare compile check package

.PHONY: prepare
prepare: prepare-dep

prepare-dep:
	git config --global http.sslVerify false
	# git config --global url.git@code.aliyun.com:.insteadOf https://code.aliyun.com/
	git submodule update --init --recursive

set-env:
	$(GO) env -w GO111MODULE=on
	$(GO) env -w CGO_ENABLED=1
	# $(GO) env -w GOPROXY=https://proxy.golang.com.cn,direct
	$(GO) env -w GONOSUMDB=\*
	# $(GO) env -w GOPRIVATE=code.aliyun.com

.PHONY: compile
compile: pre-compile post-compile

pre-compile: set-env
	# $(GO) mod tidy -v

post-compile-default: set-env
	echo "build-default"

.PHONY: proto
proto: proto-tools

proto-tools: set-env
	@for repo in \
		"github.com/cloudwego/kitex/tool/cmd/kitex@v0.11.3" \
		"github.com/cloudwego/thriftgo@v0.3.17" \
	; do \
		$(GO) install -v $$repo; \
	done
	go mod edit -droprequire=github.com/apache/thrift/lib/go/thrift
	go mod edit -replace=github.com/apache/thrift=github.com/apache/thrift@v0.13.0

.PHONY: test
test: case bench vet race coverage
# test: case bench vet race msan coverage

case: set-env
	$(GO) test -tags=unit -timeout 3h -short -v ./...

bench: set-env
	$(GO) test -v -bench=. -benchtime=10s -timeout 3h -count=5 ./...

vet: set-env
	$(GO) vet -unsafeptr=false ./...

race: set-env
	$(GO) test -race -short -v ./...

msan: set-env
	-$(GO) test -msan -short -v ./...

coverage: set-env
	$(GO) test -cover -coverprofile coverage.cov ./...
	$(GO) tool cover -func coverage.cov
	$(GO) tool cover -html=coverage.cov -o coverage.html

.PHONY: check
check: check-tools golangci-lint

.PHONY: fix
fix: check-tools golangci-lint-fix

check-tools: set-env
	$(GO) install -v "github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.1"

golangci-lint:
	PATH=$(GOBIN):$$PATH && \
	golangci-lint run \
		--timeout=10m \
		-D errcheck \
		-D gosimple \
		-D unused \
		-D gocyclo \
		-E gocritic \
		-E gofmt \
		-E goimports \
		./...

golangci-lint-fix:
	PATH=$(GOBIN):$$PATH && \
	golangci-lint run \
		--fix \
		--timeout=10m \
		-D errcheck \
		-D gosimple \
		-D unused \
		-D gocyclo \
		-E gocritic \
		-E gofmt \
		-E goimports \
		./...

.PHONY: package
package: package-bin

package-bin-default:
	mkdir -p $(OUTDIR)

.PHONY: clean
clean:
	rm -rf $(OUTDIR)
	rm -f coverage.cov coverage.html

# Overrides
# https://stackoverflow.com/a/49804748
%: %-default
	@ true
