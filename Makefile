.PHONY: init
init:
	@go mod tidy

.PHONY: test
test:
	@go test -v ./...

.PHONY: solve-day-1
solve-day-1:
	@go build -ldflags="-s -w" -o ./.bin/day-01 ./cmd/day_01/main.go
	@AOC_INPUT_PATH=$$(readlink -f ./inputs/input-day-1.txt) ./.bin/day-01

.PHONY: solve-day-2
solve-day-2:
	@go build -ldflags="-s -w" -o ./.bin/day-02 ./cmd/day_02/main.go
	@AOC_INPUT_PATH=$$(readlink -f ./inputs/input-day-2.txt) ./.bin/day-02

.PHONY: solve-day-3
solve-day-3:
	@go build -ldflags="-s -w" -o ./.bin/day-03 ./cmd/day_03/main.go
	@AOC_INPUT_PATH=$$(readlink -f ./inputs/input-day-3.txt) ./.bin/day-03

.PHONY: solve-day-4
solve-day-4:
	@go build -ldflags="-s -w" -o ./.bin/day-04 ./cmd/day_04/main.go
	@AOC_INPUT_PATH=$$(readlink -f ./inputs/input-day-4.txt) AOC_PART1_TEMPLATES_PATH=$$(readlink -f ./assets/day_04/xmas_part_01_templates.txt) AOC_PART2_TEMPLATES_PATH=$$(readlink -f ./assets/day_04/xmas_part_02_templates.txt) ./.bin/day-04

.PHONY: solve-day-5
solve-day-5:
	@go build -ldflags="-s -w" -o ./.bin/day-05 ./cmd/day_05/main.go
	@AOC_INPUT_PATH=$$(readlink -f ./inputs/input-day-5.txt) ./.bin/day-05

.PHONY: solve-day-6
solve-day-6:
	@go build -ldflags="-s -w" -o ./.bin/day-06 ./cmd/day_06/main.go
	@AOC_INPUT_PATH=$$(readlink -f ./inputs/input-day-6.txt) ./.bin/day-06

.PHONY: solve-day-7
solve-day-7:
	@go build -ldflags="-s -w" -o ./.bin/day-07 ./cmd/day_07/main.go
	@AOC_INPUT_PATH=$$(readlink -f ./inputs/input-day-7.txt) ./.bin/day-07

.PHONY: solve-day-8
solve-day-8:
	@go build -ldflags="-s -w" -o ./.bin/day-08 ./cmd/day_08/main.go
	@AOC_INPUT_PATH=$$(readlink -f ./inputs/input-day-8.txt) ./.bin/day-08

.PHONY: solve-day-9
solve-day-9:
	@go build -ldflags="-s -w" -o ./.bin/day-09 ./cmd/day_09/main.go
	@AOC_INPUT_PATH=$$(readlink -f ./inputs/input-day-9.txt) ./.bin/day-09

.PHONY: solve-day-10
solve-day-10:
	@go build -ldflags="-s -w" -o ./.bin/day-10 ./cmd/day_10/main.go
	@AOC_INPUT_PATH=$$(readlink -f ./inputs/input-day-10.txt) ./.bin/day-10

.PHONY: solve-day-11
solve-day-11:
	@go build -ldflags="-s -w" -o ./.bin/day-11 ./cmd/day_11/main.go
	@AOC_INPUT_PATH=$$(readlink -f ./inputs/input-day-11.txt) ./.bin/day-11
