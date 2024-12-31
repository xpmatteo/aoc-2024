package day14

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/xpmatteo/aoc-2024/day1"
	"github.com/xpmatteo/aoc-2024/day13"
	"github.com/xpmatteo/aoc-2024/mapping"
	"regexp"
	"strings"
)

type point day13.Point

func (p point) Plus(q point) point {
	pp := day13.Point(p)
	qq := day13.Point(q)
	return point(pp.Plus(qq))
}

type Robot struct {
	position, speed point
}

type Lobby struct {
	robots []*Robot
	size   point
}

func (l *Lobby) Simulate(seconds int) {
	for range seconds {
		for _, robot := range l.robots {
			robot.position = robot.position.Plus(robot.speed)
			for robot.position.X < 0 {
				robot.position.X += l.size.X
			}
			for robot.position.Y < 0 {
				robot.position.Y += l.size.Y
			}
			robot.position.X = robot.position.X % l.size.X
			robot.position.Y = robot.position.Y % l.size.Y
		}
	}
}

func (l *Lobby) Map() mapping.Map {
	line := strings.Repeat(".", l.size.X)
	lines := strings.Repeat(line+"\n", l.size.Y)
	lines = strings.TrimRight(lines, "\n")
	m := mapping.Map(split(lines))
	for _, robot := range l.robots {
		m.Set(robot.position.Y, robot.position.X, '1')
	}
	return m
}

func join(s ...string) string {
	return strings.Join(s, "\n")
}

func split(input string) []string {
	return strings.Split(input, "\n")
}

func parseLobby(size point, input string) Lobby {
	lines := split(input)
	var robots []*Robot
	for _, line := range lines {
		if len(line) > 0 {
			robots = append(robots, parseRobot(line))
		}
	}
	return Lobby{
		robots: robots,
		size:   size,
	}
}

func parseRobot(input string) *Robot {
	re := regexp.MustCompile(`-?\d+`)
	matches := re.FindAllString(input, -1)
	if len(matches) != 4 {
		panic(fmt.Sprintf("Unexpected # of numbers: %v", matches))
	}
	numbers := lo.Map(matches, func(s string, index int) int {
		return day1.Atoi(s)
	})
	robot := &Robot{
		position: point{numbers[0], numbers[1]},
		speed:    point{numbers[2], numbers[3]},
	}
	return robot
}
