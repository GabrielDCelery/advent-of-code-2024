package day_06

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
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
	id       int
	location Vector
	facing   Vector
}

func NewGuardian(id int, location Vector, facing Vector) *Guardian {
	return &Guardian{
		id,
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

// func (g *Guardian) determineNextAction(gameMap [][]int) Action {
// }

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
	gameMap, guardians, err := transformInputToGameState(input)

	if err != nil {
		return 0, err
	}

	gameMemory := NewGameMemory()

	err = moveGuardiansUntilTheyLeaveMap(gameMap, guardians, gameMemory)

	if err != nil {
		return 0, err
	}

	solution := gameMemory.CountNumOfUniqueCellsVisitedByGuardians()

	return solution, nil
}

func transformInputToGameState(input string) (*GameMap, []*Guardian, error) {
	guardianID := 0
	gameMapMatrix := [][]Cell{}
	guardians := []*Guardian{}

	lines := strings.Split(input, "\n")

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
				guardian := NewGuardian(guardianID, Vector{y: y, x: x}, Vector{y: -1, x: 0})
				guardians = append(guardians, guardian)
				guardianID += 1
			default:
				log.Fatalln("TODO ERROR")
			}
		}
		gameMapMatrix = append(gameMapMatrix, gameMapMatrixRow)
	}

	gameMap := NewGameMap(gameMapMatrix)

	return gameMap, guardians, nil
}

func isGuardianOnMap(gameMap *GameMap, guardian *Guardian) bool {
	guardianLocation := guardian.GetCurrentLocation()
	return guardianLocation.x >= 0 &&
		guardianLocation.y >= 0 &&
		guardianLocation.x < gameMap.Width() &&
		guardianLocation.y < gameMap.Height()
}

func moveGuardiansUntilTheyLeaveMap(gameMap *GameMap, guardians []*Guardian, gameMemory *GameMemory) error {
	start := time.Now()
	for len(guardians) > 0 {
		elapsed := time.Since(start).Milliseconds()
		if elapsed > 2 {
			return fmt.Errorf("timed out")
		}
		for _, guardian := range guardians {
			gameMemory.AppendVectorToVisitedLocations(guardian.GetCurrentLocation())
			cellInFront := gameMap.Cell(guardian.GetLocationInFront())
			switch cellInFront {
			case EMPTY:
				guardian.ExecuteAction(MOVE_FORWARD)
			case WALL:
				guardian.ExecuteAction(TURN_RIGHT)
			case NULL:
				guardian.ExecuteAction(MOVE_FORWARD)
			}
		}

		temp := []*Guardian{}
		for _, guardian := range guardians {
			if isGuardianOnMap(gameMap, guardian) {
				temp = append(temp, guardian)
			}
		}
		guardians = temp
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
