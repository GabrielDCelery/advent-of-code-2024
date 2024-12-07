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

.PHONY: solve-day-4
solve-day-4:
	@AOC_INPUT_PATH=$$(readlink -f ./inputs/input-day-4.txt) AOC_PART1_TEMPLATES_PATH=$$(readlink -f ./assets/day_04/xmas_part_01_templates.txt) AOC_PART2_TEMPLATES_PATH=$$(readlink -f ./assets/day_04/xmas_part_02_templates.txt) go run ./cmd/day_04/main.go

.PHONY: solve-day-5
solve-day-5:
	@AOC_INPUT_PATH=$$(readlink -f ./inputs/input-day-5.txt) go run ./cmd/day_05/main.go

.PHONY: solve-day-6
solve-day-6:
	@AOC_INPUT_PATH=$$(readlink -f ./inputs/input-day-6.txt) go run ./cmd/day_06/main.go
