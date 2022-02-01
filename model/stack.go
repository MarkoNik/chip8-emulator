package model

type stack struct {
	memory       [16]uint16
	stackPointer int
}

func (Stack stack) Push(value uint16) {
	Stack.stackPointer++
	Stack.memory[Stack.stackPointer] = value
}

func (Stack stack) Pop() uint16 {
	value := Stack.memory[Stack.stackPointer]
	Stack.stackPointer--
	return value
}
