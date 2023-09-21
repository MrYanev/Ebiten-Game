package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Game struct{}

// Create our empty vars
var (
	err        error
	background *ebiten.Image
)

func (g *Game) Update(screen *ebiten.Image) error {
	// Games logic goes here
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(background, op)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	background, _, err = ebitenutil.NewImageFromFile("assets/images/background.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

// Main loop
func main() {
	game := &Game{}
	ebiten.SetWindowSize(1200, 750)
	ebiten.SetWindowTitle("Some Knight Game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
