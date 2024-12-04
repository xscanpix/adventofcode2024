GO_FILES = $(shell find . -name \*.go)

test:
	@go test \
		-shuffle=on \
		-count=1 \
		-short \
		-timeout=5m \
		./...
.PHONY: test

day1:
	@cd ./cmd/day1 && go run .
.PHONY: day1