package model

var Memory [4096]byte
var Display [32][64]bool
var Stack [16]uint16
var DelayTimer byte
var SoundTimer byte

var ProgramCounter uint16
var IndexRegister uint16
var Register [16]byte
var instructionRegister uint16

func fetch() {
	instructionRegister = uint16(Memory[ProgramCounter]) << 8
	instructionRegister += uint16(Memory[ProgramCounter+1])
	ProgramCounter += 2
}

func decode() {

	// masks used for n extraction
	var maskXYN uint16 = 1<<4 - 1
	var maskNN uint16 = 1<<8 - 1
	var maskNNN uint16 = 1<<12 - 1

	var n, nX, nY, nN, nNN, nNNN uint16
	//extract nibbles
	{
		n = instructionRegister & (maskXYN << 12)
		n >>= 12

		nX = instructionRegister & (maskXYN << 8)
		nX >>= 8

		nY = instructionRegister & (maskXYN << 4)
		nY >>= 4

		nN = instructionRegister & maskXYN

		nNN = instructionRegister & maskNN

		nNNN = instructionRegister & maskNNN
	}

	switch n {
	case 0:
		{
			if nX == 0 && nY == 0xE && nN == 0 {
				Clear()
			}
		}
	case 1:
		{
			Jump(nNNN)
		}
	case 2:
		{

		}
	case 3:
		{

		}
	case 4:
		{

		}
	case 5:
		{

		}
	case 6:
		{
			SetRegister(nX, nNN)
		}
	case 7:
		{
			AddRegister(nX, nNN)
		}
	case 8:
		{

		}
	case 9:
		{

		}
	case 10:
		{
			SetIndexRegister(nNNN)
		}
	case 11:
		{

		}
	case 12:
		{

		}
	case 13:
		{

		}
	case 14:
		{

		}
	case 15:
		{

		}
	}
}
