package internal

import "github.com/hajimehoshi/ebiten/v2"

type Enemies struct {
	image      *ebiten.Image
	xPos, yPos float64
	speed      float64
}

var enemieImg *ebiten.Image
