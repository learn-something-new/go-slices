package main

import (
	"fmt"
	"github.com/obihann/GoSlices/array/src"
	//"github.com/obihann/GoSlices/copy/src"
	//"github.com/obihann/GoSlices/append/src"
)

func main() {
	stack := Stack.NewStack()
	size := 20

	for x := 0; x < size; x++ {
		if err := stack.Push(x + 1); err != nil {
			fmt.Printf("An error occured while trying to ush to the stack: %s\n", err)
			break
		} else {
			fmt.Printf("Pushing %v\n", (x + 1))
		}
	}

	for x := 0; x < size; x++ {
		if num, err := stack.Pop(); err != nil {
			fmt.Printf("An error occuured while trying to pop from the stack: %s.\n", err)
			break
		} else {
			fmt.Printf("Popped %v\n", num)
		}
	}
}
