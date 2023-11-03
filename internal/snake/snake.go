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

type SnakeStatus bool

const (
	Dead SnakeStatus = false
	Live SnakeStatus = true
)

var (
	headColor = color.RGBA{32, 255, 32, 255}
	tailColor = color.RGBA{64, 196, 64, 255}
)

type Snake struct {
	Mode      SnakeMode
	Status    SnakeStatus
	Direction [2]float32
	tail      *tail
	pause     chan bool
	dead      chan bool
	mu        sync.RWMutex
}

func NewSnake(fieldSize [2]float32, pos, dir [2]float32, len int) *Snake {
	return &Snake{
		Mode:      Pause,
		Status:    Live,
		Direction: dir,
		tail:      newTail(fieldSize, pos, dir, len),
		dead:      make(chan bool, 1),
	}
}

func (s *Snake) SetMode(mode SnakeMode) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.Mode == mode || s.Status == Dead {
		return
	}

	s.Mode = mode

	switch s.Mode {
	case Play:
		s.pause = make(chan bool)
		ticker := time.NewTicker(time.Second / 10)
		go func() {
			for {
				select {
				case <-ticker.C:
					s.Move()
				case <-s.pause:
					return
				case <-s.dead:
					s.Mode = Pause
					return
				}
			}
		}()
	case Pause:
		s.pause <- true
	}
}

func (s *Snake) Move() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.tail.move(s.Direction) == Dead {
		s.Status = Dead
		s.dead <- true
	}
}

func (s *Snake) Draw(screen *ebiten.Image) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for i, v := range s.tail.length {
		if s.tail.headIndex == i {
			vector.DrawFilledRect(screen, v[0], v[1], 1, 1, headColor, false)
		} else {
			vector.DrawFilledRect(screen, v[0], v[1], 1, 1, tailColor, false)
		}
	}
}

type tail struct {
	length    [][2]float32
	fieldSize [2]float32
	headIndex int
}

func newTail(fieldSize, pos, dir [2]float32, len int) *tail {
	t := tail{
		length:    make([][2]float32, len),
		fieldSize: fieldSize,
	}

	for i := 0; i < len; i++ {
		t.length[i][0], t.length[i][1] = pos[0]-dir[0]*float32(i), pos[1]-dir[1]*float32(i)
	}

	return &t
}

func (t *tail) move(dir [2]float32) SnakeStatus {
	targetHeadIndex := t.headIndex - 1
	if targetHeadIndex < 0 {
		targetHeadIndex = len(t.length) - 1
	}

	t.length[targetHeadIndex][0], t.length[targetHeadIndex][1] = t.length[t.headIndex][0]+dir[0], t.length[t.headIndex][1]+dir[1]
	t.headIndex = targetHeadIndex

	switch {
	case t.length[targetHeadIndex][0] > t.fieldSize[0]-1:
		t.length[targetHeadIndex][0] = 0
	case t.length[targetHeadIndex][0] < 0:
		t.length[targetHeadIndex][0] = t.fieldSize[0] - 1
	case t.length[targetHeadIndex][1] > t.fieldSize[1]-1:
		t.length[targetHeadIndex][1] = 0
	case t.length[targetHeadIndex][1] < 0:
		t.length[targetHeadIndex][1] = t.fieldSize[1] - 1
	}

	for i, v := range t.length {
		if v == t.length[t.headIndex] && i != t.headIndex {
			return Dead
		}
	}

	return Live
}
