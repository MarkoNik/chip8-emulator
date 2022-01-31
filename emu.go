package main

import "chip8/font"

var Memory [4096]byte

func main() {
	Init()

}

func Init() {
	font.InitFont()
}