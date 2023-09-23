package game

import "github.com/hajimehoshi/ebiten/v2"

type Player struct {
	image      *ebiten.Image
	xPos, yPos float64
	speed      float64
}
