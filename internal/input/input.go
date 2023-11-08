package input

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func GetDirection() (dir [2]float32) {
	dir[0], dir[1] = getAxis(ebiten.KeyA, ebiten.KeyD), getAxis(ebiten.KeyW, ebiten.KeyS)
	return
}

func getAxis(neg, pos ebiten.Key) (result float32) {
	if ebiten.IsKeyPressed(pos) && !ebiten.IsKeyPressed(neg) {
		result = 1
	} else if ebiten.IsKeyPressed(neg) && !ebiten.IsKeyPressed(pos) {
		result = -1
	}

	return
}
