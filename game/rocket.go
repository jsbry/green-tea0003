package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ROCKET_SCALE = 0.1
	ROCKET_SPEED = 2
)

var (
	rocketX float64
	rocketY float64
)

func (g *Game) drawRocket(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(ROCKET_SCALE, ROCKET_SCALE)

	for _, k := range g.input.Keys {
		switch k {
		case ebiten.KeyA, ebiten.KeyArrowLeft:
			rocketX -= ROCKET_SPEED
			continue
		case ebiten.KeyD, ebiten.KeyArrowRight:
			rocketX += ROCKET_SPEED
			continue
		case ebiten.KeyW, ebiten.KeyArrowUp:
			rocketY -= ROCKET_SPEED
			continue
		case ebiten.KeyS, ebiten.KeyArrowDown:
			rocketY += ROCKET_SPEED
			continue
		}
	}

	if rocketX <= -6 {
		rocketX = -6
	}
	// pngのwidth引いて余白のwidth足す
	if windowX-42+6 <= int(rocketX) {
		rocketX = float64(windowX - 42 + 6)
	}

	if rocketY <= -3 {
		rocketY = -3
	}
	// pngのheight引いて余白のheight足す
	if windowY-72+6 <= int(rocketY) {
		rocketY = float64(windowY - 72 + 6)
	}

	op.GeoM.Translate(rocketX, rocketY)
	screen.DrawImage(g.meImg, op)
}
