package day_06

import (
	"fmt"
	"os"
	"strings"
)

type Vector struct {
	x int
	y int
}

type Cell = int

const (
	EMPTY Cell = iota
	WALL
	NULL
)

type GameMap struct {
	matrix [][]Cell
}

func NewGameMap(matrix [][]Cell) *GameMap {
	return &GameMap{
		matrix: matrix,
	}
}

func (g *GameMap) Width() int {
	return len(g.matrix[0])
}

func (g *GameMap) Height() int {
	return len(g.matrix)
}

func (g *GameMap) Cell(vector Vector) Cell {
	isInside := vector.x >= 0 && vector.y >= 0 && vector.x < g.Width() && vector.y < g.Height()
	if !isInside {
		return NULL
	}
	return g.matrix[vector.y][vector.x]
}

type Action = int

const (
	MOVE_FORWARD Action = iota
	TURN_RIGHT
)

type Guardian struct {
	location Vector
	facing   Vector
}

func NewGuardian(location Vector, facing Vector) *Guardian {
	return &Guardian{
		location,
		facing,
	}
}

func (g *Guardian) ExecuteAction(action Action) {
	switch action {
	case MOVE_FORWARD:
		g.location.y = g.location.y + g.facing.y
		g.location.x = g.location.x + g.facing.x
	case TURN_RIGHT:
		tempY := g.facing.y
		tempX := g.facing.x
		g.facing.y = tempX
		g.facing.x = tempY * -1
	}
}

func (g *Guardian) GetCurrentLocation() Vector {
	return Vector{
		y: g.location.y,
		x: g.location.x,
	}
}

func (g *Guardian) GetLocationInFront() Vector {
	return Vector{
		y: g.location.y + g.facing.y,
		x: g.location.x + g.facing.x,
	}
}

type GameMemory struct {
	uniqueCellsVisitedByGuardians map[string]bool
}

func NewGameMemory() *GameMemory {
	return &GameMemory{
		uniqueCellsVisitedByGuardians: map[string]bool{},
	}
}

func (gm *GameMemory) AppendVectorToVisitedLocations(vector Vector) {
	key := fmt.Sprintf("%d__%d", vector.y, vector.x)
	gm.uniqueCellsVisitedByGuardians[key] = true
}

func (gm *GameMemory) CountNumOfUniqueCellsVisitedByGuardians() int {
	return len(gm.uniqueCellsVisitedByGuardians)
}

func countTheNumberOfDistinctPositionsBeforeGuardianLeaves(input string) (int, error) {
	gameMap, guardian, err := transformInputToGameState(input)

	if err != nil {
		return 0, err
	}

	gameMemory := NewGameMemory()

	err = moveGuardiansUntilTheyLeaveMap(gameMap, guardian, gameMemory)

	if err != nil {
		return 0, err
	}

	solution := gameMemory.CountNumOfUniqueCellsVisitedByGuardians()

	return solution, nil
}

func transformInputToGameState(input string) (*GameMap, *Guardian, error) {
	gameMapMatrix := [][]Cell{}

	lines := strings.Split(input, "\n")

	guardian := &Guardian{}

	for y, line := range lines {
		gameMapMatrixRow := []Cell{}
		for x, cell := range line {
			switch string(cell) {
			case ".":
				gameMapMatrixRow = append(gameMapMatrixRow, EMPTY)
			case "#":
				gameMapMatrixRow = append(gameMapMatrixRow, WALL)
			case "^":
				gameMapMatrixRow = append(gameMapMatrixRow, EMPTY)
				guardian = NewGuardian(Vector{y: y, x: x}, Vector{y: -1, x: 0})
			default:
				return &GameMap{}, &Guardian{}, fmt.Errorf("Unhandled cell type %s", string(cell))
			}
		}
		gameMapMatrix = append(gameMapMatrix, gameMapMatrixRow)
	}

	gameMap := NewGameMap(gameMapMatrix)

	return gameMap, guardian, nil
}

func moveGuardiansUntilTheyLeaveMap(gameMap *GameMap, guardian *Guardian, gameMemory *GameMemory) error {
loop:
	for true {
		gameMemory.AppendVectorToVisitedLocations(guardian.GetCurrentLocation())
		cellInFront := gameMap.Cell(guardian.GetLocationInFront())
		switch cellInFront {
		case EMPTY:
			guardian.ExecuteAction(MOVE_FORWARD)
		case WALL:
			guardian.ExecuteAction(TURN_RIGHT)
		case NULL:
			break loop
		}
	}
	return nil
}

func SolveDay6Part1() (int, error) {
	inputPath := os.Getenv("AOC_INPUT_PATH")

	file, err := os.ReadFile(inputPath)

	if err != nil {
		return 0, err
	}

	solution, err := countTheNumberOfDistinctPositionsBeforeGuardianLeaves(strings.TrimSpace(string(file)))

	if err != nil {
		return 0, err
	}

	return solution, nil
}
