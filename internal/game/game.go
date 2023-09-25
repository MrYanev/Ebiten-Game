package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	sceneManager *SceneManager
	//input        Input
}

const (
	ScreenWidth  = 1152
	ScreenHeight = 648
)

func (g *Game) Update(screen *ebiten.Image) error {
	g.MovePlayer()
	if g.sceneManager == nil {
		g.sceneManager = &SceneManager{}
		g.sceneManager.GoTo(&TitleScreen{})
	}

	/*g.Input.Update() Working around using gamepad input
	if err := g.sceneManager.Update(&g.input); err != nil {
		return err
	}*/
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1154, 650
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.sceneManager.Draw(screen)
}
