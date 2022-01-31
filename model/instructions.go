package model

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
