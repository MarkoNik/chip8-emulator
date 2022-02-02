package model

import (
	"github.com/hajimehoshi/ebiten"
)

var KeyMap = map[uint16]ebiten.Key{
	1: ebiten.Key1, 2: ebiten.Key2, 3: ebiten.Key3, 12: ebiten.Key4,
	4: ebiten.KeyQ, 5: ebiten.KeyW, 6: ebiten.KeyE, 13: ebiten.KeyR,
	7: ebiten.KeyA, 8: ebiten.KeyS, 9: ebiten.KeyD, 14: ebiten.KeyF,
	10: ebiten.KeyZ, 0: ebiten.KeyX, 11: ebiten.KeyC, 15: ebiten.KeyV,
}
