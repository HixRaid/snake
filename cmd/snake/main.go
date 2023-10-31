package main

import (
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hixraid/snake/internal/input"
	"github.com/hixraid/snake/internal/snake"
)

const (
	fieldWidth  = 40
	fieldHeight = 40
	scale       = 12
	tps         = 10
)

var fieldSize = [2]float32{float32(fieldWidth), float32(fieldHeight)}
var ticker time.Ticker

type Game struct {
	*snake.Snake
}

func (g *Game) Update() error {
	g.Direction = input.GetDirection(g.Direction)

	select {
	case <-ticker.C:
		g.Snake.Move(fieldSize)
	default:
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	g.Snake.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return fieldWidth, fieldHeight
}

func main() {
	ebiten.SetWindowSize(fieldWidth*scale, fieldHeight*scale)
	ebiten.SetWindowTitle("Snake")

	game := Game{
		Snake: snake.NewSnake([2]float32{20, 10}, [2]float32{1, 0}, 20),
	}

	ticker = *time.NewTicker(time.Second / tps)

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
