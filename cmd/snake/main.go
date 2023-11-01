package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hixraid/snake/internal/input"
	"github.com/hixraid/snake/internal/snake"
)

const (
	fieldWidth  = 40
	fieldHeight = 40
	scale       = 10
	tps         = 30
)

var fieldSize = [2]float32{float32(fieldWidth), float32(fieldHeight)}

type Game struct {
	*snake.Snake
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.SetMode(!g.Mode)
	} else {
		g.Direction = input.GetDirection(g.Direction)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.Mode {
	case snake.Play:
		g.Snake.Draw(screen)
	case snake.Pause:
		ebitenutil.DebugPrint(screen, "Pause")
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
		Snake: snake.NewSnake(fieldSize, [2]float32{20, 10}, [2]float32{1, 0}, 8),
	}

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
