package game

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	mplusbitmap "github.com/hajimehoshi/go-mplus-bitmap"
)

const (
	LayoutWidth  = 640
	LayoutHeight = 480

	MaxShotNum = 8
)

var (
	windowX int
	windowY int
)

// debug
var (
	i  int
	si string
)

type Game struct {
	meImg   *ebiten.Image
	shotImg [MaxShotNum]*ebiten.Image
	input   *Input
}

func NewGame() *Game {
	windowX, windowY = ebiten.WindowSize()
	rocketX = (float64(windowX) / 2) - 21
	rocketY = float64(windowY) - 72

	g := &Game{
		input: NewInput(),
	}

	var err error
	g.meImg, _, err = ebitenutil.NewImageFromFile("game/resource/img/rocket.png")
	if err != nil {
		log.Fatal(err)
	}

	x, y := g.meImg.Size()
	fmt.Printf("x, y : %d, %d", x, y)

	return g
}

func (g *Game) Update() error {
	g.input.Update(g)
	i++
	// PrintMemory()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// ebitenutil.DebugPrint(screen, "nihongoha tukaen")

	si = ""
	for i, v := range g.shotImg {
		if v != nil {
			si += fmt.Sprintf("%d, ", i)
		}
	}

	text.Draw(screen, fmt.Sprintf("update:%d", i), mplusbitmap.Gothic12r, 5, 13, color.White)
	text.Draw(screen, fmt.Sprintf("key:%s", g.input.Keys), mplusbitmap.Gothic12r, 5, 33, color.White)
	text.Draw(screen, fmt.Sprintf("shotImg:%#v", si), mplusbitmap.Gothic12r, 5, 53, color.White)
	text.Draw(screen, fmt.Sprintf("shotCooltime:%d", shotCooltime), mplusbitmap.Gothic12r, 5, 73, color.White)

	text.Draw(screen, fmt.Sprintf("WindowSize():%d, %d", windowX, windowY), mplusbitmap.Gothic12r, 5, 93, color.White)
	text.Draw(screen, fmt.Sprintf("rocketX, rocketY:%d, %d", int(rocketX), int(rocketY)), mplusbitmap.Gothic12r, 5, 113, color.White)

	g.drawRocket(screen)
	g.drawShot(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
	// return 320, 240
}
