package main

import (
	"log"

	"ebiten-game/internal/game"

	"github.com/hajimehoshi/ebiten"
)

/*
// Create the player class
type player struct {
	image      *ebiten.Image
	xPos, yPos float64
	speed      float64
}

// Run this code once at startup

	func (g *Game) initializeStartMenu() {
		//Load embeded font
		fontData, err := truetype.Parse(fontBytes)
		if err != nil {
			log.Fatal(err)
		}

		//Create a font face
		fontFace = truetype.NewFace(fontData, &truetype.Options{
			Size: 24,
		})

		// Measure the text size
		textWidth, textHeight = text.MeasureString("Press Enter to start", fontFace)
	}

	func (g *Game) drawStartMenu(screen *ebiten.Image) {
		screen.DrawImage(background, nil)

		if !isGameStarted {
			//Draw the message
			text := "Press ENTER to start"
			textWidht, textHeight := text.BoundString(textLayout, text) //To be fixed

			textX := (screenWidth - textWidht) / 2
			textY := (screenHeight - textHeight) / 2
			text.Draw(screen, text, arcadeFont, textX, textY, textColor)

		}
	}

	func (g *Game) UpdateStartMenu(screen *ebiten.Image) {
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			isGameStarted = true
		}
	}

	func (g *Game) MovePlayer() {
		if ebiten.IsKeyPressed(ebiten.KeyUp) {
			playerOne.yPos -= playerOne.speed
		}
		if ebiten.IsKeyPressed(ebiten.KeyDown) {
			playerOne.yPos += playerOne.speed
		}
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			playerOne.xPos -= playerOne.speed
		}
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			playerOne.xPos += playerOne.speed
		}
	}

*/

func main() {
	theGame = &game.Game{}
	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetWindowTitle("Some Knight Game")
	if err := ebiten.RunGame(theGame); err != nil {
		log.Fatal(err)
	}
}
