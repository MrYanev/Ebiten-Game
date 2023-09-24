package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	sceneManager *SceneManager
	input        Input
}

const (
	ScreenWidth  = 1152
	ScreenHeight = 648
)

func (g *Game) Update(screen *ebiten.Image) error {
	g.MovePlayer()
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
	return 1154, 650
}

func (g *Game) Draw(scanner *ebiten.Image) {
	g.screenManager.Draw(screen)
}
