.PHONY: build
build:
	go build -o build/cli ./cmd/...

.PHONY: dev
dev:
	go run ./cmd/...
