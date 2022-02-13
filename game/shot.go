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
	SHOT_SCALE    = 0.1
	SHOT_COOLTIME = 30
	SHOT_SPEED    = 3
)

var (
	shotNum      int
	shotCooltime int
)

type Shot struct {
	img        *ebiten.Image
	startFrame int
	x          float64
	y          float64
}

func (g *Game) drawShot(screen *ebiten.Image) {
	if shotCooltime == 0 {
		for _, k := range g.input.Keys {
			switch k {
			case ebiten.KeySpace:

				shotCooltime++
				img, _, err := ebitenutil.NewImageFromFile("game/resource/img/shot.png")
				if err != nil {
					log.Fatal(err)
				}
				g.shotImg[shotNum] = Shot{
					img:        img,
					startFrame: g.count,
					x:          rocketX,
					y:          rocketY,
				}

				shotNum++
				if shotNum >= MaxShotNum {
					shotNum = 0
				}
				continue
			}
		}
	}

	if shotCooltime > 0 {
		shotCooltime++
	}
	if shotCooltime > SHOT_COOLTIME {
		shotCooltime = 0
	}
}

func (g *Game) moveShot(screen *ebiten.Image) {
	for shotNum, v := range g.shotImg {
		if v.img != nil {
			g.shotImg[shotNum].y -= SHOT_SPEED
			v.y -= SHOT_SPEED

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Scale(SHOT_SCALE, SHOT_SCALE)
			op.GeoM.Translate(v.x+12, v.y-40)

			text.Draw(screen, fmt.Sprintf("shotNum: %d, shotX, shotY:%d, %d", shotNum, int(v.x+12), int(v.y-40)), mplusbitmap.Gothic12r, 5, 133+(20*shotNum), color.White)
			screen.DrawImage(v.img, op)

			if v.y < 0 {
				g.shotImg[shotNum].img.Dispose()
				g.shotImg[shotNum].img = nil
			}
		}
	}
}
