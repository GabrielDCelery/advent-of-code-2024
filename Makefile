.PHONY: init
init:
	@go mod tidy

.PHONY: test
test:
	@go test -v ./...

.PHONY: solve-day-1
solve-day-1:
	@AOC_INPUT_PATH=$$(readlink -f ./inputs/input-day-1.txt) go run ./cmd/day_01/main.go

.PHONY: solve-day-2
solve-day-2:
	@AOC_INPUT_PATH=$$(readlink -f ./inputs/input-day-2.txt) go run ./cmd/day_02/main.go

.PHONY: solve-day-3
solve-day-3:
	@AOC_INPUT_PATH=$$(readlink -f ./inputs/input-day-3.txt) go run ./cmd/day_03/main.go
