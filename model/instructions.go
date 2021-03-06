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
	var X = uint16(Register[registerX]) % 64 // column
	var Y = uint16(Register[registerY]) % 32 // row

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
	if Legacy {
		ProgramCounter = value + uint16(Register[0])
	} else {
		ProgramCounter = 256*register + value + uint16(Register[register])
	}
}

func Random(register uint16, value byte) {
	Register[register] = byte(rand.Int()) & value
}

func SkipIfKey(register uint16, pressed bool) {
	if ebiten.IsKeyPressed(KeyMap[uint16(Register[register])]) == pressed {
		ProgramCounter += 2
	}
}

func SetVXToDelay(register uint16) {
	Register[register] = DelayTimer
}

func SetDelayToVX(register uint16) {
	DelayTimer = Register[register]
}

func SetBeepToVX(register uint16) {
	SoundTimer = Register[register]
}

func AddToIndex(value uint16) {
	if IndexRegister+value > 4095 {
		Register[15] = 1
	}
	IndexRegister += value
}

func GetKey(register uint16) {
	for i := range KeyMap {
		if ebiten.IsKeyPressed(KeyMap[i]) {
			Register[register] = byte(i)
			return
		}
	}
	ProgramCounter -= 2
}

func FontCharacter(register uint16) {
	IndexRegister = LetterOffset[Register[register]]
}

func DecimalConversion(register uint16) {
	value := Register[register]
	dig1 := value % 10
	value /= 10
	dig2 := value % 10
	value /= 10
	dig3 := value

	Memory[IndexRegister] = dig3
	Memory[IndexRegister+1] = dig2
	Memory[IndexRegister+2] = dig1
}

func StoreMemory(value uint16) {
	for i := uint16(0); i <= value; i++ {
		Memory[IndexRegister+i] = Register[i]
	}
	if Legacy {
		IndexRegister += value + 1
	}
}

func LoadMemory(value uint16) {
	for i := uint16(0); i <= value; i++ {
		Register[i] = Memory[IndexRegister+i]
	}
	if Legacy {
		IndexRegister += value + 1
	}
}
