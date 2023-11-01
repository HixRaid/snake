package snake

import (
	"image/color"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type SnakeMode bool

const (
	Pause SnakeMode = false
	Play  SnakeMode = true
)

var (
	headColor = color.RGBA{255, 63, 0, 255}
	tailColor = color.RGBA{63, 255, 0, 255}
)

type Snake struct {
	Mode      SnakeMode
	Direction [2]float32
	FieldSize [2]float32
	Tail      *tail
	stop      chan bool
	mu        sync.RWMutex
}

func NewSnake(fieldSize [2]float32, pos, dir [2]float32, len int) *Snake {
	return &Snake{
		Mode:      Pause,
		Direction: dir,
		FieldSize: fieldSize,
		Tail:      newTail(pos, dir, len),
	}
}

func (s *Snake) SetMode(mode SnakeMode) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.Mode == mode {
		return
	}
	s.Mode = mode

	switch s.Mode {
	case Play:
		s.stop = make(chan bool)
		ticker := time.NewTicker(time.Second / 10)
		go func() {
			for {
				select {
				case <-ticker.C:
					s.Move()
				case <-s.stop:
					close(s.stop)
					return
				}
			}
		}()
	case Pause:
		s.stop <- true
	}
}

func (s *Snake) Move() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Tail.move(s.Direction, s.FieldSize)

}

func (s *Snake) Draw(screen *ebiten.Image) {
	s.mu.RLock()
	defer s.mu.RUnlock()
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
