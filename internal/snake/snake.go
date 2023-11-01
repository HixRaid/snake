package snake

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	headColor = color.RGBA{255, 63, 0, 255}
	tailColor = color.RGBA{63, 255, 0, 255}
)

type Snake struct {
	Direction [2]float32
	Tail      *tail
}

func NewSnake(pos, dir [2]float32, len int) *Snake {
	return &Snake{
		Direction: dir,
		Tail:      newTail(pos, dir, len),
	}
}

func (s *Snake) Move(fieldSize [2]float32) {
	s.Tail.move(s.Direction, fieldSize)
}

func (s *Snake) Draw(screen *ebiten.Image) {
	for i, v := range s.Tail.Length {
		if i > 0 {
			vector.DrawFilledRect(screen, v[0], v[1], 1, 1, tailColor, false)
		} else {
			vector.DrawFilledRect(screen, v[0], v[1], 1, 1, headColor, false)
		}
	}
}

type tail struct {
	Length [][2]float32
}

func newTail(pos, dir [2]float32, len int) *tail {
	t := tail{
		Length: make([][2]float32, len),
	}

	for i := 0; i < len; i++ {
		t.Length[i][0], t.Length[i][1] = pos[0]-dir[0]*float32(i), pos[1]-dir[1]*float32(i)
	}

	return &t
}

func (t *tail) move(dir [2]float32, fieldSize [2]float32) {

	for i := len(t.Length) - 1; i > 0; i-- {
		t.Length[i][0], t.Length[i][1] = t.Length[i-1][0], t.Length[i-1][1]
	}

	t.Length[0][0], t.Length[0][1] = t.Length[0][0]+dir[0], t.Length[0][1]+dir[1]
	switch {
	case t.Length[0][0] > fieldSize[0]-1:
		t.Length[0][0] = 0
	case t.Length[0][0] < 0:
		t.Length[0][0] = fieldSize[0] - 1
	case t.Length[0][1] > fieldSize[1]-1:
		t.Length[0][1] = 0
	case t.Length[0][1] < 0:
		t.Length[0][1] = fieldSize[1] - 1
	}
}
