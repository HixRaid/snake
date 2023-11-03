package apple

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var appleColor = color.RGBA{255, 32, 32, 255}

type Apple struct {
	Position [2]float32
}

func NewApple(pos [2]float32) *Apple {
	return &Apple{
		Position: pos,
	}
}

func (a *Apple) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, a.Position[0], a.Position[1], 1, 1, appleColor, false)
}
