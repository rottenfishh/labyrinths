COVERAGE_FILE ?= coverage.out

TARGET ?= app # CHANGE THIS TO YOUR BINARY NAME

.PHONY: build
build:
	@echo "Выполняется go build для таргета ${TARGET}"
	@mkdir -p .bin
	@go build -o ./bin/${TARGET} ./cmd/${TARGET}

## test: run all tests
.PHONY: test
test:
	@go test -coverpkg='github.com/es-debug/backend-academy-2024-go-template/...' --race -count=1 -coverprofile='$(COVERAGE_FILE)' ./...
	@go tool cover -func='$(COVERAGE_FILE)' | grep ^total | tr -s '\t'
