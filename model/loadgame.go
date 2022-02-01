package model

import (
	"chip8/utils"
	"fmt"
	"os"
)

func Load() {
	var path string
	fmt.Scanln(&path)

	dat, err := os.ReadFile(path)
	utils.Assert(err)

	for i := range dat {
		Memory[GameOffset+i] = dat[i]
	}

	ProgramCounter = uint16(GameOffset)
}
