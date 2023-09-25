package game

import (
	"bytes"
	gameImages "ebiten-game/assets/images"
	"image"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	imageBackground *ebiten.Image
	isButtonPressed bool
)

func init() {
	//Background Image
	background, _, err = image.Decode(bytes.NewReader(gameImages.Background_png))
	if err != nil {
		log.Fatal(err)
	}
	imageBackground = ebiten.NewImage(background)
}

type TitleScreen struct {
	cont int
}

func anyButtonPressed(screen *ebiten.Image) bool {
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		isButtonPressed = true
	}
	return false
}

func (s *TitleScreen) Update(state *GameState) error {
	s.cont++
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		state.SceneManager.GoTo(NewGameScene())
		return nil
	}
	return nil
}

func (s *TitleScreen) Draw(r *ebiten.Image) {
	s.drawTitleBackground(r, s.cont)
	drawLogo(r, "Knight Game")

	message := "PRESS ENTER TO START"
	x := 0
	y := ScreenHeight / 2
	drawTextWithShadowCenter(r, message, x, y, 1, color.NRGBA{0x80, 0, 0, 0xff}, ScreenWidth)
}

func (s *TitleScreen) drawTitleBackground(r *ebiten.Image, c int) {
	w, h := imageBackground.Bounds().Dx(), imageBackground.Bounds().Dy()
	op := &ebiten.DrawImageOptions{}
	for i := 0; i < (ScreenWidth/w+1)*(ScreenHeight/h+2); i++ {
		op.GeoM.Reset()
		dx := -(c / 4) % w
		dy := (c / 4) % h
		dstX := (i%(ScreenWidth/w+1))*w + dx
		dstY := (i/(ScreenWidth/w+1)-1)*h + dy
		op.GeoM.Translate(float64(dstX), float64(dstY))
		r.DrawImage(imageBackground, op)
	}
}

func drawLogo(r *ebiten.Image, str string) {
	const scale = 4
	x := 0
	y := 32
	drawTextWithShadowCenter(r, str, x, y, scale, color.NRGBA{0x00, 0x00, 0x80, 0xff}, ScreenWidth)
}
