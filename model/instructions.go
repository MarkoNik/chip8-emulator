package model

import "fmt"

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
	var X = Register[registerX] % 64
	var Y = Register[registerY] % 32

	Register[15] = 0 // set flag register to 0 (no pixels turned off)

	for i := uint16(0); i < value; i++ {
		if uint16(X)+i > 31 {
			continue
		}

		line := Register[IndexRegister+i]
		var mask byte = 1 << 7
		for j := uint16(0); j < 8; j++ {
			if uint16(Y)+j > 63 {
				continue
			}

			if line&mask > 0 { // flip bit
				if Display[uint16(X)+i][uint16(Y)+j] == true {
					Display[uint16(X)+i][uint16(Y)+j] = false
					Register[15] = 1
				} else {
					Display[uint16(X)+i][uint16(Y)+j] = true
				}
			}
		}
	}

	fmt.Println(Display)
}
