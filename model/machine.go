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
