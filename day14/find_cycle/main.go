package main

import (
	"fmt"
	"github.com/xpmatteo/aoc-2024/day1"
	"github.com/xpmatteo/aoc-2024/day14"
	"os"
)

//func main() {
//	lobby := parseLobby(point{101, 103}, day1.ReadFile("day14.txt"))
//	start := 1024
//	lobby.Simulate(start)
//	for seconds := start + 1; seconds < 10000; seconds++ {
//		lobby.Simulate(1)
//		fmt.Print("\033[H\033[2J")
//		fmt.Println("seconds:", seconds)
//		s := lobby.Map().String()
//		fmt.Println(s[:len(s)/2])
//		time.Sleep(500 * time.Millisecond)
//	}
//}

func main() {
	lobby := day14.ParseLobby(day14.Point{101, 103}, day1.ReadFile("../day14.txt"))
	for seconds := 0; seconds < 100000; seconds++ {
		memorize(lobby, seconds)
		lobby.Simulate(1)
	}
}

var lobbies = make(map[string]struct{})

// Lobby state at second 10403 is already memorized
func memorize(lobby day14.Lobby, seconds int) {
	key := lobby.Map().String()
	if _, ok := lobbies[key]; ok {
		fmt.Printf("Lobby state at second %d is already memorized.\n", seconds)
		os.Exit(0)
	} else {
		lobbies[key] = struct{}{}
	}
}
