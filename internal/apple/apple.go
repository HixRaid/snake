package apple

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/hixraid/snake/internal/snake"
)

var appleColor = color.RGBA{255, 32, 32, 255}

type Apple struct {
	pos [2]float32
}

func New() *Apple {
	return &Apple{}
}

func (a *Apple) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, a.pos[0], a.pos[1], 1, 1, appleColor, false)
}

func (a *Apple) CheckIntersection(fieldSize [2]int32) func(*[][2]float32, int) snake.TailStatus {
	a.pos = randPosition(fieldSize)
	return func(length *[][2]float32, headIndex int) snake.TailStatus {
		if (*length)[headIndex] == a.pos {
			a.pos = randPosition(fieldSize)
			*length = append(*length, [2]float32{-1, -1})
		}
		return snake.Live
	}
}

func randPosition(fieldSize [2]int32) (pos [2]float32) {
	pos[0] = float32(rand.Int31n(fieldSize[0] - 1))
	pos[1] = float32(rand.Int31n(fieldSize[1] - 1))
	return
}
