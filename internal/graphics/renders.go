package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Game struct{}

// Our game constants
const (
	screenWidth, screenHeight = 640, 480
)

// Create our empty vars
var (
	err        error
	background *ebiten.Image
	spaceShip  *ebiten.Image
	playerOne  player
)

// Create the player class
type player struct {
	image      *ebiten.Image
	xPos, yPos float64
	speed      float64
}

// Run this code once at startup
func init() {
	background, _, err = ebitenutil.NewImageFromFile("D:\\Code-projects\\Ebiten-Game\\assets\\images\\background.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	spaceShip, _, err = ebitenutil.NewImageFromFile("D:\\Code-projects\\Ebiten-Game\\assets\\images\\_Idle.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	playerOne = player{spaceShip, screenWidth / 2.0, screenHeight / 2.0, 4}
}

func (g *Game) Update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(background, op)

	playerOp := &ebiten.DrawImageOptions{}
	playerOp.GeoM.Translate(playerOne.xPos, playerOne.yPos)
	screen.DrawImage(playerOne.image, playerOp)

	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1200, 750
}

func main() {
	game := &Game{}
	ebiten.SetWindowSize(1200, 750)
	ebiten.SetWindowTitle("Some Knight Game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
