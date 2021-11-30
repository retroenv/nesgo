.PHONY: check-direnv

help:
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

generate: ## regenerate all files
	cd internal/gocc && gocc -p github.com/retroenv/nesgo/internal/gocc -a lang.bnf

lint: ## run code linters
	flint ./internal/ast/...
	flint ./internal/gocc/...
	flint ./pkg/nes/...
	golangci-lint run

test: ## run unit tests
	go test -race ./...

test-no-gui: ## run unit tests with gui disabled
	go test -tags nogui ./... -v

test-coverage: ## run unit tests and show test coverage
	go test ./... -coverprofile .testCoverage -covermode=atomic -coverpkg=./...
	go tool cover -func .testCoverage | grep total | awk '{print "Total coverage: "$$3}'
	go tool cover -html=.testCoverage

install: ## install all binaries
	go install ./cmd/...
