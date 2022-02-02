package model

type stack struct {
	memory       [16]uint16
	stackPointer int
}

func (Stack *stack) Push(value uint16) {
	Stack.memory[Stack.stackPointer] = value
	Stack.stackPointer++
}

func (Stack *stack) Pop() uint16 {
	Stack.stackPointer--
	value := Stack.memory[Stack.stackPointer]
	return value
}
