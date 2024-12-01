.PHONY: init
init:
	@go mod tidy

.PHONY: test
test:
	@go test -v ./...

.PHONY: solve-day-1
solve-day-1:
	@AOC_INPUT_PATH=$$(readlink -f ./inputs/input-day-1.txt) go run ./cmd/day_01/main.go
