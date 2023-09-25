package graphics

import (
	"ebiten-game/internal/game"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	imageGameBG   *ebiten.Image
	imageWindows  = ebiten.NewImage(game.ScreenWidth, game.ScreenHeight)
	imageGameOver = ebiten.NewImage(game.ScreenWidth, game.ScreenHeight)
)

func init() {

	//Player Image
	knight, _, err = ebitenutil.NewImageFromFile("D:\\Code-projects\\Ebiten-Game\\assets\\images\\_Idle.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	playerOne = player{knight, screenWidth / 2.0, screenHeight / 2.0, 4}

	//Gameover screen
	imageGameOver.Fill(color.NRGBA{0x00, 0x00, 0x00, 0x80})
	y := (game.ScreenHeight - knightHeight) / 2
	drawTextWithShadowCenter(imageGameOver, "GAME OVER\n\nPRESS ENTER", 0, y, 1, color.White, game.ScreenWidth)
}

type GameScene struct {
	gameover bool
}
