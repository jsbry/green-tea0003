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
					x:          g.meImg.x,
					y:          g.meImg.y,
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

	g.moveShot(screen)
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

			isBreak := g.Intersect(v)
			screen.DrawImage(v.img, op)

			if isBreak || v.y < 0 {
				g.shotImg[shotNum].img.Dispose()
				g.shotImg[shotNum].img = nil
			}
		}
	}
}

func (g *Game) Intersect(shot Shot) bool {
	bx, by := shot.x, shot.y
	bw, bh := shot.img.Size()

	for enemyNum, v := range g.enemyImg {
		if v.img != nil {
			ax, ay := v.x, v.y
			aw, ah := v.img.Size()

			// https://qiita.com/zenwerk/items/7123d878a9ad4ecc291f
			if (ax < bx+float64(bw)*SHOT_SCALE && ay < by+float64(bh)*SHOT_SCALE) && (ax+float64(aw)*ENEMY_SCALE > bx && ay+float64(ah)*ENEMY_SCALE > by) {
				g.enemyImg[enemyNum].img.Dispose()
				g.enemyImg[enemyNum].img = nil

				return true
			}
		}
	}
	return false
}
