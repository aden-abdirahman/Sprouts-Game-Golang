package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	screenWidth  = 1000
	screenHeight = 480
)

var (
	running = true
	bgColor = rl.NewColor(147, 211, 196, 255)
)

// function for drawing a scene
func drawScene() {}

// function for making an input
func input() {}

// function for making a change based on that input
func update() {
	running = !rl.WindowShouldClose()
}

// function for drawing it to the screen
func render() {
	rl.BeginDrawing()

	rl.ClearBackground(bgColor)

	drawScene()

	rl.EndDrawing()
}

// init function that handles initializing rl screen and fps. Init function automatically executes when a package is loaded
func init() {
	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - basic window")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)
}

// quiting function that handles window close
func quit() {
	rl.CloseWindow()
}

func main() {

	// main loop for game, while running is true the loop will continue
	for running {
		input()
		update()
		render()
	}

	quit()
}
