REFLEX ?= github.com/cespare/reflex

.PHONY: dev
dev:
	go get github.com/cespare/reflex
	go run $(REFLEX) -R "\\.vscode|vendor|tests" -r "\\.go" -s -- sh -c "go run --race ./cmd/app/..."


.PHONY: tests
tests:
	go test ./internal/domain -v

.PHONY: lint
lint:
	golangci-lint run --fast
