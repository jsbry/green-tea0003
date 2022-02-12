package game

import (
	"fmt"
	"image/color"
	"log"
	"time"

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

func (g *Game) drawShot(screen *ebiten.Image) {
	if shotCooltime == 0 {
		for _, k := range g.input.Keys {
			switch k {
			case ebiten.KeySpace:
				var err error

				shotCooltime++
				g.shotImg[shotNum], _, err = ebitenutil.NewImageFromFile("game/resource/img/shot.png")
				if err != nil {
					log.Fatal(err)
				}
				go g.moveShot(screen, shotNum, rocketX)

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

func (g *Game) moveShot(screen *ebiten.Image, shotNum int, x float64) {
	for y := rocketY; y > 0; y -= SHOT_SPEED {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(SHOT_SCALE, SHOT_SCALE)
		op.GeoM.Translate(x+12, y-40)

		text.Draw(screen, fmt.Sprintf("shotNum: %d, shotX, shotY:%d, %d", shotNum, int(x+12), int(y-40)), mplusbitmap.Gothic12r, 5, 133+(20*shotNum), color.White)

		screen.DrawImage(g.shotImg[shotNum], op)
		time.Sleep(1 * time.Millisecond)
	}
	g.shotImg[shotNum].Dispose()
	g.shotImg[shotNum] = nil
}
