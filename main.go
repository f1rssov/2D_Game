package main

import (
	"log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

)

// Game implements ebiten.Game interface.
type myGame struct{
	windowWidth int
	windowHeight int
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *myGame) Update() error {
    // Write your game's logical update.
    return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *myGame) Draw(screen *ebiten.Image) {
    // Write your game's rendering.
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *myGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return 320, 240
}
/////////////////////////////////////////////////////////////////////////////////////////////////


func main() {
    game := &myGame{}
	game.windowWidth = 600
	game.windowHeight = 600
    // Specify the window size as you like. Here, a doubled size is specified.
    ebiten.SetWindowSize(game.windowWidth, game.windowHeight)
    ebiten.SetWindowTitle("2D_Game")
    // Call ebiten.RunGame to start your game loop.
    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
}