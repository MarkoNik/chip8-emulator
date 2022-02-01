package model

type stack struct {
	memory       [16]uint16
	stackPointer int
}

func Push(stack stack, value uint16) {
	stack.stackPointer++
	stack.memory[stack.stackPointer] = value
}

func Pop(stack stack) (value uint16) {
	value = stack.memory[stack.stackPointer]
	stack.stackPointer--
	return value
}
