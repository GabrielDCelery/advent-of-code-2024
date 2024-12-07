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

func (g *GameMap) Clone() *GameMap {
	clonedMatrix := [][]Cell{}
	for y := range len(g.matrix) {
		clonedMatrixLine := []Cell{}
		for x := range len(g.matrix[0]) {
			clonedMatrixLine = append(clonedMatrixLine, Cell(g.matrix[y][x]))
		}
		clonedMatrix = append(clonedMatrix, clonedMatrixLine)
	}
	return NewGameMap(clonedMatrix)
}

func (g *GameMap) SetCellToType(y int, x int, cellType Cell) {
	g.matrix[y][x] = cellType
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

func (g *Guardian) Clone() *Guardian {
	return NewGuardian(Vector{y: g.location.y, x: g.location.x}, Vector{y: g.facing.y, x: g.facing.x})
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

func (g *Guardian) GetFacing() Vector {
	return Vector{
		y: g.facing.y,
		x: g.facing.x,
	}
}

func (g *Guardian) GetLocationInFront() Vector {
	return Vector{
		y: g.location.y + g.facing.y,
		x: g.location.x + g.facing.x,
	}
}

type GameMemory struct {
	isStuckInInfiniteLoop        bool
	uniqueLocationsWithFacingMap map[string]bool
}

func NewGameMemory() *GameMemory {
	return &GameMemory{
		isStuckInInfiniteLoop:        false,
		uniqueLocationsWithFacingMap: map[string]bool{},
	}
}

func (gm *GameMemory) AppendGuardianSnapshot(guardian *Guardian) {
	location := guardian.GetCurrentLocation()
	facing := guardian.GetFacing()
	currLocationWithFacingKey := fmt.Sprintf("%d_%d_%d_%d", location.y, location.x, facing.y, facing.x)
	_, ok := gm.uniqueLocationsWithFacingMap[currLocationWithFacingKey]
	if ok {
		gm.isStuckInInfiniteLoop = true
	}
	gm.uniqueLocationsWithFacingMap[currLocationWithFacingKey] = true
}

func (gm *GameMemory) CountNumOfUniqueCellsVisitedByGuardian() int {
	uniqueLocationKeyMap := map[string]bool{}
	for key := range gm.uniqueLocationsWithFacingMap {
		keyComponents := strings.Split(key, "_")
		y, x := keyComponents[0], keyComponents[1]
		uniqueLocationKey := fmt.Sprintf("%s_%s", y, x)
		uniqueLocationKeyMap[uniqueLocationKey] = true
	}
	return len(uniqueLocationKeyMap)
}

func (gm *GameMemory) IsStuckInInfiniteLoop() bool {
	return gm.isStuckInInfiniteLoop
}

func countTheNumberOfDistinctPositionsBeforeGuardianLeaves(input string) (int, error) {
	gameMap, guardian, err := transformInputToGameState(input)

	if err != nil {
		return 0, err
	}

	gameMemory := NewGameMemory()

	err = moveGuardianUntilItLeavesMap(gameMap, guardian, gameMemory)

	if err != nil {
		return 0, err
	}

	solution := gameMemory.CountNumOfUniqueCellsVisitedByGuardian()

	return solution, nil
}

func countNumberOfWaysGuardianCanBeLockedIntoInfiniteLoop(input string) (int, error) {
	gameMap, guardian, err := transformInputToGameState(input)

	if err != nil {
		return 0, err
	}

	numOfGuardiansStuckInALoop := 0

	for y := range gameMap.Height() {
		for x := range gameMap.Width() {
			clonedGuardian := guardian.Clone()
			location := clonedGuardian.GetCurrentLocation()
			if location.y == y && location.x == x {
				continue
			}
			clonedGameMap := gameMap.Clone()
			clonedGameMap.SetCellToType(y, x, WALL)
			gameMemory := NewGameMemory()

			err = moveGuardianUntilItLeavesMap(clonedGameMap, clonedGuardian, gameMemory)

			if err != nil {
				return 0, err
			}

			if gameMemory.IsStuckInInfiniteLoop() {
				numOfGuardiansStuckInALoop += 1
			}
		}
	}

	return numOfGuardiansStuckInALoop, nil
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

func moveGuardianUntilItLeavesMap(gameMap *GameMap, guardian *Guardian, gameMemory *GameMemory) error {
loop:
	for true {
		gameMemory.AppendGuardianSnapshot(guardian)
		//  exit early if guardian is stuck in an infinite loop
		if gameMemory.IsStuckInInfiniteLoop() {
			return nil
		}
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

func SolveDay6Part2() (int, error) {
	inputPath := os.Getenv("AOC_INPUT_PATH")

	file, err := os.ReadFile(inputPath)

	if err != nil {
		return 0, err
	}

	solution, err := countNumberOfWaysGuardianCanBeLockedIntoInfiniteLoop(strings.TrimSpace(string(file)))

	if err != nil {
		return 0, err
	}

	return solution, nil
}
