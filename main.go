package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	screenWidth  = 1000
	screenHeight = 480
)

var (
	running = true
	bgColor = rl.NewColor(147, 211, 196, 255)

	// setting our texture variables, vars will be initialized with a png asset on init
	grassSprite  rl.Texture2D
	playerSprite rl.Texture2D

	playerSrc  rl.Rectangle
	playerDest rl.Rectangle

	playerSpeed float32 = 3
)

// function for drawing a scene
func drawScene() {

	// drawTexture takes in a texture, position x, position y, and a tint
	rl.DrawTexture(grassSprite, 100, 50, rl.White)

	// this func takes the texture, the source of the img(where in the img we need to draw), the destination, origin, rotation, and tint
	rl.DrawTexturePro(playerSprite, playerSrc, playerDest, rl.NewVector2(playerDest.Width, playerDest.Height), 0, rl.White)

}

// function for making an input
func input() {
	// making movements for player
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		// changing the players y coord by the players speed in px
		playerDest.Y -= playerSpeed
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		playerDest.Y += playerSpeed
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		playerDest.X -= playerSpeed
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		playerDest.X += playerSpeed
	}
}

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
	// this func takes 3 args, the height, width, and title of game
	rl.InitWindow(screenWidth, screenHeight, "Sprout Kingdom")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)

	grassSprite = rl.LoadTexture("assets/Tilesets/grass.png")
	playerSprite = rl.LoadTexture("assets/Characters/BasicCharakterSpritesheet.png")

	// initializing playersrc to a new rectangle with a width and height of 48px
	playerSrc = rl.NewRectangle(0, 0, 48, 48)
	// init playerDest(this is how big we want the player to be 100px)
	playerDest = rl.NewRectangle(200, 200, 100, 100)
}

// quiting function that handles window close
func quit() {
	// unloading textures from gpu memory
	rl.UnloadTexture(grassSprite)
	rl.UnloadTexture(playerSprite)
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
