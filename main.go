package main

import (
	"github.com/hajimehoshi/ebiten"
)

func main() {
	ebiten.SetWindowSize(internal.screenWidth, internal.screenHeight)
	ebiten.SetWindowTitle("Broforce V2 on Ebiten")
	if err := ebiten.RunGame(internal.update, internal.screenWidth, internal.screenHeight, 2, "Broforce V2 on Ebiten"); err != nil {
		panic(err)
	}

}
