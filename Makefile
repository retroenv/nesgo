help: ## show help, shown by default if no target is specified
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

generate: ## regenerate all files
	cd internal/gocc && gocc -p github.com/retroenv/nesgo/internal/gocc -a lang.bnf

lint: ## run code linters
	golangci-lint run

build-all: ## build code with all 3 GUI mode settings
	go build ./...
	go build -tags noopengl,sdl ./...
	go build -tags nogui ./...

test: install run-tests ## run tests
	go test -timeout 10s -race ./...

run-tests:
	nesgo -q -o ./examples/blue/main.nes ./examples/blue/main.go
	nesgo -q -o ./examples/debugprint/main.nes ./examples/debugprint/main.go

test-no-gui: install-no-gui run-tests ## run unit tests with gui disabled
	go test -timeout 10s -tags nogui ./...

test-coverage: ## run unit tests and create test coverage
	go test -timeout 10s -tags nogui ./... -coverprofile .testCoverage -covermode=atomic -coverpkg=./...

test-coverage-web: test-coverage ## run unit tests and show test coverage in browser
	go tool cover -func .testCoverage | grep total | awk '{print "Total coverage: "$$3}'
	go tool cover -html=.testCoverage

install: ## install all binaries
	go install ./cmd/...

install-no-gui: ## install all binaries with gui disabled
	go install -tags nogui ./cmd/...

install-linters: ## install all used linters
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.51.2

release: ## build release binaries for current git tag and publish on github
	goreleaser release

release-snapshot: ## build release binaries from current git state as snapshot
	goreleaser release --snapshot --rm-dist
