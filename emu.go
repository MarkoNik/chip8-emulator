package main

import (
	"chip8/font"
	"chip8/model"
	"chip8/utils"
	"os"
)

func main() {
	initialize()
}

func initialize() {
	model.Load()
	go model.Run()
	font.InitFont()
	initSettings()
	initEbiten()
}

func initSettings() {
	dat, err := os.ReadFile("settings.txt")
	utils.Assert(err)
	if dat[0] == '1' {
		model.Legacy = true
	}
}
