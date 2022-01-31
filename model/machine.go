package model

var Memory [4096]byte
var Display [32][64]bool
var Stack [16]int16
var DelayTimer byte
var SoundTimer byte

var ProgramCounter int16
var IndexRegister int16
var Register [16]byte
