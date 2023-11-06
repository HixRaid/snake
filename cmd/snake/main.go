package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hixraid/snake/internal/apple"
	"github.com/hixraid/snake/internal/input"
	"github.com/hixraid/snake/internal/snake"
)

const (
	fieldWidth  = 40
	fieldHeight = 40
	scale       = 10
	tps         = 30
)

var (
	backgroundColor = color.RGBA{16, 16, 16, 255}
	fieldSize       = [2]float32{float32(fieldWidth), float32(fieldHeight)}
)

type Game struct {
	snake *snake.Snake
	apple *apple.Apple
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		if g.snake.GetStatus() == snake.Dead {
			g.Start()
		}
		g.snake.SetMode(!g.snake.GetMode())
	} else {
		g.snake.Direction = input.GetDirection(g.snake.Direction)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)
	switch g.snake.GetMode() {
	case snake.Play:
		g.snake.Draw(screen)
		g.apple.Draw(screen)
	case snake.Pause:
		if g.snake.GetStatus() == snake.Live {
			ebitenutil.DebugPrint(screen, "ESC\nPause")
		} else {
			ebitenutil.DebugPrint(screen, "Dead\n8")
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return fieldWidth, fieldHeight
}

func (g *Game) Start() {
	g.apple = apple.New()
	checkIntersection := g.apple.CheckIntersection([2]int32{int32(fieldWidth), int32(fieldHeight)})

	tailMove := []snake.TailMove{checkFieldBoundaries, checkIntersection}
	g.snake = snake.NewSnake(tailMove, [2]float32{20, 10}, [2]float32{1, 0}, 8)
}

func checkFieldBoundaries(length *[][2]float32, headIndex int) snake.TailStatus {
	switch {
	case (*length)[headIndex][0] > fieldSize[0]-1:
		(*length)[headIndex][0] = 0
	case (*length)[headIndex][0] < 0:
		(*length)[headIndex][0] = fieldSize[0] - 1
	case (*length)[headIndex][1] > fieldSize[1]-1:
		(*length)[headIndex][1] = 0
	case (*length)[headIndex][1] < 0:
		(*length)[headIndex][1] = fieldSize[1] - 1
	}

	return snake.Live
}

func main() {
	ebiten.SetWindowSize(fieldWidth*scale, fieldHeight*scale)
	ebiten.SetWindowTitle("Snake")
	ebiten.SetTPS(tps)

	game := Game{}
	game.Start()

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
