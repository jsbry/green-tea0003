package main

import (
	"log"
	"tententenntenp/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := game.NewGame()
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!にほんご")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
