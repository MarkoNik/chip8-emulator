package font

import (
	"bufio"
	"chip8/model"
	"chip8/utils"
	"os"
	"strings"
)

var Offset int = 50

func InitFont() {

	file, err := os.Open("font/font.txt")
	utils.Assert(err)

	scanner := bufio.NewScanner(file)
	address := Offset

	for scanner.Scan() {

		line := scanner.Text()
		split := strings.Split(line, ",")

		for i := 0; i < 5; i++ {
			model.Memory[address] = utils.HexToByte(split[i])
			address++
		}
	}

	err = file.Close()
	utils.Assert(err)

}
