package utils

import (
	"encoding/binary"
	"strconv"
	"strings"
)

func HexToByte(input string) byte {
	input = strings.TrimSpace(input)
	dec, err := strconv.ParseInt(input, 0, 64)
	Assert(err)
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(dec))
	return b[0]
}
