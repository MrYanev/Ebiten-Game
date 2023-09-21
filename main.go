package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/font"
)

const (
	screenWidth   = 640
	screenHeight  = 480
	tileSize      = 32
	titleFontSize = fontSize * 1.5
	fontSize      = 24
	smallFontSize = fontSize / 2
)

var (
	knightImage            *ebiten.Image
	platformImage          *ebiten.Image
	titleArcadeFont        font.Face
	arcadeFont             font.Face
	smallArcadeFont        font.Face
	charecterX, charecterY float64 = 100, 100
)

type Game struct{}

func (g *Game) Update(screen *ebiten.Image) error {
	if !internal.isGameStarted {
		internal.updateStartMenu(screen)
	} else {
		//Game logic
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Broforece V2")
	if !internal.isGameStarted {
		internal.drawStartMenu(screen)
	} else {
		// Draw game elements here
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	game := &Game{}

	ebiten.SetWindowSize(1000, 650)
	ebiten.SetWindowTitle("Broforce V2")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
