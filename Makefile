help:
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

generate: ## regenerate all files
	cd internal/gocc && gocc -p github.com/retroenv/nesgo/internal/gocc -a lang.bnf

lint: ## run code linters
	flint ./cmd/...
	flint ./internal/ast/...
	flint ./internal/gocc/...
	flint ./pkg/nes/...
	golangci-lint run

build-all: ## build code with all 3 GUI settings
	go build ./...
	go build -tags noopengl,sdl ./...
	go build -tags nogui ./...

test: install ## run tests
	go test -race ./...
	nesgo -q -f ./examples/blue/main.go -o ./examples/blue/main.nes
	nesgo -q -f ./examples/debugprint/main.go -o ./examples/debugprint/main.nes
	nesgodisasm -f examples/blue/main.nes -o examples/blue/disasm.asm -v
	nesgodisasm -f examples/debugprint/main.nes -o examples/debugprint/disasm.asm -v

test-no-gui: ## run unit tests with gui disabled
	go test -tags nogui ./... -v

test-coverage: ## run unit tests with test coverage
	go test -tags nogui ./... -coverprofile .testCoverage -covermode=atomic -coverpkg=./...

test-coverage-web: test-coverage ## run unit tests and show test coverage in browser
	go tool cover -func .testCoverage | grep total | awk '{print "Total coverage: "$$3}'
	go tool cover -html=.testCoverage

install: ## install all binaries
	go install ./cmd/...

install-linters: ## install all linters
	go install github.com/fraugster/flint@v0.1.1
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.46.2
