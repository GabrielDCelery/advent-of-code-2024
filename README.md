# Advent of Code 2024

## What is this repository for?

This is a repo to solve the puzzles published on [Advent of Code 2024](https://adventofcode.com/2024).

## What do you need to run this?

Have `go 1.22` and `make` installed.

## Running the app 

The puzzles' inputs are stored in the `inputs` folder. As the data is specific to each competitor, if you want this repository to solve your puzzle, you must first replace the files with your own.

To obtain the solution for a specific day, execute the corresponding make command.

Example:

```sh
make solve-day-1 
```

Outputs:

```sh
Day 1, part 1 solution: 1941353
Day 1, part 2 solution: 22539317
```

## Using the dev environment

To install dependencies:

```sh
make init
```

To run tests:

```sh
make test
```

