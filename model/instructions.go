package model

import (
	"github.com/hajimehoshi/ebiten"
	"math/rand"
)

func Clear() {
	for i := 0; i < 32; i++ {
		for j := 0; j < 64; j++ {
			Display[i][j] = false
		}
	}
}

func Jump(address uint16) {
	ProgramCounter = address
}

func SetRegister(register uint16, value uint16) {
	Register[register] = byte(value)
}

func AddRegister(register uint16, value uint16) {
	Register[register] += byte(value)
}

func SetIndexRegister(address uint16) {
	IndexRegister = address
}

func DisplayInstruction(registerX uint16, registerY uint16, value uint16) {

	// coordinates to draw to
	var X = uint16(Register[registerX]-1) % 64 // column
	var Y = uint16(Register[registerY]-1) % 32 // row

	Register[15] = 0 // set flag register to 0 (no pixels turned off)

	for i := uint16(0); i < value; i++ {
		if Y+i > 31 {
			continue
		}

		line := Memory[IndexRegister+i]
		var mask byte = 1 << 7
		for j := uint16(0); j < 8; j++ {
			if X+j > 63 {
				continue
			}

			if line&mask > 0 { // flip bit
				if Display[Y+i][X+j] == true {
					Display[Y+i][X+j] = false
					Register[15] = 1
				} else {
					Display[Y+i][X+j] = true
				}
			}
			mask >>= 1
		}
	}
}

func StartSubroutine(address uint16) {
	Stack.Push(ProgramCounter)
	Jump(address)
}

func EndSubroutine() {
	ProgramCounter = Stack.Pop()
}

func Skip(value1, value2 uint16, equal bool) {
	if value1 == value2 == equal {
		ProgramCounter += 2
	}
}

func OrInstruction(register uint16, value byte) {
	Register[register] |= value
}

func AndInstruction(register uint16, value byte) {
	Register[register] &= value
}

func XorInstruction(register uint16, value byte) {
	Register[register] ^= value
}

func AddInstruction(register uint16, value byte) {
	Register[15] = 0
	if Register[register]+value > 255 {
		Register[15] = 1
	}
	Register[register] += value
}

func SubtractInstruction(register uint16, value byte) {
	Register[15] = 1
	if Register[register]-value < 0 {
		Register[15] = 0
	}
	Register[register] -= value
}

func ShiftInstruction(register uint16, value byte, right bool) {
	if Legacy {
		Register[register] = value
	}

	Register[15] = 0
	var mask byte
	if right {
		mask = 1
		if Register[register]&mask > 0 {
			Register[15] = 1
		}
		Register[register] >>= 1
	} else {
		mask = 1 << 7
		if Register[register]&mask > 0 {
			Register[15] = 1
		}
		Register[register] <<= 1
	}
}

func JumpOffset(register uint16, value uint16) {
	ProgramCounter = value + uint16(Register[register])
}

func Random(register uint16, value byte) {
	Register[register] = byte(rand.Int()) & value
}

func SkipIfKey(value uint16, pressed bool) {
	if ebiten.IsKeyPressed(KeyMap[value]) == pressed {
		ProgramCounter += 2
	}
}
