package main

import (
	"chip8/font"
	"chip8/model"
)

func Init() {
	font.InitFont()
	initEbiten()
}

func main() {
	model.Load()
	go model.Run()
	Init()
}
