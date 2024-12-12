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

day4:
	@cd ./cmd/day4 && go run .
.PHONY: day4

day5:
	@cd ./cmd/day5 && go run .
.PHONY: day5

day6:
	@cd ./cmd/day6 && go run .
.PHONY: day6

day7:
	@cd ./cmd/day7 && go run .
.PHONY: day7

day8:
	@cd ./cmd/day8 && go run .
.PHONY: day8

day9:
	@cd ./cmd/day9 && go run .
.PHONY: day9

day10:
	@cd ./cmd/day10 && go run .
.PHONY: day10