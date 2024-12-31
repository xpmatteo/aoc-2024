package day15

import (
	"github.com/xpmatteo/aoc-2024/mapping"
	"strings"
)

type Move int32

const (
	moveUp    = '^'
	moveRight = '>'
	moveDown  = 'v'
	moveLeft  = '<'

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

func nextPos(oldPos mapping.Coord, move Move) mapping.Coord {
	var result mapping.Coord
	switch move {
	case moveRight:
		result = oldPos.East()
	case moveLeft:
		result = oldPos.West()
	case moveUp:
		result = oldPos.North()
	case moveDown:
		result = oldPos.South()
	default:
		panic("unknown move: " + string(move))
	}
	return result
}

func attemptMove(inputMap mapping.Map, move Move, oldPos mapping.Coord) (succeed bool, newPos mapping.Coord) {
	newPos = nextPos(oldPos, move)
	objectAtNewPos := inputMap.At(newPos)
	if objectAtNewPos == objectWall {
		succeed = false
		newPos = oldPos
		return
	}
	if objectAtNewPos == objectBox {
		succeed, _ = attemptMove(inputMap, move, newPos)
		if !succeed {
			newPos = oldPos
			return
		}
	}
	if isHorizontal(move) && isLargeBox(objectAtNewPos) {
		succeed, _ = attemptMove(inputMap, move, newPos)
		if !succeed {
			newPos = oldPos
			return
		}
	}
	if isVertical(move) && isLargeBox(objectAtNewPos) {
		cloneMap := inputMap.Clone()
		otherNewPos := posOfOtherHalf(newPos, objectAtNewPos)
		succeedLeft, _ := attemptMove(cloneMap, move, newPos)
		succeedRight, _ := attemptMove(cloneMap, move, otherNewPos)
		if succeedLeft && succeedRight {
			attemptMove(inputMap, move, newPos)
			attemptMove(inputMap, move, otherNewPos)
		} else {
			newPos = oldPos
			return
		}
	}
	singleMove(inputMap, oldPos, newPos)
	succeed = true
	return
}

func posOfOtherHalf(pos mapping.Coord, object int32) mapping.Coord {
	switch object {
	case objectBoxLeft:
		return pos.East()
	case objectBoxRight:
		return pos.West()
	default:
		panic("unknown object " + string(object))
	}
}

func isHorizontal(move Move) bool {
	return move == moveLeft || move == moveRight
}

func isVertical(move Move) bool {
	return !isHorizontal(move)
}

func isLargeBox(object int32) bool {
	return object == objectBoxLeft || object == objectBoxRight
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
		if value == objectBox || value == objectBoxLeft {
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
