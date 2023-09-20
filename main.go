package main

import (
	"bytes"
	"image"
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
	gopherImage     *ebiten.Image
	tilesImage      *ebiten.Image
	titleArcadeFont font.Face
	arcadeFont      font.Face
	smallArcadeFont font.Face
)

type Game struct{}

func inint() {
	img, _, err := image.Decode(bytes.NewReader(images.run_png))
	if err != nil {
		log.Fatal(err)
	}
	runnerImage := ebiten.NewImageFromImage(img)

	img, _, err := image.Decode(bytes.NewReader(images.idle_png))
	if err != nil {
		log.Fatal(err)
	}
	idleImage := ebiten.NewImageFromImage(img)
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Broforece V2")
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
