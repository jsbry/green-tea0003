package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Input struct {
	Keys []ebiten.Key
}

func NewInput() *Input {
	return &Input{}
}

func (i *Input) Update(g *Game) {
	i.Keys = inpututil.AppendPressedKeys(i.Keys[:0])
}
