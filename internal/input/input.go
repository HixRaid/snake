package input

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func GetDirection(cur [2]float32) (result [2]float32) {
	result[0], result[1] = getAxis(ebiten.KeyA, ebiten.KeyD), getAxis(ebiten.KeyW, ebiten.KeyS)
	if (result[0] != 0 || result[1] != 0) &&
		result[0]+cur[0] != 0 && result[1]+cur[1] != 0 &&
		math.Abs(float64(result[0])) != math.Abs(float64(result[1])) {
		return
	}
	return cur
}

func getAxis(neg, pos ebiten.Key) (result float32) {
	if ebiten.IsKeyPressed(pos) && !ebiten.IsKeyPressed(neg) {
		result = 1
	} else if ebiten.IsKeyPressed(neg) && !ebiten.IsKeyPressed(pos) {
		result = -1
	}

	return
}
