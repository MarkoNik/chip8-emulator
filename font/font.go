package font

import (
	"bufio"
	"chip8/model"
	"chip8/utils"
	"os"
	"strings"
)

func InitFont() {

	file, err := os.Open("font/font.txt")
	utils.Assert(err)

	scanner := bufio.NewScanner(file)
	address := model.FontOffset

	j := byte(0)
	for scanner.Scan() {

		line := scanner.Text()
		split := strings.Split(line, ",")
		model.LetterOffset[j] = uint16(address)
		j++

		for i := 0; i < 5; i++ {
			model.Memory[address] = utils.HexToByte(split[i])
			address++
		}
	}

	err = file.Close()
	utils.Assert(err)
}
