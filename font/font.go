package font

import (
	"bufio"
	"chip8/model"
	"chip8/utils"
	"os"
	"strings"
)

var Offset int = 80
var LetterOffset = make(map[byte]byte)

func InitFont() {

	file, err := os.Open("font/font.txt")
	utils.Assert(err)

	scanner := bufio.NewScanner(file)
	address := Offset

	j := byte(0)
	for scanner.Scan() {

		line := scanner.Text()
		split := strings.Split(line, ",")
		LetterOffset[j] = byte(address)
		j++

		for i := 0; i < 5; i++ {
			model.Memory[address] = utils.HexToByte(split[i])
			address++
		}
	}

	err = file.Close()
	utils.Assert(err)

}
