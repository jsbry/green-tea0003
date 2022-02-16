package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ROCKET_SCALE = 0.1
	ROCKET_SPEED = 2
)

func (g *Game) drawRocket(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(ROCKET_SCALE, ROCKET_SCALE)

	for _, k := range g.input.Keys {
		switch k {
		case ebiten.KeyA, ebiten.KeyArrowLeft:
			g.meImg.x -= ROCKET_SPEED
			continue
		case ebiten.KeyD, ebiten.KeyArrowRight:
			g.meImg.x += ROCKET_SPEED
			continue
		case ebiten.KeyW, ebiten.KeyArrowUp:
			g.meImg.y -= ROCKET_SPEED
			continue
		case ebiten.KeyS, ebiten.KeyArrowDown:
			g.meImg.y += ROCKET_SPEED
			continue
		}
	}

	if g.meImg.x <= -6 {
		g.meImg.x = -6
	}
	// pngのwidth引いて余白のwidth足す
	if windowX-42+6 <= int(g.meImg.x) {
		g.meImg.x = float64(windowX - 42 + 6)
	}

	if g.meImg.y <= -3 {
		g.meImg.y = -3
	}
	// pngのheight引いて余白のheight足す
	if windowY-72+6 <= int(g.meImg.y) {
		g.meImg.y = float64(windowY - 72 + 6)
	}

	op.GeoM.Translate(g.meImg.x, g.meImg.y)
	screen.DrawImage(g.meImg.img, op)
}
