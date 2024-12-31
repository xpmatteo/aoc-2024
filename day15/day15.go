package day15

import (
	"github.com/xpmatteo/aoc-2024/mapping"
	"strings"
)

type Move int32

const (
	moveUp         = '^'
	moveRight      = '>'
	moveDown       = 'v'
	moveLeft       = '<'
	objectNone     = '.'
	objectRobot    = '@'
	objectWall     = '#'
	objectBox      = 'O'
	objectBoxLeft  = '['
	objectBoxRight = ']'
)

func predictRobot(inputMap mapping.Map, moves []Move) mapping.Map {
	robotPos := findRobot(inputMap)
	for _, move := range moves {
		_, robotPos = attemptMove(inputMap, move, robotPos)
	}
	return inputMap
}

func attemptMove(inputMap mapping.Map, move Move, oldPos mapping.Coord) (succeed bool, newPos mapping.Coord) {
	switch move {
	case moveRight:
		newPos = oldPos.East()
	case moveLeft:
		newPos = oldPos.West()
	case moveUp:
		newPos = oldPos.North()
	case moveDown:
		newPos = oldPos.South()
	default:
		panic("unknown move: " + string(move))
	}
	if inputMap.At(newPos) == objectWall {
		succeed = false
		newPos = oldPos
		return
	}
	if inputMap.At(newPos) == objectBox {
		succeed, _ = attemptMove(inputMap, move, newPos)
		if !succeed {
			newPos = oldPos
			return
		}
	}
	singleMove(inputMap, oldPos, newPos)
	succeed = true
	return
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
	for _, m := range strings.ReplaceAll(moves, "\n", "") {
		result = append(result, Move(m))
	}
	return result
}

func gpsTotal(m mapping.Map) int {
	var total int
	m.ForEachCoord(func(c mapping.Coord, value int32) {
		if value == objectBox {
			total += c.Row*100 + c.Col
		}
	})
	return total
}

func enlarge(inputMap mapping.Map) mapping.Map {
	var result mapping.Map
	for _, row := range inputMap {
		enlargedRow := ""
		for _, object := range row {
			switch object {
			case objectNone:
				enlargedRow += string(object)
				enlargedRow += string(object)
			case objectWall:
				enlargedRow += string(object)
				enlargedRow += string(object)
			case objectRobot:
				enlargedRow += string(objectRobot)
				enlargedRow += string(objectNone)
			case objectBox:
				enlargedRow += string(objectBoxLeft)
				enlargedRow += string(objectBoxRight)
			default:
				panic("unknown object: " + string(object))
			}
		}
		result = append(result, enlargedRow)
	}
	return result
}
