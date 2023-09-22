package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Player struct { //Here for testing purposes
	image      *ebiten.Image
	xPos, yPos float64
	speed      float64
}

type Game struct{}

const (
	screenWidth, screenHeight = 1200, 750
)

var (
	err        error
	background *ebiten.Image
	enemieImg  *ebiten.Image
	playerImg  *ebiten.Image
	playerOne  Player
)

func (g *Game) Draw(screen *ebiten.Image) {
	background, _, err = ebitenutil.NewImageFromFile("assets/images/background.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	playerImg, _, err = ebitenutil.NewImageFromFile("assets/images/_Idle.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	playerOne = Player{playerImg, screenHeight / 2.0, screenWidth / 2.0, 4} //To be checked

	enemieImg, _, err = ebitenutil.NewImageFromFile("assets/images/_IdleEnem.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

}

func (g *Game) Update(screen *ebiten.Image) error {
	// Games logic goes here
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(background, op)

	playerOp := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(playerOne.image, playerOp) //Related to the previous

	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1200, 750
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
