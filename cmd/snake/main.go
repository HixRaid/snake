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

var backgroundColor = color.RGBA{16, 16, 16, 255}
var fieldSize = [2]float32{float32(fieldWidth), float32(fieldHeight)}

type Game struct {
	snake *snake.Snake
	apple *apple.Apple
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		if g.snake.Status == snake.Dead {
			g.snake = snake.NewSnake(fieldSize, [2]float32{20, 10}, [2]float32{1, 0}, 8)
		}
		g.snake.SetMode(!g.snake.Mode)
	} else {
		g.snake.Direction = input.GetDirection(g.snake.Direction)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)
	switch g.snake.Mode {
	case snake.Play:
		g.snake.Draw(screen)
		g.apple.Draw(screen)
	case snake.Pause:
		if g.snake.Status == snake.Live {
			ebitenutil.DebugPrint(screen, "ESC\nPause")
		} else {
			ebitenutil.DebugPrint(screen, "Dead\n8")
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return fieldWidth, fieldHeight
}

func main() {
	ebiten.SetWindowSize(fieldWidth*scale, fieldHeight*scale)
	ebiten.SetWindowTitle("Snake")
	ebiten.SetTPS(tps)

	game := Game{
		snake: snake.NewSnake(fieldSize, [2]float32{20, 10}, [2]float32{1, 0}, 8),
		apple: apple.NewApple([2]float32{30, 30}),
	}

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
