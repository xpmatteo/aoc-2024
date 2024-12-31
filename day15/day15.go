package day15

import "github.com/xpmatteo/aoc-2024/mapping"

type Move int32

const (
	moveUp      = '^'
	moveRight   = '>'
	moveDown    = 'v'
	moveLeft    = '<'
	objectNone  = '.'
	objectRobot = '@'
	objectWall  = '#'
	objectBox   = 'O'
)

func predictRobot(inputMap mapping.Map, moves []Move) mapping.Map {
	robotPos := findRobot(inputMap)
	for _, move := range moves {
		attemptMove(inputMap, move, robotPos)
	}
	return inputMap
}

func attemptMove(inputMap mapping.Map, move Move, agentPos mapping.Coord) {
	switch move {
	case moveRight:
		newPos := agentPos.East()
		if inputMap.At(newPos) == objectWall {
			break
		}
		if inputMap.At(newPos) == objectBox {
			attemptMove(inputMap, move, newPos)
		}
		singleMove(inputMap, agentPos, newPos)
	case moveLeft:
		newPos := agentPos.West()
		singleMove(inputMap, agentPos, newPos)
	default:
		panic("unknown move: " + string(move))
	}
}

func singleMove(inputMap mapping.Map, robotPos mapping.Coord, newPos mapping.Coord) {
	theObject := inputMap.At(robotPos)
	inputMap.SetCoord(robotPos, objectNone)
	inputMap.SetCoord(newPos, theObject)
}

func findRobot(m mapping.Map) mapping.Coord {
	var result mapping.Coord
	m.ForEachCoord(func(c mapping.Coord, value int32) {
		if value == objectRobot {
			result = c
		}
	})
	return result
}

func parseMoves(moves string) []Move {
	var result []Move
	for _, m := range moves {
		result = append(result, Move(m))
	}
	return result
}
