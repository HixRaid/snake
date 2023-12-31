package snake

import (
	"image/color"
	"math"
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
	headColor = color.RGBA{32, 255, 32, 255}
	tailColor = color.RGBA{64, 196, 64, 255}
)

type Snake struct {
	direction [2]float32
	mode      SnakeMode
	tail      *tail
	pause     chan bool
	dead      chan bool
	mu        sync.RWMutex
}

func NewSnake(tailMove []TailMove, pos, dir [2]float32, len int) *Snake {
	return &Snake{
		direction: dir,
		mode:      Pause,
		tail:      newTail(append(tailMove, checkCollision), pos, dir, len),
		dead:      make(chan bool, 1),
	}
}

func (s *Snake) SetMode(mode SnakeMode) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.mode == mode || s.tail.status == Dead {
		return
	}

	s.mode = mode

	switch s.mode {
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
					s.mode = Pause
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

	if s.tail.move(s.direction) == Dead {
		s.tail.status = Dead
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

func (s *Snake) SetDirection(dir [2]float32) {
	if (dir[0] != 0 || dir[1] != 0) &&
		dir[0]+s.direction[0] != 0 && dir[1]+s.direction[1] != 0 &&
		math.Abs(float64(dir[0])) != math.Abs(float64(dir[1])) {
		s.direction = dir
	}
}

func (s *Snake) GetStatus() TailStatus {
	return s.tail.status
}

func (s *Snake) GetMode() SnakeMode {
	return s.mode
}

func (s *Snake) GetLength() int {
	return len(s.tail.length)
}

func checkCollision(length *[][2]float32, headIndex int) TailStatus {
	for i, v := range *length {
		if v == (*length)[headIndex] && i != headIndex {
			return Dead
		}
	}
	return Live
}
