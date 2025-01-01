package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
	"os"

	"github.com/xpmatteo/aoc-2024/day1"
	"github.com/xpmatteo/aoc-2024/day14"
	"strings"
)

const squareSide = 8
const lobbyWidth = 101
const lobbyHeight = 103
const threshold = 0.5
const maxSeconds = 10403

func main() {
	pixelgl.Run(run)
}

func run() {
	// if number of arguments is not 1, print usage and exit
	// the easter egg is found on second 7572
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <seconds>\n", os.Args[0])
		os.Exit(1)
	}
	seconds := day1.Atoi(os.Args[1])

	lobby := day14.ParseLobby(day14.Point{lobbyWidth, lobbyHeight}, day1.ReadFile("../day14.txt"))
	lobby.Simulate(seconds)

	cfg := pixelgl.WindowConfig{
		Title:  "Lobby Simulation",
		Bounds: pixel.R(0, 0, squareSide*lobbyWidth, squareSide*lobbyHeight), // Adjust the window size based on the lobby size
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.Clear(color.White)
	for !win.Closed() {
		win.Clear(color.White)
		drawLobby(win, lobby.Map().String())
		win.Update()
	}
}

func drawLobby(win *pixelgl.Window, lobbyMap string) {
	lines := strings.Split(lobbyMap, "\n")
	for y, line := range lines {
		for x, char := range line {
			if char != '.' {
				drawSquare(win, float64(x*squareSide), float64(len(lines)-y-1)*squareSide)
			}
		}
	}
}

func drawSquare(win *pixelgl.Window, x, y float64) {
	imd := pixel.NewSprite(nil, pixel.R(x, y, x+2, y+2))
	imd.Draw(win, pixel.IM.Moved(pixel.V(x+1, y+1)))
}

func drawSeconds(seconds int, proportion float64) {
	fmt.Printf("\r%04d proportion %.3f", seconds, proportion)
	//basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	//basicTxt := text.New(pixel.V(10, 100), basicAtlas)
	//
	//_, _ = fmt.Fprintf(basicTxt, "%04d", seconds)
}
