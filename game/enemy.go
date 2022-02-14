package game

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	mplusbitmap "github.com/hajimehoshi/go-mplus-bitmap"
)

const (
	ENEMY_SCALE = 0.1
	ENEMY_SPEED = 1
)

var (
	enemyNum int
)

type Enemy struct {
	img        *ebiten.Image
	startFrame int
	x          float64
	y          float64
}

func (g *Game) drawEnemy(screen *ebiten.Image) {

	img, _, err := ebitenutil.NewImageFromFile("game/resource/img/ufo.png")
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())

	if g.count == 1 || g.count%100 == 0 {
		g.enemyImg[enemyNum] = Enemy{
			img:        img,
			startFrame: g.count,
			x:          float64(rand.Intn(100) - 150),
			y:          float64(rand.Intn(300) + 50),
		}

		enemyNum++
		if enemyNum >= MaxEnemyNum {
			enemyNum = 0
		}
	}

	g.moveEnemy(screen)
}

func (g *Game) moveEnemy(screen *ebiten.Image) {
	for enemyNum, v := range g.enemyImg {
		if v.img != nil {
			g.enemyImg[enemyNum].x += ENEMY_SPEED
			v.x += ENEMY_SPEED

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Scale(ENEMY_SCALE, ENEMY_SCALE)
			op.GeoM.Translate(v.x+12, v.y-40)

			text.Draw(screen, fmt.Sprintf("enemyNum: %d, enemyX, enemyY:%d, %d", enemyNum, int(v.x+12), int(v.y-40)), mplusbitmap.Gothic12r, 350, 133+(20*enemyNum), color.White)
			screen.DrawImage(v.img, op)

			if float64(windowX) < v.x {
				g.enemyImg[enemyNum].img.Dispose()
				g.enemyImg[enemyNum].img = nil
			}
		}
	}
}
