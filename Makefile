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

day2:
	@cd ./cmd/day2 && go run .
.PHONY: day2

day3:
	@cd ./cmd/day3 && go run .
.PHONY: day3