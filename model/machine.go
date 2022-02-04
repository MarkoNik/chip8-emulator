package model

const GameOffset int = 512
const FontOffset int = 80

var LetterOffset = make(map[byte]uint16)

var Legacy bool = false
var DisplayChanged = true

var Memory [4096]byte
var Stack stack
var Display [32][64]bool
var DelayTimer byte
var SoundTimer byte

var ProgramCounter uint16
var IndexRegister uint16
var Register [16]byte
var instructionRegister uint16

func Run() {
	fetch()
	//fmt.Println(fmt.Sprintf("%x", instructionRegister))
	decode()
}

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
				DisplayChanged = true
			}
			if nX == 0 && nY == 0xE && nN == 0xE {
				EndSubroutine()
			}
		}
	case 1:
		{
			Jump(nNNN)
		}
	case 2:
		{
			StartSubroutine(nNNN)
		}
	case 3:
		{
			Skip(uint16(Register[nX]), nNN, true)
		}
	case 4:
		{
			Skip(uint16(Register[nX]), nNN, false)
		}
	case 5:
		{
			Skip(uint16(Register[nX]), uint16(Register[nY]), true)
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
			if nN == 0 {
				SetRegister(nX, uint16(Register[nY]))
			}
			if nN == 1 {
				OrInstruction(nX, Register[nY])
			}
			if nN == 2 {
				AndInstruction(nX, Register[nY])
			}
			if nN == 3 {
				XorInstruction(nX, Register[nY])
			}
			if nN == 4 {
				AddInstruction(nX, Register[nY])
			}
			if nN == 5 {
				SubtractInstruction(nX, Register[nY])
			}
			if nN == 7 {
				SubtractInstruction(nY, Register[nX])
			}
			if nN == 6 {
				ShiftInstruction(nX, Register[nY], true)
			}
			if nN == 10 {
				ShiftInstruction(nX, Register[nY], false)
			}
		}
	case 9:
		{
			Skip(uint16(Register[nX]), uint16(Register[nY]), false)
		}
	case 10:
		{
			SetIndexRegister(nNNN)
		}
	case 11:
		{
			if Legacy {
				JumpOffset(0, nNNN)
			} else {
				JumpOffset(nX, nNN)
			}
		}
	case 12:
		{
			Random(nX, byte(nNN))
		}
	case 13:
		{
			DisplayInstruction(nX, nY, nN)
			DisplayChanged = true
		}
	case 14:
		{
			if nY == 9 && nN == 14 {
				SkipIfKey(nX, true)
			}
			if nY == 10 && nN == 1 {
				SkipIfKey(nX, false)
			}
		}
	case 15:
		{
			if nY == 0 && nN == 7 {
				SetVXToDelay(nX)
			}
			if nY == 1 && nN == 5 {
				SetDelayToVX(nX)
			}
			if nY == 1 && nN == 8 {
				SetBeepToVX(nX)
			}
			if nY == 1 && nN == 14 {
				AddToIndex(nX)
			}
			if nY == 0 && nN == 10 {
				GetKey(nX)
			}
			if nY == 2 && nN == 9 {
				FontCharacter(nX)
			}
			if nY == 3 && nN == 3 {
				DecimalConversion(nX)
			}
			if nY == 5 && nN == 5 {
				StoreMemory(nX)
			}
			if nY == 6 && nN == 5 {
				LoadMemory(nX)
			}
		}
	}
}
